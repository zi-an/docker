FROM scratch
#debian12.x --force-rm

ADD rootfs.tar.xz /
ENV TZ='CST-8' LANG='zh_CN.UTF-8'
# 笔记本盒盖问题+apt清华源
# 安装openssh-server时要确认,使用echo管道
# 解决sshd时中文显示错误问题(apt安装语言包locales),bash下可以不用
# 中文汉化,通过find /etc -mtime -1找到dpkg-reconfigure locales的不同,确定是/etc/default/locale文件
# 添加docker用户不设密码
RUN echo "HandleLidSwitch=lock" >> /etc/systemd/logind.conf &&\
    sed -i "s|deb.debian.org|mirrors.tuna.tsinghua.edu.cn|g" /etc/apt/sources.list.d/debian.sources &&\
    apt update &&\
    echo 'y' | apt install -y bash-completion curl iproute2 iputils-ping locales net-tools openssh-server sudo vim-tiny wget &&\
    mkdir /run/sshd &&\
    apt clean &&\
    echo "zh_CN.UTF-8 UTF-8" > /etc/locale.gen &&\
    echo "LANG=zh_CN.UTF-8" > /etc/default/locale &&\
    locale-gen &&\
    echo "source /etc/profile" >> /root/.bashrc &&\
    echo "clear" >> /root/.bashrc &&\
    useradd -ms/bin/bash -k/etc/skel/ docker &&\
    usermod -aG sudo docker &&\
    echo docker:. | chpasswd

# 里面有部分配置文件
ADD config.tgz /

# 修改系统所有文件时间,(软链接/lib与/lib64需要添加-h如 touch -h /sbin -t 197001010800.00),方便以后查找
RUN find /* -path /sys -prune -o -path /proc -prune -o -exec touch -ht 197001010000.00 {} \;

# 这里不能像centos7直接加参数
CMD /usr/sbin/sshd -D