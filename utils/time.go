package lacia

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
********************功能概览********************
CompareTwoStrTime(t1, t2 string)      比较两个时间
GetPreferTimeByTimeStrAndDifference(timeStr string, y, m, d int)    获取指定日期之前/之后的日期
GetTimeByStr(timeStr string)          日期字符串转日期
GetTsByTimeStr(timeStr string)        日期字符串转时间戳
GetTimeByTs(ts int64)                 时间戳转日期
GetPreDayTimeStampByNow()             获取当前时间前一天的秒级时间戳
GetTimeStrOfDayByTs(ts int64)         时间戳转对应日期字符串
GetTimeStrOfDayZeroByTs(ts int64)     时间戳转对应日期零时字符串
GetTimeStrOfDayTimeByTs(ts int64)     时间戳转对应日期的实际时间
GetTimeStrOfDifferenceDayByTs(ts int64,n int)   时间戳转n日前/后的实际日时间
GetTimeStrOfDifferenceDayZeroByTs(ts int64,n int)  时间戳转n日前/后的零时时间
GetTimeStrOfDifferenceDayTimeByTs(ts int64,n int)  时间戳转n日前/后的实际时间
GetDaysInMonth(yearInt, month int)   给定年月值获取该月有多少天
GetNMonthAgoOrAfterThisDayByN(thisT *time.Time, thisTStr string, n int)   给定一个日期或日期字符串，返回n月前/后的当天日期
GetAllDaysStrByMonths(startMonthT, endMonthT string, ifDelayOneDay bool)  给定两个月份，获取期间的所有天粒度的字符串列表（支持获取整体延后一天的时间列表）
GetTimeListByTwoDayStr(startT, endT string, ifDelayOneDay bool)   给定一个时间区间，根据传入的时间精确到月/天，用于获取该区间内所有的天粒度/月粒度的时间字符串列表
CheckTimeStrListIsCorrect(strList []string)    检查一批日期字符串是否是正常日期
GetTimeListByTwoDayStr2(startT, endT string, ifDelayOneDay bool)   旧版。建议使用替代函数GetTimeListByTwoDayStr或GetAllDaysStrByMonths，功能同上述两函数
GetDaysNumByTwoDays(yearInt, startMonthInt, startDayInt, endMonthInt, endDayInt int)       输入两个YYYYMMDD日期,返回期间的天数总长度
GetYearMonthListByTwoDay(inputStart, inputEnd string)   给定两个YYYYMM或YYYYMMDD日期,返回期间的月份列表数组
GetYearMonthListByTwoDayDelayOneDay(inputStart, inputEnd string)   给定两个YYYYMM或YYYYMMDD日期,返回期间的月份列表数组(整体顺延一月)
IsJumpMonth(yearInt int, monthInt int, dayAdd int)      以给定的年月日判断dayAdd+1后是否到了下月
IsNum(s string)   判断一个字符串是否是纯数字
ConvertTime2Str(inputTime string, isStart bool)    时间转换方法：转换时间20200504为"2020-05-04 00:00:00"格式
*/

var CstZone = time.FixedZone("CST", 8*3600)
var TimeTemplate1 = "2006-01-02 15:04:05"
var TimeTemplate2 = "2006/01/02 15:04:05"
var TimeTemplate3 = "2006-01-02"
var TimeTemplate5 = "20060102 150405"
var TimeTemplate6 = "20060102"

// ***************************日期比较start***************************
/*
传入两个字符串类型的日期,格式可以是YYYYMMDD、YYYYMMDD HH:mm:ss、YYYY-MM-DD、YYYY/MM/DD之一
不论是什么格式，t1和t2两者格式必须相同
return:
	t1=t2  		return 0
	t1>t2  		return 1
	t1<t2  		return 2
	not compare return -1
*/
func CompareTwoStrTime(t1, t2 string) int {
	if len(t1) != len(t2) {
		return -1
	}

	t1, t2 = RemoveX(t1, ` `), RemoveX(t2, ` `)
	count := len(t1)
	for i := 0; i < count; i++ {
		if t1[i] == t2[i] {
			continue
		}
		if !IsNum(string(t1[i])) || !IsNum(string(t2[i])) {
			continue
		}
		n1, err := strconv.Atoi(string(t1[i]))
		if err != nil {
			return -1
		}
		n2, err := strconv.Atoi(string(t2[i]))
		if err != nil {
			return -1
		}

		if n1 > n2 {
			return 1
		}
		if n2 > n1 {
			return 2
		}
	}

	return 0
}

