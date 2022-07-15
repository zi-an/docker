FROM centos7.8

RUN  yum install dovecot postfix cyrus-sasl cyrus-sasl-plain -y &&\
     sed -i "s|inet_interfaces = localhost|inet_interfaces = all|g" /etc/postfix/main.cf &&\
     echo "mydomain = mail.zi-an.ml" >> /etc/postfix/main.cf &&\
     echo "myhostname = mail.zi-an.ml" >> /etc/postfix/main.cf &&\
     /usr/libexec/postfix/aliasesdb &&\
     /usr/libexec/postfix/chroot-update &&\
     echo "disable_plaintext_auth = no" >> /etc/dovecot/conf.d/10-auth.conf &&\
     sed -i "s|ssl = required|ssl = no|g" /etc/dovecot/conf.d/10-ssl.conf &&\
     echo "protocols = pop3" >> /etc/dovecot/dovecot.conf &&\
     echo "mail_location = mbox:/var/mail:INBOX=/var/mail/%u" >> /etc/dovecot/dovecot.conf &&\
     echo "mail_access_groups = mail" >> /etc/dovecot/conf.d/10-mail.conf &&\
     echo "smtpd_sasl_auth_enable = yes" >> /etc/postfix/main.cf &&\
     echo "smtpd_sasl_local_domain = mail.zi-an.ml" >> /etc/postfix/main.cf &&\
     echo "smtpd_sasl_security_options = noanonymous" >> /etc/postfix/main.cf &&\
     echo "smtpd_recipient_restrictions = permit_sasl_authenticated" >> /etc/postfix/main.cf &&\
     adduser haiyang  &&\
     echo "haiyang:1" | chpasswd

CMD ["/usr/sbin/sshd -D"]

# postfix start &&\
#     dovecot &&\
#     saslauthd -a pam