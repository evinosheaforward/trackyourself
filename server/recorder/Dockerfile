FROM golang:1.10
EXPOSE 9299

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
# ENV GOPATH=/go
# ENV PATH=$GOPATH/bin:$PATH

ADD ./src/vendor/ /go/src/
ADD ./src/record/ /go/src/record/
ADD ./src/ /go/src/

WORKDIR /go/src/

CMD go run main.go
