FROM debian12.x

WORKDIR /root/
RUN apt install -y gcc wget zlib1g-dev &&\
    wget https://download.oracle.com/graalvm/17/latest/graalvm-jdk-17_linux-x64_bin.tar.gz &&\
    tar -zxvf graalvm-jdk-17_linux-x64_bin.tar.gz -C / &&\
    rm graalvm-jdk-17_linux-x64_bin.tar.gz &&\
    mv /graalvm* /graalvm &&\
    echo "JAVA_HOME=/graalvm/" >> /etc/profile &&\
    echo "PATH=\$JAVA_HOME/bin/:$PATH" >> /etc/profile

ADD Test.txt /root/Test.java

