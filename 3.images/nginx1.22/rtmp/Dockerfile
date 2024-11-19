FROM debian12.x

RUN apt install -y gcc libpcre3-dev libssl-dev zlib1g-dev make unzip
RUN wget https://nginx.org/download/nginx-1.26.2.tar.gz &&\
    tar -zxvf nginx-1.26.2.tar.gz &&\
    cd nginx-1.26.2/ &&\
    wget gitdl.cn/github.com/arut/nginx-rtmp-module/archive/refs/heads/master.zip &&\
    unzip master.zip &&\
    ./configure --add-module=./nginx-rtmp-module-master &&\
    make install

RUN mkdir -p /tmpfs/hls &&\
    sed -i "19ialias nginx=/usr/local/nginx/sbin/nginx" /root/.bashrc
ADD nginx.conf /usr/local/nginx/conf/nginx.conf

# ffmpeg -re -stream_loop -1 -i heishenhua.mp4 -c copy -f flv rtmp://215.mm:1935/hls/999
# http://215.mm/hls/999.m3u8

# 推送电脑画面(没声音,因为是mpg格式转的flv)只能用rtmp://213.mm:1935/hls/888观看
#./ffmpeg.exe -f gdigrab -i desktop -f flv rtmp://213.mm:1935/hls/888
CMD ["/usr/sbin/sshd","-D"]