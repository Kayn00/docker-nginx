# docker-nginx

没有vi或者vim命令，解决办法：apt-get update 完成之后 apt-get install vim

docker commit xxx xxx
docker exec -it xxx bash

docker run --name "docker-nginx" -v /Users/lx/docker-nginx/log:/var/log/nginx -v /Users/lx/docker-nginx/nginx.conf:/etc/nginx/nginx.conf -v /Users/lx/docker-nginx/conf.d:/etc/nginx/conf.d -v /Users/lx/docker-nginx/resource:/u01/resource -p 8080:80 -d docker-nginx