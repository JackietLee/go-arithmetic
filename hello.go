package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	//scans()
	//deal()
	//out()
	copy()
}

func scans() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func deal() {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

func out() {
	var buf [7]byte
	var content []byte
	read, _ := os.Stdin.Read(buf[:])
	content = append(content, buf[:read]...)
	fmt.Print(string(buf[:]))
	//os.Stdin.WriteString(string(buf[:]))
	file, err := os.Create("D:\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	s := "hello world"
	compare := strings.Compare(s, string(buf[:]))
	println(compare)
	file.WriteString(string(buf[:]))
	file.WriteString(s)
}

func copy() {
	open, err := os.Open("D:\\test.txt")
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(open)
	if err != nil {
		fmt.Println(err)
		return
	}
	create, err := os.Create("D:\\test1.txt")
	defer func(create *os.File) {
		err := create.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(create)
	if err != nil {
		fmt.Println(err)
	}
	bytes := make([]byte, 1024)
	for {
		read, err := open.Read(bytes)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		create.Write(bytes[:read])
	}
}

type A interface {
}
