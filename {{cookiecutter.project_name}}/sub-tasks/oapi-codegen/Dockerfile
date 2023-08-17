FROM golang:alpine

WORKDIR /work

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

ENTRYPOINT [ "oapi-codegen" ]
