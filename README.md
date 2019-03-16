# GO-FF

A modern microservices Fly For Fun V15 Emulator.

## Quick Start

To get a working environment you'll just need `docker`, `kubectl` & either docker-desktop or minikube

* Docker : https://docs.docker.com/install/
* Kubernetes : https://kubernetes.io/docs/setup/

And setup basic component to the Kubernetes cluster

```bash
make install
```

Then run ...

```bash
make run
```

---

Retreive the IP of the VM that executes docker & put it in your client.
For modern Flyff clients with custom launch arguments you can specify this IP at launch

```bash
neuz.exe sunkist 127.0.0.1
```

## Architecture

Here's the complete architecture overview of the project (Open-it with Draw.io)

* https://drive.google.com/file/d/1ouAbHP_F27fSzD8aAxr3OK-l1cX32Mvg/view?usp=sharing

### Services

These descriptions relates to ACTUAL state of the services. It might evolve & grow ! Refer to the Architecture diagram to get a more complete description of their usage.

---

* Asynchronous Messaging - Message Broker
  * Engine : RabbitMQ
* Persistent SQL Storage
  * Engine : PostgreSQL
  * Using : Storing player data for a long period
* In-memory Caching SQL Storage
  * Engine : MariaDB (MEMORY Db Engine)
  * Using : Storing player in-game state for the time they're connected
* Database Administration Tool
  * Name : Adminer
  * Port : 80
* Game Services
  * Out of Game
    * Login
      * Language : Golang
      * Port : 23000
      * Stateless : `yes`
      * Using : Handling player account login (actually static data are used)
    * Cluster
      * Language : Golang
      * Port : 28000
      * Stateless : `yes`
      * Using : Handling player character management
  * In Game
    * ConnectionServer
      * Language : Golang
      * Port : 5400
      * Stateless : `not yet`
      * Using
        * Handling player connection
        * Handling player packets
        * Dispatching incoming player packets to the Broker
        * Receiving ougoing packets to be broadcaster to players
    * Entity
      * Language : Golang
      * Stateless : `yes`
      * Using
        * Handling player connection & disconnection
        * Broadcasting spawn & despawn events
    * Moving
      * Language : Golang
      * Stateless : `yes`
      * Using
        * Handling player moves (WASD, Mouse)
        * Handling player motions (Jump etc...)
    * Chat
      * Language : Golang
      * Stateless : `yes`
      * Using
        * Handling player public chat
    * Action
      * Language : Golang
      * Stateless : `yes`
      * Using
        * Handling player inventory actions
