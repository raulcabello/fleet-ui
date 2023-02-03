package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/raulcabello/fleet-ui/internal/k8s"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type HTTP struct {
	client *k8s.Client
	router *httprouter.Router
}

func NewHttp(client *k8s.Client) HTTP {
	return HTTP{client, httprouter.New()}
}

func (s *HTTP) Start() {
	s.router.GET("/gitrepos/:namespace", s.getGitRepos)
	s.router.GET("/bundles/:namespace", s.getBundles)
	s.router.GET("/bundles/:namespace/:name", s.getBundle)
	s.router.POST("/gitrepo", s.createGitRepo)
	s.router.DELETE("/gitrepos/:namespace", s.deleteGitRepos)
	s.router.GET("/gitrepo/:namespace/:name", s.getGitRepo)
	s.router.GET("/ws/gitrepo/:namespace/:name", s.wsGitRepo)
	s.router.GET("/ws/bundles/:namespace/:repoName", s.wsGitRepoBundles)

	// TODO handle CORS
	handler := cors.AllowAll().Handler(s.router)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func writeResponseOK(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}

func (s *HTTP) getGitRepos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetGitRepoList(ps.ByName("namespace"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	data, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	writeResponseOK(w, data)
}

func (s *HTTP) getBundles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetBundleList(ps.ByName("namespace"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	data, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	writeResponseOK(w, data)
}

func (s *HTTP) getBundle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetBundle(ps.ByName("namespace"), ps.ByName("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	data, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeResponseOK(w, data)
}

func (s *HTTP) createGitRepo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	gitRepoRequest := &k8s.GitRepoRequest{}
	err := json.NewDecoder(r.Body).Decode(gitRepoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.client.CreateGitRepo(gitRepoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (s *HTTP) deleteGitRepos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var gitRepoNames []string
	err := json.NewDecoder(r.Body).Decode(&gitRepoNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.client.DeleteGitRepos(gitRepoNames, ps.ByName("namespace"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (s *HTTP) getGitRepo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetGitRepo(ps.ByName("namespace"), ps.ByName("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResponseOK(w, data)
}

//TODO
/*
This is creating a new watch per websocket connection, which is not very efficient. There should be just one k8s watch
that broadcast changes to all the websocket connections.
*/
// wsGitRepo establishes a websocket connection, and it leaves the connection opened while watching for GitRepo changes.
// If a change is detected, a new message is sent in the websocket.
func (s *HTTP) wsGitRepo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	upgrader := websocket.Upgrader{}
	//TODO handle CORS!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer ws.Close()

	watcher, err := s.client.WatchGitRepo(ps.ByName("namespace"), ps.ByName("name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
Watch:
	for {
		select {
		case event := <-watcher.ResultChan():
			v1alpha1GitRepo, ok := event.Object.(*v1alpha1.GitRepo)
			if !ok {
				http.Error(w, err.Error(), http.StatusBadRequest)
				break Watch
			}
			gitRepo := k8s.ConvertGitRepo(v1alpha1GitRepo, &v1alpha1.BundleList{})
			bytes, err := json.Marshal(gitRepo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				break Watch
			}
			err = ws.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				watcher.Stop()
				break Watch
			}
		}
	}
}

// wsGitRepo establishes a websocket connection, and it leaves the connection opened while watching for Bundles changes
// for a specific GitRepo identified by the name. If a change is detected, a new message is sent in the websocket.
func (s *HTTP) wsGitRepoBundles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	upgrader := websocket.Upgrader{}
	//TODO handle CORS!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer ws.Close()

	watcher, err := s.client.WatchGitRepoBundles(ps.ByName("namespace"), ps.ByName("repoName"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
Watch:
	for {
		select {
		case event := <-watcher.ResultChan():
			v1alpha1Bundle, ok := event.Object.(*v1alpha1.Bundle)
			if !ok {
				http.Error(w, err.Error(), http.StatusBadRequest)
				break Watch
			}
			bundle := k8s.ConvertBundle(v1alpha1Bundle)
			bytes, err := json.Marshal(bundle)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				break Watch
			}
			err = ws.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				watcher.Stop()
				break Watch
			}
		}
	}
}
