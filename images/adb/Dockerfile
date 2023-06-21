FROM centos7.8

RUN mkdir .android
ADD .android/adbkey /root/.android/
ADD auto.sh /root/

RUN wget https://dl.google.com/android/repository/platform-tools_r34.0.3-linux.zip

RUN yum install unzip -y &&\
    unzip platform-tools_r34.0.3-linux.zip &&\
    mv platform-tools/* /bin/ &&\
    chmod 750 ~/.android &&\
    chmod 600 ~/.android/adbkey &&\
    chmod 600 /root/auto.sh &&\
    sed -i "s/\r//" /root/auto.sh &&\
    rm platform* -rf

CMD ["/usr/sbin/sshd -D"]