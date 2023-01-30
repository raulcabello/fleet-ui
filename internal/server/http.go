package server

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
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
	//s.router.GET("/gitrepos/:namespace/:name", s.getGitRepos)
	s.router.GET("/bundles/:namespace", s.getBundles)
	s.router.GET("/bundles/:namespace/:name", s.getBundle)
	// Add CORS support (Cross Origin Resource Sharing)
	handler := cors.Default().Handler(s.router)

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

/*
func (s *HTTP) getGitRepo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list, err := s.client.GetGitRepo(ps.ByName("namespace"))
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
*/
