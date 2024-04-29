FROM debian12.x

RUN apt install -y squid &&\
    sed -i "s|http_access deny all|http_access allow all|g" /etc/squid/squid.conf

CMD ["squid", "-N"]