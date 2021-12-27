FROM scratch
COPY go-rest-client /
ENTRYPOINT ["/go-rest-client"]
