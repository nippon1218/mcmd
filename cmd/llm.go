package cmd

import "github.com/spf13/cobra"

// newLinuxCmd 创建 Linux 命令
func newLLMCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "llm",
		Short:   "llm 相关命令",
		Long:    "显示llm系统相关的常用命令和操作",
		Message: "使用 --subcmd 参数查看具体命令，可选：llama, llm",
		SubCommands: map[string]string{
			"llm":   `llm：`,
			"llama": `llama：`,
		},
	})
}
