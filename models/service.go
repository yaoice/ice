package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type Service struct {
	AppName string
	//    Protocol string
	Port       int32
	TargetPort int32
	//    ServiceType string
}

var raw_servicejson = `
{
	"kind": "Service",
	"apiVersion": "v1",
	"metadata": {
		"name": "%s",
		"creationTimestamp": null
	},
	"spec": {
		"ports": [{
			"protocol": "TCP",
			"port": %d,
			"targetPort": %d
		}],
		"selector": {
			"app": "%s"
		}
	},
	"status": {
		"loadBalancer": {}
	}
}
`

func CreateService(project string, s *Service) (result *v1.Service, err error) {
	serviceClient := clientset.CoreV1().Services(project)

	servicejson := fmt.Sprintf(raw_servicejson, s.AppName, s.Port, s.TargetPort, s.AppName)

	//    fmt.Printf("%s", servicejson)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, e := decode([]byte(servicejson), nil, nil)
	if err != nil {
		logs.Error("service json decode err: %v", e)
	}

	service := obj.(*v1.Service)
	result, err = serviceClient.Create(service)
	return result, err
}

func GetService(project string, serviceName string) (result *v1.Service, err error) {
	serviceClient := clientset.CoreV1().Services(project)
	result, err = serviceClient.Get(serviceName, metav1.GetOptions{})
	return result, err
}

func DeleteService(project string, serviceName string) (err error) {
	serviceClient := clientset.CoreV1().Services(project)
	deletePolicy := metav1.DeletePropagationForeground
	err = serviceClient.Delete(serviceName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	return err
}
