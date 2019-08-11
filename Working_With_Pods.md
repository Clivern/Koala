<p align="center">
    <img alt="Kubernets Logo" src="https://cdn.worldvectorlogo.com/logos/kubernets.svg" height="150" />
</p>


Working With Pods
-----------------

Get a full YAML of a deployed pod

```
$ kubectl get pods $podname -o yaml
$ kubectl get pods $podname -o json
```

A simple YAML descriptor for a pod

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: koala
  labels:
    app.kubernetes.io/name: mysql
    app.kubernetes.io/instance: wordpress-abcxzy
    app.kubernetes.io/version: "5.7.21"
    app.kubernetes.io/component: database
    app.kubernetes.io/part-of: wordpress
    app.kubernetes.io/managed-by: helm
spec:
  containers:
    -
      image: clivern/koala
      name: front
      ports:
        -
          containerPort: 8080
          protocol: TCP
```

To create a pod from yaml file

```
$ kubectl create -f koala.yaml
```

To get all pods

```
$ kubectl get pods
$ kubectl get pods --show-labels
$ kubectl get pods -L app.kubernetes.io/name,app.kubernetes.io/version

# Add label to pod
$ kubectl label pods ${podName} app.kubernetes.io/author=clivern

# Overwrite a pod label
$ kubectl label pods ${podName} app.kubernetes.io/author=clivern --overwrite
```

To get pod logs

```
$ kubectl logs ${podName}
$ kubectl logs ${podName} -c ${containerName}
```

To talk to a pod without going through a service, you can forward a local port to a port of the pod

```
# this will forward local port 8000 to port 8080 of a pod
$ kubectl port-forward ${podName} 8000:8080
```

Listing pods using a label selector

```
$ kubectl get pods -l environment=production,tier=frontend
$ kubectl get pods -l 'environment in (production),tier in (frontend)'
$ kubectl get pods -l 'environment in (production, qa)'
$ kubectl get pods -l 'environment,environment notin (frontend)'

// Field selectors let you select Kubernetes resources based on the value of one or more resource fields.
$ kubectl get pods --field-selector status.phase=Running
$ kubectl get pods --field-selector metadata.name=my-service
$ kubectl get pods --field-selector metadata.namespace!=default
$ kubectl get pods --field-selector status.phase=Pending
```
