package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
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
	s.router.GET("/bundles/:namespace", s.getBundles)
	s.router.GET("/bundles/:namespace/:name", s.getBundle)
	s.router.POST("/gitrepo", s.createGitRepo)
	s.router.DELETE("/gitrepos", s.deleteGitRepos)
	s.router.GET("/gitrepo/:namespace/:name", s.getGitRepo)
	s.router.GET("/ws", s.ws)

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

var upgrader = websocket.Upgrader{}

func (s *HTTP) ws(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	watcher, err := s.client.WatchGitRepo("fleet-default")
	if err != nil {
		fmt.Fprintf(w, "error!") //TODO
		return
	}
	for {
		select {
		case event := <-watcher.ResultChan():
			fmt.Println(event)
			bytes, err := json.Marshal(event)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ws.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				// TODO exit for!
				fmt.Println(err.Error())
				return
			}
		}
	}

	/*ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for {
		select {
		case <-ticker.C:
			fmt.Println("meess")
			err := ws.WriteMessage(websocket.TextMessage, []byte("message "+strconv.Itoa(r1.Int())))
			fmt.Println(err)
			//close connection when write: broken pipe ?
		case <-quit:
			ticker.Stop()
			return
		}
	}*/
}
