# Simple Object Storage

## Table of Contents
- [Overview](#overview)
- [Install](#install)
- [Usage](#usage)
- [Implementation](#implementation)
- [Dependencies](#dependencies)
- [References](#references)
- [License](#license)

## Overview
In this research class, we are implementing a distributed object storage for small objects based 
on Facebook’s Haystack paper. The project aims to address the challenges associated with efficiently 
storing and fetching a large volume of small objects in a distributed architecture.

Key objectives of the project are following
* Designing and implementing a distributed object storage based on the ideas from the Facebook haystack paper 
* Optimizing the object storage for large volumes of small files
* Minimizing the metadata operations on small files, thus improving the read performance of the system

## Install
To build the application run the following command.

```shell
go build src/sos.go
```

To start the primary node run the following command
```shell
./sos primary --port=8080 --grpc_port=1234
```


To start to data node run the following command
```shell
./sos data --vol_dir="tmp/node1" --primary_node="localhost:1212" --port="8081" --node_id="1"
```


## Usage

## Implementation

## Dependencies
This project uses the following libraries.
- [Logrus logging library](https://github.com/sirupsen/logrus)
- [Echo library for rest api](https://echo.labstack.com/)
- [Cobra library to build cli](https://github.com/spf13/cobra)

## References
- [Finding a needle in Haystack: Facebook’s photo storage](papers/Haystack.pdf)
- [Ceph Architecture Guide](https://access.redhat.com/documentation/en-us/red_hat_ceph_storage/4/html/architecture_guide/the-ceph-architecture_arch)
- [CFS: A Distributed File System for Large Scale Container Platforms](papers/1911.03001.pdf)
- [Echo library](https://echo.labstack.com/docs)
- [Heartbeats in Golang](https://medium.com/geekculture/heartbeats-in-golang-1a12c4c366f)
- [GRPC Docs](https://grpc.io/docs/languages/go/basics/)

## License
The project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.