FROM golang:alpine3.18 as builder
RUN apk add --no-cache --update git build-base

WORKDIR /app
COPY . .
RUN go build \
    -a \
    -trimpath \
    -o sendtome \
    -ldflags "-s -w -buildid=" \
    "./cmd/sendtome" && \
    ls -lah

FROM alpine:3.18 as runner
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/sendtome .
VOLUME /app/log
#EXPOSE 8080

ENTRYPOINT ["./sendtome"]