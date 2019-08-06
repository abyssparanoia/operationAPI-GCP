## what is this

```
operationAPI is api server for web service operation on GCP
```

- operationAPI-GCP for running on container
- using cloud build for CI/CD

## how to develop

```bash

## build
> docker-compose build

## start
> docker-compose up -d

## down
> docker-compose down

```

## upload secret files

```bash
> gsutil cp ./.env gs://{GCP_PROJECT}/operation/.env
> gsutil cp ./serviceAccount.json gs://{GCP_PROJECT}/operation/serviceAccount.json
```
