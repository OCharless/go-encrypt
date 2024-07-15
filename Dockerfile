FROM golang:bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /app/cmd

# Second stage: create the final image
FROM alpine:latest as go-encrypt

WORKDIR /root/

# Copy the executable from the first stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]