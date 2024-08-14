FROM debian12.x

RUN apt install nginx libnginx-mod-stream dnsutils -y &&\
    rm -rf /etc/nginx/sites-enabled/default

ADD dot.conf /
ADD cert.txt /
RUN cat /dot.conf >> /etc/nginx/nginx.conf
RUN openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout /etc/nginx/cert.key -out /etc/nginx/cert.pem < /cert.txt



#https://garywu520.github.io/2019/11/15/%E9%83%A8%E7%BD%B2%E8%87%AA%E5%B7%B1%E7%9A%84DoT-DNS/

# dig @127.1 +tls qq.com