package cmd

import "github.com/spf13/cobra"

func newecosdaCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "ecosda",
		Short:   "ecosda 相关命令",
		Long:    "显示 ecosda 相关的信息和常用命令",
		Message: "使用 --subcmd 参数查看具体命令，可选：url, who, env, all",
		SubCommands: map[string]string{
			"url": `ecosda svn, 看资料的地方：
svn://192.168.4.145/mynote/Arc_Group
svn://192.168.4.145/mynote/Com_Group
svn://192.168.4.145/mynote/Soft_Group
svn://192.168.4.145/mynote/dfsx_code/Soft

192.168.4.132  # 内网gitlab
192.168.15.49:80 # docker的harbor
192.168.15.49:80/library/ubuntu22.04_20250107:V1.0

## 资料文档
Soft_Group/01_工作区/01.02软件架构组/软件栈设计/子系统HLD
/d/work/svn/Soft_Group/01_工作区/01.02软件架构组/AI应用分析

## git bash
/d/work
/d/Users/admin/Desktop

`,
			"who": `# 软件组:
陈小刚: 系统集成和交付技术专家
安李： 配置管理
李冠华： 模拟器负责人
刘奕： 软件组PM


# IT
钱锦波： IT，钱老师，任何问题都可以找
任卫星： IT，信息安全专员


`,
			"env": `make 4.3, cmake 3.22.1 gcc/g++ 11.4, pytorch 2.4, python3.12 
torch 2.4.0+cpu(pip)
`,
			"work": `1. 维护qemu景象以及docker image ?NO!!
2.  测试用例的维护
`,
			"ecosda": `qemu(Quick Emulator）是一个开源的虚拟化技术，它可以用来模拟不同的计算机架构和运行多个操作系统。QEMU 主要用于硬件仿真、虚拟化、系统开发和测试等场景。
ldap: 存储在LDAP（轻量级目录访问协议）目录服务中的用户账户信息
"office":"chenwenyi@ecosda.com, ecosda-1, "
`,
		},
	})
}

func newLLMCmd() *cobra.Command {
	return createCommandWithSubcommands(CommandConfig{
		Use:     "llm",
		Short:   "llm 相关命令",
		Long:    "显示llm系统相关的常用命令和操作",
		Message: "使用 --subcmd 参数查看具体命令，可选：llama, llm",
		SubCommands: map[string]string{
			"llm": `LLM(Large Language Model)：通过大规模数据训练出的自然语言处理模型 
微调：将预训练模型微调为适合特定任务的模型

`,
			"llama":          `llama：`,
			"models":         `GPT3-175b, BERT-Large, BERT-Base, LLAMA2-7B, LLAMA2-13B, LLAMA2-34B, LLAMA3-7B, LLAMA3-70B`,
			"transformer":    `transformer：`,
			"prompt":         `prompt：`,
			"chatglm":        `chatglm：`,
			"MoE":            `MoE: Mixture of Expers：`,
			"NLP":            "自然语言处理",
			"RDMA":           "",
			"分布式训练方式":        "数据并行，模型并行， 张量并行，通信并行，计算并行",
			"数据并行":           "多个计算节点同时处理同一份数据，可以提高计算性能，但是会带来数据不一致问题，需要进行同步操作",
			"注意力(Attention)": "它借鉴了人类的注意力机制，让模型在处理信息时能够重点关注重要的部分,它允许模型根据输入序列的不同部分给予不同的关注程度，从而提高对关键信息的捕捉能力，尤其是在长序列的上下文中。",
		},
	})
}

func newGPUCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "gpu",
		Short: "gpu",
		Long:  "显示gpu系统相关的常用命令和操作",
		Message: `GPU
常用GPU：A100, H100, Ascend 910B
竞争对手GPU：A100等

TPS: Tokens Per Second, 表示每秒能处理的词元(token)数量, 这是用于衡量大语言模型(LLM)处理速度的指标
, TPS受多个因素影响，包括模型大小、硬件配置、批处理大小等

TFLOPS: Tera Floating Point Operations Per Second,表示每秒可以执行的浮点运算次数，以万亿次(tera)为单位
,是衡量计算设备(如GPU、TPU)原始计算能力的指标,1 TFLOPS = 1万亿次浮点运算/秒,分为理论TFLOPS(硬件理论峰值)和实际TFLOPS(实际应用中达到的性能)

MFU: Model FLOPs Utilization,  模型FLOP利用率，衡量实际使用中对硬件计算能力的利用程度,计算公式：MFU = 实际TFLOPS / 理论峰值TFLOPS × 100%
,例如，如果一个GPU的理论峰值是100 TFLOPS，而实际运行时达到60 TFLOPS，则MFU为60%,MFU反映了模型和硬件的匹配程度，以及优化的空间

"die和tile：一个die是GPU的基本单位，包含多个tile，tile是GPU的最小计算单元，通常由多个核心组成，每个tile可以同时执行多个线程，用于并行处理任务。"

`,
	})
}
