
##############################
# Build Stage
##############################
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Install make and Node.js (with npm, which provides npx)
RUN apk add --no-cache make nodejs npm
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the remaining source code
COPY . .

RUN make build

##############################
# Run Stage
##############################
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/bin/main /app/main
COPY --from=builder /app/static /app/static

EXPOSE 8080

CMD ["/app/main"]

