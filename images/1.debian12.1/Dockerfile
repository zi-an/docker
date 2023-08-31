FROM scratch
#debian12.1

#https://raw.githubusercontent.com/debuerreotype/docker-debian-artifacts/dist-amd64/bookworm/rootfs.tar.xz
ADD rootfs.tar.xz /

# 笔记本盒盖问题+apt清华源
RUN echo "HandleLidSwitch=lock" >> /etc/systemd/logind.conf &&\
    sed -i "s|deb.debian.org|mirrors.tuna.tsinghua.edu.cn|g" /etc/apt/sources.list.d/debian.sources

ENV TZ='CST-8'

#安装openssh-server时要确认,使用echo管道
RUN apt update &&\
    apt install -y bash-completion curl iproute2 iputils-ping locales net-tools vim-tiny wget zip &&\
    mkdir /run/sshd &&\
    echo 'y' | apt install -y openssh-server

# vim-minimal在centos8.4的配置文件,对应在debian中的vim-tiny,解决i模式下的复制以及光标移动问题
ADD etc/virc /etc/vim/vimrc.tiny

# zip -P zian keys.tar.gz *ed25519*
ADD keys.zip /etc/ssh/
RUN unzip -P zian -d /etc/ssh/ -o /etc/ssh/keys.zip

RUN echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOnFp3JLA1fwdHjEs8NSpgQ4k0eAhvaQbecKXl03mhsm" >> ~/.ssh/authorized_keys &&\
    chmod 700 ~/.ssh &&\
    chmod 644 ~/.ssh/authorized_keys &&\
    chmod 600 /etc/ssh/ssh_host_ed25519_key &&\
    echo "source /etc/profile" >> /root/.bashrc &&\
    echo "clear" >> /root/.bashrc

# 解决sshd时中文显示错误问题(apt安装语言包locales),bash下可以不用
# 中文汉化,通过find /etc -mtime -1找到dpkg-reconfigure locales的不同,确定是/etc/default/locale文件
# 其中
ENV LANG='zh_CN.UTF-8'
RUN apt install locales -y && \
    echo "zh_CN.UTF-8 UTF-8" > /etc/locale.gen &&\
    echo "LANG=zh_CN.UTF-8" > /etc/default/locale &&\
    locale-gen

#这里不能像centos7直接加参数
CMD ["/usr/sbin/sshd","-D"]