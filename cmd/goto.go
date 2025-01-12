package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// newGotoCommand creates a command for quick navigation and command execution
func newGotoCommand() *cobra.Command {
	var targetDir string
	var initShell bool
	var condaEnv string

	cmd := &cobra.Command{
		Use:     "goto",
		Short:   "Quick navigation and command execution",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			if initShell {
				// 输出shell函数定义
				fmt.Println(`_mcmd_goto_cd() {
    if [ "$#" -eq 0 ]; then
        mcmd goto --help
        return
    fi
    local output
    # 检查是否是直接路径还是别名
    if [[ "$1" == /* ]] || [[ "$1" == .* ]] || [[ "$1" == ~* ]]; then
        output=$(mcmd goto -d "$1")
    else
        output=$(mcmd goto -d "$1")
    fi
    if [[ $output == "MCMD_CD:"* ]]; then
        cd "${output#MCMD_CD:}"
    else
        echo "$output"
    fi
}

_mcmd_goto_conda() {
    if [ "$#" -eq 0 ]; then
        mcmd goto --help
        return
    fi
    local output
    output=$(mcmd goto -r "$1")
    if [[ $output == CONDA_CMD:* ]]; then
        eval "${output#CONDA_CMD:}"
    else
        echo "$output"
    fi
}

# 定义主命令
mcmd_goto() {
    if [ "$1" = "-d" ] || [ "$1" = "--dir" ]; then
        _mcmd_goto_cd "$2"
    elif [ "$1" = "-r" ] || [ "$1" = "--run" ]; then
        _mcmd_goto_conda "$2"
    else
        mcmd goto "$@"
    fi
}

# 设置别名
alias goto=_mcmd_goto_cd
alias mcmd-goto=mcmd_goto
alias activate=_mcmd_goto_conda`)
				return
			}

			if len(args) == 0 && targetDir == "" && condaEnv == "" {
				fmt.Println("Goto command usage:")
				fmt.Println("After initialization, you can use:")
				fmt.Println("1. Directory navigation:")
				fmt.Println("   goto /path/to/dir        # Quick navigation with absolute path")
				fmt.Println("   goto work                # Quick navigation to /data/work")
				fmt.Println("   goto go                  # Quick navigation to /data/work/go_work")
				fmt.Println("   mcmd-goto -d /path/to/dir # Alternative way")
				fmt.Println("\nAvailable shortcuts:")
				// 打印所有可用的目录快捷方式
				fmt.Println("   Directory shortcuts:")
				for shortcut, path := range DirectoryMapping {
					fmt.Printf("   - %-10s -> %s\n", shortcut, path)
				}
				fmt.Println("\n2. Quick commands:")
				fmt.Println("   mcmd-goto -r pytorch_env # Activate pytorch environment")
				fmt.Println("   mcmd-goto -r base       # Activate base environment")
				fmt.Println("   mcmd-goto -r edit       # Edit config file")
				fmt.Println("   mcmd-goto -r zsh        # Reload zsh configuration")
				fmt.Println("\n3. System commands:")
				fmt.Println("   mcmd-goto sys   # Show system info (uname -a)")
				fmt.Println("   mcmd-goto mem   # Show memory info (free -h)")
				fmt.Println("   mcmd-goto net   # Show network ports (netstat -tuln)")
				fmt.Println("\n4. Other commands:")
				fmt.Println("   mcmd-goto 'custom command'  # Run custom command")
				fmt.Println("   mcmd-goto version          # Show version")
				fmt.Println("   mcmd-goto help             # Show this help")
				fmt.Println("\nTo enable shell integration, add this to your .bashrc or .zshrc:")
				fmt.Println("eval \"$(mcmd goto --init-shell)\"")
				return
			}

			// Handle conda environment activation
			if condaEnv != "" {
				if cmd, exists := GetOtherCommand(condaEnv); exists {
					fmt.Printf("CONDA_CMD:%s\n", cmd)
					return
				}
				fmt.Printf("Unknown conda environment: %s\n", condaEnv)
				return
			}

			// Handle directory navigation
			if targetDir != "" {
				// 首先检查是否是目录映射
				if mappedPath, exists := GetMappedDirectory(targetDir); exists {
					targetDir = mappedPath
				}

				absPath, err := filepath.Abs(targetDir)
				if err != nil {
					fmt.Printf("Error resolving path: %v\n", err)
					return
				}

				// Check if directory exists
				if _, err := os.Stat(absPath); os.IsNotExist(err) {
					fmt.Printf("Directory does not exist: %s\n", absPath)
					return
				}

				// Check if it's a directory
				fileInfo, err := os.Stat(absPath)
				if err != nil || !fileInfo.IsDir() {
					fmt.Printf("Not a directory: %s\n", absPath)
					return
				}

				// Output special format for directory change
				fmt.Printf("MCMD_CD:%s\n", absPath)
				return
			}

			// Handle command execution
			if len(args) > 0 {
				content := args[0]
				switch content {
				case "version":
					fmt.Println(cmd.Version)
				case "help":
					cmd.Help()
				default:
					// Check if it's a mapped command
					shellCmd := content
					mappedCmd, exists := SysCommandMapping[content]
					if exists {
						shellCmd = mappedCmd
					} else {
						// Remove surrounding quotes if present
						shellCmd = strings.Trim(content, "'\"")
					}

					fmt.Printf("执行shell命令: %s\n", shellCmd)
					output, err := exec.Command("sh", "-c", shellCmd).CombinedOutput()
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(string(output))
				}
			}
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&targetDir, "dir", "d", "", "Target directory to navigate to")
	cmd.Flags().StringVarP(&condaEnv, "run", "r", "", "Conda environment to activate")
	cmd.Flags().BoolVar(&initShell, "init-shell", false, "Initialize shell integration")

	return cmd
}
