package lacia

import (
	"regexp"
	"strconv"
	"strings"
)

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
