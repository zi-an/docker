FROM debian13.x

RUN apt clean &&\
    apt update &&\
    apt install -y apache2-utils nginx unzip &&\
    rm -rf /var/lib/apt

RUN mkdir /home/nginx/ &&\
    wget https://zian.netlify.app/debian/goweb.tgz &&\
    tar -zxf goweb.tgz -C /usr/bin/ &&\
    rm goweb.tgz &&\
    sed -i 's/# gzip_proxied/gzip_proxied/' /etc/nginx/nginx.conf &&\
    sed -i 's/# gzip_vary on;/gzip_vary on;/' /etc/nginx/nginx.conf &&\
    sed -i 's/# gzip_comp_level 6;/gzip_comp_level 6;/' /etc/nginx/nginx.conf &&\
    sed -i 's/# gzip_types text\/plain/gzip_types text\/plain/' /etc/nginx/nginx.conf &&\
    sed -i 's|txt;|conf log md out reg sh txt;|g' /etc/nginx/mime.types &&\
    rm /etc/nginx/sites-enabled/default &&\
    sed -i 's|www-data|root|g' /etc/nginx/nginx.conf &&\
    htpasswd -bc /passwd zian zian

RUN wget https://zian.netlify.app/debian/https.tgz &&\
    tar -zxf https.tgz -C /etc/nginx/ &&\
    rm https.tgz

ADD /shell /usr/bin/
ADD /autoindex/* /var/www/autoindex/
ADD /conf/* /etc/nginx/sites-enabled

RUN chmod 777 /usr/bin/gitupdate /usr/bin/m3u8downloader /usr/bin/goweb
# 0 */2 * * * docker exec zian gitupdate

