############################
# STEP 1 build executable binary
############################
FROM golang:1.20.5-alpine3.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o  main .

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/main /app/main
COPY --from=builder /app/html /html
COPY --from=builder /app/static /static

EXPOSE 8080
CMD ["/app/main"]