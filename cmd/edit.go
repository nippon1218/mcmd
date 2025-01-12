package cmd

import "github.com/spf13/cobra"

// newVimCmd 创建 Vim 命令
func newVimCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "vim",
		Short: "Vim 编辑器命令",
		Long:  "显示 Vim 编辑器相关的配置和常用命令",
		Message: `sudo apt install vim-gtk3 # 这个vim 其实有clickboard
# 如何编译vim
sudo apt remove vim vim-runtime vim-common vim-gtk3 # 先卸载vim
sudo apt install -y ncurses-dev # chatgpt说的，解决no terminal library found
./configure --with-features=huge --enable-multibyte --with-x --enable-gui=auto # 1.9以后可以这样编译
make && sudo make install

:PlugInstall # 安装vim插件
:PlugUpdate # 更新vim插件
:PlugClean # 清理vim插件

		`,
	})
}

// newTmuxCmd 创建 Tmux 命令
func newTmuxCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "tmux",
		Short:   "Tmux 终端复用器命令",
		Long:    "显示 Tmux 终端复用器相关的配置和常用命令",
		Message: "使用 --subcmd 参数查看具体命令，可选：basic, config, session",
		SubCommands: map[string]string{
			"basic": `Tmux 基本操作：
# 会话管理
tmux new -s name     - 创建新会话
tmux ls             - 列出会话
tmux attach -t name - 进入会话
tmux kill-session -t name - 关闭会话

# 窗口操作 (Prefix = Ctrl+b)
Prefix c    - 创建新窗口
Prefix n    - 下一个窗口
Prefix p    - 上一个窗口
Prefix &    - 关闭当前窗口
Prefix ,    - 重命名窗口

# 面板操作
Prefix %    - 垂直分割
Prefix "    - 水平分割
Prefix x    - 关闭面板
Prefix 方向键 - 切换面板
Prefix z    - 最大化/还原面板`,

			"config": `Tmux 配置说明：
# 配置文件位置
~/.tmux.conf

# 常用配置示例
# 修改前缀键
set -g prefix C-a
unbind C-b
bind C-a send-prefix

# 开启鼠标支持
set -g mouse on

# 窗口编号从1开始
set -g base-index 1

# 自动重新编号
set -g renumber-windows on

# 状态栏配置
set -g status-bg black
set -g status-fg white
set -g status-interval 1`,

			"session": `Tmux 会话管理：
# 会话操作
Prefix d    - 分离当前会话
Prefix s    - 选择会话
Prefix $    - 重命名会话
Prefix (    - 切换到上一个会话
Prefix )    - 切换到下一个会话

# 保存和恢复会话
# 需要安装 tmux-resurrect 插件
Prefix Ctrl+s - 保存会话
Prefix Ctrl+r - 恢复会话

# 会话脚本示例
#!/bin/bash
tmux new-session -d -s dev
tmux split-window -h
tmux split-window -v
tmux new-window -n server
tmux select-window -t dev:1
tmux attach -t dev`,
		},
	})
}
