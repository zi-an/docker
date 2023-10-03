FROM debian12.1

RUN apt install -y squid &&\
    sed -i "s|http_access deny all|http_access allow all|g" /etc/squid/squid.conf &&\
    sed -i "s|http_port 3128|http_port 8080|g" /etc/squid/squid.conf

CMD ["/usr/sbin/sshd","-D"]