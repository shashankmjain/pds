FROM golang
ADD  work/src/github.com/tylertreat/ /go/src/github.com/tylertreat/
ADD  work/src/newtest /go/src/newtest
ADD  work/src/github.com/caio/go-tdigest /go/src/github.com/caio/go-tdigest
RUN go install newtest
ENTRYPOINT /go/bin/newtest
EXPOSE 8000 