// ***************************日期转换-->当日start***************************
// 获取指定日期之前/之后的日期  timeStr格式:YYYYMMDD
// y,m,d可为正数或负数  如d=1则获取timeStr的一天之后的日期，如d=-1则获取timeStr的前一天的日期
func GetPreferTimeByTimeStrAndDifference(timeStr string, y, m, d int) (time.Time, error) {
	thisT, err := GetTimeByStr(timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return thisT.AddDate(y, m, d), nil
}

// 日期字符串转日期  输入格式支持YYYYMMDD、YYYY-MM-DD之一,如20201001、2020-10-01
func GetTimeByStr(timeStr string) (time.Time, error) {
	timeStr = DeletePreAndSufSpace(timeStr)
	if !strings.Contains(timeStr, "-") {
		if len(timeStr) != 8 || !IsNum(timeStr) {
			return time.Time{}, errors.New("input timeStr format error")
		}
		timeStr = timeStr[:4] + "-" + timeStr[4:6] + "-" + timeStr[6:8]
	} else if len(timeStr) != 10 || len(strings.Split(timeStr, "-")[0]) != 4 {
		return time.Time{}, errors.New("input timeStr format error")
	}
	thisT, err := time.Parse(TimeTemplate3, timeStr)
	if err != nil {
		ts, err := GetTsByTimeStr(timeStr, TimeTemplate3)
		if err != nil {
			return time.Time{}, err
		}
		thisT = GetTimeByTs(ts)
	}
	return thisT, nil

}

// 时间戳转日期
func GetTimeByTs(ts int64) time.Time {
	return time.Unix(ts, 0)
}

/*
日期字符串转时间戳
timeTemplate:日期模板，可选TimeTemplate1, TimeTemplate2, TimeTemplate3, TimeTemplate5, TimeTemplate6之一，对应如下：
	var TimeTemplate1 = "2006-01-02 15:04:05"
	var TimeTemplate2 = "2006/01/02 15:04:05"
	var TimeTemplate3 = "2006-01-02"
	var TimeTemplate5 = "20060102 150405"
	var TimeTemplate6 = "20060102"
*/
func GetTsByTimeStr(timeStr, timeTemplate string) (int64, error) {
	if ExistsInListString([]string{TimeTemplate1, TimeTemplate2, TimeTemplate3,
		TimeTemplate5, TimeTemplate6}, timeTemplate, true)[0] == -1 {
		return -1, errors.New("error layout,please check")
	}
	thisT, err := time.ParseInLocation(timeTemplate, timeStr, time.Local)
	if err != nil {
		return -1, err
	}
	return thisT.Unix(), nil
}

// 获取当前时间前一天的秒级时间戳
func GetPreDayTimeStampByNow() int64 {
	return time.Now().In(CstZone).AddDate(0, 0, -1).Unix()
}

// 时间戳转对应日期字符串 返回格式2020-09-03
func GetTimeStrOfDayByTs(ts int64) string {
	curTm := time.Unix(ts, 0)
	return curTm.In(CstZone).Format(TimeTemplate3)
}

// 时间戳转对应日期零时字符串 返回格式2020-09-03 00:00:00
func GetTimeStrOfDayZeroByTs(ts int64) string {
	curTm := time.Unix(ts, 0)
	return curTm.In(CstZone).Format(TimeTemplate3) + " 00:00:00"
}

// 时间戳转对应日期的实际时间 返回格式2020-09-03 11:31:59
func GetTimeStrOfDayTimeByTs(ts int64) string {
	curTm := time.Unix(ts, 0)
	return curTm.In(CstZone).Format(TimeTemplate1)
}

// ***************************日期转换-->前一日start***************************
// 时间戳转n日前/后的实际时间 返回格式2020-09-03
func GetTimeStrOfDifferenceDayByTs(ts int64, n int) string {
	curTm := time.Unix(ts, 0)
	beforeDayTime := curTm.In(CstZone).AddDate(0, 0, n)
	return beforeDayTime.In(CstZone).Format(TimeTemplate3)
}

// 时间戳转n日前/后零时时间 返回格式2020-09-03 00:00:00
func GetTimeStrOfDifferenceDayZeroByTs(ts int64, n int) string {
	curTm := time.Unix(ts, 0)
	beforeDayTime := curTm.In(CstZone).AddDate(0, 0, n)
	return beforeDayTime.In(CstZone).Format("2006-01-02") + " 00:00:00"
}

// 时间戳转n日前/后的实际时间 返回格式2020-09-03 11:31:59
func GetTimeStrOfDifferenceDayTimeByTs(ts int64, n int) string {
	curTm := time.Unix(ts, 0)
	beforeDayTime := curTm.In(CstZone).AddDate(0, 0, n)
	return beforeDayTime.In(CstZone).Format(TimeTemplate1)
}

// ***************************日期获取&计算start***************************
// 判断指定日期A离B是否走过了指定距离  传入时间戳
func CalcAreaByTimeStr(tsA, tsB, step int64) bool {
	if tsB-tsA >= step {
		return true
	}

	return false
}

//给定年月值,获取该月有多少天
func GetDaysInMonth(yearInt, month int) (days int) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
		return
	case 4, 6, 9, 11:
		days = 30
		return
	case 2:
		if (yearInt%4 == 0 && yearInt%100 != 0) || yearInt%400 == 0 {
			days = 29
			return
		}
		days = 28
		return
	}

	return 31
}

