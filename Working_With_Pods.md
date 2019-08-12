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

Labeling a node

```
$ kubectl get nodes
$ kubectl label node ${nodeName} gpu=true

# show nodes with gpu is true
$ kubectl get nodes -l gpu=true

# show nodes with gpu label
$ kubectl get nodes -L gpu
```

Scheduling pods to a specific nodes

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    app.kubernetes.io/component: database
    app.kubernetes.io/instance: wordpress-abcxzy
    app.kubernetes.io/managed-by: helm
    app.kubernetes.io/name: mysql
    app.kubernetes.io/part-of: wordpress
    app.kubernetes.io/version: "5.7.21"
  name: koala
spec:
  containers:
    -
      image: clivern/koala
      name: front
      ports:
        -
          containerPort: 8080
          protocol: TCP
  nodeSelector:
    gpu: "true"
```

Annontations can contain large data (up to 256 KB in total) unlike labels.

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  annotations:
    app.kubernetes.io/created-by: "{\"kind\": \"..\": \"apiVersion\": \"1.0.0\", \"Ref\": \"....\"}"
  labels:
    app.kubernetes.io/component: database
    app.kubernetes.io/instance: wordpress-abcxzy
    app.kubernetes.io/managed-by: helm
    app.kubernetes.io/name: mysql
    app.kubernetes.io/part-of: wordpress
    app.kubernetes.io/version: "5.7.21"
  name: koala
spec:
  containers:
    -
      image: clivern/koala
      name: front
      ports:
        -
          containerPort: 8080
          protocol: TCP
  nodeSelector:
    gpu: "true"
```

```
$ kubectl annotate pod ${podName} app.kubernetes.io/created-by="foo bar"
```

Namespaces used to separate objects so you can have the same resource names multiple times across different namespaces.

```
$ kubectl get namespaces
```

if you list pods without namespace, it will use the default namespace

```
$ kubectl get pods --namespace kuber-system
```

To create a namespace

```
apiVersion: v1
kind: Namespace
metadata:
  name: custom-namespace
```

```
$ kubectl create -f custom-namespace.yaml

// OR

$ kubectl create namespace custom-namespace
```

To create a resource on a namespace, either add a `namespace: custom-namespace` entry to `metadata`

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: koala
  namespace: custom-namespace
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

or specify a namespace when creating the resource

```
$ kubectl create -f koala.yaml -n custom-namespace
```

To delete a pod

```
$ kubectl delete pods ${podName}

# Delete all pods that has creation_method is manual
$ kubectl delete pods -l creation_method=manual
```

Deleting a namespace will delete all the pods attached automatically

```
$ kubectl delete namespace custom-namespace
```

To delete all pods on the current namespace

```
$ kubectl delete pods --all
```

To delete all pods, services and ReplicationController withing the current namespace

```
$ kubectl delete all --all
```
