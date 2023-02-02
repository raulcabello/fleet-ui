import { defineStore } from 'pinia'

//TODO namespace should be selected in a dropdown
//TODO url from env
export const configStore = defineStore("config", {
    state: () => ({
        namespace: "fleet-default",
        url: "localhost:8080"
    })
})
