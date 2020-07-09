FROM golang:alpine AS builder

WORKDIR /opt/cccrypto
COPY main.go .
COPY go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/cccrypto
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

FROM scratch
COPY --from=builder /go/bin/cccrypto /cccrypto
COPY --from=builder /etc_passwd /etc/passwd

COPY web /web

USER nobody

EXPOSE 3100

ENTRYPOINT ["./cccrypto"]