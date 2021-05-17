package lacia

import (
	"regexp"
	"strconv"
	"strings"
)

// 去除字符串中所有指定字符(含去除所有空格)
func RemoveX(str string, x string) string {
	var res string
	for i := 0; i < len(str); i++ {
		if string(str[i]) != x {
			res = res + string(str[i])
		}
	}
	return res
}

// 删除字符串前后空格
func DeletePreAndSufSpace(str string) string {
	strList := []byte(str)
	spaceCount, count := 0, len(strList)
	for i := 0; i <= len(strList)-1; i++ {
		if strList[i] != 32 {
			break
		}
		spaceCount++
	}

	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] != 32 {
			break
		}
		spaceCount++
	}

	return string(strList[:count-spaceCount])
}

// 检查一组字符串中每个字符串的长度,只要有一个字符串的长度不为n则返回false
func CheckEveryStringLengthIsN(strList []string, n int) bool {
	for _, s := range strList {
		if len(s) != n {
			return false
		}
	}
	return true // all str length equals n, return true
}

// 计算给定src字符串中出现cond的次数   src范围:仅限于字母和数字   通用时使用CalcStrFrequencyWith
func CalcStrFrequency(src, cond string) (n int) {
	condCount, srcCount := len(cond), len(src)
	if condCount > len(src) {
		return -1
	}

	if condCount == 1 {
		for i := 0; i < srcCount; i++ {
			if string(src[i]) == cond {
				n += 1
			}
		}
		return
	}

	for i := 0; i < srcCount; {
		if i+1 >= srcCount || i+condCount > srcCount {
			break
		}

		if src[i:i+condCount] == cond {
			n += 1
			i += condCount
			continue
		}

		i += 1
	}

	return
}

// 计算给定src字符串中出现cond的次数
func CalcStrFrequencyWith(str, cond string) (n int) {
	src := []rune(str)
	condCount, srcCount := len([]rune(cond)), len(src)
	if condCount > srcCount {
		return -1
	}

	if condCount == 1 {
		for i := 0; i < srcCount; i++ {
			if string(src[i]) == cond {
				n += 1
			}
		}
		return
	}

	for i := 0; i < srcCount; {
		if i+1 >= srcCount || i+condCount > srcCount {
			break
		}

		if string(src[i:i+condCount]) == cond {
			n += 1
			i += condCount
			continue
		}

		i += 1
	}

	return
}

var (
	numberRe = regexp.MustCompile("[0-9]+")
	wordRe   = regexp.MustCompile("[a-z]+")
)

// 版本号比较
// eg:参数1 > 参数2 的版本时，返回true  反之为false
func CompareVersion(a, b string) bool {
	return thanT(a, b)
}

func stripMetadata(v string) string {
	split := strings.Split(v, "+")
	if len(split) > 1 {
		return split[0]
	}
	return v
}

var thanStr = []string{".", "-"}

func thanT(a, b string) bool {
	a = stripMetadata(a)
	b = stripMetadata(b)

	a = strings.TrimLeft(a, "v")
	b = strings.TrimLeft(b, "v")

	aSplit := SplitByManyStrWith(a, thanStr)
	bSplit := SplitByManyStrWith(b, thanStr)

	if len(bSplit) > len(aSplit) {
		return !thanT(b, a) && a != b
	}

	for i := 0; i < len(aSplit); i++ {
		if i == len(bSplit) {
			if _, err := strconv.Atoi(aSplit[i]); err == nil {
				return true
			}
			return false
		}
		aWord := wordRe.FindString(aSplit[i])
		bWord := wordRe.FindString(bSplit[i])
		if aWord != "" && bWord != "" {
			if strings.Compare(aWord, bWord) > 0 {
				return true
			}
			if strings.Compare(bWord, aWord) > 0 {
				return false
			}
		}
		aMatch := numberRe.FindString(aSplit[i])
		bMatch := numberRe.FindString(bSplit[i])
		if aMatch == "" || bMatch == "" {
			if strings.Compare(aSplit[i], bSplit[i]) > 0 {
				return true
			}
			if strings.Compare(bSplit[i], aSplit[i]) > 0 {
				return false
			}
		}
		aNum, _ := strconv.Atoi(aMatch)
		bNum, _ := strconv.Atoi(bMatch)
		if aNum > bNum {
			return true
		}
		if bNum > aNum {
			return false
		}
	}

	return false
}

// 按多个指定字符分割字符串,自定义需按切割的字符存放在flagList;用法详见单元测试TestSplitByManyStrWith
func SplitByManyStrWith(src string, flagList []string) []string {
	return strings.FieldsFunc(src, func(r rune) bool {
		res := false
		for _, s := range flagList {
			if string(r) == s {
				res = true
			}
		}
		return res
	})
}
