FROM scratch

ADD centos-7.8.2003-x86_64-docker.tar.xz /

# 笔记本盒盖问题
RUN echo "HandleLidSwitch=lock" >> /etc/systemd/logind.conf

# 改清华旧源
RUN mkdir /etc/yum.repos.d/backup/ &&\
    mv /etc/yum.repos.d/*.repo /etc/yum.repos.d/backup/
ADD qinghua.repo /etc/yum.repos.d/
RUN yum install bash-completion wget openssh-server -y &&\
    rm -rf /var/cache/yum/

# 北京时区与ssh免密码登录echo "Asia/shanghai" > /etc/timezone  export
ENV TZ='CST-8' LANG=zh_CN.UTF-8
RUN mkdir ~/.ssh &&\
    echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOnFp3JLA1fwdHjEs8NSpgQ4k0eAhvaQbecKXl03mhsm" >> ~/.ssh/authorized_keys &&\
    chmod 700 ~/.ssh &&\
    chmod 644 ~/.ssh/authorized_keys &&\
    sshd-keygen -A &&\
    echo "root:1" | chpasswd

# 以openssh服务开启
CMD ["/usr/sbin/sshd -D"]
