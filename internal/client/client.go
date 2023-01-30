package client

import (
	fleetcontrollers "github.com/rancher/fleet/pkg/generated/controllers/fleet.cattle.io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Client struct {
	factory *fleetcontrollers.Factory
}

func NewClient() (*Client, error) {
	config := ctrl.GetConfigOrDie()
	factory, err := fleetcontrollers.NewFactoryFromConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{factory: factory}, nil
}

// TODO filter by namespace!
func (c *Client) GetBundleList(namespace string) (*BundleList, error) {
	list, err := c.factory.Fleet().V1alpha1().Bundle().List(namespace, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return convertBundleList(list), nil
}

func (c *Client) GetBundle(namespace, name string) (*Bundle, error) {
	bundle, err := c.factory.Fleet().V1alpha1().Bundle().Get(namespace, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return convertBundle(bundle), nil
}

func (c *Client) GetGitRepoList(namespace string) (*GitRepoList, error) {
	list, err := c.factory.Fleet().V1alpha1().GitRepo().List(namespace, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return convertGitRepoList(list), nil
}
