FROM scratch

ADD rootfs.tar.gz /

ENV TZ='Asia/Taipei' LANG='zh_CN.UTF-8'
# 笔记本盒盖问题+apt清华源(因版本太新,暂时改为阿里)
# 安装openssh-server时要确认,使用echo管道
# 解决sshd时中文显示错误问题(apt安装语言包locales),bash下可以不用
# ssh远程不输出最后登录时间及横幅
# 禁用控制台发声
# 中文汉化,通过find /etc -mtime -1找到dpkg-reconfigure locales的不同,确定是/etc/default/locale文件
# 添加docker用户并设密
RUN echo 'HandleLidSwitch=lock' >> /etc/systemd/logind.conf &&\
    rm -rf /etc/apt/sources.list.d &&\
    echo 'deb http://mirrors.ustc.edu.cn/debian trixie main' > /etc/apt/sources.list &&\
    apt update &&\
    echo 'y' | apt install -y bash-completion curl iproute2 iputils-ping locales net-tools openssh-server sudo vim-tiny wget &&\
    rm -rf /var/lib/apt &&\
    echo 'zh_CN.UTF-8 UTF-8' > /etc/locale.gen &&\
    echo 'LANG=zh_CN.UTF-8' > /etc/default/locale &&\
    locale-gen &&\
    rm /etc/motd /etc/update-motd.d/10-uname &&\
    echo 'PrintLastLog no' >> /etc/ssh/sshd_config &&\
    echo 'set bell-style none' >> /etc/inputrc &&\
    sed -i '1i source /etc/profile' /etc/skel/.bashrc &&\
    cat /etc/skel/.bashrc > /root/.bashrc &&\
    useradd -ms/bin/bash -k/etc/skel/ docker &&\
    usermod -aG sudo docker &&\
    echo docker:. | chpasswd

# 里面有部分配置文件
ADD config.tgz /

# 修改系统所有文件时间,(软链接/lib与/lib64需要添加-h如 touch -h /sbin -t 197001010800.00),方便以后查找
RUN find /* -path /sys -prune -o -path /proc -prune -o -exec touch -ht 197001010000.00 {} \;

CMD /usr/sbin/sshd -D