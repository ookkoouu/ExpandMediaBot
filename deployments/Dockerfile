FROM golang:1.17-alpine AS builder
WORKDIR /src
RUN apk --no-cache add curl && \
  curl -sL -o app.tar.gz "$( \
  curl -sL https://api.github.com/repos/ookkoouu/ExpandMediaBot/releases/latest \
  | grep tarball \
  | awk -F '"' '{print $4}')" && \
  tar -zxvf app.tar.gz --strip-components 1 && \
  apk del curl
RUN CGO_ENABLED=0 go build -trimpath -o app ./cmd/main.go

FROM alpine:latest
WORKDIR /
RUN apk --no-cache add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata
COPY --from=builder /src/app .
ENTRYPOINT [ "./app" ]