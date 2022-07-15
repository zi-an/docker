FROM centos7.8

RUN yum install java-11-openjdk-devel -y &&\
    rm -rf /var/cache/yum/

CMD ["/usr/sbin/sshd -D"]