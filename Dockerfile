FROM golang:1.14-alpine

RUN apk add --no-cache git
RUN apk add --update make
RUN mkdir -p /app

# Move to working directory /build
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY Makefile .
RUN go mod download

# Copy the code into the container
COPY . .

RUN make all