FROM alpine
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
ADD mikudos_message_deliver-srv /mikudos_message_deliver-srv
WORKDIR /
ENTRYPOINT [ "/mikudos_message_deliver-srv" ]
