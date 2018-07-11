FROM scratch
MAINTAINER Brad Tumy

EXPOSE 9091

COPY ./server ./
COPY config/config.json ./config/config.json

ENTRYPOINT [ "./server" ]