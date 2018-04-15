package models

import (
	"github.com/yaoice/ice/client"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	v1beta1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//    "github.com/astaxie/beego/logs"
)

var (
	clientset = client.K8sDefaultClient
	Flavors   map[string]*Flavor
)

type Flavor struct {
	CPU              int64
	Memory           int64
	EphemeralStorage int64
}

type App struct {
	AppName       string
	Replicas      int32
	Image         string
	ContainerPort int32
	FlavorRef     string
}

func init() {
	Flavors = make(map[string]*Flavor)
	Flavors["0"] = &Flavor{1, 1 * 1024 * 1024 * 1024, 1 * 1024 * 1024 * 1024}
	Flavors["1"] = &Flavor{1, 1 * 1024 * 1024 * 1024, 10 * 1024 * 1024 * 1024}
	Flavors["2"] = &Flavor{2, 2 * 1024 * 1024 * 1024, 20 * 1024 * 1024 * 1024}
	Flavors["3"] = &Flavor{4, 4 * 1024 * 1024 * 1024, 40 * 1024 * 1024 * 1024}
	Flavors["4"] = &Flavor{8, 8 * 1024 * 1024 * 1024, 80 * 1024 * 1024 * 1024}
}

func int32Ptr(i int32) *int32                         { return &i }
func quantity(q *resource.Quantity) resource.Quantity { return *q }

func CreateApp(project string, app *App) (result *v1beta1.Deployment, err error) {
	deploymentsClient := clientset.AppsV1beta1().Deployments(project)
	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: app.AppName,
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr(app.Replicas),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": app.AppName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  app.AppName,
							Image: app.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          app.AppName,
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: app.ContainerPort,
								},
							},
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									"cpu": quantity(resource.NewQuantity(
										Flavors[app.FlavorRef].CPU,
										resource.DecimalSI,
									)),
									"memory": quantity(resource.NewQuantity(
										Flavors[app.FlavorRef].Memory,
										resource.DecimalSI,
									)),
									"ephemeral-storage": quantity(resource.NewQuantity(
										Flavors[app.FlavorRef].EphemeralStorage,
										resource.DecimalSI,
									)),
								},
							},
						},
					},
				},
			},
		},
	}
	result, err = deploymentsClient.Create(deployment)
	return result, err
}

func GetApp(project string, appName string) (app *v1beta1.Deployment, err error) {
	deploymentsClient := clientset.AppsV1beta1().Deployments(project)
	app, err = deploymentsClient.Get(appName, metav1.GetOptions{})
	return app, err
}

func GetAllApp(project string) (list *v1beta1.DeploymentList, err error) {
	deploymentsClient := clientset.AppsV1beta1().Deployments(project)
	list, err = deploymentsClient.List(metav1.ListOptions{})
	return list, err
}

func DeleteApp(project string, appName string) (err error) {
	deploymentsClient := clientset.AppsV1beta1().Deployments(project)
	deletePolicy := metav1.DeletePropagationForeground
	err = deploymentsClient.Delete(appName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	return err
}
