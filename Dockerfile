FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

COPY . ./guess
WORKDIR ./guess

RUN go build -mod=vendor -mod=readonly -o /guessout .

FROM alpine:latest

COPY --from=builder /guessout /guess

EXPOSE 8080

ENTRYPOINT ["/guess", "-http"]
