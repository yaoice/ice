package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8sClient() *kubernetes.Clientset {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "conf/config", "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

var K8sDefaultClient = NewK8sClient()
