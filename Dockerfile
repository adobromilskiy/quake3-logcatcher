FROM golang:1.19-alpine AS builder
RUN apk update && apk add git
ADD . /build
WORKDIR /build
RUN cd app && go build -v -race -mod=vendor -ldflags "-X main.revision=$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse --short HEAD)" -o /app
FROM alpine
COPY --from=builder /app /app
ENTRYPOINT [ "/app" ]