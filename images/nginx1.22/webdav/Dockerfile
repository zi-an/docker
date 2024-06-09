FROM debian12.x
# http://nginx.org/en/docs/http/ngx_http_dav_module.html

RUN apt install nginx libnginx-mod-http-dav-ext/stable apache2-utils -y &&\
    rm /var/www/html/* -rf &&\
    chmod  777 /var/www/html/ &&\
    rm /etc/nginx/sites-enabled/default &&\
    htpasswd -bc /passwd zian zian

ADD webdav.conf /etc/nginx/conf.d

# 密码:zian/zian