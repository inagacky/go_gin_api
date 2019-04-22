From golang:1.12.4

MAINTAINER inagacky

ARG USER="go_sample"
ARG UID="1000"
ARG GROUP="go_sample"
ARG GID="1000"

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

COPY api /go/src/github.com/inagacky/go_gin_sample/src/api

WORKDIR /go/src/github.com/inagacky/go_gin_sample/src/api

RUN apt-get update \
  && apt-get install -y git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

EXPOSE 8080

CMD ["go", "run", "main.go"]
