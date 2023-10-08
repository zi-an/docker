FROM debian12.x
# 此处使用apt安装
ENV JENKINS_PORT=80

RUN echo "deb https://pkg.jenkins.io/debian-stable binary/" > /etc/apt/sources.list.d/jenkins.list &&\
    apt update --allow-insecure-repositories &&\
    apt install -y jenkins openjdk-17-jdk --allow-unauthenticated &&\
    echo "export JENKINS_PORT=80" >> /etc/profile &&\
    echo "cat /root/.jenkins/secrets/initialAdminPassword" >> /root/.bashrc
#启动命令jenkins,一般使用nohup jenkins &

RUN apt install -y git maven
# mvn --version获取Maven home: /usr/share/maven
# http://127.0.0.1/manage/configureTools/填入Maven home

# 必备插件(初始化系统不要选择推荐插件,太慢)
# Maven Integration plugin
# Git plugin
# Localization: Chinese (Simplified)