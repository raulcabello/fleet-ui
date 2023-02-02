import { defineStore } from 'pinia'
import axios from "axios";

export const gitRepoStore = defineStore("gitRepo", {
    state: () => ({
        gitRepo: {},
        bundles: []
    }),
    getters: {
        getGitRepo(state){
            return state.gitRepo
        },
        getBundles(state){
            return state.bundles
        },
        getResources(state){
            return state.gitRepo.resources
        },
        getConditions(state){
            return state.gitRepo.conditions
        }
    },
    actions: {
        async fetchGitRepo(name) {
            try {
                const data = await axios.get('http://localhost:8080/gitrepo/fleet-default/'+name)
                this.gitRepo = data.data

                this.bundles = this.gitRepo.bundles
                // watch for gitrepo changes
                let socket = new WebSocket("ws://localhost:8080/ws/gitrepo/"+name);

                socket.onopen = function(e) {
                    console.log("connected to ws")
                };

                socket.onmessage = (event) => {
                    console.log(`ws received: ${event.data}`)
                    this.gitRepo = JSON.parse(event.data)
                };

                socket.onclose =  (event) => {
                    if (event.wasClean) {
                        console.log("clean closed ws")
                    } else {
                        console.log("close error ws")
                    }
                };

                socket.onerror =  (event) => {
                    console.log("error ws")
                };

                // watch for gitrepo changes
                let socketBundles = new WebSocket("ws://localhost:8080/ws/bundles/"+name);

                socketBundles.onopen = function(e) {
                    console.log("bconnected to ws")
                };

                socketBundles.onmessage = (event) => {
                    console.log(`bws received: ${event.data}`)
                    let bundleFound = false
                    let newBundle = JSON.parse(event.data)
                    this.bundles.forEach((bundle, i)=> {
                        if (bundle.name === newBundle.name) {
                            this.bundles[i] = newBundle
                            bundleFound = true
                        }
                    })
                    if (!bundleFound) {
                        this.bundles.push(JSON.parse(event.data))
                    }
                };

                socketBundles.onclose =  (event) => {
                    if (event.wasClean) {
                        console.log("bclean closed ws")
                    } else {
                        console.log("bclose error ws")
                    }
                };

                socketBundles.onerror =  (event) => {
                    console.log("berror ws")
                };


            }
            catch (error) {
                alert(error)
                console.log(error)
            }
        }
    },
})
