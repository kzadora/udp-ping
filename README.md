# udp-ping
A test of UDP-based communication between various Azure regions.

## Prerequisites
1. Go language (the server and the clinet are written in Go).
2. Azure CLI and [Docker with Compose CLI](https://docs.docker.com/engine/context/aci-integration/#install-the-docker-compose-cli-on-linux) for building the client and server containers and deploying them to different Azure regions.


## Local test

1.  **Start the server**
    ```shell
    go run server/server.go <port>
    ```
    Substitute `<port>` with desired port value, e.g. 17335.

2.  **Run the client**
    ```shell
    go run client/client.go 127.0.0.1:<server-port> <optional-payload>
    ```
    Use the same server port value as in step 1.


## Build the server container

```shell
cd server
docker build -t udp-ping-server:201127a .
```