FROM debian12.x

RUN echo "X11UseLocalhost no" >> /etc/ssh/sshd_config &&\
    echo "alias idea='/usr/idea/bin/idea.sh &'" >> /root/.bashrc &&\
    echo "alias java=/usr/idea/jbr/bin/java" >> /root/.bashrc &&\
    apt install -y fonts-wqy-microhei libxtst6 libxi6 xterm

#RUN wget https://download.jetbrains.com/idea/ideaIU-2024.1.6.tar.gz &&\
#    tar -zxvf ideaIU-2024.1.6.tar.gz -C /usr/ &&\
#    rm ideaIU-2024.1.6.tar.gz

# 2024.1 之后只能本地构建后上传到服务器，需要安装docker才能使用，所以放到服务器使用
# 最后能远程构建的是2023.3
# https://github.com/zi-an/soft/blob/windows/idea/jetbra-2023.3.4.zip