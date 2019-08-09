<p align="center">
    <img alt="Koala Logo" src="https://raw.githubusercontent.com/Clivern/Koala/master/assets/img/koala.png" height="80" />
    <h3 align="center">Koala</h3>
    <p align="center">Kubernetes Playground for Testing Purposes.</p>
</p>

## Documentation

### Testing Application v1:

Config the application

```console
$ export KOALA_PORT="8080"
$ export KOALA_REDIS_HOST="127.0.0.1"
$ export KOALA_REDIS_PORT="6379"
$ export KOALA_REDIS_PASSWORD=
```

Application Endpoints

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


### Testing Application Deployment v1:


### Chaos Maker v1:


## References

- [Kubernetes in Action.](https://www.manning.com/books/kubernetes-in-action)
- [How To Deploy a PHP Application with Kubernetes on Ubuntu 16.04.](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-php-application-with-kubernetes-on-ubuntu-16-04)
- [Kubernetes ConfigMaps and Secrets.](https://medium.com/google-cloud/kubernetes-configmaps-and-secrets-68d061f7ab5b)
- [Understanding the Container Storage Interface (CSI).](https://medium.com/google-cloud/understanding-the-container-storage-interface-csi-ddbeb966a3b)


## License

© 2019, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Koala** is authored and maintained by [@Clivern](http://github.com/clivern).
