package preciousnotes

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewPreciousNotesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preciousnotes",
		Short: "this is my first project in go ",
		Long:  "",
		// 命令出错时，不打印帮助信息。
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
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
	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	fmt.Println("Hello PreciousNotes!")
	return nil
}
