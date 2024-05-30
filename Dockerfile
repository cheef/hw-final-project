FROM golang:1.21 as base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux make build

EXPOSE 44044

CMD ["/bin/bfa-protection-server"]