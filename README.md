# Cloud-APP-Purchase-Platform
A full-stack web application for software purchases with React, Golang, GCE, Elasticsearch database, and Stripe API. 


Video Demo: https://recordit.co/XVxzKphjlM

## Features:
- Upload: The upload API allows users to submit an app which can include a title, description, and a media file.
- Checkout: The checkout API allows users to buy the published apps from AppStore and make a payment through Stripe API.
- Search: The search API allows users to search other users' apps based on the title and description.
- Authentication:  The authentication API allows users to sign up and sign in. It uses token-based authentication to protect upload, checkout and search services.

## Frontend:

Developed with React.

## Backend:

Backend code diagram: https://drive.google.com/file/d/1248sPO6l15aH9DUThKUid0rou1Ni1qTk/view?usp=sharing
![image txt](https://github.com/haijing-zhang/Cloud-APP-Purchase_Platform/blob/main/codediagram2.png)

Backend part is written in Go language and directly developed on GCE remote service.

- "/handler": handle client requests

    - "/handler/router.go": Use **Mux** library to achieve HTTP routing (in place of annotation-based http routing in spring).
    - "/handler/app.go": Define UploadHandler, SignupHandler, SigninHandler, SearchHandler and CheckoutHandler. 

- "/service": Define all logic functions.

- "/backend" : 
    - "/backend/elasticsearch.go": **Elasticsearch** is a NoSQL database providing a quick keywordbased search for all apps. It stores user information and app text information (title and description).
    
    - "/backend/stripe.go": **Stripe** is a suite of APIs powering online payment processing and commerce solutions(https://stripe.com/docs/api). It is used for creating product, price and checkout here. 
    - "/backend/gcs.go": **GCS** is used for storing app's media files. 

## Deploy on Google APP Engine

- Step 1: Create Docker file 
```
vim Dockerfile

FROM golang:1.17
MAINTAINER zhanghaijing0601@gmail.com


WORKDIR /go/src/appstore
ADD . / /go/src/appstore/


RUN go get cloud.google.com/go/storage
RUN go get github.com/auth0/go-jwt-middleware
RUN go get github.com/form3tech-oss/jwt-go
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/olivere/elastic/v7
RUN go get github.com/stripe/stripe-go/v74
RUN go get github.com/pborman/uuid
RUN go get gopkg.in/yaml.v2


EXPOSE 8080
CMD ["/usr/local/go/bin/go", "run", "main.go"]

```
- Step 2: Create app.yaml
```
vim app.yaml

runtime: custom
env: flex
network:
  name: default
  subnetwork_name: default
automatic_scaling:
  min_num_instances: 1
  max_num_instances: 2

```

- Step 3: Download the whole appstore directory and upload to GCE instance.

- Step 4: Deploy service to GAE

```
gcloud app deploy
```
