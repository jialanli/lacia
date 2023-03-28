package lacia

import (
	"fmt"
	"testing"
)

/*
func TestRemoveX				去除指定字符
func TestDeletePreAndSufSpace	去除首尾空格
func TestCalcStrFrequency		指定字符在字符串中出现的次数（src范围:仅限于字母和数字）
func TestCalcStrFrequencyWith	指定字符在字符串中出现的次数（src范围:所有通用）
func TestVersionGreaterThanT	字符串版本号比较
func TestSplitByManyStrWith		按多个指定字符切割字符串
func TestTrimLineFeed			去除字符串中所有多余换行符(换行的时候只保留一个换行符),首尾换行符也一并去掉
func TestTrimAllFeeds			去除字符串中所有换行符
func TestTrimHeadTailFeedsOnce	字符串中首尾单次换行符
func TestTrimHeadTailFeeds		字符串中首尾所有换行符
func TestTrimHeadTailFeedsAndSpaces		去除字符串中首尾所有换行符、首尾所有空格
*/
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
	s1 := "a  a a"
	fmt.Println("s1原：", s1)               //  a  a a
	fmt.Println(DeletePreAndSufSpace(s1)) //a  a a
	s2 := " 和  a 和   "
	fmt.Println("s2原：", s2)               //  和  a 和
	fmt.Println(DeletePreAndSufSpace(s2)) //和  a 和
	s3 := "   "                           //都是空格时全部消除
	fmt.Println("s3原：", s3)               //
	fmt.Println(DeletePreAndSufSpace(s3)) //
}

func TestCalcStrFrequency(t *testing.T) {
	fmt.Println(CalcStrFrequency("abcd ef g h", " "))  // 3
	fmt.Println(CalcStrFrequency("abcd ef g h ", " ")) // 4
	fmt.Println(CalcStrFrequency("  ", " "))           // 2
	fmt.Println(CalcStrFrequency("abcdef", "f"))       // 1
	fmt.Println(CalcStrFrequency("abcaef", "a"))       // 2
	fmt.Println(CalcStrFrequency("abaaf", "aa"))       // 1
	fmt.Println(CalcStrFrequency("abaacaaf", "aa"))    // 2
	fmt.Println(CalcStrFrequency("abaaacaaf", "aa"))   // 2
	fmt.Println(CalcStrFrequency("abaaaacaaf", "aa"))  // 3
	fmt.Println(CalcStrFrequency("abc", "abc"))        // 1
	fmt.Println(CalcStrFrequency("ab123ac", "abc"))    // 0
	fmt.Println(CalcStrFrequency("ab123ac", "a"))      // 2
	fmt.Println(CalcStrFrequency("ab123ac", "1"))      // 1
	fmt.Println(CalcStrFrequency("ab123ac", "b1"))     // 1
	fmt.Println(CalcStrFrequency("ab123ac", "12"))     // 1
}

func TestCalcStrFrequencyWith(t *testing.T) {
	fmt.Println(CalcStrFrequencyWith("abcd ef g h", " "))  // 3
	fmt.Println(CalcStrFrequencyWith("abcd ef g h ", " ")) // 4
	fmt.Println(CalcStrFrequencyWith("  ", " "))           // 2
	fmt.Println(CalcStrFrequencyWith("abcdef", "f"))       // 1
	fmt.Println(CalcStrFrequencyWith("abcaef", "a"))       // 2
	fmt.Println(CalcStrFrequencyWith("abaaf", "aa"))       // 1
	fmt.Println(CalcStrFrequencyWith("abaacaaf", "aa"))    // 2
	fmt.Println(CalcStrFrequencyWith("abaaacaaf", "aa"))   // 2
	fmt.Println(CalcStrFrequencyWith("abaaaacaaf", "aa"))  // 3
	fmt.Println(CalcStrFrequencyWith("abc", "abc"))        // 1
	fmt.Println(CalcStrFrequencyWith("ab123ac", "abc"))    // 0
	fmt.Println(CalcStrFrequencyWith("ab123ac", "a"))      // 2
	fmt.Println(CalcStrFrequencyWith("ab123ac", "1"))      // 1
	fmt.Println(CalcStrFrequencyWith("ab123ac", "b1"))     // 1
	fmt.Println(CalcStrFrequencyWith("ab123ac", "12"))     // 1
	fmt.Println(CalcStrFrequencyWith("我爱中国", "我"))         // 1
	fmt.Println(CalcStrFrequencyWith("我爱中国我", "我"))        // 2
	fmt.Println(CalcStrFrequencyWith("ab我ab爱中国", "a"))     // 2
	fmt.Println(CalcStrFrequencyWith("ab我abc爱中国", "abc"))  // 1
	fmt.Println(CalcStrFrequencyWith("我爱你i", "我爱你"))       // 1
	fmt.Println(CalcStrFrequencyWith("我爱你i我爱你我爱你", "我爱你")) // 3
	fmt.Println(CalcStrFrequencyWith("我爱你i我爱你我爱你", "i"))   // 1
}