// 给定一个日期或日期字符串，返回n月前/后的当天日期
// 传入日期或日期字符串两者之一即可。日期字符串格式：YYYYMMDD
// 应用:月同比
func GetNMonthAgoOrAfterThisDayByN(thisT *time.Time, thisTStr string, n int) (resT time.Time, resTStr string, err error) {
	if thisTStr == "" && thisT == nil {
		err = errors.New("params format error")
		return
	}

	if n == 0 {
		return
	}
	thisYearInt, thisMonthInt, thisDayInt := 0, 0, 0
	resYearStr, resMonthStr, resDayStr := "", "", ""
	if thisT != nil {
		fmt.Println("打印：", thisT.YearDay())
		thisMonth := thisT.Month()
		thisYearInt, thisDayInt = thisT.Year(), thisT.Day()
		thisMonthInt = int(thisMonth)
	} else if thisTStr != "" {
		if len(thisTStr) != 8 {
			err = errors.New("params str format error")
			return
		}
		thisYearStr, thisMonthStr, thisDayStr := thisTStr[:4], thisTStr[4:6], thisTStr[6:]
		thisYearInt, err = strconv.Atoi(thisYearStr)
		if err != nil {
			return
		}
		thisMonthInt, err = strconv.Atoi(thisMonthStr)
		if err != nil {
			return
		}
		thisDayInt, err = strconv.Atoi(thisDayStr)
		if err != nil {
			return
		}
	}

	newYearInt, newMonthInt := thisYearInt, thisMonthInt
	switch n > 0 {
	case true:
		c := thisMonthInt + n
		if n >= 12 {
			if c/12 > 1 {
				newYearInt = thisYearInt + (c)/12
			} else {
				newYearInt = thisYearInt + n/12
			}
			n = n % 12
		}

		if c > 12 {
			if c/12 > 1 && c%12 != 0 {
				c = c - (c/12)*12
			} else {
				c = c - 12
				newYearInt = thisYearInt + 1
			}

			newMonthInt = c
			resMonthStr = strconv.Itoa(newMonthInt)
			resYearStr = strconv.Itoa(newYearInt)
		} else {
			newMonthInt = thisMonthInt + n
			resYearStr, resMonthStr = strconv.Itoa(newYearInt), strconv.Itoa(newMonthInt)
		}

		if newMonthInt < 10 {
			resMonthStr = "0" + resMonthStr
		}

		newMonthDays := GetDaysInMonth(newYearInt, newMonthInt)
		resDayStr = strconv.Itoa(newMonthDays)
		if thisDayInt < 10 {
			resDayStr = "0" + strconv.Itoa(thisDayInt)
		} else if newMonthDays < thisDayInt {
			resDayStr = strconv.Itoa(newMonthDays) // 取前一天
		} else {
			resDayStr = strconv.Itoa(thisDayInt)
		}
		resTStr = resYearStr + resMonthStr + resDayStr
	case false:
		c := thisMonthInt + n // -15
		if n <= -12 {
			newYearInt = thisYearInt - (Abs(n)+thisMonthInt)/12
		}

		if c <= 0 {
			if c/12 <= -1 && c%12 != 0 {
				c = 12 + (c - (c/12)*12)
			} else {
				c = 12 + c
				newYearInt = thisYearInt - 1
			}
			newMonthInt = Abs(c)
			resMonthStr = strconv.Itoa(newMonthInt)
			//newYearInt = thisYearInt - 1
			resYearStr = strconv.Itoa(newYearInt)
		} else {
			newMonthInt = thisMonthInt + n
			resYearStr, resMonthStr = strconv.Itoa(newYearInt), strconv.Itoa(newMonthInt)
		}

		if newMonthInt < 10 {
			resMonthStr = "0" + resMonthStr
		}

		newMonthDays := GetDaysInMonth(newYearInt, newMonthInt)
		resDayStr = strconv.Itoa(newMonthDays)
		if thisDayInt < 10 {
			resDayStr = "0" + strconv.Itoa(thisDayInt)
		} else if newMonthDays < thisDayInt {
			resDayStr = strconv.Itoa(newMonthDays) // 取前一天
		} else {
			resDayStr = strconv.Itoa(thisDayInt)
		}
		resTStr = resYearStr + resMonthStr + resDayStr
	}
	resT, err = GetTimeByStr(resTStr)
	return
}

