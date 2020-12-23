package lacia

import (
	"regexp"
	"strconv"
	"strings"
)

// 指定字符去除(含去除所有空格)
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
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	return string(strList[:count-spaceCount])
}

var (
	numberRe = regexp.MustCompile("[0-9]+")
	wordRe   = regexp.MustCompile("[a-z]+")
)

// 版本号比较
// usages:参数1 > 参数2 的版本时，返回true  反之为false
func VersionGreaterThanT(a, b string) bool {
	return GreaterThan(a, b)
}

func stripMetadata(v string) string {
	split := strings.Split(v, "+")
	if len(split) > 1 {
		return split[0]
	}
	return v
}

func periodDashSplit(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case '.', '-':
			return true
		}
		return false
	})
}

func GreaterThan(a, b string) bool {
	a = stripMetadata(a)
	b = stripMetadata(b)

	a = strings.TrimLeft(a, "v")
	b = strings.TrimLeft(b, "v")

	aSplit := periodDashSplit(a)
	bSplit := periodDashSplit(b)

	if len(bSplit) > len(aSplit) {
		return !GreaterThan(b, a) && a != b
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

// 按多个指定字符分割字符串,需指定的字符按常用默认的来判断
func SplitByManyStr(src string) []string {
	return strings.FieldsFunc(src, func(r rune) bool {
		return r == ':' || r == '=' || r == '+' || r == '-' || r == '*' || r == '$' ||
			r == '&' || r == '@' || r == '!' || r == '！' || r == '%' || r == '^' || r == '_' ||
			r == '/' || r == '\\' || r == '{' || r == '}'
	})
}
