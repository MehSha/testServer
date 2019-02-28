#BUILDER
FROM golang as builder 
RUN mkdir /build 
ADD . /build 
WORKDIR /build
RUN go build


#ACTUAL
<<<<<<< HEAD
FROM debian
RUN apt update  && apt install -y ca-certificates
# RUN apk update  && apk add -f ca-certificates


#add tini
ENV TINI_VERSION v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]
=======
FROM foodora/debian
>>>>>>> fullcompose

#copy 
COPY --from=builder /build/ci_test /usr/local/bin/ci_test
RUN chmod +x /usr/local/bin/ci_test

CMD /usr/local/bin/ci_test