// 给定两个月份，获取期间的所有天粒度的字符串列表
// ifDelayOneDay为是否取延迟一天的标志位，ifDelayOneDay=true，则获取的天粒度数据全部往后推迟一天
// 入参示例：202007 202008;或只传一个月：startT=endT=202006;返回示例请查看对应单元测试TestGetAllDaysDelayOneDayByMonths
func GetAllDaysStrByMonths(startMonthT, endMonthT string, ifDelayOneDay bool) ([]string, []string) {
	preStr := startMonthT[:4]         // 2020
	startMonthStr := startMonthT[4:6] // 06
	endMonthStr := endMonthT[4:6]     // 07
	yearInt, _ := strconv.Atoi(preStr)
	startMonthInt, _ := strconv.Atoi(startMonthStr)
	endMonthInt, _ := strconv.Atoi(endMonthStr)
	startStr, endStr := "", ""
	currentMonthDays := GetDaysInMonth(yearInt, startMonthInt)
	if endMonthInt == startMonthInt {
		startStr, endStr = preStr+startMonthStr+"01", preStr+startMonthStr+strconv.Itoa(currentMonthDays)
	} else {
		currentMonthDays = GetDaysInMonth(yearInt, endMonthInt)
		startStr, endStr = preStr+startMonthStr+"01", preStr+endMonthStr+strconv.Itoa(currentMonthDays)
	}

	timeList, timeListNoUnderline := GetTimeListByTwoDayStr(startStr, endStr, ifDelayOneDay)
	return timeList, timeListNoUnderline
}

