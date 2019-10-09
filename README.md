# validator

## Requirements

* go v1.10+
* helm
* docker
* kubectl

## Config

To update config variables connect to consul

```
kubectl port-forward config-consul-0 8500:8500
```

go to http://localhost:8500/ui and edit the key/values you need


## Release

First you must have access to the consul in the environment you wish to deploy to.

### access consul

```

kubectl port-forward config-consul-0 8500:8500

```

### deploy

To release a new image simply run `make` && `make deploy` which will build and push a new docker image then release the assosiated helm chart.

If you have new dependencies run `make deps` before releasing.
