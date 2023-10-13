FROM debian12.x

RUN apt install nginx -y
#RUN curl https://mj.pki.net.cn/build1/static/js/main.3edcddfc.js > /var/www/html/maomi.js &&\
#    curl  https://mj.pki.net.cn/build1/static/css/main.7ed28a66.css > /var/www/html/maomi.css
RUN sed -i "s|location /|location /default|g" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location / {try_files '' /index.html;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/index.html {index  index.html;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/maomi.js {index maomi.js;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/maomi.css {index maomi.css;}" /etc/nginx/sites-enabled/default &&\
    sed -i "42 i location =/ {rewrite / /page/shipin/13 redirect;}" /etc/nginx/sites-enabled/default

ADD index.html.bak /var/www/html/index.html
ADD maomi.js.bak /var/www/html/maomi.js
ADD maomi.css.bak /var/www/html/maomi.css

CMD nginx -g "daemon off;"