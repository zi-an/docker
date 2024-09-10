# debian docker镜像

## 下载到本地
* https://github.com/debuerreotype/docker-debian-artifacts/tree/dist-amd64/
* 这是powershell中的wget,与linux的不同
* -Uri与-OutFile必须指定才能下载
* 最好在笔记本上使用代理加速下载
```shell
curl -o virc https://gitdl.cn/https://github.com/zi-an/zi-an.github.io/blob/master/debian/virc
curl -o keys.zip https://gitdl.cn/https://github.com/zi-an/zi-an.github.io/blob/master/debian/keys.zip
curl -o rootfs.tar.xz https://gitdl.cn/https://github.com/debuerreotype/docker-debian-artifacts/blob/dist-amd64/bookworm/rootfs.tar.xz
```