import { defineStore } from 'pinia'
import axios from "axios";
import {configStore} from "./config";

export const gitRepoStore = defineStore("gitRepo", {
    state: () => ({
        gitRepo: {},
        bundles: []
    }),
    getters: {
        getConditions(state){
            return state.gitRepo.conditions
        },
        getResources(state){
            return state.gitRepo.resources
        }
    },
    actions: {
        async fetchGitRepo(name) {
            const config = configStore()
            try {
                const data = await axios.get('http://'+config.url+'/gitrepo/'+localStorage.namespace+'/'+name)
                this.gitRepo = data.data
                this.bundles = this.gitRepo.bundles

                // watch for gitrepo changes
                let socket = new WebSocket('ws://'+config.url+'/ws/gitrepo/'+localStorage.namespace+'/'+name);

                socket.onmessage = (event) => {
                    this.gitRepo = JSON.parse(event.data)
                };
                //TODO reconnect!
                socket.onclose =  (event) => {};
                socket.onerror =  (event) => {};

                // watch for gitrepo changes
                let socketBundles = new WebSocket('ws://'+config.url+'/ws/bundles/'+localStorage.namespace+'/'+name);
                socketBundles.onmessage = (event) => {
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

                //TODO reconnect!
                socketBundles.onclose =  (event) => {};
                socketBundles.onerror =  (event) => {};
            }
            catch (error) {
                // TODO handle errors!
                alert(error.message)
            }
        }
    },
})
