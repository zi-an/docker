FROM centos7.8

COPY seafile-server_5.1.4_x86-64.tar.gz /

RUN yum install -y which python-setuptools python-pillow &&\
    mkdir /seafile &&\
    tar -zxvf /seafile-server_5.1.4_x86-64.tar.gz -C /seafile/

# 拼凑密码，answer第三行为服务器域名或者ip
RUN echo > /answer &&\
    echo "zian">> /answer &&\
    echo "192.168.10.210">> /answer &&\
    echo >> /answer &&\
    echo >> /answer &&\
    sed -i "s/= ask_admin_email()/= 'a@a.a'/g" /seafile/seafile-server-5.1.4/check_init_admin.py &&\
    sed -i "s/= ask_admin_password()/= 'a'/g" /seafile/seafile-server-5.1.4/check_init_admin.py

RUN /seafile/seafile-server-5.1.4/setup-seafile.sh < /answer &&\
    /seafile/seafile-server-5.1.4/seafile.sh start

CMD ["/usr/sbin/sshd -D"]


# /seafile/seafile-server-5.1.4/seahub/seahub/settings.py
# 中FILE_PREVIEW_MAX_SIZE 30m限制可以解除
# 登录页默认密码
# sed -i 's/"password" value=""/"password" value="a"/g' /seafile/seafile-server-5.1.4/seahub/seahub/templates/registration/login.html
# sed -i 's/"login" value=""/"login" value="a@a.a"/g' /seafile/seafile-server-5.1.4/seahub/seahub/templates/registration/login.html
