FROM alpine:3.7 AS compile
RUN apk add --no-cache bash curl git go musl-dev libpcap-dev
RUN git clone https://github.com/mongodb/mongo-tools.git && \
    cd mongo-tools && \
    git checkout tags/r4.0.0 && \
    . ./set_gopath.sh && \
    # go build -o bin/mongodump mongodump/main/mongodump.go
    ./build.sh

FROM alpine:3.7
RUN apk add --no-cache bash curl

COPY --from=compile mongo-tools/bin/* /usr/bin/

ADD ./backup.sh /backup.sh
ADD ./entry.sh /entry.sh
ADD ./dropbox_uploader/dropbox_uploader.sh /dropbox_uploader.sh

RUN chmod +x /entry.sh
RUN mkdir /var/log/cron && touch /var/log/cron/cron.log

VOLUME /config

ENTRYPOINT [ "/entry.sh" ]
