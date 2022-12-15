FROM ubuntu:22.10

EXPOSE 8080

COPY ./microservice ./

ENTRYPOINT ["./microservice"]
CMD ["postgres://user:password@host:5432/database"]
