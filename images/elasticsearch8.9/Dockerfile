FROM ubuntu:latest
LABEL authors="root"


RUN echo "deb [signed-by=/usr/share/keyrings/elasticsearch-keyring.gpg] https://artifacts.elastic.co/packages/8.x/apt stable main" > /etc/apt/sources.list.d/elastic-8.x.list

RUN apt update --allow-insecure-repositories
RUN apt install elasticsearch
# /usr/share/elasticsearch/bin/elasticsearch
# 不能以管理员登录
ENTRYPOINT ["top", "-b"]