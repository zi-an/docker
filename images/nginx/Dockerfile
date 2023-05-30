FROM centos7.8

# nginx在epel库内
RUN yum install nginx iproute -y

# 增加nginx代理规则
# RUN sed -i "50 i location / {proxy_pass http://192.168.10.210;}" /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*
ADD autoindex/* /usr/share/nginx/html/
RUN sed -i "50 i location /post {proxy_pass http://127.0.0.1:8888;client_max_body_size 8192m;}" /etc/nginx/nginx.conf
RUN sed -i "50 i location /json/{autoindex on;autoindex_format json;alias /home/nginx/;index json;}" /etc/nginx/nginx.conf
WORKDIR /home/nginx/

CMD ["/usr/sbin/sshd -D"]