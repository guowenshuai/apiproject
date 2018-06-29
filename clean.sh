#########################################################################
# File Name: clean.sh
# Author: WenShuai
# mail: guowenshuai8207@163.com
# Created Time: Fri 29 Jun 2018 05:48:06 PM CST
#########################################################################
#!/bin/bash


docker-compose stop && docker-compose rm -f
bash resource/server/clean.sh