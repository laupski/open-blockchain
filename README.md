# Open Blockchain [![CircleCI](https://circleci.com/gh/laupski/open-blockchain.svg?style=svg)](https://circleci.com/gh/laupski/open-blockchain) [![Go Report Card](https://goreportcard.com/badge/github.com/laupski/open-blockchain)](https://goreportcard.com/report/github.com/laupski/open-blockchain)
Simple proof of concept for blockchain technology written in Go.

## Requirements
* `go` installed
* `gcc` installed 

## Installation
Install `go`: 

To install `gcc` on Windows: https://golang.org/doc/install
```
choco install mingw
# verify installation
gcc --version
```
If you don't have `choco`, search MinGW for installation.

To install `gcc` on Linux:
```
sudo apt update
sudo apt install build-essential
# verify installation
gcc --version
```

## Build 
`go build -o obc`

## Run
`obc`

## Future Features
* Persistence
* P2P
* CLI
* API
* Dockerize
* UI