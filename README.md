<p align="center">
    <img alt="Kubernets Logo" src="https://cdn.worldvectorlogo.com/logos/kubernets.svg" height="150" />
    <h3 align="center">Koala</h3>
    <p align="center">Kubernetes Playground for Testing Purposes</p>
    <p align="center">
        <a href="https://hub.docker.com/r/clivern/koala"><img src="https://img.shields.io/badge/Docker-Latest-orange"></a>
    </p>
</p>

## Documentation

### Application Endpoints

```console
# Current host info
$ curl http://localhost:8080/ -v

# Current host health check
$ curl http://localhost:8080/_health -v

# Change all hosts state
$ curl http://localhost:8080/_change -v

# Get current host state
$ curl http://localhost:8080/_state -v

# Get current host down
$ curl http://localhost:8080/_hostup -v

# Get current host up
$ curl http://localhost:8080/_hostdown -v

# Get all hosts up
$ curl http://localhost:8080/_kindup -v

# Get all hosts down
$ curl http://localhost:8080/_kinddown -v
```


## Third-party

- [Chaoskube:](https://github.com/linki/chaoskube) Periodically kills random pods on the cluster.
- [DO Container Storage Interface (CSI):](https://github.com/digitalocean/csi-digitalocean) Driver for DigitalOcean Block Storage.


## References

- [Kubernetes Documentation.](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)
- [Kubernetes in Action.](https://www.manning.com/books/kubernetes-in-action)
- [Kubernetes: Up and Running.](https://www.oreilly.com/library/view/kubernetes-up-and/9781491935668/)
- [How To Deploy a PHP Application with Kubernetes on Ubuntu 16.04.](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-php-application-with-kubernetes-on-ubuntu-16-04)
- [Kubernetes ConfigMaps and Secrets.](https://medium.com/google-cloud/kubernetes-configmaps-and-secrets-68d061f7ab5b)
- [How to write a Container Storage Interface (CSI) plugin.](https://arslan.io/2018/06/21/how-to-write-a-container-storage-interface-csi-plugin/)
- [Architecting Applications for Kubernetes.](https://www.digitalocean.com/community/tutorials/architecting-applications-for-kubernetes)
- [Understanding the Container Storage Interface (CSI).](https://medium.com/google-cloud/understanding-the-container-storage-interface-csi-ddbeb966a3b)


## License

© 2019, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Koala** is authored and maintained by [@Clivern](http://github.com/clivern).
