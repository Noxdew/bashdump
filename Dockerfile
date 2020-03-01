FROM golang:1.14 as builder
RUN apt-get install ca-certificates -y
COPY . /bashdump
WORKDIR /bashdump
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bashdump/bashdump /bashdump
ENTRYPOINT [ "/bashdump" ]
