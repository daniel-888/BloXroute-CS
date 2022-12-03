# BloXroute-CS  Client-Server
This application contains two apps
* server
* client

Client app can be executed by cli, infinitely sends messages to RabbitMQ queue.
* `client-app AddItem`
* `client-app RemoveItem`
* `client-app GetItem`
* `client-app GetAllItems`

Server polls messages from the queue and handles commands. It uses custom linkedhashmap structure.

## Prerequisites
You have to have installed:
* Docker, Docker Compose
* make

## How to run
### Simulation purpose
> make run-demo-docker-compose

It will start locally:
* RabbitMQ
* Server application
* 4 client applications
  * first will send AddItem command
  * second will send GetItem command
  * third will send RemoveItem command
  * fourth will send GetAllItems command


### Manual start
Run each one in different terminal.
1. `make run-rabbit-mq`
2. `make run-server`
3. `make run-client-add-item`
4. `make run-client-remove-item`
5. `make run-client-get-item`
6. `make run-client-get-all-items`

### Run tests
`make run-tests`