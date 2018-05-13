FROM alpine:3.7

RUN apk add --no-cache bash curl mongodb-tools

ADD ./backup.sh /backup.sh
ADD ./entry.sh /entry.sh
ADD ./dropbox_uploader/dropbox_uploader.sh /dropbox_uploader.sh

RUN chmod +x /entry.sh
RUN mkdir /var/log/cron && touch /var/log/cron/cron.log

VOLUME /config

ENTRYPOINT [ "/entry.sh" ]
