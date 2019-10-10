
# go-reverse-proxy

A Go app, which can easily be deployed to Heroku.

## Running Locally
Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ git clone https://github.com/heroku/go-reverse-proxy.git
$ cd go-reverse-proxy
$ go build -o bin/go-reverse-proxy -v . # or `go build -o bin/go-reverse-proxy.exe -v .` in git bash
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
