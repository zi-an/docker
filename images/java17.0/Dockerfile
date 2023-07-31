FROM debian12.1

RUN apt update &&\
    apt install -y openjdk-17-jdk

CMD ["/usr/sbin/sshd","-D"]