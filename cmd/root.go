package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "qmpmeddler",
	Short: "Meddle with the Qemu Machine Protocol (QMP)",
	Long: `This program can communicate with the Qemu Machine Protocol (QMP).
This program requires direct access to the socket you give, so if the domain is managed
by libvirt this will not work, unless libvirtd is shutdown.

http://wiki.qemu.org/Documentation/QMP

`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
