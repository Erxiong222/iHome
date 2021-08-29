#! /bin/bash

# redis fastdfs tracker storage consul
sudo redis-server /etc/redis/redis.conf

sudo fdfs_trackerd /etc/fdfs/tracker.conf

sudo fdfs_storaged /etc/fdfs/storage.conf

sudo /usr/local/nginx/sbin/nginx

consul agent -dev

# micro-service start

go run /home/lijia/GoProject/ihome/service/userOrder/main.go

go run /home/lijia/GoProject/ihome/service/user/main.go

go run /home/lijia/GoProject/ihome/service/register/main.go

go run /home/lijia/GoProject/ihome/service/house/main.go

go run /home/lijia/GoProject/ihome/service/getImg/main.go

go run /home/lijia/GoProject/ihome/service/getArea/main.go

