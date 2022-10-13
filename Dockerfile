FROM golang:1.19 AS builder
ADD . /build
WORKDIR /build
RUN cd app && go build -v -race -mod=vendor -ldflags "-X main.revision=$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse --short HEAD)" -o /app
FROM ubuntu
COPY --from=builder /app /app
ENTRYPOINT [ "/app" ]