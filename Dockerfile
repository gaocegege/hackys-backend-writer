FROM ubuntu

RUN apt-get install -y ca-certificates

COPY ./app /
CMD ["/app"]
