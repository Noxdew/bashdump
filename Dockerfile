FROM golang:1.14 as builder
COPY . /bashdump
WORKDIR /bashdump
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=builder /bashdump/bashdump /bashdump
ENTRYPOINT [ "/bashdump" ]
