FROM debian12.x
RUN echo "StrictHostKeyChecking no" >> /etc/ssh/ssh_config
CMD ["ssh","-ND","0.0.0.0:55555","docker@172.19.0.1"]