package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// CommandConfig 定义命令的配置结构
type CommandConfig struct {
	Use         string            // 命令名称
	Short       string            // 短描述
	Long        string            // 长描述
	Message     string            // 默认消息
	SubCommands map[string]string // 子命令及其对应的消息
}

// createBasicCommand 创建基本命令
// 用于创建没有子命令的简单命令
func createBasicCommand(config CommandConfig) *cobra.Command {
	return &cobra.Command{
		Use:   config.Use,
		Short: config.Short,
		Long:  config.Long,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config.Message)
		},
	}
}

// runCmdByContent 根据命令的内容来执行不同的shell命令
func runCmdByContent(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("请提供要执行的命令")
		return
	}

	content := args[0]
	switch content {
	case "version":
		fmt.Println(cmd.Version)
	case "help":
		cmd.Help()
	default:
		fmt.Printf("执行shell命令: %s\n", content)
		output, err := exec.Command("sh", "-c", content).CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	}
}

// createCommandWithSubcommands 创建带子命令的命令
// 用于创建包含多个子命令的复杂命令
func createCommandWithSubcommands(config CommandConfig) *cobra.Command {
	var subcmd string

	cmd := &cobra.Command{
		Use:   config.Use,
		Short: config.Short,
		Long:  config.Long,
		Run: func(cmd *cobra.Command, args []string) {
			if subcmd == "" {
				fmt.Println(config.Message)
				return
			}

			if msg, exists := config.SubCommands[subcmd]; exists {
				fmt.Println(msg)
			} else {
				fmt.Printf("未知的子命令: %s\n", subcmd)
				fmt.Println("可用的子命令:")
				for key := range config.SubCommands {
					fmt.Printf("  - %s\n", key)
				}
			}
		},
	}

	// 只有当有子命令时才添加 subcmd 标志
	if len(config.SubCommands) > 0 {
		cmd.Flags().StringVarP(&subcmd, "subcmd", "s", "", "要执行的子命令")
	}

	return cmd
}

// NewRootCommand 创建并返回根命令
func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "mcmd",
		Short: "mcmd 是一个简单的命令行工具",
		Long: `mcmd 是一个命令行工具集合，包含了常用的开发和系统管理命令。
例如：Go、Python、Docker、Git等开发工具的常用命令，
以及 VPN、Linux 等系统管理命令。`,
		Version: "1.0.0",
	}
}

// RegisterCommands 注册所有子命令
func RegisterCommands(rootCmd *cobra.Command) {
	// 开发工具命令
	rootCmd.AddCommand(
		newGoCmd(),
		newPythonCmd(),
		newDockerCmd(),
		newPodmanCmd(),
		newGitCmd(),
		newRepoCmd(),
		newOthersCmd(),
	)

	// 系统工具命令
	rootCmd.AddCommand(
		newVpnCmd(),
		newLinuxCmd(),
		newWindowsCmd(),
		newNetworkCmd(),
		newGotoCommand(), // 添加goto命令
	)

	// 编辑器命令
	rootCmd.AddCommand(
		newVimCmd(),
		newTmuxCmd(),
		newWindsurfCmd(),
	)

	// ecosda 命令
	rootCmd.AddCommand(
		newecosdaCmd(),
		newGPUCmd(),
		newLLMCmd(),
	)
}
