FROM debian12.1

RUN apt install -y gcc libpcre3-dev libssl-dev zlib1g-dev make
RUN wget https://nginx.org/download/nginx-1.24.0.tar.gz &&\
    tar -zxvf nginx-1.24.0.tar.gz &&\
    cd nginx-1.24.0/ &&\
    wget https://github.com/winshining/nginx-http-flv-module/archive/refs/heads/master.zip &&\
    unzip master.zip &&\
    ./configure --add-module=./nginx-http-flv-module-master &&\
    make install

RUN sed -i "19ialias nginx=/usr/local/nginx/sbin/nginx" /root/.bashrc

#使用nginx-http-flv-module开源项目的推荐配置文件
RUN cd /nginx-1.24.0/nginx-http-flv-module-master/ &&\
    num=`grep "worker_processes  1" README.md -n | awk -F ':' '{print $1}'` &&\
    cat README.md | tail -n +$num > /usr/local/nginx/conf/nginx.conf


# nohup ffmpeg -re -stream_loop -1 -i heishenhua.mp4 -c copy -f flv rtmp://214.mm:1935/hls/999 &
# http://213.mm/hls/999.m3u8

# 推送电脑画面(没声音,因为是mpg格式转的flv)只能用rtmp://214.mm:1935/hls/888观看
#./ffmpeg.exe -f gdigrab -i desktop -f flv rtmp://214.mm:1935/hls/888
CMD ["/usr/sbin/sshd","-D"]