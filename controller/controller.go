package controller

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func Version(data string) {
	if data == "true" {
		cmd := exec.Command("lsb_release", "-a")
		con, _ := cmd.Output()
		fmt.Println(string(con))
	}
}

//go:embed replace/ubuntu16.04.txt
var ubuntu16 string

//go:embed replace/ubuntu18.04.txt
var ubuntu18 string

//go:embed replace/ubuntu20.04.txt
var ubuntu20 string

func ReplaceOsSource(data string) {
	if data == "true" {
		cmd := exec.Command("lsb_release", "-r")
		con, _ := cmd.Output()
		switch string(con[9 : len(con)-1]) {
		case "20.04":
			file, _ := os.OpenFile("/etc/apt/sources.list", os.O_WRONLY|os.O_TRUNC, 0777)
			defer file.Close()
			write, err := file.Write([]byte(ubuntu20))
			if err != nil {
				fmt.Println("写入错误：", err)
			} else {
				cmd1 := exec.Command("sudo", "apt", "update", "-y")
				con1, _ := cmd1.StdoutPipe()
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					defer wg.Done()
					reader := bufio.NewReader(con1)
					for {
						resString, err := reader.ReadString('\n')
						if err == io.EOF {
							return
						}
						fmt.Printf(resString)
					}
				}()
				err = cmd1.Start()
				wg.Wait()
				fmt.Println(con1)
				fmt.Printf("写入成功，本地成功写入%d字节\n", write)
			}
		case "18.04":
			file, _ := os.OpenFile("/etc/apt/sources.list", os.O_WRONLY|os.O_TRUNC, 0777)
			defer file.Close()
			write, err := file.Write([]byte(ubuntu18))
			if err != nil {
				fmt.Println("写入错误：", err)
			} else {
				cmd1 := exec.Command("sudo", "apt", "update", "-y")
				con1, _ := cmd1.StdoutPipe()
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					defer wg.Done()
					reader := bufio.NewReader(con1)
					for {
						resString, err := reader.ReadString('\n')
						if err == io.EOF {
							return
						}
						fmt.Printf(resString)
					}
				}()
				err = cmd1.Start()
				wg.Wait()
				fmt.Println(con1)
				fmt.Printf("写入成功，本地成功写入%d字节\n", write)
			}
		case "16.04":
			file, _ := os.OpenFile("/etc/apt/sources.list", os.O_WRONLY|os.O_TRUNC, 0777)
			defer file.Close()
			write, err := file.Write([]byte(ubuntu16))
			if err != nil {
				fmt.Println("写入错误：", err)
			} else {
				cmd1 := exec.Command("sudo", "apt", "update", "-y")
				con1, _ := cmd1.StdoutPipe()
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					defer wg.Done()
					reader := bufio.NewReader(con1)
					for {
						resString, err := reader.ReadString('\n')
						if err == io.EOF {
							return
						}
						fmt.Printf(resString)
					}
				}()
				err = cmd1.Start()
				wg.Wait()
				fmt.Println(con1)
				fmt.Printf("写入成功，本地成功写入%d字节。\n", write)
			}
		default:
			fmt.Println("本工具目前仅支持：ubuntu20，18,16版本系统")
			fmt.Println("您的系统版本为：", con)
		}
	}
}

// FindPort 查询端口占用
func FindPort(data string) {
	cmd := exec.Command("netstat", "-antup", "|", "grep", data)
	con, err := cmd.Output()
	if err != nil {
		fmt.Println("命令不存在，正在安装...")
		cmd1 := exec.Command("apt-get", "install", "net-tools")
		err := cmd1.Start()
		if err != nil {
			fmt.Println("安装错误：", err)
		}
	} else {
		fmt.Println(string(con))
	}
}
