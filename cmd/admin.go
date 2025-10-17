package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nav-server/app/admin"
	"nav-server/app/admin/config"
	"os"
)

func adminCMD() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "admin",
		Short: "启动admin服务",
		Run: func(cmd *cobra.Command, args []string) {
			//启动时要执行的操作写在这里
			admin.StartServer()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			//启动前要执行的方法写在这里，例如加载配置文件、初始化数据库连接等
			var (
				etcdEndpoints      = os.Getenv("etcdEndpoints")
				etcdUsername       = os.Getenv("etcdUsername")
				etcdPassword       = os.Getenv("etcdPassword")
				etcdConfigFilename = os.Getenv("etcdConfigFilename")
			)
			//解析配置
			if len(etcdEndpoints) != 0 {
				fmt.Println("尝试从etcd读取配置")
				config.ParseAndWatchFromEtcd(etcdEndpoints, etcdUsername, etcdPassword, etcdConfigFilename)
			} else {
				fmt.Println("尝试从本地文件读取配置")
				config.ParseAndWatchFromFile(cfgPath)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "配置文件路径")
	return cmd
}

func init() {
	RootCMD.AddCommand(adminCMD())
}
