package cmd

import (
	"fmt"
	"github.com/sgsullivan/qmpmeddler"
	"github.com/spf13/cobra"
	"os"
)

var qmpSocketFile string
var qmpTimeout int
var jsonPayload string

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "execute arbitrary QMP JSON commands",
	Long: `execute arbitrary QMP JSON commands:

http://wiki.qemu.org/Documentation/QMP
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := qmpmeddler.Meddle(&qmpSocketFile, &qmpTimeout, &jsonPayload); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(executeCmd)
	executeCmd.Flags().StringVarP(&qmpSocketFile, "sockfile", "", "", "file location to the QMP socket file; use this option if virsh doesn't have the socket.")
	executeCmd.Flags().IntVarP(&qmpTimeout, "timeout", "", 5, "QMP timeout")
	executeCmd.Flags().StringVarP(&jsonPayload, "json", "", "", "the raw JSON to send to QMP.")
}
