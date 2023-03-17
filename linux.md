# 给linux改名
`vi /etc/hostname` 改完后重启
# centos7 上下载CMAKE
wget https://cmake.org/files/v3.12/cmake-3.12.4.tar.gz
* 下载后解压：
tar -zxvf cmake-3.12.4.tar.gz
* 进入文件目录：
cd cmake-3.12.4
* 运行某文件：
./bootstrap
* 运行命令：
gmake
* 运行命令
sudo gmake install
* 安装完毕，检查版本
cake —version