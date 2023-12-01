FROM golang:alpine

WORKDIR /opt/sagoo-iot-server/

COPY ./config/ /opt/sagoo-iot-server/config

COPY ./curl.sh /opt/sagoo-iot-server/
COPY ./entrypoint.sh /opt/sagoo-iot-server/
COPY ./sagoo-admin /opt/sagoo-iot-server/

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
RUN chmod +x /opt/sagoo-iot-server/sagoo-admin

EXPOSE 8200

CMD ["/bin/sh", "-c", "sleep 50 && /entrypoint.sh"]