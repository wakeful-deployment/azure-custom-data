FROM golang:1.5.2-onbuild

RUN apt-get update \
 && apt-get install zip -y

RUN go get github.com/aktau/github-release

CMD /go/src/app/release
