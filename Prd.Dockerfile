FROM golang:1.12-alpine AS build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/abyssparanoia/operationAPI-GCP

ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .

FROM google/cloud-sdk:alpine AS server

ENV PORT 8080
EXPOSE 8080

RUN apk add ca-certificates
COPY --from=server_builder /go/bin/operationAPI-GCP /bin/operationAPI-GCP
COPY --from=server_builder /go/src/github.com/abyssparanoia/operationAPI-GCP/.env /go/src/github.com/abyssparanoia/operationAPI-GCP/.env
COPY --from=server_builder /go/src/github.com/abyssparanoia/operationAPI-GCP/serviceAccount.json /go/src/github.com/abyssparanoia/operationAPI-GCP/serviceAccount.json

WORKDIR /go/src/github.com/abyssparanoia/operationAPI-GCP

RUN apk --no-cache --update upgrade \
    && apk add --no-cache git alpine-sdk \
    && gcloud auth activate-service-account --key-file ./serviceAccount.json \
    && gcloud components install alpha \
    && gcloud components update 

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/bin/operationAPI-GCP"]