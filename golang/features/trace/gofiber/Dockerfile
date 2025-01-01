FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
# Build the binary statically
RUN CGO_ENABLED=0 go build -o fiber-app

FROM scratch
WORKDIR /app
COPY --from=builder /app/fiber-app /fiber-app
EXPOSE 3000
CMD ["/fiber-app"]
