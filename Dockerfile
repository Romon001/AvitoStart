FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./


RUN go get -d -v ./...
# build go app
RUN go mod download
RUN go build -o avitostart ./cmd/main.go

CMD ["./avitostart"]

# #build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go mod download
# RUN go build -o /go/bin/app -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=avitostart Version=0.0.1
# EXPOSE 5436
