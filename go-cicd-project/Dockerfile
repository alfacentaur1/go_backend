FROM golang:1.22 AS builder
WORKDIR /go-cicd-project
COPY . .
RUN go build -o main ./cmd/api

FROM alpine
COPY --from=builder /go-cicd-project/main .
CMD ["./main"]