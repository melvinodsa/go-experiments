FROM melvinodsa/go-web-application:latest
LABEL maintainer="melvinodsa@gmail.com"

WORKDIR $GOPATH/src/github.com/go-experiments/cache

# go get the dependencies and clone the repo
COPY . $GOPATH/src/github.com/go-experiments/cache
RUN cd $GOPATH/src/github.com/go-experiments/cache \
    && dep ensure

ENTRYPOINT ["go", "test", "-bench", "."]