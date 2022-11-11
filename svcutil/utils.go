package svcutil

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kardianos/service"
)

type Svc struct {
	SvcName    string
	SvcDisName string
	SvcDes     string
	ExecPath   string
	SvcFun     func()
}

func (svc *Svc) Init() error {
	serConfig := &service.Config{
		Name:        svc.SvcName,
		DisplayName: svc.SvcDisName,
		Description: svc.SvcDes,
	}
	if runtime.GOOS == "linux" {
		serConfig.Dependencies = []string{"After=network.target syslog.target"}
		serConfig.Option = service.KeyValue{"Restart": "always"}
	}
	serConfig.Executable = svc.ExecPath

	s, err := service.New(svc, serConfig)
	if err != nil {
		fmt.Println(err, "service.New() err")
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println("install err", err)
			} else {
				fmt.Println("install success")
			}
			return nil
		}

		if os.Args[1] == "remove" {
			s.Stop()
			fmt.Println("正在停止服务")
			err = s.Uninstall()
			if err != nil {
				fmt.Println("uninstall err", err)
			} else {
				fmt.Println("uninstall success")
			}
			return nil
		}

	}

	err = s.Run()
	if err != nil {
		fmt.Println("s.Run err", err)
		return err
	}
	return nil
}

func (svc *Svc) run() {
	//go server.Init_proxy("0.0.0.0:" + port)
	go svc.SvcFun()
	fmt.Println("开机自启动服务 - run")
}

func (svc *Svc) Start(s service.Service) error {
	fmt.Println("server start")
	go svc.run()
	return nil
}

func (svc *Svc) Stop(s service.Service) error {
	fmt.Println("server stop")
	return nil
}
