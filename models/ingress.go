package models

import (
    v1beta1 "k8s.io/api/extensions/v1beta1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes/scheme"
    "github.com/astaxie/beego/logs"
    "fmt"
)

type Ingress struct {
    AppName    string
    Host        string
    ServicePort int32
    Path        string
}

var raw_ingressjson = `
{
	"apiVersion": "extensions/v1beta1",
	"kind": "Ingress",
	"metadata": {
		"name": "%s",
		"namespace": "%s"
	},
	"spec": {
		"rules": [{
			"host": "%s",
			"http": {
				"paths": [{
					"backend": {
						"serviceName": "%s",
						"servicePort": %d
					},
					"path": "%s"
				}]
			}
		}]
	}
}
`

func CreateIngress(project string, i *Ingress) (result *v1beta1.Ingress, err error) {
    ingressClient := clientset.ExtensionsV1beta1().Ingresses(project)

    ingressjson := fmt.Sprintf(raw_ingressjson, i.AppName,
        project, i.Host, i.AppName, i.ServicePort, i.Path)

//    fmt.Printf("%s", ingressjson)
    decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, e := decode([]byte(ingressjson), nil, nil)
	if err != nil {
        logs.Error("ingress json decode err: %v", e)
	}

    ingress := obj.(*v1beta1.Ingress)
    result, err = ingressClient.Create(ingress)
    return result, err
}

func GetIngress(project string, ingressName string) (result *v1beta1.Ingress, err error) {
    ingressClient := clientset.ExtensionsV1beta1().Ingresses(project)
    result, err = ingressClient.Get(ingressName, metav1.GetOptions{})
    return result, err
}

func DeleteIngress(project string, ingressName string) (err error) {
    ingressClient := clientset.ExtensionsV1beta1().Ingresses(project)
    deletePolicy := metav1.DeletePropagationForeground
    err = ingressClient.Delete(ingressName, &metav1.DeleteOptions{
        PropagationPolicy: &deletePolicy,
    })
    return err
}

