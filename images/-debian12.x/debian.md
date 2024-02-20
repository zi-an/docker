# debian docker镜像

## 下载到本地
* 这是powershell中的wget,与linux的不同
* -Uri与-OutFile必须指定才能下载
* 最好在笔记本上使用代理加速下载
```shell
wget -Uri https://zi-an.github.io/debian/virc -OutFile virc
wget -Uri https://zi-an.github.io/debian/keys.zip -OutFile keys.zip
wget -Uri https://github.com/debuerreotype/docker-debian-artifacts/raw/fd3102585b5636ea53fbfcf122a1d373720540fc/bookworm/rootfs.tar.xz -OutFile rootfs.tar.xz
```
