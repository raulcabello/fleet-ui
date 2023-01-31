package client

import (
	"github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GitRepoListItem struct {
	State         string `json:"state"`
	Name          string `json:"name"`
	RepoName      string `json:"repoName"`
	RepoCommit    string `json:"repoCommit"`
	ClustersReady string `json:"clustersReady"`
	Resources     string `json:"resources"`
	Age           string `json:"age"`
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

func convertGitRepoList(v1alpha1GitRepoList *v1alpha1.GitRepoList) *GitRepoList {
	gitRepoList := &GitRepoList{}
	for _, item := range v1alpha1GitRepoList.Items {
		gitRepoListItem := GitRepoListItem{
			State:         "active", //TODO! check rancher code!
			Name:          item.Name,
			RepoName:      item.Spec.Repo,
			RepoCommit:    item.Spec.Branch + "@" + item.Spec.Revision,
			ClustersReady: "11", //TODO
			Resources:     "2",  //TODO string(len(item.Status.Resources)),
			Age:           item.CreationTimestamp.String(),
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
