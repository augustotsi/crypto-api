FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/crypto-api /opt/bin/crypto-api
ADD ./static /opt/static

CMD ["/opt/bin/crypto-api"]

