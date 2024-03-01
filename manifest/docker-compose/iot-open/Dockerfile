FROM golang:alpine AS builder
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /opt/sagoo-iot-server/

COPY ./config/ /opt/sagoo-iot-server/config

COPY ./curl.sh /opt/sagoo-iot-server/
COPY ./entrypoint.sh /opt/sagoo-iot-server/
COPY ./sagooiot /opt/sagoo-iot-server/

COPY ./entrypoint.sh /opt/sagoo-iot-server/
RUN chmod +x /opt/sagoo-iot-server/entrypoint.sh
RUN chmod +x /opt/sagoo-iot-server/sagooiot

COPY ./resource/public /opt/sagoo-iot-server/resource/public/
COPY ./resource/rsa /opt/sagoo-iot-server/resource/rsa/
COPY ./resource/app.lic /opt/sagoo-iot-server/resource

EXPOSE 8200

CMD ["/bin/sh", "-c", "sleep 50 && /opt/sagoo-iot-server/entrypoint.sh"]