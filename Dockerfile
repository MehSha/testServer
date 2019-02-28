#BUILDER
FROM golang as builder 
RUN mkdir /build 
ADD . /build 
WORKDIR /build
RUN go build


#ACTUAL
FROM foodora/debian

#copy 
COPY --from=builder /build/ci_test /usr/local/bin/ci_test
RUN chmod +x /usr/local/bin/ci_test

CMD /usr/local/bin/ci_test