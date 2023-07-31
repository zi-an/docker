FROM debian12.1

RUN sed -i "s|bookworm-updates|bookworm-updates bullseye|g" /etc/apt/sources.list.d/debian.sources
RUN apt update &&\
    apt install -y openjdk-11-jdk

CMD ["/usr/sbin/sshd","-D"]