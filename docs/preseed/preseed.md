# debian自动化安装

* https://www.debian.org/releases/stable/amd64/apbs02.zh-cn.html#preseed-loading
* demo https://www.debian.org/releases/bookworm/example-preseed.txt
* https://debian-handbook.info/get/now/
* https://blog.csdn.net/sinolover/article/details/120778961

# 安装必备软件+解压包+下载配置文件
* preseed.cfg为配置文件,主要修改
```
apt -y update
apt -y install xorriso isolinux syslinux-utils p7zip-full

wget https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-12.1.0-amd64-netinst.iso
7z x debian-12.1.0-amd64-netinst.iso -oiso
chmod +w -R iso/install.amd/
gunzip iso/install.amd/initrd.gz
wget https://www.debian.org/releases/bookworm/example-preseed.txt -O preseed.cfg
```

#编写菜单(可跳过,Windows下执行也可以)
* boot/grub/grub.cfg快捷键e
* 可以加定时*set timeout=5*,但是个人用没必要
* initrd.gz直接丢在U盘根目录就好
```
menuentry --hotkey=3 '按3自动安装' {
 set background_color=black
 linux /install.amd/vmlinuz vga=788 --- quiet 
 initrd /initrd.gz
}
```


# 内嵌+压缩
```
echo preseed.cfg | cpio -H newc -o -A -F iso/install.amd/initrd
gzip iso/install.amd/initrd
chmod -w -R iso/install.amd
```


#打包
* 也可以不打包,U盘刻录镜像后替换boot/grub/grub.cfg和iso/install.amd/initrd两个文件即可
```
xorriso -as mkisofs \
-isohybrid-mbr /usr/lib/ISOLINUX/isohdpfx.bin \
-c isolinux/boot.cat \
-b isolinux/isolinux.bin \
-no-emul-boot \
-boot-load-size 4 \
-boot-info-table \
-eltorito-alt-boot \
-e boot/grub/efi.img \
-no-emul-boot \
-isohybrid-gpt-basdat \
-o debian12.iso \
iso
```
==========================================================================================================================================

# 导出安装好的配置对比
```

```


# 优化版 9-15测试通过
* 本地化,域名
* 本土化,换为国内地址
* root密码+不创建用户
* 流行问卷取消
* 安装并下载脚本并且运行
* pkgsel/include安装包选择/包括
* 不更新软件
```
wget https://www.debian.org/releases/bookworm/example-preseed.txt -O preseed.cfg

sed -i "s|US/Eastern|Asia/Shanghai|g" preseed.cfg
sed -i "s|http.us.debian.org|mirrors.tuna.tsinghua.edu.cn|g" preseed.cfg
sed -i "s|#d-i apt-setup/security_host string security.debian.org|d-i apt-setup/security_host string mirrors.tuna.tsinghua.edu.cn|g" preseed.cfg
sed -i "s|#d-i debian-installer/locale string en_GB|d-i debian-installer/locale string zh_CN|g" preseed.cfg
sed -i "s|#d-i clock-setup/ntp-server string ntp.example.com|d-i clock-setup/ntp-server string ntp.aliyun.com|g" preseed.cfg


sed -i "s|somehost|5.mm|g" preseed.cfg
sed -i "s|unassigned-hostname|k610d|g" preseed.cfg
sed -i "s|#d-i netcfg/hostname string somehost|d-i netcfg/hostname string k610d|g" preseed.cfg
sed -i "s|#d-i netcfg/dhcp_hostname string radish|d-i netcfg/dhcp_hostname string k610d|g" preseed.cfg


sed -i "s|#d-i passwd/root-password password r00tme|d-i passwd/root-password password 1|g" preseed.cfg
sed -i "s|#d-i passwd/root-password-again password r00tme|d-i passwd/root-password-again password 1|g" preseed.cfg


sed -i "s|#d-i passwd/make-user|d-i passwd/make-user|g" preseed.cfg
sed -i "s|#popularity-contest|popularity-contest|g" preseed.cfg
sed -i "373ctasksel tasksel/first multiselect standard ssh-server" preseed.cfg

#sed -i "s|#d-i pkgsel/include string openssh-server|d-i pkgsel/include string openssh-server #|g" preseed.cfg
sed -i "s|#d-i pkgsel/upgrade select none|d-i pkgsel/upgrade select none|g" preseed.cfg
echo "d-i preseed/late_command string in-target wget https://zi-an.github.io/install/auto.sh -O /usr/sbin/auto;in-target chmod 700 /usr/sbin/auto;in-target auto" >> preseed.cfg
```

# 暂无改硬盘的方法,硬盘有数据不会自动选择
```
sed -i "s|#d-i partman-auto/disk string /dev/sda|d-i partman-auto/disk string /dev/sda|g" preseed.cfg
sed -i "221id-i partman-auto/expert_recipe string boot-root :: 40 50 100 ext4 $primary{ } $bootable{ } method{ format } format{ } use_filesystem{ } filesystem{ ext4 } mountpoint{ /boot } . 500 10000 1000000000 ext4 method{ format } format{ } use_filesystem{ } filesystem{ ext4 } mountpoint{ / } ." preseed.cfg
```

# 选择包关系
* https://wiki.debian.org/tasksel
* ssh-server
```
tasksel --list-tasks 
tasksel --task-packages standard ssh-server
```

# 简单版
* 需改时只需要把iso/install.amd/initrd替换成新的就可以了
* wget 7.mm/initrd.gz;gunzip initrd.gz
```
apt -y update
apt -y install debconf-utils xorriso isolinux syslinux-utils p7zip-full nginx
debconf-get-selections --installer > /var/www/html/1.txt
wget https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-12.1.0-amd64-netinst.iso
7z x debian-12.1.0-amd64-netinst.iso -oiso
chmod +w -R iso/install.amd/
cp iso/install.amd/initrd.gz .

gunzip initrd.gz
echo preseed.cfg | cpio -H newc -o -A -F initrd
gzip initrd
mv initrd.gz /var/www/html/
```
* 下载地址 http://5.mm/initrd.gz
