package client

import (
	"github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/rancher/wrangler/pkg/genericcondition"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type GitRepoListItem struct {
	State                 string `json:"state"`
	Name                  string `json:"name"`
	RepoName              string `json:"repoName"`
	RepoCommit            string `json:"repoCommit"`
	ClustersReady         string `json:"clustersReady"`
	ResourcesDesiredReady int    `json:"resourcesDesiredReady"`
	ResourcesReady        int    `json:"resourcesReady"`
	Age                   string `json:"age"`
}

type GitRepoList struct {
	Items []GitRepoListItem `json:"items"`
}

type GitRepoRequest struct {
	Name                  string               `json:"name"`
	RepoUrl               string               `json:"repoUrl"`
	BranchName            string               `json:"branchName,omitempty"`
	Revision              string               `json:"revision,omitempty"`
	RepoName              string               `json:"repoName,omitempty"`
	GitSecretName         string               `json:"state,omitempty"`
	HelmSecretName        string               `json:"helmSecretName,omitempty"`
	CABundle              string               `json:"CABundle,omitempty"`
	InsecureSkipTLSVerify bool                 `json:"insecureSkipTLSVerify,omitempty"`
	TargetNamespace       string               `json:"targetNamespace,omitempty"`
	ServiceAccount        string               `json:"serviceAccount,omitempty"`
	Paths                 []string             `json:"paths,omitempty"`
	Targets               []v1alpha1.GitTarget `json:"targets,omitempty"`
}

type ResourceCount struct {
	DesiredReady int `json:"desiredReady"`
	Missing      int `json:"missing"`
	Modified     int `json:"modified"`
	NotReady     int `json:"notReady"`
	Orphaned     int `json:"orphaned"`
	Ready        int `json:"ready"`
	Unknown      int `json:"Unknown"`
	WaitApplaied int `json:"waitApplaied"`
}

type GitRepoResources struct {
	State      string `json:"state"`
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"` //TODO
	Cluster    string `json:"cluster"`   //TODO
}

type GitRepo struct {
	Name                string                              `json:"name"`
	Age                 string                              `json:"age"`
	DisplayBundlesReady string                              `json:"displayBundlesReady"`
	ResourceCount       ResourceCount                       `json:"resourceCount"`
	Resources           []GitRepoResources                  `json:"resources"`
	Bundles             []*Bundle                           `json:"bundles"`
	Conditions          []genericcondition.GenericCondition `json:"conditions"`
}

func ConvertGitRepo(v1alpha1GitRepo *v1alpha1.GitRepo, bundles *v1alpha1.BundleList) *GitRepo {
	resources := []GitRepoResources{}
	for _, resource := range v1alpha1GitRepo.Status.Resources {
		namespace, _, _ := strings.Cut(resource.ID, "/")
		resources = append(resources, GitRepoResources{
			State:      resource.State,
			APIVersion: resource.APIVersion,
			Kind:       resource.Kind,
			Name:       resource.Name,
			Namespace:  namespace,
			Cluster:    "", //TODO
		})
	}
	bundlesList := []*Bundle{}
	for _, bundle := range bundles.Items {
		bundlesList = append(bundlesList, ConvertBundle(&bundle))
	}
	return &GitRepo{
		Name:                v1alpha1GitRepo.Name,
		Age:                 v1alpha1GitRepo.CreationTimestamp.String(),
		DisplayBundlesReady: v1alpha1GitRepo.Status.Display.ReadyBundleDeployments,
		ResourceCount: ResourceCount{
			DesiredReady: v1alpha1GitRepo.Status.ResourceCounts.DesiredReady,
			Missing:      v1alpha1GitRepo.Status.ResourceCounts.Missing,
			Modified:     v1alpha1GitRepo.Status.ResourceCounts.Modified,
			NotReady:     v1alpha1GitRepo.Status.ResourceCounts.NotReady,
			Orphaned:     v1alpha1GitRepo.Status.ResourceCounts.Orphaned,
			Ready:        v1alpha1GitRepo.Status.ResourceCounts.Ready,
			Unknown:      v1alpha1GitRepo.Status.ResourceCounts.Unknown,
			WaitApplaied: v1alpha1GitRepo.Status.ResourceCounts.WaitApplied,
		},
		Resources:  resources,
		Bundles:    bundlesList,
		Conditions: v1alpha1GitRepo.Status.Conditions,
	}
}

func convertGitRepoList(v1alpha1GitRepoList *v1alpha1.GitRepoList) *GitRepoList {
	gitRepoList := &GitRepoList{}
	for _, item := range v1alpha1GitRepoList.Items {
		state := item.Status.Display.State
		if item.Status.DesiredReadyClusters == item.Status.ReadyClusters && item.Status.Display.State == "" {
			state = "Active"
		}
		gitRepoListItem := GitRepoListItem{
			State:                 state,
			Name:                  item.Name,
			RepoName:              item.Spec.Repo,
			RepoCommit:            item.Spec.Branch + "@" + item.Spec.Revision,
			ClustersReady:         item.Status.Display.ReadyBundleDeployments,
			ResourcesDesiredReady: item.Status.ResourceCounts.DesiredReady,
			ResourcesReady:        item.Status.ResourceCounts.DesiredReady,
			Age:                   item.CreationTimestamp.String(),
		}
		gitRepoList.Items = append(gitRepoList.Items, gitRepoListItem)
	}

	return gitRepoList
}

func convertGitRepoRequest(request *GitRepoRequest) *v1alpha1.GitRepo {
	gitrepo := &v1alpha1.GitRepo{
		ObjectMeta: v1.ObjectMeta{
			Name:      request.Name,
			Namespace: "fleet-default",
		},
		Spec: v1alpha1.GitRepoSpec{
			Repo:                  request.RepoUrl,
			Branch:                request.BranchName,
			Revision:              request.Revision,
			TargetNamespace:       request.TargetNamespace,
			ClientSecretName:      request.GitSecretName,
			HelmSecretName:        request.HelmSecretName,
			CABundle:              []byte(request.CABundle),
			InsecureSkipTLSverify: request.InsecureSkipTLSVerify,
			Paths:                 request.Paths,
			Paused:                false, //TODO?
			ServiceAccount:        request.ServiceAccount,
			Targets:               request.Targets,
		},
	}

	return gitrepo
}
