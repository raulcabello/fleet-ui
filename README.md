## Fleet UI

> **Warning**
> Still under development. Not production ready!

### How to install it? 

> This is for testing the alpha release, it will change!

- Install `fleet`. See the [quick start](https://fleet.rancher.io/quickstart) for more info.
- Install `fleet-ui` chart
```
helm install fleet-ui https://github.com/raulcabello/fleet-ui-helm/releases/download/v0.0.1-alpha1/fleet-ui-0.0.1-alpha1.tgz  -n fleet-ui --create-namespace
```
You should see two pods in the `fleet-ui` namespace. One is for the backend, and the other for the frontend

- Use port-forward to access the fleet-ui.
```
kubectl port-forward svc/fleet-ui-svc -n fleet-ui http
```
and
```
kubectl port-forward svc/fleet-ui-backend-svc -n fleet-ui http
```

You can now access the fleet UI in http://localhost:9090

### Architecture

The backend is a go app that runs an http server on the port 8080.
It provides rest APIs for the frontend to create/list/delete `GitRepos`. 
It uses the wrangler API inside fleet to retrieve the `GitRepos` and `Bundles`. 

It provides a `WebSocket` API for the frontend to receive real time information. 
It uses the `Watch` feature of wrangler, then it pushes the changes detected to the `WebSocket`.

For now, I had to copy a few files in `fleet/pkg/apis`, this will be removed in the future!

The frontend is a Vue 3 app that uses bootstrap for the css.