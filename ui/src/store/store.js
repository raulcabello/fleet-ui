import { defineStore } from 'pinia'
import axios from "axios";

export const gitRepoStore = defineStore("gitRepo", {
    state: () => ({
        gitRepo: {},
    }),
    getters: {
        getGitRepo(state){
            return state.gitRepo
        },
        getBundles(state){
            return state.gitRepo.bundles
        },
        getResources(state){
            return state.gitRepo.resources
        }
    },
    actions: {
        async fetchGitRepo(name) {
            try {
                const data = await axios.get('http://localhost:8080/gitrepo/fleet-default/'+name)
                this.gitRepo = data.data
                console.log(this.gitRepo)
                setTimeout(() => {
                    this.gitRepo = {}
                    console.log("Delayed for 5 second.");
                }, "5000")
            }
            catch (error) {
                alert(error)
                console.log(error)
            }
        }
    },
})
