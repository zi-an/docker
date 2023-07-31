FROM scratch
#debian12.1

#https://raw.githubusercontent.com/debuerreotype/docker-debian-artifacts/dist-amd64/bookworm/rootfs.tar.xz
ADD rootfs.tar.xz /

# 笔记本盒盖问题+apt清华源
RUN echo "HandleLidSwitch=lock" >> /etc/systemd/logind.conf &&\
    sed -i "s|deb.debian.org|mirrors.tuna.tsinghua.edu.cn|g" /etc/apt/sources.list.d/debian.sources

ENV TZ='CST-8' LC_ALL='C.UTF-8'

#安装openssh-server时要确认,使用echo管道
RUN apt update &&\
    apt install -y bash-completion curl iproute2 iputils-ping net-tools vim-tiny wget &&\
    mkdir /run/sshd &&\
    echo 'y' | apt install -y openssh-server

RUN echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOnFp3JLA1fwdHjEs8NSpgQ4k0eAhvaQbecKXl03mhsm" >> ~/.ssh/authorized_keys &&\
    chmod 700 ~/.ssh &&\
    chmod 644 ~/.ssh/authorized_keys &&\
    echo "source /etc/profile" >> /root/.bashrc &&\
    echo "clear" >> /root/.bashrc

#这里不能像centos7直接加参数
CMD ["/usr/sbin/sshd","-D"]