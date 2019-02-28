#BUILDER
FROM golang as builder 
RUN mkdir /build 
ADD . /build 
WORKDIR /build
RUN go build -o svc1
RUN ls


#ACTUAL
FROM foodora/debian

#copy 
COPY --from=builder /build/svc1 /usr/local/bin/svc1
RUN chmod +x /usr/local/bin/svc1

CMD /usr/local/bin/svc1