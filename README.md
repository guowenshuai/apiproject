# apiproject


## how to run

Use `docker-compose` deploy the application

```
$ docker-compose up -d

```
Waiting finished, `docker ps` check containers
```
$ docker ps
CONTAINER ID        IMAGE                       COMMAND                  CREATED             STATUS              PORTS                      NAMES
0a9fc304b2f1        wenshuai/apiproject         "/bin/sh -c apipro..."   5 seconds ago       Up 3 seconds        0.0.0.0:38080->8080/tcp    apiproject_server
0f8877ec4a87        wenshuai/apiprojectclient   "nginx -g 'daemon ..."   6 seconds ago       Up 5 seconds        0.0.0.0:30080->80/tcp      apiproject_client
4b170c7691ce        mysql:5.7                   "docker-entrypoint..."   6 seconds ago       Up 4 seconds        3306/tcp                   apiproject_db

```
http API server Running on http://localhost:38080
webSocket client Running on http://localhost:30080

## how to test

### use postman to test RESUful API

import `resource/postman/trainForRancher.postman_collection.json` to postman

the default port is 38080
you can change the port if needed


### test webSocket
open http://localhost:30080
it is a chat Room, join then send message

## clean all
It will clean all running containers and remove apiproject images
```
$ bash clean.sh
```

---------------------
OLD VERSION README


## how to run apiproject

you can run source code, binary file or as a docker service

**recommend run as a docker service**

### run source code

- install bee tool

```
go get github.com/beego/bee

```

- add your `$GOPATH/bin` to PATH

```
bee run
```
http server Running on http://localhost:8080

### run binary file

```
$ mkdri test && cd test
$ cp ../apiproject.tar.gz .
$ tar -zxvf apiproject.tar.gz
$ ./apiproject
```
http server Running on http://localhost:8080

### docker service

- build `wenshuai/apiproject` image

```
$ cd resource/server
$ bash build.sh
```

- run as a docker service named `apiproject-server`

you should run `apiproject-mysql` service before start `apiproject-server`, otherwise the `apiproject-server` can't started

```
$ bash run.sh
$ docker ps
CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS                      NAMES
712bd0edd2a6        wenshuai/apiproject      "/bin/sh -c apipro..."   5 seconds ago       Up 4 seconds        0.0.0.0:38080->8080/tcp    apiproject-server
```
http server Running on http://localhost:38080



## `apiproject-mysql`

`apiproject-mysql` is a db service for `apiproject-server`

### run as a docker service

- start
```
$ cd resource/sql
$ bash run.sh
$ docker ps
CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS                      NAMES
7ebb2651dd7c        mysql:5.7                "docker-entrypoint..."   45 seconds ago      Up 45 seconds       0.0.0.0:33306->3306/tcp    apiproject-mysql
```

- import data to mysql

```
$ mysql -h 127.0.0.1 -P 33306 -u root -proot < test_work_jobs.sql
```

mysql Running on http://localhost:33306


### the default parameters of mysql

- user: root
- password: root
- port: `33306`
- database: test_work


## how to test

### use postman to test RESUful API

import `resource/postman/trainForRancher.postman_collection.json` to postman

the default port is 38080
you can change the port if needed



### test websocket

`apiproject-client` is a docker service used to test websocket

`apiproject-client` is a service based on a nginx container

```
$ cd resource/client
$ bash run.sh
$ docker ps
CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS              PORTS                      NAMES
54b78334620f        nginx                    "nginx -g 'daemon ..."   23 minutes ago      Up 21 minutes       0.0.0.0:30080->80/tcp      apiproject-client
```

websocket clinet Running on http://localhost:30080