// 给定一个时间区间，根据传入的时间精确到月/天，用于获取该区间内所有的天粒度/月粒度的时间字符串列表
// ifDelayOneDay为是否取延迟一天的标志位，ifDelayOneDay=true，则获取的天粒度数据全部往后推迟一天
// eg:想获取20200807、20200808、20200809三天的天粒度字符串，输入20200807 20200809，则返回为[]string{20200808,20200809,20200810}
// eg:想获取20200730、20200731、20200801三天的天粒度字符串，输入20200730 20200801，则返回为[]string{20200731,20200801,20200802}
// eg:想获取2020年6月、2020年7月 两月的月粒度字符串，输入202007 202008，则返回为[]string{202008,202009}
func GetTimeListByTwoDayStr(startT, endT string, ifDelayOneDay bool) (timeList, timeListNoUnderline []string) {
	if startT == endT { // not allow one day
		return
	}

	if !CheckEveryStringLengthIsN([]string{startT, endT}, 8) {
		return
	}

	// check
	if ok, _ := CheckTimeStrListIsCorrect([]string{startT, endT}); !ok {
		return
	}

	timeList = append(timeList, startT[:4]+"-"+startT[4:6]+"-"+startT[6:8])
	timeListNoUnderline = append(timeListNoUnderline, startT)
	nextT, stop := startT, 1
	for {
		if CompareTwoStrTime(nextT, endT) == 0 {
			stop--
		}
		currentT, err := GetPreferTimeByTimeStrAndDifference(nextT, 0, 0, 1)
		if err != nil {
			fmt.Println("GetPreferTimeByTimeStrAndDifference error:", err.Error())
			timeListNoUnderline2 := GetTimeListByTwoDayStr2(nextT, endT, ifDelayOneDay)
			timeListNoUnderline = append(timeListNoUnderline, timeListNoUnderline2...)
			return
		}
		nextT = GetTimeStrOfDayByTs(currentT.Unix())
		timeList = append(timeList, nextT)
		nextT = RemoveX(nextT, "-")
		timeListNoUnderline = append(timeListNoUnderline, nextT)
		if stop == 0 {
			break
		}
	}

	count, countNoUnderline := len(timeList), len(timeListNoUnderline)
	if !ifDelayOneDay {
		timeList, timeListNoUnderline = timeList[:count-1], timeListNoUnderline[:countNoUnderline-1]
		return
	}

	timeList, timeListNoUnderline = timeList[1:], timeListNoUnderline[1:]
	return
}

// 检查一批日期字符串是否是正常日期
// eg: 20201001 return true,  20201131 return false, 20201331 return false
func CheckTimeStrListIsCorrect(strList []string) (bool, error) {
	for _, thisTStr := range strList {
		if len(thisTStr) != 8 {
			return false, errors.New(fmt.Sprintf("time str [%v] length error, please check", thisTStr))
		}
		thisYearStr, thisMonthStr, thisDayStr := thisTStr[:4], thisTStr[4:6], thisTStr[6:]
		thisYearInt, err := strconv.Atoi(thisYearStr)
		if err != nil {
			return false, errors.New(fmt.Sprintf("time year_str [%v] format error, please check", thisYearStr))
		}

		if thisYearInt < 1970 {
			return false, errors.New(fmt.Sprintf("time year_str [%v] area error, please check", thisYearStr))
		}

		thisMonthInt, err := strconv.Atoi(thisMonthStr)
		if err != nil {
			return false, errors.New(fmt.Sprintf("time month_str [%v] format error, please check", thisMonthStr))
		}

		if thisMonthInt > 12 || thisMonthInt <= 0 {
			return false, errors.New(fmt.Sprintf("time month_str [%v] area error, please check", thisMonthStr))
		}

		thisDayInt, err := strconv.Atoi(thisDayStr)
		if err != nil {
			return false, errors.New(fmt.Sprintf("time day_str [%v] format error, please check", thisDayStr))
		}

		days := GetDaysInMonth(thisYearInt, thisMonthInt)
		if thisDayInt > days || thisDayInt <= 0 {
			return false, errors.New(fmt.Sprintf("time day_str [%v] area error, please check", thisDayStr))
		}
	}

	return true, nil
}

