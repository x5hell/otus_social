FROM alpine:3.11

COPY run /run

RUN apk add mysql-client && \
    chmod +x /run/*.sh

ENTRYPOINT ["sh", "/run/run.sh"]