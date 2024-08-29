# Jenkins部署

# war安装+启动
* export JENKINS_PORT=80 无效
```
apt install openjdk-17-jdk -y
wget http://mirrors.jenkins.io/war-stable/2.346.3/jenkins.war
java -jar jenkins.war --httpPort=80
sed -i "s|https://updates.jenkins.io/|https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/|g" /root/.jenkins/hudson.model.UpdateCenter.xml
```
# 测试
需要插件
* http://222.mm/pluginManager/ 安装 Maven Integration plugin + Git
* http://222.mm/configureTools/ 配置maven(/usr/share/maven)和git(默认不需要配置)路径
* 创建项目->源码管理git填入"https://gitee.com/y_project/RuoYi.git",保存后启动即可
```  
apt -y install git maven
# mvn --version获取Maven home: /usr/share/maven
# http://127.0.0.1/manage/configureTools/填入Maven home
```

==================================================================================================================================================
# apt安装(推荐)
```
echo "deb http://pkg.jenkins.io/debian-stable binary/" > /etc/apt/sources.list.d/jenkins.list
wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | apt-key add -
apt update --allow-insecure-repositories
apt install -y jenkins openjdk-17-jdk
```

# 默认密码
```
cat /root/.jenkins/secrets/initialAdminPassword
sed -i "s|https://updates.jenkins.io/|https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/|g" /root/.jenkins/hudson.model.UpdateCenter.xml
```

 