func GetTimeListByTwoDayStr2(startT, endT string, ifDelayOneDay bool) (timeList []string) {
	/*
		count=8,精确到天;
		count=6,精确到月;
	*/
	count := len(startT)
	switch count {
	case 8: // 日
		preStr := startT[:4]         // 2020
		startMonthStr := startT[4:6] // 08
		endMonthStr := endT[4:6]     // 08
		startDayStr := startT[6:]    // 07
		endDayStr := endT[6:]        // 09

		yearInt, _ := strconv.Atoi(preStr)
		startMonthInt, _ := strconv.Atoi(startMonthStr)
		endMonthInt, _ := strconv.Atoi(endMonthStr)
		startDayInt, _ := strconv.Atoi(startDayStr)                // 07
		endDayInt, _ := strconv.Atoi(endDayStr)                    // 09
		dayAdd, monthAdd, dayStr := startDayInt, startMonthInt, "" // 如需天数顺延一天，则dayAdd = startDayInt + 1 即可
		if ifDelayOneDay {
			dayAdd = dayAdd + 1
		}
		listLen := GetDaysNumByTwoDays(yearInt, startMonthInt, startDayInt, endMonthInt, endDayInt)
		for i := 0; i < listLen; {
			monthStr := strconv.Itoa(monthAdd)
			if dayAdd < 10 {
				dayStr = "0" + strconv.Itoa(dayAdd)
			} else {
				dayStr = strconv.Itoa(dayAdd)
			}
			timeList = append(timeList, preStr+monthStr+dayStr)
			if IsJumpMonth(yearInt, monthAdd, dayAdd) {
				monthAdd++
				dayAdd = 0
			}
			dayAdd++
			if monthAdd == endMonthInt && dayAdd == endDayInt+2 {
				break
			}
			i++
		}

		return

	case 6: // 月,202006、202007
		preStr := startT[:4]         // 2020
		startMonthStr := startT[4:6] // 06
		endMonthStr := endT[4:6]     // 07
		yearInt, _ := strconv.Atoi(preStr)
		startMonthInt, _ := strconv.Atoi(startMonthStr)
		endMonthInt, _ := strconv.Atoi(endMonthStr)
		startStr, endStr := "", ""
		currentMonthDays := GetDaysInMonth(yearInt, startMonthInt)
		if endMonthInt == startMonthInt {
			startStr, endStr = preStr+startMonthStr+"01", preStr+startMonthStr+strconv.Itoa(currentMonthDays)
		} else {
			currentMonthDays = GetDaysInMonth(yearInt, endMonthInt)
			startStr, endStr = preStr+startMonthStr+"01", preStr+endMonthStr+strconv.Itoa(currentMonthDays)
		}
		fmt.Println("转换后的8位str：", startStr, endStr)
		GetTimeListByTwoDayStr2(startStr, endStr, ifDelayOneDay)
		return
	case 4:
		fmt.Println("暂不支持以年查询！")
		return
	}

	return
}

// 输入两个YYYYMMDD日期,返回期间的天数总长度，如传入20200730，20200801，则返回3; 如传入20200630，20200801，则返回33
func GetDaysNumByTwoDays(yearInt, startMonthInt, startDayInt, endMonthInt, endDayInt int) (days int) {
	if startMonthInt == endMonthInt { // 当月
		days = endDayInt - startDayInt + 1
		return
	}

	months := endMonthInt - startMonthInt // 2
	days = GetDaysInMonth(yearInt, startMonthInt) - startDayInt + 1
	currentMonth := startMonthInt // 6
	for i := 0; i < months; i++ { // 跳月
		currentMonth++
		if currentMonth == endMonthInt {
			days = days + endDayInt
			break
		}
		days = days + GetDaysInMonth(yearInt, currentMonth) // 32
	}

	return
}

