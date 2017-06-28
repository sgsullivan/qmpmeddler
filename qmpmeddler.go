package qmpmeddler

import (
	"fmt"
	"github.com/digitalocean/go-qemu/qmp"
	"time"
)

func Meddle(qmpSocketFile *string, qmpTimeout *int, jsonPayload *string) error {
	if *qmpSocketFile == "" {
		return fmt.Errorf("QMP socket file was not passed; check your usage!")
	}
	var monitor *qmp.SocketMonitor
	var err error
	switch {
	case *qmpSocketFile != "":
		fmt.Printf("Using QMP socketfile [%s]...\n", *qmpSocketFile)
		monitor, err = qmp.NewSocketMonitor("unix", *qmpSocketFile, time.Duration(*qmpTimeout)*time.Second)
	default:
		err = fmt.Errorf("unknown switch condition, check usage")
	}
	if err != nil {
		return err
	}
	if monErr := monitor.Connect(); monErr != nil {
		return monErr
	}
	defer monitor.Disconnect()
	cmd := []byte(*jsonPayload)
	raw, monRunErr := monitor.Run(cmd)
	if monRunErr != nil {
		return monRunErr
	}
	fmt.Printf("%s\n", raw)
	return nil
}
