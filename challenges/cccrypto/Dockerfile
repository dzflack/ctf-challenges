FROM golang:alpine AS builder1
WORKDIR /opt/cccrypto
COPY main.go .
COPY go.mod .

RUN apk add --no-cache git
RUN go install mvdan.cc/garble@latest
RUN GOOS=js GOARCH=wasm garble -tiny -literals build -o /main.wasm


FROM golang:alpine AS builder2
WORKDIR /opt/cccrypto
COPY server . 
RUN echo $(ls -laht . )

COPY --from=builder1 /main.wasm ./web/src/main.wasm
RUN echo $(ls -laht  ./web/src)

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/cccrypto
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd


FROM scratch
COPY --from=builder2 /go/bin/cccrypto /cccrypto
COPY --from=builder2 /etc_passwd /etc/passwd

USER nobody

EXPOSE 3100

ENTRYPOINT ["./cccrypto"]