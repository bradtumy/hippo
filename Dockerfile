FROM alpine
MAINTAINER Brad Tumy

EXPOSE 8080

COPY ./hippo ./
COPY config/config.json ./config/config.json

ENTRYPOINT [ "./hippo" ]