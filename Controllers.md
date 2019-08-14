<p align="center">
    <img alt="Kubernets Logo" src="https://cdn.worldvectorlogo.com/logos/kubernets.svg" height="150" />
    <h2 align="center">Controllers</h2>
</p>


### ReplicationController

A `ReplicationController`’s job is to make sure that an exact number of pods always matches its label selector. If it doesn’t, the `ReplicationController` takes the appropriate action to reconcile the actual with the desired number.

```yaml
---
apiVersion: v1
kind: ReplicationController
metadata:
  name: koala
spec:
  replicas: 3
  selector:
    app: koala
template:
  metadata:
    labels:
      app: koala
  spec:
    containers:
      -
        image: clivern/koala
        name: koala
        ports:
          -
            containerPort: 8080
```

No need to specify a pod selector when defining a `ReplicationController`. Let kubernetes extract it from the pod template. This will keep your YAML shorter and simpler.

```yaml
---
apiVersion: v1
kind: ReplicationController
metadata:
  name: koala
spec:
  replicas: 3
template:
  metadata:
    labels:
      app: koala
  spec:
    containers:
      -
        image: clivern/koala
        name: koala
        ports:
          -
            containerPort: 8080
```

```
$ kubectl create -f koala-rc.yaml
```

Now, let’s see what information the `kubectl get` command shows for replication controllers

```
$ kubectl get rc
```

```
$ kubectl describe rc koala
```

if you overrite one of the pods label, replication controller will spin another pod to reach the desired state.

```
# changing the labels of a managed pod
$ kubectl label pod ${podName} app=foo --overwrite
$ kubectl get pods --show-labels
```

Editing the `ReplicationController` definition or scaling up or scaling down

```
$ kubectl edit rc koala
```

Delete replication controller without the managed pods

```
$ kubectl delete rc koala --cascade=false
```

Delete replication controller with the managed pods

```
$ kubectl delete rc koala
```

Initially `ReplicationControllers` were the only Kubernetes component for replicating pods and rescheduling them when nodes failed. Later, a similar resource called a `ReplicaSet` was introduced. It’s a new generation of `ReplicationController` and replaces it completely (`ReplicationControllers` will eventually be deprecated).


### ReplicaSet

