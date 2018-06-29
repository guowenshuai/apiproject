FROM ubuntu:16.04

MAINTAINER wenshuai.guo

COPY apiproject.tar.gz /app/

WORKDIR /app


RUN tar -zxf apiproject.tar.gz \
    && chmod +x apiproject \
    && rm apiproject.tar.gz
#    && ln -s /app/apiproject /bin/apiproject
#RUN apt-get update \
#    && apt-get install -y iputils-ping mysql-client-5.7
ENV PATH $PATH:/app

EXPOSE 8080

ENTRYPOINT apiproject
