#########################################################################
# File Name: clean.sh
# Author: WenShuai
# mail: guowenshuai8207@163.com
# Created Time: Mon 25 Jun 2018 05:22:50 PM CST
#########################################################################
#!/bin/bash


docker stop apiproject-server && docker rm apiproject-server
docker stop apiproject-client && docker rm apiproject-client
docker stop apiproject-mysql && docker rm apiproject-mysql

## remove docker image
docker rmi wenshuai/apiproject wenshuai/apiprojectclient