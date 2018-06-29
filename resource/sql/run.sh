#########################################################################
# File Name: run.sh
# Author: WenShuai
# mail: guowenshuai8207@163.com
# Created Time: Tue 26 Jun 2018 10:08:09 AM CST
#########################################################################
#!/bin/bash

##### create mysql
# password root
# database test_work
docker run --name apiproject-mysql -it -d -p 33306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_work mysql:5.7


## import data
#mysql -h 127.0.0.1 -P 33306 -u root -proot < test_work_jobs.sql


