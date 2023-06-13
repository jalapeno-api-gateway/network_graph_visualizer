# Network Graph Visualizer API
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/marcom4rtinez/network_graph_api/main?style=for-the-badge)

This part of the Network Graph Visualizer is the API written in Golang.

## What does the API do

The API checks the Jalapeño API Gateway for the network data and topology and provides it via a RESTful API.

## How to run the API

You can either build the docker container from the repository or run it directly.

```
cd build 
docker build -t network_graph_api:latest -f Dockerfile ../
docker run -p 8080:8080 network_graph_api:latest
```

## How to run it directly from the repository

Navigate to the files and run the following:

```
go mod tidy
go run main.go
```

Make sure to set the JAGW in configs/config.yml

## Configuration

Make sure to configure the API under `configs/config.yml`