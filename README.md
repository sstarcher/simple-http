# Simple Golang HTTP Server

Runs a simple golang http server on port 8080.  Responds to requests with `Hello, World`, but when the `Accept` header specifies `application/json` as supported it will respond with a json formatted message.  If the `-debug` parameter is specified it will log each request URL.

## Prerequisites
* Golang 1.12.4 - https://golang.org/doc/install
* Curl

## Setup
* `go get github.com/sstarcher/simple-http`
* `cd $GOPATH/src/github.com/sstarcher/simple-http`
* `GO111MODULE=on go mod vendor`

## Run
* `go run .`

## Tests
* `go test ./...`

## Options
* `-debug` - Will enable debug logging and output each request

## Tests
* Non-json responses of `<p>Hello, World</p>`
  * `curl 'http://localhost:8080'`
  * `curl 'http://localhost:8080/'`
  * `curl -H 'Accept: bob' 'http://localhost:8080'`
* Json responses of `{"message": "Good Morning"}`
  * `curl -H 'Accept: application/json' 'http://localhost:8080'`
  * `curl -H 'Accept: application/json,text/html' 'http://localhost:8080'`
  * `curl -H 'Accept: text/html,application/json' 'http://localhost:8080'`
* Empty responses
  * `curl 'http://localhost:8080/blah/other?and=things&other=stuff'`
  * `curl -H 'Accept: application/json' 'http://localhost:8080/blah/other?and=things&other=stuff'`
  * `curl 'http://localhost:8080/blah/other?and=things&other=stuff'`

## Assumptions
* GET and POST methods are both accepted.  POST is not specified as to what it should do so it's blank.
* GET method only returns data on `/` and is ignored with any parameters.
* Debug logging should happen on all HTTP methods not just GET.
* When `application/json` is specified as one of the `Accept` headers we respond with json.  We behave the same way if multiple comma seperated `Accept` headers are specified.
