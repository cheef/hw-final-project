FROM golang:1.21 as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go get -d -v ./...
COPY . .
RUN go build -v -o server ./cmd/bfa-protection & \
    go build -v -o cli  ./cmd/bfa-protection-cli

FROM ubuntu:latest
WORKDIR /app
COPY --from=build /app/bin/. .
EXPOSE 44044
ENTRYPOINT [ "/app/main"]