FROM golang

ADD . /go/src/github.com/lobiCode/3fs

RUN go get  github.com/ttacon/libphonenumber
RUN go install github.com/lobiCode/3fs
 
ENTRYPOINT /go/bin/3fs

EXPOSE 1234 

