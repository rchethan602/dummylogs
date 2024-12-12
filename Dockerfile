FROM golang:1.22 as builder

WORKDIR /app

# Copy and build the Go application
COPY . .
RUN go build -o logbuf main.go

# Use a minimal image to run the application
FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /app/logbuf /logbuf

# Command to run the executable
CMD ["/logbuf"]