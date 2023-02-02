package client

import "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"

type BundleListItem struct {
	State       string `json:"state"`
	Name        string `json:"name"`
	Deployments string `json:"deployments"`
	Age         string `json:"age"`
}

type BundleList struct {
	Items []BundleListItem `json:"items"`
}

type BundleResource struct {
	State      string `json:"state"`
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

type Bundle struct {
	State       string           `json:"state"`
	Name        string           `json:"name"`
	Deployments string           `json:"deployments"`
	Date        string           `json:"date"`
	LastUpdated string           `json:"lastUpdated"`
	Resources   []BundleResource `json:"resources"`
}

func convertBundleList(v1alpha1BundleList *v1alpha1.BundleList) *BundleList {
	bundleList := &BundleList{}
	for _, item := range v1alpha1BundleList.Items {
		bundleListItem := BundleListItem{
			State:       "active", //TODO! check rancher code!
			Name:        item.Name,
			Deployments: "10",                            //TODO! check rancher code!
			Age:         item.CreationTimestamp.String(), //TODO! check rancher code!
		}
		bundleList.Items = append(bundleList.Items, bundleListItem)
	}

	return bundleList
}

func ConvertBundle(v1alpha1Bundle *v1alpha1.Bundle) *Bundle {
	state := v1alpha1Bundle.Status.Display.State
	if v1alpha1Bundle.Status.Summary.DesiredReady == v1alpha1Bundle.Status.Summary.Ready && v1alpha1Bundle.Status.Display.State == "" {
		state = "Active"
	}
	bundle := &Bundle{
		State:       state,
		Name:        v1alpha1Bundle.Name,
		Deployments: v1alpha1Bundle.Status.Display.ReadyClusters,
		Date:        v1alpha1Bundle.CreationTimestamp.String(),
	}

	for _, resource := range v1alpha1Bundle.Spec.Resources {
		bundleResource := BundleResource{
			State:      "Ready", //TODO! check rancher code!
			APIVersion: "v1?",   //TODO! check rancher code!
			Kind:       "Kind",  //TODO! check rancher code!
			Name:       resource.Name,
			Namespace:  "namespace", //TODO! check rancher code!
		}
		bundle.Resources = append(bundle.Resources, bundleResource)
	}

	return bundle
}
