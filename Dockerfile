#BUILDER
FROM golang as builder 
RUN mkdir /build 
ADD . /build 
WORKDIR /build
RUN go build -o testServer


#ACTUAL
FROM ubuntu

#install k6
RUN apt update && apt install -y gnupg2 ca-certificates vim curl etcd-client


#add tini
ENV TINI_VERSION v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

#copy 
RUN mkdir -p /mnt/perf/
COPY --from=builder /build/testServer /usr/local/bin/testServer
RUN chmod +x /usr/local/bin/testServer

CMD /usr/local/bin/testServer