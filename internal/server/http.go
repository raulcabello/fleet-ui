package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/raulcabello/fleet-ui/internal/client"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type HTTP struct {
	client *client.Client
	router *httprouter.Router
}

func NewHttp(client *client.Client) HTTP {
	return HTTP{client, httprouter.New()}
}

func (s *HTTP) Start() {
	s.router.GET("/gitrepos/:namespace", s.getGitRepos)
	s.router.GET("/bundles/:namespace", s.getBundles)
	s.router.GET("/bundles/:namespace/:name", s.getBundle)
	s.router.POST("/gitrepo", s.createGitRepo)
	s.router.DELETE("/gitrepos", s.deleteGitRepos)
	s.router.GET("/gitrepo/:namespace/:name", s.getGitRepo)
	s.router.GET("/ws/gitrepo/:name", s.wsGitRepo)
	s.router.GET("/ws/bundles/:repoName", s.wsBundles)

	// TODO Add CORS support (Cross Origin Resource Sharing)
	handler := cors.AllowAll().Handler(s.router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func (s *HTTP) getGitRepos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetGitRepoList(ps.ByName("namespace"))
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data) //TODO handle error
}

func (s *HTTP) getBundles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetBundleList(ps.ByName("namespace"))
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data) //TODO handle error
}

func (s *HTTP) getBundle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetBundle(ps.ByName("namespace"), ps.ByName("name"))
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data) //TODO handle error
}

func (s *HTTP) createGitRepo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	gitRepoRequest := &client.GitRepoRequest{}
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

func (s *HTTP) deleteGitRepos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var gitRepoNames []string
	err := json.NewDecoder(r.Body).Decode(&gitRepoNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.client.DeleteGitRepos(gitRepoNames)
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
		fmt.Fprintf(w, "error!") //TODO
	}
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(data) //TODO handle error

}

func (s *HTTP) wsGitRepo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	upgrader := websocket.Upgrader{}

	//TODO handle CORS!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
		return
	}
	defer ws.Close()

	//filter by label name!
	watcher, err := s.client.WatchGitRepo("fleet-default", ps.ByName("name"))
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
		return
	}
	for {
		select {
		case event := <-watcher.ResultChan():
			fmt.Println(event)

			v1alpha1GitRepo, ok := event.Object.(*v1alpha1.GitRepo)
			if !ok {
				// TODO exit for!
				fmt.Println(err.Error())
				return
			}

			gitRepo := client.ConvertGitRepo(v1alpha1GitRepo, &v1alpha1.BundleList{})
			bytes, err := json.Marshal(gitRepo)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//check event modified
			err = ws.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				// TODO exit for!
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func (s *HTTP) wsBundles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	upgrader := websocket.Upgrader{}

	//TODO handle CORS!
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
		return
	}
	defer ws.Close()

	watcher, err := s.client.WatchBundles("fleet-default", ps.ByName("repoName"))
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
		return
	}
	for {
		select {
		case event := <-watcher.ResultChan():
			fmt.Println(event)

			v1alpha1Bundle, ok := event.Object.(*v1alpha1.Bundle)
			if !ok {
				// TODO exit for!
				fmt.Println(err.Error())
				return
			}

			bundle := client.ConvertBundle(v1alpha1Bundle)
			bytes, err := json.Marshal(bundle)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//check event modified
			err = ws.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				// TODO exit for!
				fmt.Println(err.Error())
				return
			}
		}
	}
}
