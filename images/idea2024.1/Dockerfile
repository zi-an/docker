FROM debian12.x

RUN wget https://download.jetbrains.com/idea/ideaIU-2024.1.6.tar.gz &&\
    tar -zxvf ideaIU-2024.1.6.tar.gz -C /usr/ &&\
    rm ideaIU-2024.1.6.tar.gz

RUN echo "X11UseLocalhost no" >> /etc/ssh/sshd_config &&\
    echo "alias idea=/usr/idea-IU-241.19072.14/bin/idea.sh" >> .bashrc &&\
    apt install -y fonts-wqy-microhei libxtst6 libxi6 xterm

# 2024.1 之后只能本地构建后上传到服务器，需要安装docker才能使用，所以放到服务器使用
# 最后能远程构建的是2023.3