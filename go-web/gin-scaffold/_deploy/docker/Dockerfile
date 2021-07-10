FROM golang:stretch as builder

WORKDIR /opt/app/

COPY config/Config.yaml /opt/app/config/Config.yaml

# RUN cd gin-scaffold && GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /opt/app/gin-scaffold main.go

RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/goAlltime.git && \
	cd goAlltime/go-web/gin-scaffold/ && \
    CGO_ENABLED=0 GOOS=linux go build -o /opt/app/gin-scaffold-cgi main.go

FROM ubuntu:16.04 as prod

WORKDIR /root/

COPY --from=0 /opt/app/config/Config.yaml  /root/config/Config.yaml
COPY --from=0 /opt/app/gin-scaffold-cgi  /root/gin-scaffold-cgi

RUN ls -all
RUN chmod 777 gin-scaffold-cgi

CMD ["/bin/sh", "-c", "/root/gin-scaffold-cgi"]