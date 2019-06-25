# Go Cloud

A Golang boilerplate code to benchmark cloud platforms to host a Go applications.

## Deploy

Set up [Go](https://golang.org) on your device and run this command. It will automatically check if there is a PORT configured. Otherwhise, it will us `3000` as default PORT.

```sh
go run main.go
```

You can deploy on IBM Cloud in two different ways: using Go buildpack (default Golang) or using Static buildpack (NGINX).

### Go

```sh
ibmcloud app push -b go_buildpack
```

### Staticfile

```sh
ibmcloud app push
```

Support platforms:

- [IBM Cloud](https://cloud.ibm.com/devops/setup/deploy?repository=https://github.com/victorshinya/go-cloud)
- [Heroku](https://heroku.com/deploy?template=https://github.com/victorshinya/go-cloud/tree/master)

## License

MIT License

Copyright (c) 2019 Victor Kazuyuki Shinya
