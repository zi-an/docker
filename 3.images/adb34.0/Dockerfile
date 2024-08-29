FROM debian12.x

RUN wget https://dl.google.com/android/repository/platform-tools-latest-linux.zip &&\
    apt install -y unzip &&\
    unzip platform-tools-latest-linux.zip &&\
    mv platform-tools/* /usr/bin/ &&\
    rm platform* -rf

RUN mkdir ~/.android/
ADD adbkey /root/.android/
ADD auto.sh /root/
RUN echo "alias auto='nohup /root/auto.sh &'" >> /etc/profile &&\
    chmod 750 ~/auto.sh &&\
    sed -i "s/\r//g" ~/auto.sh &&\
    chmod 750 ~/.android/ &&\
    chmod 600 ~/.android/adbkey

#客户端shell里要有执行source /etc/profile
