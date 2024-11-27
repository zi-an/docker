FROM debian12.x

RUN apt install -y apache2-utils libnginx-mod-http-dav-ext nginx unzip

RUN mkdir /home/nginx/ &&\
    #wget gitdl.cn/github.com/zi-an/zi-an.github.io/blob/master/debian/goweb -O /usr/bin/goweb &&\
    wget http://120.76.158.149/debian/goweb -O /usr/bin/goweb &&\
    rm /etc/nginx/sites-enabled/default &&\
    htpasswd -bc /passwd zian zian

ADD /shell /usr/bin/
ADD /autoindex/* /var/www/autoindex/
ADD /conf/* /etc/nginx/sites-enabled

RUN chmod 777 /usr/bin/updategit /usr/bin/m3u8downloader /usr/bin/goweb
# 0/2 * * * * docker exec zian updategit