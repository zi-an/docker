FROM debian12.1

# debian自带的nginx内的--with-stream=dynamic,是动态的,要重新编译,可能无法使用straem进行tcp转发
RUN apt install nginx -y
RUN mkdir /home/nginx/ &&\
    sed -i "42 i client_max_body_size 8192m;" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location /post {proxy_pass http://127.0.0.1:8888;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location /json/{autoindex on;autoindex_format json;alias /home/nginx/;}" /etc/nginx/sites-enabled/default
ADD autoindex/* /var/www/html/
RUN chmod 777 /var/www/html/.go

CMD ["/usr/sbin/sshd","-D"]