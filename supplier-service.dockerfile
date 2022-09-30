# build tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY supplierApp /app

CMD ["app/supplierApp"]