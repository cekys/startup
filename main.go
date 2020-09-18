package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 读取配置文件
	err := readJSON(&conf, DefaultConfig)
	if err != nil {
		// 如果配置文件不存在则写入默认配置文件
		if errors.Is(err, os.ErrNotExist) {
			err = writeJSON(&defaultConf, DefaultConfig)
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}
	// 读取配置文件中每一项指令对象
	for _, i := range conf {
		if i.Enabled {
			// 参数按空格分词
			args := strings.Fields(i.Args)
			// 组合指令
			cmd := exec.Command(i.Command, args...)
			// 打印组合后的指令
			fmt.Println(cmd)
			// 运行指令,不等待指令运行结束
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
