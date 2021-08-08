# docker build -t roandocker/gin-scaffold:0.0.1 Dockerfile .
# docker build --target builder -t roandocker/gin-scaffold:0.0.1 .

FROM golang:stretch as builder

WORKDIR /opt/app/

COPY gin-scaffold/config/Config.yaml /opt/app/config/Config.yaml
COPY gin-scaffold/web /opt/app/web/

RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/goAlltime.git && \
	cd goAlltime/web/gin/gin-scaffold/ && \
    CGO_ENABLED=0 GOOS=linux go build -o /opt/app/gin-scaffold main.go


FROM ubuntu:16.04 as prod

WORKDIR /root/

COPY --from=0 /opt/app/gin-scaffold /root
COPY --from=0 /opt/app/config/Config.yaml  /root/config/Config.yaml
COPY --from=0 /opt/app/web /root/web/

RUN ls -all
RUN chmod 777 gin-scaffold

# health check
# HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
#   CMD curl -fs http://localhost:8080 || exit 1

CMD ["/bin/sh", "-c", "/root/gin-scaffold"]