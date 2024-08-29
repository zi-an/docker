FROM debian12.x

# https://www.elastic.co/guide/en/elasticsearch/reference/7.17/docker.html
# 目前最新7.17.13,docker pull elasticsearch:7.17.13
# https://blog.csdn.net/holdlg/article/details/89378722
RUN echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" > /etc/apt/sources.list.d/e7.list &&\
    apt update --allow-insecure-repositories &&\
    apt install -y elasticsearch sudo --allow-unauthenticated

#[discovery.seed_hosts, discovery.seed_providers, cluster.initial_master_nodes]
#RUN sed -i "s|#network.host: 192.168.0.1|network.host: 0.0.0.0|g" /etc/elasticsearch/elasticsearch.yml
#RUN echo 'cluster.name: "docker-cluster"' >> /etc/elasticsearch/elasticsearch.yml

ENV discovery.type=single-node
#USER elasticsearch

#CMD sudo -u elasticsearch /usr/share/elasticsearch/bin/elasticsearch