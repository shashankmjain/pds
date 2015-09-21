FROM golang
RUN git clone https://github.com/tylertreat/BoomFilters/
RUN git clone  https://github.com/shashankmjain/pds/
RUN git clone https://github.com/caio/go-tdigest
ADD  github.com/tylertreat/BoomFilters/ /go/src/github.com/tylertreat/
ADD  github.com/shashankmjain/pds/webserver.go  /go/src/newtest
ADD  github.com/caio/go-tdigest /go/src/github.com/caio/go-tdigest
RUN go install newtest
ENTRYPOINT /go/bin/newtest
EXPOSE 8000 
