package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func readConfigMap(clientset *kubernetes.Clientset) {

	const namespace = "k8s-df"
	const configMapName = "k8s-df-config"

	configMapClient := clientset.CoreV1().ConfigMaps(namespace)

	configMap, err := configMapClient.Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	excludePattern := configMap.Data["exclude-pattern"]
	fmt.Println("configured exclude pattern:", excludePattern)
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig) // TODO: add in-cluster support, see "From a cluster" https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	readConfigMap(clientset)

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceAll)

	// list deployments
	fmt.Println("listing deployments:")
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s\n", d.Name)
	}

	// watch deployments
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
