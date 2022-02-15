# docker-nginx

没有vi或者vim命令，解决办法：apt-get update 完成之后 apt-get install vim

docker commit xxx xxx
docker exec -it xxx bash

docker run --name "docker-nginx" --link=beego1:t1 --link=beego2:t2 -v /Users/lx/docker-nginx/log:/var/log/nginx -v /Users/lx/docker-nginx/nginx.conf:/etc/nginx/nginx.conf -v /Users/lx/docker-nginx/conf.d:/etc/nginx/conf.d -v /Users/lx/docker-nginx/resource:/u01/resource -p 8080:80 -d docker-nginx

修改配置后 docker restart xxx 重启容器生效

启动测试服务器
docker run --name "beego1" -p 8081:8081 -v /Users/lx/docker-nginx/beego1/src/:/go/src/ -d golang bash -c "cd src/beegoDemo/ && go run main.go"

docker run --name "beego2" -p 8082:8082 -v /Users/lx/docker-nginx/beego2/src/:/go/src/ -d golang bash -c "cd src/beegoDemo/ && go run main.go"