FROM debian13.x

RUN apt clean &&\
    apt update &&\
    apt install -y apache2-utils libnginx-mod-http-dav-ext nginx unzip &&\
    rm -rf /var/lib/apt

RUN mkdir /home/nginx/ &&\
    wget https://zian.netlify.app/debian/goweb.tgz &&\
    tar -zxf goweb.tgz -C /usr/bin/ &&\
    rm goweb.tgz &&\
    sed -i 's|txt;|txt md reg;|g' /etc/nginx/mime.types &&\
    rm /etc/nginx/sites-enabled/default &&\
    sed -i 's|www-data|root|g' /etc/nginx/nginx.conf &&\
    htpasswd -bc /passwd zian zian

ADD /shell /usr/bin/
ADD /autoindex/* /var/www/autoindex/
ADD /conf/* /etc/nginx/sites-enabled

RUN chmod 777 /usr/bin/updategit /usr/bin/m3u8downloader /usr/bin/goweb
# 0 */2 * * * docker exec zian updategit