#########################################################################
# File Name: run.sh
# Author: WenShuai
# mail: guowenshuai8207@163.com
# Created Time: Wed 04 Jul 2018 05:42:16 PM CST
#########################################################################
#!/bin/bash
bee pack

mv apiproject.tar.gz resource/server/

docker-compose up -d
