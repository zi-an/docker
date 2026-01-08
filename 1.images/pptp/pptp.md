docker run -d --privileged --name pptp -p 1723:1723/tcp stargriv/pptpd





andrey / andrey12
docker exec -it pptp sh
tail -fn20 /var/log/ppp/pptpd.log


sudo modprobe nf_conntrack_pptp