// 给定两个YYYYMM或YYYYMMDD日期,返回期间的月份列表数组(无顺延);
// 输入20200101 20200303 返回月份列表数组string{202001,202002,202003}
func GetYearMonthListByTwoDay(inputStart, inputEnd string) []string {
	year := inputStart[:4] // 年月inputStart[:6]  月inputStart[4:6]
	yearMon := inputStart[:6]
	startMon := inputStart[4:6]
	endMon := inputEnd[4:6]
	startMonthInt, _ := strconv.Atoi(startMon)
	endMonthInt, _ := strconv.Atoi(endMon)

	if startMonthInt == endMonthInt {
		return []string{yearMon}
	}

	res := []string{yearMon}
	currMon := startMonthInt
	num := endMonthInt - startMonthInt
	for i := 0; i < num; i++ {
		currMon = currMon + 1
		currMonthStr := strconv.Itoa(currMon)
		if currMon < 10 {
			currMonthStr = "0" + currMonthStr
		}
		res = append(res, year+currMonthStr)
	}

	return res
}

// 给定两个YYYYMM或YYYYMMDD日期,返回期间的月份列表数组(整体顺延一月);
// 输入20200101 20200303 返回string{202002,202003,202004}
// 输入20200101 20200303 返回string{202002,202003,202004}
// 输入202008 202008 返回string{202009}
// 输入20200705 20200708 返回string{202008}
// 输入20200705 20200808 返回string{202008,202009}
func GetYearMonthListByTwoDayDelayOneDay(inputStart, inputEnd string) []string {
	year := inputStart[:4] // 2020
	yearMon := inputStart[:6]
	startMon := inputStart[4:6]
	endMon := inputEnd[4:6]
	yearMonthInt, _ := strconv.Atoi(yearMon)
	startMonthInt, _ := strconv.Atoi(startMon)
	endMonthInt, _ := strconv.Atoi(endMon)

	yearMonthInt = yearMonthInt + 1
	startMonthInt = startMonthInt + 1
	endMonthInt = endMonthInt + 1
	newYearMonthStr := strconv.Itoa(yearMonthInt)
	if startMonthInt == endMonthInt { // 只需要一个月的数据
		return []string{newYearMonthStr}
	}

	res := []string{newYearMonthStr}
	currMon := startMonthInt
	num := endMonthInt - startMonthInt
	for i := 0; i < num; i++ {
		currMon = currMon + 1
		currMonthStr := strconv.Itoa(currMon)
		if currMon < 10 {
			currMonthStr = "0" + currMonthStr
		}
		res = append(res, year+currMonthStr)
	}

	return res
}

// 以给定的年月日判断dayAdd+1后是否到了下月,返回true则表示后续处理需要将月数+1，返回false时不需要跳月还是在当月
func IsJumpMonth(yearInt int, monthInt int, dayAdd int) bool {
	dayTotal := GetDaysInMonth(yearInt, monthInt)
	if dayAdd+1 > dayTotal { // 如出现32 > 31,返回false
		return true
	}
	return false
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//时间转换方法：转换时间20200504为"2020-05-04 00:00:00"格式
func ConvertTime2Str(inputTime string, isStart bool) (resTime string, err error) {
	if len(inputTime) != 6 && len(inputTime) != 8 {
		fmt.Println("时间参数长度不合法！")
		err = errors.New("时间参数长度不合法！")
		return
	}

	if len(inputTime) == 8 {
		resTime = inputTime[:4] + "-" + inputTime[4:6] + "-" + inputTime[6:8] + " 00:00:00"
		return
	}

	if len(inputTime) == 6 {
		if isStart {
			resTime = inputTime[:4] + "-" + inputTime[4:6] + "-01 00:00:00" // 以1号开始
		} else {
			yearStr := inputTime[:4]   // 2020
			monthStr := inputTime[4:6] // 06
			yearInt, _ := strconv.Atoi(yearStr)
			monthInt, _ := strconv.Atoi(monthStr)
			days := GetDaysInMonth(yearInt, monthInt)
			resTime = inputTime[:4] + "-" + inputTime[4:6] + "-" + strconv.Itoa(days) + " 00:00:00"
		}
	}

	return
}
