FROM google/cloud-sdk:alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/bin:$PATH

# Install Go
ENV GO_VERSION 1.12
RUN curl -Lso go.tar.gz "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz
ENV PATH /usr/local/go/bin:$PATH

WORKDIR /go/src/github.com/abyssparanoia/operationAPI-GCP/
COPY . .

ENV GO111MODULE=on

ENV PORT 8080
EXPOSE 8080

RUN apk --no-cache --update upgrade \
    && apk add --no-cache git alpine-sdk \
    && go get github.com/pilu/fresh \
    && gcloud auth activate-service-account --key-file ./serviceAccount.json \
    && gcloud components install alpha \
    && gcloud components update 

CMD fresh 
