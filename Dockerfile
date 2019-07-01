FROM golang:1.9

#update system timezone
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime; \
    echo "Asia/Shanghai" >> /etc/timezone

WORKDIR /go/src
ADD wangyi.stock.tar.gz $WORKDIR
COPY PPGo_amaze/run /usr/bin/
RUN chmod 777 /usr/bin/run

#build
RUN cd PPGo_amaze; go build PPGo_amaze
#RUN go get -d -v ./...
#RUN go install -v ./...

ENV WORKDIR /go/src
CMD ["run"]
