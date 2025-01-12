package cmd

// DirectoryMapping maps shortcut names to actual directory paths
var DirectoryMapping = map[string]string{
	"work":   "/data/work",
	"go":     "/data/work/go_work",
	"conda":  "/data/work/python_work/pytorch_work",
	"llm":    "/data/work/llm_work",
	"ebpf":   "/data/work/githubwork/ebpf",
	"mcmd":   "/data/work/go_work/ws/mcmd",
	"python": "/data/work/python_work",
}

// SysCommandMapping maps system commands to their shell equivalents
var SysCommandMapping = map[string]string{
	"sys": "uname -a",
	"mem": "free -h",
	"net": "netstat -tuln",
}

// CondaCommandMapping maps conda environment names to activation commands
var OtherCommandMapping = map[string]string{
	"pytorch_env": "conda activate pytorch_env",
	"base":        "conda activate base",
	"edit":        "vim /data/work/go_work/ws/mcmd/cmd/config.go",
	"zsh":         "source ~/.zshrc",
}

// GetMappedDirectory returns the mapped directory path for a given shortcut
func GetMappedDirectory(shortcut string) (string, bool) {
	path, exists := DirectoryMapping[shortcut]
	return path, exists
}

// GetCondaCommand returns the conda activation command for a given environment
func GetOtherCommand(env string) (string, bool) {
	cmd, exists := OtherCommandMapping[env]
	return cmd, exists
}
