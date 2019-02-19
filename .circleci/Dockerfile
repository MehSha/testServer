FROM debian
RUN apt update  && apt install -y ca-certificates


#add tini
ENV TINI_VERSION v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

#copy 
ADD ci_test /usr/local/bin/ci_test
RUN chmod +x /usr/local/bin/ci_test
CMD /usr/local/bin/ci_test