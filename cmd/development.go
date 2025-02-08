package cmd

import "github.com/spf13/cobra"

// newGoCmd 创建 Go 命令
func newGoCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "go",
		Short: "Go 相关命令",
		Long:  "显示 Go 相关的信息和工具",
		Message: `Go 常用命令和操作：

# 下载安装
https://go.dev/dl/
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz    

# 环境配置
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on

# 常用命令
go mod init <project-name>
go mod tidy
go run main.go
go build
go test ./...`,
	})
}

// newPythonCmd 创建 Python 命令
func newPythonCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "python",
		Short:   "Python 相关命令",
		Long:    "显示 Python 相关的信息和工具",
		Message: "使用 --subcmd 参数查看具体命令，可选：conda, pip, venv",
		SubCommands: map[string]string{
			"conda": `Conda 常用命令：
conda --version
conda update conda
conda info --envs
conda create --name myenv python=3.9
conda create --name myenv --clone python_env
conda activate myenv
conda deactivate
conda list
conda install package_name
conda remove package_name
conda remove --name myenv package_name
conda remove --name myenv --all
`,

			"pytest": `pytest：
pytest --version`,
			"allure": `allure：
allure generate allure-results -o allure-report --clean,
allure serve allure-report`,
			"pip": `Pip 常用命令：
pip install package_name
pip uninstall package_name
pip list
pip freeze > requirements.txt
pip install -r requirements.txt
pip show package_name`,

			"venv": `Python venv 使用：
python -m venv myenv
source myenv/bin/activate  # Linux
myenv\Scripts\activate.bat  # Windows
deactivate`,
		},
	})
}

// newDockerCmd 创建 Docker 命令
func newDockerCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "docker",
		Short: "Docker 相关命令",
		Long:  "显示 Docker 相关的信息和常用命令",
		Message: `
# permission: Docker 权限问题解决：
sudo usermod -aG docker $USER
newgrp docker
sudo systemctl restart docker`,
	})
}

func newPodmanCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "podman",
		Short: "Docker 相关命令",
		Long:  "显示 Docker 相关的信息和常用命令",
		Message: `
podman load -i image.tar # 加载镜像
podman images
podman run -itd -v /data:/data --shm-size=8g --name Name image_name bash
podman exec -it Name bash
podman stop Name
podman rm Name
`,
	})
}

// newGitCmd 创建 Git 命令
func newGitCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "git",
		Short:   "Git 相关命令",
		Long:    "显示 Git 相关的信息和常用命令",
		Message: "使用 --subcmd 参数查看具体命令，可选：basic, branch, remote",
		SubCommands: map[string]string{
			"basic": `Git 基本命令：
git init
git add .
git commit -m "message"
git status
git log
## 假设git 已经提交了文件,需要移除的化这么操作
git rm -r --cached . # 清空缓存
git add . 
#########
git diff`,

			"branch": `Git 分支操作：
git branch --show-current
git checkout -b new_branch
git checkout branch_name
git merge branch_name
git branch -d branch_name`,

			"remote": `Git 远程操作：		
git remote remove origin # 本地删除远程
git remote add origin url
git push -u origin main
git pull origin main
git fetch
git clone url`,
		},
	})
}

// newGitCmd 创建 Git 命令
func newRepoCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "repo",
		Short: "Repo 相关命令",
		Long:  "显示 Repo 相关的信息和常用命令",
		Message: `Repo 基本命令：
repo init --repo-url "https://mirrors.tuna.tsinghua.edu.cn/git/git-repo" -u git@github.com:nippon1218/manifests.git -m test.xml -b main --no-repo-verify
repo sync -j 4`,
	})
}
