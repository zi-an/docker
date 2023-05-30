FROM centos7.8

RUN yum install xorg-x11-xauth -y &&\
    echo "X11UseLocalhost no" >> /etc/ssh/sshd_config &&\
    yum install -y java-11-openjdk-devel &&\
    yum install -y wqy-microhei-fonts &&\
    yum clean all

# 使用不包含jdk的idea,使用idea的命令为idea
RUN wget https://download.jetbrains.com/idea/ideaIU-2020.1.1-no-jbr.tar.gz &&\
    tar -zxvf ideaIU-2020.1.1-no-jbr.tar.gz -C /usr/local/ &&\
    echo "alias idea='/usr/local/idea-IU-201.7223.91/bin/idea.sh &'" >> ~/.bashrc &&\
    rm ideaIU-2020.1.1-no-jbr.tar.gz

# 插件
# https://plugins.jetbrains.com/files/13710/96075/zh.201.16.jar
# https://raw.githubusercontent.com/zi-an/soft/master/soft/idea/jetbrains-agent.jar
CMD ["/usr/sbin/sshd -D"]