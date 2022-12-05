FROM golang:alpine AS builder

WORKDIR /opt/obfuscatego
COPY main.go .
COPY go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/obfuscatego
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/passwd

FROM scratch
COPY --from=builder /go/bin/obfuscatego /obfuscatego
COPY --from=builder /etc/passwd /etc/passwd

COPY web /web

USER nobody

EXPOSE 80

ENTRYPOINT ["./obfuscatego"]