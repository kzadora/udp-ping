# udp-ping
A test of UDP-based communication between various Azure regions.

## Prerequisites
1. Go language (the server and the client are written in Go).
2. Azure CLI and Docker Desktop for Windows/Mac or [Docker with Compose CLI (Linux)](https://docs.docker.com/engine/context/aci-integration/#install-the-docker-compose-cli-on-linux) -- for building the client and server containers and deploying them to different Azure regions.


## Test the program locally (as needed)

1.  **Start the server**
    ```shell
    go run server/server.go 17335
    ```
    `17335` is the port number that the server will use. You can use a different one, but you will need to change the server Dockerfile accordingly for Azure deployment.

2.  **Run the client**
    ```shell
    go run client/client.go 127.0.0.1:17335 <optional-payload>
    ```

## Build the server container

```shell
cd server
docker build -t udp-ping-server:201127a .
```

## Prepare Azure resources for server deployment

We'll use Azure Container Registry (ACR) here, but Docker Hub should work equally well.

If necessary, do `az login` followed by `az account show` to verify that correct Azure account is active. Use `az account set` to change it.

1.  **Create Azure resource groups for ACR and server deployments**

In this experiment we will use Australia East, Brazil South and West Europe regions. 

```shell
az group create --name udp-ping-australia --location australiaeast
az group create --name udp-ping-brazil --location brazilsouth
az group create --name udp-ping-westeurope --location westeurope
```

2.  **Create Azure Account Registry**

```shell
z acr create --name udppingacr201127a --resource-group udp-ping-westeurope --location westeurope --sku basic
```

3.  **Create Docker contexts for deploying server to Azure regions**

Do `docker login azure` as necessary to enable Docker-Azure interaction.

```shell
docker context create aci udp-ping-australia --resource-group udp-ping-australia
docker context create aci udp-ping-brazil --resource-group udp-ping-brazil
docker context create aci udp-ping-westeurope --resource-group udp-ping-westeurope
```

4.  **Push the server image to account registry**

```shell
az acr login --name udppingacr201127a
docker tag udp-ping-server:201127a udppingacr201127a.azurecr.io/udp-ping-server:201127a
docker push udppingacr201127a.azurecr.io/udp-ping-server:201127a
```

## Deploy server to Azure regions and ping it with the client

```shell
docker --context udp-ping-westeurope run -d -p 17335/udp udppingacr201127a.azurecr.io/udp-ping-server:201127a
docker --context udp-ping-westeurope ps
# Note the IP address of the server
go run client/client.go <server-IP-address>:17335
```

Repeat the same with `udp-ping-australia` and `udp-ping-brazil` Docker contexts/Azure regions. Note that the image to run (`udppingacr201127a.azurecr.io/udp-ping-server:201127a`) does not change, only the context changes.

## When finished, clean up Azure resources
(TBD)