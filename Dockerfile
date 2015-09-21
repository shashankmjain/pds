FROM golang
RUN git clone https://github.com/shashankmjain/BoomFilters/ /go/src/github.com/tylertreat/
RUN git clone  https://github.com/shashankmjain/pds/ /go/src/newtest
RUN git clone https://github.com/caio/go-tdigest /go/src/github.com/caio/go-tdigest
#ADD  BoomFilters/ /go/src/github.com/tylertreat/
#ADD  pds/webserver.go  /go/src/newtest
#ADD  go-tdigest/ /go/src/github.com/caio/go-tdigest
RUN go install newtest
ENTRYPOINT /go/bin/newtest
EXPOSE 8000 
