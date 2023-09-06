FROM guacamole/guacd:1.5.3

USER root
ENV TZ='Asia/Shanghai'
# alpine时间识别不准,是因为没有/usr/share/zoneinfo/下的文件,文件由tzdata提供
RUN apk update && apk add tzdata


RUN wget https://github.com/dushixiang/next-terminal/releases/download/v1.3.9/next-terminal.tar.gz
RUN tar -zxvf next-terminal.tar.gz && rm next-terminal.tar.gz

CMD  /opt/guacamole/sbin/guacd;/next-terminal/next-terminal