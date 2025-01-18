package cmd

import "github.com/spf13/cobra"

// newVpnCmd 创建 VPN 命令
func newVpnCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "vpn",
		Short: "VPN 相关命令",
		Long:  "显示 VPN 相关的配置和常用命令",
		Message: `systemctl enable trojan
sudo systemctl start trojan
sudo systemctl stop trojan
/usr/src/trojan #log在这里
`,
	})
}

func newWindowsCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "windows",
		Short: "Windows 系统命令",
		Long:  "显示 Windows 系统相关的常用命令和操作",
		Message: `与 Linux 命令对照	
功能	Linux 命令	Windows 命令
显示当前目录	pwd	cd
列出目录内容	ls	dir
切换目录	cd	cd
创建目录	mkdir	mkdir
删除文件	rm	del 或 erase
删除目录	rm -r 或 rmdir	rmdir /s
复制文件	cp	copy 或 xcopy
移动/重命名文件	mv	move
查看文件内容	cat	type
显示当前活动进程	ps	tasklist
杀死进程	kill	taskkill
查找文件或内容	grep	findstr
更改文件权限/权限管理	chmod	icacls 或 attrib
网络连接信息	ifconfig	ipconfig 或 netstat
下载文件	curl 或 wget	curl 或 Invoke-WebRequest
查看磁盘使用情况	df	dir 或 wmic logicaldisk
显示系统信息	uname	systeminfo
清空终端屏幕	clear	cls
退出终端	exit	exit
# 系统信息
wmic os get caption, version
wmic qfe get hotfixid, description

# others
mkdir -m 755 -p test # 创建文件夹,并设置权限

# 磁盘信息
wmic diskdrive get deviceid, name, size

# 网络信息
ipconfig
netstat -ano

# 网络连接测试
ping host
telnet host port
nc -zv host port
curl host	

# 抓包工具
tcpdump -i eth0
wireshark	

# DNS
nslookup domain
dig domain	

# 路由追踪
traceroute host
mtr host	

# 端口扫描
nmap host`,
	})
}

// newLinuxCmd 创建 Linux 命令
func newLinuxCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "linux",
		Short:   "Linux 系统命令",
		Long:    "显示 Linux 系统相关的常用命令和操作",
		Message: "使用 --subcmd 参数查看具体命令，可选：system, network, disk",
		SubCommands: map[string]string{
			"system": `系统管理命令：
# 系统信息
uname -a
cat /etc/os-release # 查看系统版本信息 方法A
lsb_release -a # 查看系统版本信息 方法B
top -H1

# sudo问题
vim /etc/sudoers # 配置 sudo

# 服务管理
systemctl status service_name
systemctl start service_name
systemctl enable service_name

# ssh
ssh-keygen -t rsa -b 4096 -C "yourname@ecosda.com"
cat xxx_id_rsa.pub >> ~/.ssh/authorized_keys # 将公钥添加到 ~/.ssh/authorized_keys

# 用户管理
useradd username
usermod -aG group username
passwd username`,

			"network": `网络管理命令：
# 网络信息
ip addr
ifconfig
netstat -tunlp
ss -tunlp

# 防火墙
sudo ufw status
sudo ufw allow 22
sudo ufw enable`,

			"disk": `磁盘管理命令：
# 磁盘信息
df -h
du -sh *
lsblk

# 磁盘操作
mount /dev/sda1 /mnt
umount /mnt
fdisk -l`,
		},
	})
}

// newNetworkCmd 创建网络工具命令
func newNetworkCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "network",
		Short: "网络工具命令",
		Long:  "显示网络相关的工具和命令",
		Message: `# 网络连接测试
ping host
telnet host port
nc -zv host port
curl host

# 抓包工具
tcpdump -i eth0
wireshark

# DNS
nslookup domain
dig domain

# 路由追踪
traceroute host
mtr host

# 端口扫描
nmap host`,
	})
}
