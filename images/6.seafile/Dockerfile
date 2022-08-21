FROM centos7.8

COPY seafile-server_5.1.4_x86-64.tar.gz /
COPY seafile.sh /

RUN yum install -y which python-setuptools python-pillow &&\
    mkdir /seafile &&\
    tar -zxvf /seafile-server_5.1.4_x86-64.tar.gz -C /seafile/ &&\
    chmod 777 seafile.sh

# 拼凑密码，answer第三行为服务器域名或者ip
RUN echo > /answer &&\
    echo "zian">> /answer &&\
    echo "192.168.10.210">> /answer &&\
    echo >> /answer &&\
    echo >> /answer &&\
    sed -i "s/= ask_admin_email()/= 'a@a.a'/g" /seafile/seafile-server-5.1.4/check_init_admin.py &&\
    sed -i "s/= ask_admin_password()/= 'a'/g" /seafile/seafile-server-5.1.4/check_init_admin.py

# 优化体验,解除30m播放限制,免密码等
RUN sed -i "s|= 30 |= 1024 |g" /seafile/seafile-server-5.1.4/seahub/seahub/settings.py &&\
    sed -i 's/"password" value=""/"password" value="a"/g' /seafile/seafile-server-5.1.4/seahub/seahub/templates/registration/login.html &&\
    sed -i 's/"login" value=""/"login" value="a@a.a"/g' /seafile/seafile-server-5.1.4/seahub/seahub/templates/registration/login.html

# 优化启动顺序,先用10s启动seafile,20s后启动seahub,方便挂在数据卷上
RUN sed -i '3iif [[ ! -e "/seafile/seafile-data/seafile.db" ]];then'  /seafile/seafile-server-5.1.4/seafile.sh &&\
    sed -i '4i/seafile/seafile-server-5.1.4/setup-seafile.sh</answer && sleep 10;fi;'  /seafile/seafile-server-5.1.4/seafile.sh &&\
    sed -i '3isleep 15' /seafile/seafile-server-5.1.4/seahub.sh


CMD ["/usr/sbin/sshd -D"]


