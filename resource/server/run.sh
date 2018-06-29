#########################################################################
# File Name: run.sh
# Author: WenShuai
# mail: guowenshuai8207@163.com
# Created Time: Mon 25 Jun 2018 05:22:39 PM CST
#########################################################################
#!/bin/bash


docker run --name apiproject-server -it -d --env-file envFile -p 38080:8080 wenshuai/apiproject
