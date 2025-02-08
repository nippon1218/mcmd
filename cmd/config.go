package cmd

// DirectoryMapping maps shortcut names to actual directory paths
var DirectoryMapping = map[string]string{
	"work":   "/data/work",
	"tmp":   "/data/tmp",
	"go":     "/data/work/go_work",
	"conda":  "/data/work/python_work/pytorch_work",
	"llm":    "/data/work/llm_work",
	"ebpf":   "/data/work/githubwork/ebpf",
	"mcmd":   "/data/work/go_work/ws/mcmd",
	"python": "/data/work/python_work",
	"pytorch":"/data/work/python_work/pytorch_work",
}

// SysCommandMapping maps system commands to their shell equivalents
var SysCommandMapping = map[string]string{
	"sys": "uname -a",
	"mem": "free -h",
	"net": "netstat -tuln",
}

// CommandMap 定义了常用命令的快捷方式映射
var CommandMap = map[string]string{
	// 环境相关
	"base": "conda activate base", // 激活 conda base 环境
	"pytorch": "conda activate pytorch_env", // 激活 pytorch_env环境
	"pytorch_env": "conda activate pytorch_env", // 激活 pytorch_env环境
	"zsh":  "source ~/.zshrc",     // 重新加载 zsh 配置

	// 编辑器相关
	"edit": "vim /data/work/go_work/ws/mcmd/cmd/config.go", // 编辑配置文件

	// 文档查看
	"dfsx": "google-chrome /home/vincent/Documents/dfsx.png", // 打开东方四象文档
	"rjjg": "google-chrome /home/vincent/Documents/dfsx.png", // 打开软件架构文档
}

// GetMappedDirectory returns the mapped directory path for a given shortcut
func GetMappedDirectory(shortcut string) (string, bool) {
	path, exists := DirectoryMapping[shortcut]
	return path, exists
}

// GetCommand returns the command for a given shortcut
func GetCommand(env string) (string, bool) {
	cmd, exists := CommandMap[env]
	return cmd, exists
}
