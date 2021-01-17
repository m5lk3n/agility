package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func readConfigMap(clientset *kubernetes.Clientset) {

	configMapClient := clientset.CoreV1().ConfigMaps(*namespace)

	configMap, err := configMapClient.Get(context.TODO(), "k8s-df-config", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	excludePatternNamespace := configMap.Data["exclude-pattern-namespace"]
	fmt.Println("configured exclude pattern namespace:", excludePatternNamespace)
}

func handleDeployments(clientset *kubernetes.Clientset) {

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceAll)

	fmt.Println("listing deployments:")
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s\n", d.Name)
	}

	fmt.Println("watching deployments:")
	watcher, err := deploymentsClient.Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	ch := watcher.ResultChan()
	for event := range ch {
		deployment, ok := event.Object.(*v1.Deployment)
		if !ok {
			log.Fatalln("unexpected type")
		}
		switch event.Type {
		case watch.Added:
			fmt.Printf(" * %s deployed\n", deployment.Name)
		}
	}
}

var namespace *string

func main() {
	namespace = flag.String("namespace", "k8s-df", "namespace in which deployment watcher is deployed")
	flag.Parse()

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	readConfigMap(clientset)
	handleDeployments(clientset)
}
