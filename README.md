## Introduction

Goload is a simple CLI tool to spin up a load balancer given a Dockerfile

Usually creating a load balancer takes a lot of configuration with tools such as nginx or HAPROXY. This simple CLI tool lets you spin up a simple load balancer with minimum config. The goal of this will be to allow users to start up a simple round robin load balancer very quick, but also be able to configure the load balancer to fit their needs.

## Usage

### Create a new project:

`goload init --name mynewproject`</br>
or specify a directory:</br>
`goload init --name mynewproject --dir ./mynodeproject`

### Run a load balancer with 4 replicas:

`goload run --replicas 4`

### Rebuild when changes are made:

`goload rebuild`

## Resources

[Commands](https://github.com/sharithg/goload/blob/master/COMMANDS.md)
[Todos](https://github.com/sharithg/goload/blob/master/TODO.md)
