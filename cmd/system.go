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
cat /etc/os-release
top
htop

# 服务管理
systemctl status service_name
systemctl start service_name
systemctl enable service_name

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
