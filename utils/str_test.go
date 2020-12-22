package lacia

import (
	"fmt"
	"testing"
)

func TestRemoveX(t *testing.T) {
	var r = `\`
	var x = `abc\scs\\c\..\\\.`
	res := RemoveX(x, r)
	fmt.Println(res) // abcscsc...

	var r1 = `{`
	var x1 = `abs.{{ff,{qp{vf`
	res1 := RemoveX(x1, r1)
	fmt.Println(res1) // abs.ff,qpvf

	var r2 = `-`
	var x2 = `h-e---ll--o`
	res2 := RemoveX(x2, r2)
	fmt.Println(res2) // hello

	var r3 = ` `
	var x3 = `h  el   l o`
	res3 := RemoveX(x3, r3)
	fmt.Println(res3) // hello
}

func TestDeletePreAndSufSpace(t *testing.T) {
	s1 := " a  a a  "
	fmt.Println("s1原：", s1)               //  a  a a
	fmt.Println(DeletePreAndSufSpace(s1)) //a  a a
	s2 := " 和  a 和   "
	fmt.Println("s2原：", s2)               //  和  a 和
	fmt.Println(DeletePreAndSufSpace(s2)) //和  a 和
	s3 := "   "                           //都是空格时全部消除
	fmt.Println("s3原：", s3)               //
	fmt.Println(DeletePreAndSufSpace(s3)) //
}
