# debian docker镜像

## 下载到本地
* https://github.com/debuerreotype/docker-debian-artifacts/tree/dist-amd64/
* 这是powershell中的wget,与linux的不同
* -Uri与-OutFile必须指定才能下载
* 最好在笔记本上使用代理加速下载
```shell
wget -Uri https://zi-an.github.io/debian/virc -OutFile virc
wget -Uri https://zi-an.github.io/debian/keys.zip -OutFile keys.zip
```

```shell
wget -Uri https://gitdl.cn/https://raw.githubusercontent.com/debuerreotype/docker-debian-artifacts/dist-amd64/bookworm/rootfs.tar.xz -OutFile rootfs.tar.xz
```


```加速器
https://gitdl.cn/https://raw.githubusercontent.com/debuerreotype/docker-debian-artifacts/dist-amd64/bookworm/rootfs.tar.xz
```