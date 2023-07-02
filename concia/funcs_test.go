package lacia

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

// eg2
var ch1 chan int
var ch2 chan TestObj

func init() {
	if ch1 == nil {
		ch1 = make(chan int)
	}
	if ch2 == nil {
		ch2 = make(chan TestObj)
	}
}

type TestObj struct {
	A string
	B int
}

const (
	file1 = "../src_file/t1.txt"
	file2 = "../src_file/t2.txt"
)

// eg1
func TestCConcurrencyGos(t *testing.T) {
	err := addContent(file1, file2, "write hello 2 file!")
	fmt.Println(err)
}

// eg2
func TestCConcurrencyGos0(t *testing.T) {
	err := fc00()
	fmt.Println(err) // 无报错时<nil>     下设人造错误时：fc01执行异常 结果B1A2B3A4B5A6B7A8B9A10B11
}

func fc00() error {
	return ConcurrencyGos(
		// A、B、C...
		fc01, fc02, fc03,
		//	Your more func...
	)
}

func fc01() (err error) {
	for i := 0; i <= 20; i++ {
		//if i == 11 { // 人造错误
		//	err = errors.New("fc01执行异常")
		//	break
		//}
		ch1 <- i
	}
	ch1 <- -1
	close(ch1)
	return
}

func fc02() error {
fc2Lp:
	for {
		select {
		case data := <-ch1:
			if data == -1 {
				break fc2Lp
			}
			input := TestObj{A: "A", B: data + 1}
			if data%2 == 0 {
				input.A = "B"
			}
			ch2 <- input
		}
	}
	ch2 <- TestObj{A: "stop"}
	close(ch2)
	return nil
}

func fc03() error {
	str := ""
fc3Lp:
	for {
		select {
		case data := <-ch2:
			if data.A == "stop" {
				break fc3Lp
			}
			str += data.A + strconv.Itoa(data.B)
		}
	}
	fmt.Println(str)
	return writeTo(file1, str)
}

func writeTo(fileName, content string) error {
	dstFile, err := os.Create(fileName)
	// 制造错误，如果有错误，将把错误处理成一个切片
	//if fileName == file1Name {
	//err = errors.New("人造错误") //错误切片：[人造错误, 人造错误]
	//}

	if err != nil {
		//fmt.Println("写入报错：", err.Error())
		return err
	}
	defer dstFile.Close()
	dstFile.WriteString(content + "\n")
	//fmt.Println("写入完成：", fileName)
	return nil
}

func addContent(file1Name, file2Name, content string) error {
	return ConcurrencyGos(
		// A
		func() error {
			return writeTo(file1Name, content) // do A sth
		},
		// B
		func() error {
			return writeTo(file2Name, content) // do B sth
		},
		//	Your more func...
	)
}
