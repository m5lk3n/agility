package main

import (
	"context"
	"flag"
	"regexp"
	"time"

	logflag "github.com/elankath/logflag"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// flags (other than -log)
var (
	namespace *string
)

// based upon config map keys
var (
	excludeNames      *regexp.Regexp // see exclude-pattern-names
	excludeNamespaces *regexp.Regexp // see exclude-pattern-namespaces
)

func readConfigMap(clientset *kubernetes.Clientset) {

	configMapClient := clientset.CoreV1().ConfigMaps(*namespace)

	configMap, err := configMapClient.Get(context.TODO(), "k8s-df-config", metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	excludeNames = createRegexp(configMap.Data["exclude-regexp-pattern-names"])
	excludeNamespaces = createRegexp(configMap.Data["exclude-regexp-pattern-namespaces"])
}

func createRegexp(pattern string) *regexp.Regexp {

	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Warnf("could not create regexp for pattern '%s', skipping!", pattern)
		return nil
	}

	return r
}

func isIncluded(deployment *v1.Deployment) bool {

	if excludeNames != nil {
		if excludeNames.MatchString(deployment.Name) {
			return false
		}
	}

	if excludeNamespaces != nil {
		if excludeNamespaces.MatchString(deployment.Namespace) {
			return false
		}
	}

	return true
}

func handleDeployments(clientset *kubernetes.Clientset) {

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceAll)

	log.Debugln("listing deployments:")
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, d := range list.Items {
		log.Debugf(" * %s\n", d.Name)
	}

	log.Info("watching deployments:")
	watcher, err := deploymentsClient.Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	ch := watcher.ResultChan()
	for event := range ch {
		deployment, ok := event.Object.(*v1.Deployment)
		if !ok {
			log.Fatalln("unexpected type")
		}
		switch event.Type {
		case watch.Added:
			if isIncluded(deployment) {
				log.Infof(" + %s deployed in namespace %s at %d\n", deployment.Name, deployment.Namespace, time.Now().UnixNano())
			} else {
				log.Infof(" - %s deployed in namespace %s but excluded from watching as configured\n", deployment.Name, deployment.Namespace)
			}
		}
	}
}

func main() {
	namespace = flag.String("namespace", "k8s-df", "namespace in which deployment watcher is deployed")
	flag.Parse()
	logflag.Parse() // Call after regular flag.Parse()

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	readConfigMap(clientset)
	handleDeployments(clientset)
}
