#!/bin/bash

name="stock_wangyi"
tag="latest"
dockerfile="Dockerfile"

build_img(){
  echo -e "[Start to build image]\nname:$name\ntag:$tag"
  rm -rf wangyi.stock.tar.gz
  tar -zcf wangyi.stock.tar.gz PPGo_amaze/
  docker build -t "$name:$tag" . < $dockerfile
}

restart_container(){
  echo -e "stop container analize"
  docker rm -f $(docker ps | grep analize | awk '{print $1}')
  echo -e "start container analize"
  docker run --privileged -tid --restart=always -h analize --name=analize --net=host tushare-stock:latest python ./TushareStock/manage.py runserver 0.0.0.0:8000
}
build_img
#restart_container
