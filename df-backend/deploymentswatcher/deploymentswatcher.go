package deploymentswatcher

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
	nodeexporter "lttl.dev/k8s-df/df-backend/nodeexporter"
)

// flags (other than -log)
var (
	namespace     *string
	configmapName *string
)

// based upon config map keys
var (
	excludeNames      *regexp.Regexp // see exclude-regexp-pattern-names
	excludeNamespaces *regexp.Regexp // see exclude-regexp-pattern-namespaces
)

func readConfigMap(clientset *kubernetes.Clientset) {
	configMapClient := clientset.CoreV1().ConfigMaps(*namespace)

	configMap, err := configMapClient.Get(context.TODO(), *configmapName, metav1.GetOptions{})
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

func listDeployments(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceAll)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Debugln("listing deployments:")
	for _, d := range list.Items {
		log.Debugf(" * %s", d.Name)
	}
}

func watchDeployments(clientset *kubernetes.Clientset) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceAll)
	watcher, err := deploymentsClient.Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Info("watching deployments:")
	ch := watcher.ResultChan()
	for event := range ch {
		deployment, ok := event.Object.(*v1.Deployment)
		if !ok {
			log.Warningln("unexpected type in result channel")
			continue
		}

		switch event.Type {
		case watch.Added:
			if isIncluded(deployment) {
				log.Infof(" + %s deployed in namespace %s at %d", deployment.Name, deployment.Namespace, time.Now().UnixNano())
				nodeexporter.Add(nodeexporter.Deployment{Name: deployment.Name, Namespace: deployment.Namespace})
			} else {
				log.Infof(" - %s deployed in namespace %s but excluded from watching as configured", deployment.Name, deployment.Namespace)
			}
		case watch.Error:
			log.Errorln("error watching deployments:")
			log.Errorf("%+v\n", event.Object)
		}
	}
}

// Start a deploymentswatcher
func Start() {
	namespace = flag.String("namespace", "agility", "namespace in which deployment watcher is deployed")
	configmapName = flag.String("configmapName", "agility-configmap", "name of configmap used by deployment watcher")
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
	for {
		watchDeployments(clientset) // blocking call until an error occured (or Stop() is called on the underlying channel)
		log.Errorln("deployments watching stopped, restarting...")
	}
}
