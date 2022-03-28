# Brincation Go

Just for fun weekend project about microservices and RPC communication.

## How To 

To run this project, you will need of [docker-compose](https://docs.docker.com/compose/install/) on your machine.

After installation, you can run:

```bash
$ docker-compose up
```

The command above will build and run the docker images, because we have two services.

#### clientapi

When the project starts, it loads all the items from the [ports.json](./ports.json) file into the `portdomainservice` project via RPC calls. It also has a Rest API that allows you to search ports using your ID.

#### portdomainservice

The service is responsible for saving the ports into database. It uses a simple memory database that can manage concurrent calls.

### Tech stack

- Golang
- RPC using Twirp
- Docker and Docker Compose

### Project structure

#### cmd folder

This folder contains the two services (clientapi, portdomainservice) that are used in the project. Here a binary is assembled with all its dependencies, depending on the project and needs.

#### internal folder

This folder we have two others folders that contain the portdomainservice repository logic (repo folder) and the portdomainservice server implementation (server folder).

#### pkg folder

This folder we build the memory database used by portdomainservice.

#### rpc folder

This folder we build the .proto file used by portdomainservice. Here is only allowed edit the .proto file. The *pb.go and *twirp.go are build using this command:

```bash
$ protoc --go_out=paths=source_relative:. --twirp_out=paths=source_relative:. rpc/ports.proto
```


