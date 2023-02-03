import { defineStore } from 'pinia'

//TODO url from env
export const configStore = defineStore("config", {
    state: () => ({
        url: "localhost:8080",
        loading: false
    })
})