func TestRunee(t *testing.T) {
	s := "我爱a中国bc"
	fmt.Println("s的长度", len(s))
	r := []rune(s)
	fmt.Println("r的长度", len(r))
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r[i]))
	}
}

func TestVersionGreaterThanT(t *testing.T) {
	flag := CompareVersion("1.19.2", "1.3.1")
	fmt.Println(flag) // true
	flag1 := CompareVersion("v1.1", "v1.7")
	fmt.Println(flag1) // false
	flag2 := CompareVersion("abc.1.11", "abc.1.12")
	fmt.Println(flag2) // false
	flag4 := CompareVersion("v.1.b.6", "v.1.b.5")
	fmt.Println(flag4) // true
	flag5 := CompareVersion("v.1.b.5", "v.1.a.6")
	fmt.Println(flag5) // true
	flag6 := CompareVersion("v.1.7.105", "v.1.7.115")
	fmt.Println(flag6) // false
}

func TestSplitByManyStrWith(t *testing.T) {
	fmt.Println(SplitByManyStrWith("ab+c*de+f/gh", []string{`*`, `+`, `/`}))    // [ab c de f gh]
	fmt.Println(SplitByManyStrWith("a%%b&c&d%h+fg-h", []string{`&`, `%`, `-`})) // [a b c d h+fg h]
	fmt.Println(SplitByManyStrWith("a:b=c", []string{`=`, `:`}))                // [a b c]
}

// 去除字符串中所有多余换行符(换行的时候只保留一个换行符),首尾换行符也一并去掉
func TestTrimLineFeed(t *testing.T) {
	var x = `

x


yx


z

`
	t.Log("---------------------")
	for i := 0; i < len(x); i++ {
		t.Log(x[i])
	}
	t.Log("---------------------")

	y := TrimLineFeed(x)
	t.Log("|", y, "|")
	t.Log(x == y)
	t.Log("---------------------")
	for i := 0; i < len(y); i++ {
		t.Log(y[i])
	}
	t.Log("---------------------")
}

// 去除字符串中所有换行符
func TestTrimAllFeeds(t *testing.T) {
	var x = `

x


yx



`
	t.Log("---------------------")
	for i := 0; i < len(x); i++ {
		t.Log(x[i])
	}
	t.Log("---------------------")

	y := TrimAllFeeds(x)
	t.Log("|", y, "|")
	t.Log(x == y)
	t.Log("---------------------")
	for i := 0; i < len(y); i++ {
		t.Log(y[i])
	}
	t.Log("---------------------")
}

// 去除字符串中首尾单次换行符
func TestTrimHeadTailFeedsOnce(t *testing.T) {
	var x = `

x


yx



`
	t.Log("---------------------")
	for i := 0; i < len(x); i++ {
		t.Log(x[i])
	}
	t.Log("---------------------")

	y := TrimHeadTailFeedsOnce(x)
	t.Log("|", y, "|")
	t.Log(x == y)
	t.Log("---------------------")
	for i := 0; i < len(y); i++ {
		t.Log(y[i])
	}
	t.Log("---------------------")
}

// 去除字符串中首尾所有换行符
func TestTrimHeadTailFeeds(t *testing.T) {
	var x = `

x


yx



`
	t.Log("---------------------")
	for i := 0; i < len(x); i++ {
		t.Log(x[i])
	}
	t.Log("---------------------")

	y := TrimHeadTailFeeds(x)
	t.Log("|", y, "|")
	t.Log(x == y)
	t.Log("---------------------")
	for i := 0; i < len(y); i++ {
		t.Log(y[i])
	}
	t.Log("---------------------")
}

// 去除字符串中首尾所有换行符、首尾所有空格
func TestTrimHeadTailFeedsAndSpaces(t *testing.T) {
	var x = `
 
x


yx 


z 

`
	t.Log("---------------------")
	for i := 0; i < len(x); i++ {
		t.Log(x[i])
	}
	t.Log("---------------------")

	y := TrimHeadTailFeedsAndSpaces(x)
	t.Log("|", y, "|")
	t.Log(x == y)
	t.Log("---------------------")
	for i := 0; i < len(y); i++ {
		t.Log(y[i])
	}
	t.Log("---------------------")
}
