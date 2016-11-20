# go-rabbitmq
![Build Status](https://travis-ci.org/sklarsa/go-rabbitmq.svg?branch=master)  ![Go Report Card](https://goreportcard.com/badge/github.com/sklarsa/go-rabbitmq)
#### An API Client written in Go for RabbitMQ's built-in REST API
This is an API client used to communicate with w remote RabbitMQ instance.  If you are just trying to query a node for information, or perhaps manipulate its behavior from a remote machine (without rabbitmqctl or another tool installed), this is the tool for you!

Included will be a full-featured API written in (hopefully) idiomatic Golang, as well as a command line tool for general use.

This is just a small side-project for me (and my first real foray into the world of Golang), but if anyone is interested in helping out, pull requests are more than welcome!

[Link to builds](https://travis-ci.org/sklarsa/go-rabbitmq)
[Link to go report card](https://goreportcard.com/report/github.com/sklarsa/go-rabbitmq)

## API Methods Implemented:
| Endpoint        | Methods           |
| --------------  |:-----------------:|
|~~/api/overview~~ | ~~GET~~ |
|/api/cluster-name | ~~GET~~, PUT |
|~~/api/node~~ | ~~GET~~
|~~/api/nodes/name~~ | ~~GET~~
|/api/extensions | GET
|/api/definitions | GET, POST
|/api/definitions/vhost | GET, POST
|/api/connections | GET
|/api/vhosts/vhost/connections | GET
|/api/connections/name | GET, DELETE
|/api/connections/name/channels | GET
|/api/channels | GET
|/api/vhosts/vhost/channels | GET
|/api/channels/channel | GET
|/api/consumers | GET
|/api/consumers/vhost | GET
|~~/api/exchanges~~ | ~~GET~~
|~~/api/exchanges/vhost~~ | ~~GET~~
|/api/exchanges/vhost/name | GET, PUT, DELETE
|/api/exchanges/vhost/name/bindings/source | GET
|/api/exchanges/vhost/name/bindings/destination | GET
|/api/exchanges/vhost/name/publish | POST
|/api/queues | GET
|/api/queues/vhost | GET
|/api/queues/vhost/name | GET, PUT, DELETE
|/api/queues/vhost/name/bindings | GET
|/api/queues/vhost/name/contents | DELETE
|/api/queues/vhost/name/actions | POST
|/api/queues/vhost/name/get | POST
|/api/bindings | GET
|/api/bindings/vhost | GET
|/api/bindings/vhost/e/exchange/q/queue | GET, POST
|/api/bindings/vhost/e/exchange/q/queue/props | GET, DELETE
|/api/bindings/vhost/e/source/e/destination | GET, POST
|/api/bindings/vhost/e/source/e/destination/props | GET, DELETE
|/api/vhosts | GET
|/api/vhosts/name | GET, PUT, DELETE
|/api/vhosts/name/permissions | GET
|/api/users | GET
|/api/users/name | GET, PUT, DELETE
/api/users/user/permissions | GET
/api/whoami | GET
/api/permissions | GET
/api/permissions/vhost/user | GET, PUT, DELETE
/api/parameters | GET
/api/parameters/component | GET
/api/parameters/component/vhost | GET
/api/parameters/component/vhost/name | GET, PUT, DELETE
/api/policies | GET
/api/policies/vhost | GET
/api/policies/vhost/name | GET, PUT, DELETE
/api/aliveness-test/vhost | GET
/api/healthchecks/node | GET
/api/healthchecks/node/node | GET
