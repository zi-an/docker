FROM debian12.x

RUN echo "deb https://nginx.org/packages/mainline/debian bookworm nginx" > /etc/apt/sources.list.d/nginx.list &&\
    apt update --allow-insecure-repositories &&\
    apt install -y nginx --allow-unauthenticated
RUN echo "events {worker_connections 1024;}" > /etc/nginx/nginx.conf &&\
    echo "stream {" >> /etc/nginx/nginx.conf &&\
    echo "upstream tcp80 {server 172.18.0.1:80;}" >> /etc/nginx/nginx.conf &&\
    echo "server {listen 80;listen [::]:80;proxy_pass tcp80;}}" >> /etc/nginx/nginx.conf

RUN apt install -y git
RUN sed -i "19icd /home/nginx/" /root/.bashrc &&\
    cp /etc/ssh/ssh_host_ed25519_key ~/.ssh/id_ed25519 &&\
    ssh -o "StrictHostKeyChecking=no" git@github.com ;exit 0
RUN git clone git@github.com:zi-an/zi-an.github.io.git --branch=ipv6 /zian
# 添加定时任务 */5 * * * * docker exec ipv6 /zian/ipv6update.sh

#debian自带的nginx内的--with-stream=dynamic,是动态的,要重新编译
#nginx源里的nginx自带--with-stream,不用编译
CMD nginx ; /usr/sbin/sshd -D