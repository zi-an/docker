FROM centos7.8

RUN wget https://mirrors.tuna.tsinghua.edu.cn/epel/7/x86_64/Packages/a/android-tools-20130123git98d0789-5.el7.x86_64.rpm --no-check-certificate

RUN rpm -i android-tools-20130123git98d0789-5.el7.x86_64.rpm

RUN adb shell

ADD y5s.sh /

CMD ["/usr/sbin/sshd -D"]