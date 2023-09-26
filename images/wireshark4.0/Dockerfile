FROM debian12.1

RUN apt install -y wireshark tcpdump fonts-wqy-zenhei &&\
    echo "X11UseLocalhost no" >> /etc/ssh/sshd_config

CMD ["/usr/sbin/sshd","-D"]