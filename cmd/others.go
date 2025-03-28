package cmd

import "github.com/spf13/cobra"

func newOthersCmd() *cobra.Command {
	return createBasicCommand(CommandConfig{
		Use:   "others",
		Short: "others 命令",
		Long:  "显示others",
		Message: `Others
jenkins 普通admin密码: 40b4eb13df9b4c75a4d62ce1bb824e74
jenkins root admin密码：a0a9d307ca614a49b70c34fe2a16300e
java -jar jenkins.war # /data/work/jenkins_work/jenkin_war
java -jar agent.jar -url http://localhost:8080/ -secret xxxx -name "slave-ubuntu" -webSocket -workDir "/data/jenkins/agent" # /data/work/jenins_work/jenkins_war/agent_normal
ecosda:jenkins-slave: ubuntu@192.168.15.28 passwd: 1qaz@WSX
ecosda:jenkins-master: chenwenyi@192.168.15.23 passwd: chenwenyi@123
`,
	})
}
