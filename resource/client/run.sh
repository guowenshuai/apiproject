#########################################################################
# File Name: run.sh
# Author: WenShuai
# mail: guowenshuai8207@163.com
# Created Time: Tue 26 Jun 2018 09:58:40 AM CST
#########################################################################
#!/bin/bash


docker run --name apiproject-client -it -d -p 30080:80 -v `pwd`/index.html:/usr/share/nginx/html/index.html nginx:1.14

# `pwd`/index.html
