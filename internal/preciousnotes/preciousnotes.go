package preciousnotes

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"precious-notes/internal/pkg/log"
)

var cfgFile string

func NewPreciousNotesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preciousnotes",
		Short: "this is my first project in go ",
		Long:  "",
		// 命令出错时，不打印帮助信息。
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			//verflag.PrintAndExitIfRequested()

			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中

			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)
	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the this project configuration file. Empty string for no configuration file.")
	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	// 打印所有的配置项及其值
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))
	// 打印 db -> username 配置项的值
	log.Infow(viper.GetString("db.username"))
	return nil
}
