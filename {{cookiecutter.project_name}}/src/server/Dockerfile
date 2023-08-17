FROM golang:1.20-alpine

ARG HTTPS_PROXY
ARG VERSION

ENV API_VERSION=${VERSION}

WORKDIR /app
COPY ./ /app/

RUN go build -o main .

CMD ["/app/main"]
