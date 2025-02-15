# hadoop3.4.1
FROM debian12.x

ENV HADOOP_VERSION=3.4.1

RUN wget https://dlcdn.apache.org/hadoop/common/hadoop-${HADOOP_VERSION}/hadoop-${HADOOP_VERSION}.tar.gz &&\
    sed -i "s|bookworm-updates|bookworm-updates bullseye|g" /etc/apt/sources.list.d/debian.sources &&\
    apt update && apt install openjdk-11-jdk-headless -y

RUN sed -i '$aexport HADOOP_VERSION=\$HADOOP_VERSION' /etc/profile &&\
    sed -i '$aexport JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64' /etc/profile &&\
    sed -i '$aexport HADOOP_HOME=/opt/hadoop' /etc/profile &&\
    sed -i '$aPATH=\$PATH:\$JAVA_HOME/bin/:\$HADOOP_HOME/bin:\$HADOOP_HOME/sbin' /etc/profile &&\
    sudo useradd -ms/bin/bash -k/etc/skel/ hadoop &&\
    sudo usermod -aG sudo hadoop &&\
    echo hadoop:. | chpasswd &&\
    tar -zxvf hadoop-${HADOOP_VERSION}.tar.gz -C /opt/ --transform="s|hadoop-${HADOOP_VERSION}|hadoop|g" &&\
    rm hadoop-${HADOOP_VERSION}.tar.gz &&\
    sed -i "1 i JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64" /opt/hadoop/etc/hadoop/hadoop-env.sh &&\
    find /opt/hadoop/ -name *cmd -delete &&\
    mkdir /opt/hadoop/data &&\
    chown hadoop /opt -R

RUN find /opt/ -exec touch -ht 197001010000.00 {} \;