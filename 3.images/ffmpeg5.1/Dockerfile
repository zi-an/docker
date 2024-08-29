FROM debian12.x

RUN apt install -y ffmpeg nginx
RUN sed -i "51i autoindex on;alias /home/nginx/;" /etc/nginx/sites-enabled/default &&\
    mkdir /home/nginx/ &&\
    sed -i "19i cd /home/nginx/" /root/.bashrc