FROM debian12.1

#debian编译的nginx可能无法使用straem进行tcp转发
RUN apt install nload iftop lsof

CMD ["/usr/sbin/sshd","-D"]