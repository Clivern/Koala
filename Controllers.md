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

We’ll rewrite the `ReplicationController` into a `ReplicaSet`

```yaml
---
apiVersion: apps/v1beta2
kind: ReplicaSet
metadata:
  name: koala
spec:
  replicas: 3
  selector:
    matchLabels:
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

The main improvements of `ReplicaSets` over `ReplicationControllers` are their more expressive label selectors. You intentionally used the simpler `matchLabels` selector in the first `ReplicaSet` example to see that `ReplicaSets` are no different from `ReplicationControllers`.

Now, you’ll rewrite the selector to use the more powerful `matchExpressions` property, as shown in the following listing.

```yaml
selector:
  matchExpressions:
    -
      key: app
      operator: In
      values:
        - koala
```

You can add additional expressions to the selector. As in the example, each expression must contain a key, an operator, and possibly (depending on the operator) a list of values. You’ll see four valid operators:

- `In` Label’s value must match one of the specified values.
- `NotIn` Label’s value must not match any of the specified values.
- `Exists` Pod must include a label with the specified key (the value isn’t important). When using this operator, you shouldn’t specify the values field.
- `DoesNotExist` Pod must not include a label with the specified key. The values property must not be specified.

If you specify multiple expressions, all those expressions must evaluate to true for the selector to match a pod. If you specify both matchLabels and matchExpressions, all the labels must match and all the expressions must evaluate to true for the pod to match the selector.


You can examine the `ReplicaSet` with `kubectl get` and `kubectl describe`

```
$ kubectl get rs
```

You can delete the `ReplicaSet` the same way you’d delete a `ReplicationController`

```
$ kubectl delete rs koala
```

Deleting the `ReplicaSet` should delete all the pods. List the pods to confirm that’s the case.
