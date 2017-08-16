FROM alpine:3.6
MAINTAINER David vonThenen <davidvonthenen@gmail.com>

ADD kubernetes-scaleio-prom /kubernetes-scaleio-prom
EXPOSE 80

ENTRYPOINT /kubernetes-scaleio-prom
