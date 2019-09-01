<p align="center">
    <img alt="Kubernets Logo" src="https://cdn.worldvectorlogo.com/logos/kubernets.svg" height="150" />
    <h2 align="center">Services, Load Balancing and Networking</h2>
</p>

`Service` is a resource you create to make a single, constant point of entry to a group of pods providing the same service. Each service has an IP address and port that never change while the service exists. Clients can open connections to that IP and port, and those connections are then routed to one of the pods backing that service. This way, clients of a service don’t need to know the location of individual pods providing the service, allowing those pods to be moved around the cluster at any time.

Creating services

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: kubia
spec:
  sessionAffinity: ClientIP
  ports:
    -
      port: 80
      targetPort: 8080
  selector:
    app: kubia
```

```
$ kubectl create -f kubia-svc.yaml
```

You’re defining a `service` called `kubia`, which will accept connections on port 80 and route each connection to port 8080 of one of the pods matching the `app=kubia` label selector.

You can list all Service resources in your namespace and see that an internal cluster IP has been assigned to your service:

```
$ kubectl get svc
```

The `kubectl exec` command allows you to remotely run arbitrary commands inside an existing container of a pod.

```
$ kubectl get pods
$ kubectl exec ${podName} bash
```

If you want all requests made by a certain client to be redirected to the same pod every time, you can set the service’s `sessionAffinity` property to `ClientIP` (instead of `None`, which is the default).

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: kubia
spec:
  sessionAffinity: ClientIP
  ports:
    -
      port: 80
      targetPort: 8080 // Port 80 is mapped to the pods’ port 8080.
  selector:
    app: kubia
```

This makes the service proxy redirect all requests originating from the same client IP to the same pod.

Services can also support multiple ports

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: kubia
spec:
  ports:
    -
      name: http
      port: 80
      targetPort: 8080 // Port 80 is mapped to the pods’ port 8080.
    -
      name: https
      port: 443
      targetPort: 8443 // Port 443 is mapped to the pods’ port 8443.
  selector:
    app: kubia
```

To use named ports, you can specify port names in a pod definition

```yaml
---
kind: Pod
spec:
  containers:
    -
      name: kubia
      ports:
        -
          containerPort: 8080
          name: http
        -
          containerPort: 8443
          name: https
```

You can then refer to those ports by name in the service spec

```yaml
---
apiVersion: v1
kind: Service
spec:
  ports:
    -
      name: http
      port: 80
      targetPort: http
    -
      name: https
      port: 443
      targetPort: https
```

The biggest benefit of doing so is that it enables you to change port numbers later without having to change the service spec.