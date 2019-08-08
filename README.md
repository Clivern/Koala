<p align="center">
    <img alt="Koala Logo" src="https://raw.githubusercontent.com/Clivern/Koala/master/assets/img/koala.png" height="80" />
    <h3 align="center">Koala</h3>
    <p align="center">Kubernetes Playground for Testing Purposes.</p>
</p>

## Documentation

### Application Version 1.0.0:

Config the application

```console
$ export KOALA_PORT="8080"
$ export KOALA_REDIS_HOST="127.0.0.1"
$ export KOALA_REDIS_PORT="6379"
$ export KOALA_REDIS_PASSWORD=
```

Application Endpoints

```console
# App info
$ curl http://localhost:8080/ -v

# Health check
$ curl http://localhost:8080/_health -v

# Change app state
$ curl http://localhost:8080/_change -v

# Get current state
$ curl http://localhost:8080/_state -v

# Get the host down
$ curl http://localhost:8080/_hostup -v

# Get the host up
$ curl http://localhost:8080/_hostdown -v

# Get the all hosts up
$ curl http://localhost:8080/_kindup -v

# Get all hosts down
$curl http://localhost:8080/_kinddown -v
```


### Deployment Version 1.0.0:



## References

- [Kubernetes in Action.](https://www.manning.com/books/kubernetes-in-action)
- [How To Deploy a PHP Application with Kubernetes on Ubuntu 16.04.](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-php-application-with-kubernetes-on-ubuntu-16-04)
- [Kubernetes ConfigMaps and Secrets.](https://medium.com/google-cloud/kubernetes-configmaps-and-secrets-68d061f7ab5b)


## License

© 2019, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Koala** is authored and maintained by [@Clivern](http://github.com/clivern).
