use mysql;
create database zian default character set utf8mb4;
create user 'zian'@'%' identified by 'zian';
grant all privileges on *.* to 'zian'@'%';
flush privileges;