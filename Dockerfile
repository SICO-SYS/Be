FROM golang:alpine

MAINTAINER sine "sinerwr@gmail.com"

RUN apk --update add git && \
    go-wrapper download github.com/SiCo-Ops/Be && \
    apk del git && \
    cd $GOPATH/src/github.com/SiCo-Ops/Be && \
    go-wrapper install && \
    cd / &&\
    rm -rf $GOPATH/src

WORKDIR $GOPATH/bin

VOLUME $GOPATH/bin

CMD ["Be"]