package client

import "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"

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

func convertGitRepoList(v1alpha1GitRepoList *v1alpha1.GitRepoList) *GitRepoList {
	gitRepoList := &GitRepoList{}
	for _, item := range v1alpha1GitRepoList.Items {
		gitRepoListItem := GitRepoListItem{
			State:         "active", //TODO! check rancher code!
			Name:          item.Name,
			RepoName:      item.Spec.Repo,
			RepoCommit:    item.Spec.Branch + "@" + item.Spec.Revision,
			ClustersReady: "11", //TODO
			Resources:     string(rune(len(item.Status.Resources))),
			Age:           item.CreationTimestamp.String(),
		}
		gitRepoList.Items = append(gitRepoList.Items, gitRepoListItem)
	}

	return gitRepoList
}
