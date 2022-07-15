FROM scratch

ADD centos-7.8.2003-x86_64-docker.tar.xz /

# 笔记本盒盖问题
RUN echo "HandleLidSwitch=lock" >> /etc/systemd/logind.conf

# 改清华旧源
RUN rm -rf /var/cache/yum/ &&\
    mkdir /etc/yum.repos.d/backup/ &&\
    mv /etc/yum.repos.d/*.repo /etc/yum.repos.d/backup/ &&\
    echo "[qinghua]" >> /etc/yum.repos.d/qinghua.repo &&\
    echo "name=qinghua" >> /etc/yum.repos.d/qinghua.repo &&\
    echo "baseurl=https://mirrors.tuna.tsinghua.edu.cn/centos-vault/7.8.2003/os/x86_64/" >> /etc/yum.repos.d/qinghua.repo &&\
    echo "gpgcheck=0" >> /etc/yum.repos.d/qinghua.repo &&\
    yum install bash-completion wget openssh-server -y &&\
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
