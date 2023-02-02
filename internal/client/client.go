package client

import (
	"fmt"
	fleetcontrollers "github.com/rancher/fleet/pkg/generated/controllers/fleet.cattle.io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/watch"
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
	return ConvertBundle(bundle), nil
}

func (c *Client) GetGitRepo(namespace, name string) (*GitRepo, error) {
	gitrepo, err := c.factory.Fleet().V1alpha1().GitRepo().Get(namespace, name, v1.GetOptions{})
	labelSelector := v1.LabelSelector{MatchLabels: map[string]string{"fleet.cattle.io/repo-name": name}}
	bundles, err := c.factory.Fleet().V1alpha1().Bundle().List(namespace, v1.ListOptions{LabelSelector: labels.Set(labelSelector.MatchLabels).String()})

	if err != nil {
		return nil, err
	}
	return ConvertGitRepo(gitrepo, bundles), nil
}

func (c *Client) GetGitRepoList(namespace string) (*GitRepoList, error) {
	list, err := c.factory.Fleet().V1alpha1().GitRepo().List(namespace, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return convertGitRepoList(list), nil
}

func (c *Client) CreateGitRepo(request *GitRepoRequest) error {
	gitRepo := convertGitRepoRequest(request)
	_, err := c.factory.Fleet().V1alpha1().GitRepo().Create(gitRepo)
	return err
}

func (c *Client) DeleteGitRepos(gitRepoNames []string) error {
	for _, name := range gitRepoNames {
		err := c.factory.Fleet().V1alpha1().GitRepo().Delete("fleet-default", name, &v1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) WatchGitRepo(namespace, name string) (watch.Interface, error) {
	return c.factory.Fleet().V1alpha1().GitRepo().Watch(namespace, v1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", name),
	})
}

func (c *Client) WatchBundles(namespace, repoName string) (watch.Interface, error) {
	labelSelector := v1.LabelSelector{MatchLabels: map[string]string{"fleet.cattle.io/repo-name": repoName}}

	return c.factory.Fleet().V1alpha1().Bundle().Watch(namespace, v1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})
}
