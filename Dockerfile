FROM golang:1.14 as build
ADD . /graph-api/
WORKDIR /graph-api/
RUN make
ENV PORT 8888
CMD ["./bin/graph"]
