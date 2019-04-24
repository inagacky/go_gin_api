From golang:1.12.4

MAINTAINER inagacky

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

COPY api /go/src/github.com/inagacky/go_gin_api/src/api
RUN mkdir -p /go/src/github.com/inagacky/go_gin_api/log
WORKDIR /go/src/github.com/inagacky/go_gin_api/src/api

RUN apt-get update \
  && apt-get install -y git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

# debug
# RUN go get github.com/derekparker/delve/cmd/dlv
# ホットリロード
# RUN go get github.com/oxequa/realize

EXPOSE 8080

#RUN go build -gcflags "all=-N -l" -o /server main.go
#RUN realize start --run --no-config

CMD ["go", "run", "main.go"]
#CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/server"]"
#CMD ["dlv", "debug", "--listen=:2345", "--headless=true", "--api-version=2", "main.go"]"
