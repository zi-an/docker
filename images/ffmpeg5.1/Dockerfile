FROM debian12.1

RUN apt install -y ffmpeg nginx
RUN sed -i "51i autoindex on;alias /home/nginx/;" /etc/nginx/sites-enabled/default &&\
    mkdir /home/nginx/ &&\
    sed -i "19i cd /home/nginx/" /root/.bashrc

CMD ["/usr/sbin/sshd","-D"]