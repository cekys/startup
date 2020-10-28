package main

import (
	"errors"
	"fmt"
	"github.com/cekys/gopkg"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	DefaultConfig  = "config.json"
	DefaultSimple  = "simple"
	DefaultForking = "forking"
)

func main() {
	conf := config{}
	defaultConf := make(config, 1)
	// 读取配置文件
	err := conf.ReadFromFile(DefaultConfig)
	if err != nil {
		// 如果配置文件不存在则写入默认配置文件
		if errors.Is(err, os.ErrNotExist) {
			err = defaultConf.WriteToFile(DefaultConfig)
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}
	// 读取配置文件中每一项指令对象
	for n, i := range conf {
		if i.Enabled {
			// 参数按空格分词
			args := strings.Fields(i.Args)
			// 组合指令
			cmd := exec.Command(i.Command, args...)
			// 打印组合后的指令
			fmt.Printf("%d: %v\n", n+1, cmd)
			// simple模式运行指令,等待指令运行结束
			if i.Type == "" || i.Type == DefaultSimple {
				err = gopkg.ProgramRealtimeOutput(cmd)
			}
			// forking模式运行指令,不等待指令运行结束
			if i.Type == DefaultForking {
				err = cmd.Start()
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
