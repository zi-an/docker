FROM debian13.x

RUN apt-get update &&\
    apt install dnsmasq -y &&\
    echo "log-queries" > /etc/dnsmasq.conf &&\
    echo "local-ttl=600" >> /etc/dnsmasq.conf &&\
    echo "log-facility=/dns.log" >> /etc/dnsmasq.conf &&\
    echo "address=/update.miui.com/127.0.0.1" >> /etc/dnsmasq.conf &&\
    echo "stop-dns-rebind" >> /etc/dnsmasq.conf

CMD ["dnsmasq","-d"]
#dnsmasq --help
#https://oss.segetech.com/intra/srv/dnsmasq.conf
#tail -fn20 /dns.log | grep replyZZ