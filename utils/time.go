package lacia

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// ***************************日期转换start***************************
//获取当前时间前一天的时间戳
func GetPreDayUnix() int64 {
	var cstZone = time.FixedZone("CST", 8*3600)
	prevDay := time.Now().In(cstZone).AddDate(0, 0, -1)
	return prevDay.Unix()
}

// 时间戳转日期字符串，具体到天数2020-09-03
func Unix2TimeDayRoughStr(timeStamp int64) string {
	var cstZone = time.FixedZone("CST", 8*3600)
	curTm := time.Unix(timeStamp, 0)
	dayStr := curTm.In(cstZone).Format("2006-01-02")
	return dayStr
}

// 时间戳转对应日的零时时间2020-09-03 00:00:00
func Unix2TimeDayZeroStr(timeStamp int64) string {
	var cstZone = time.FixedZone("CST", 8*3600)
	curTime := time.Unix(timeStamp, 0)
	//nextTime := curTime.In(cstZone).AddDate(0, 0, 1)
	dayStr := curTime.In(cstZone).Format("2006-01-02") + " 00:00:00"
	return dayStr
}

// 时间戳转对应日的实际时间2020-09-03 11:31:59
func Unix2TimeStr(timeStamp int64) string {
	var cstZone = time.FixedZone("CST", 8*3600)
	curTime := time.Unix(timeStamp, 0)
	//nextTime := curTime.In(cstZone).AddDate(0, 0, 1)
	dayStr := curTime.In(cstZone).Format("2006-01-02 15:04:05")
	return dayStr
}

// ***************************日期获取&判断start***************************
/*
	给定年月(YYYYMM),获取该月有多少天
*/
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

// 日结无顺延,入参示例：202007，202008  或只传一个月：startT=endT=202006
func GetAllDaysByMonths(startT, endT string) []string {
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
	return GetInputTimes(startStr, endStr)
}

// 本函数用于执行后付费查询
// eg:想获取20200807、20200808、20200809三天的数据，输入20200807 20200809，则返回为[]string{20200808,20200809,20200810}
// eg:想获取20200730、20200731、20200801三天的数据，输入20200730 20200801，则返回为[]string{20200731,20200801,20200802}
// eg:想获取2020年6月、2020年7月 两月的数据，输入202007 202008，则返回为[]string{202008,202009}
func GetInputTimes(startT, endT string) (timeList []string) {
	/*
		count=8,精确到天;
		count=6,精确到月;
		count=4时,暂不支持直接按年
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
		startDayInt, _ := strconv.Atoi(startDayStr)                  // 07
		endDayInt, _ := strconv.Atoi(endDayStr)                      // 09
		dayAdd, monthAdd, dayStr := startDayInt+1, startMonthInt, "" // 如需天数顺延一天，则dayAdd = startDayInt + 1 即可
		listLen := GetDaysNumByTwoDays(yearInt, startMonthInt, endMonthInt, startDayInt, endDayInt)
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
			dayAdd++ // 务必放在判断跳月之后，否则数据不准
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
		GetInputTimes(startStr, endStr)
		return
	case 4:
		fmt.Println("暂不支持以年查询！")
		return
	}

	return
}

// 输入两个YYYYMMDD日期,返回期间的天数总长度，如传入20200730，20200801，则返回3; 如传入20200630，20200801，则返回33
func GetDaysNumByTwoDays(yearInt, startMonthInt, endMonthInt, startDayInt, endDayInt int) (days int) {
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

// 年月inputStart[:6]  月inputStart[4:6]
// 输入两个YYYYMMDD日期,返回期间的月份列表数组; 输入20200101 20200303 返回月份列表数组string{202001,202002,202003}
func GetYearMonthListByTwoDay(inputStart, inputEnd string) []string {
	year := inputStart[:4]
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

// 输入20200101 20200303 返回string{202002,202003,202004}
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
	//currStr := ""
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

// 判断当前是几月,判断dayAdd+1后是否到了下月,返回true则表示后续处理需要将月数+1，返回false时不需要跳月还是在当月
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

func GetChBufNum(num int) int {
	if num < 200 {
		return num
	}

	if num > 1000 && num < 5000 {
		return num / 5
	}

	if num > 5000 && num < 10000 {
		return num / 10
	}

	if num > 10000 && num < 30000 {
		return num / 20
	}

	return 1000
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

// 解析sdkappid_input.csv,应用维度的数据，放在common目录下，机器上执行脚本时，将文件放到与脚本同级的目录下即可。
func ParseSdkAppIdListCsvFile(fileName string) ([]int, error) {
	fileName = "./" + fileName
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(bs)))
	ss, _ := r2.ReadAll()
	//fmt.Println(ss)
	sz := len(ss)
	var resList []int
	for i := 0; i < sz; i++ {
		//fmt.Println(ss[i])    // [1400084549]
		//fmt.Println(ss[i][0]) // 1400084549 读取到的每行的数据  可以作为map的数据的值
		sdkAppID, err := strconv.Atoi(ss[i][0]) // strconv.Atoi: parsing "\ufeff1400358110": invalid syntax,和excel有关系，特殊处理
		/*  该err不能忽略。
		err != nil的情况：只有给出的csv文件不正常时才会报错，第一行的sdkAppID值有杂乱字符会解析不出来。
		*/
		if err != nil { // 需要手动处理
			fmt.Println(fmt.Sprintf("转换%v出错,err：%v", ss[i][0], err))
		}
		resList = append(resList, sdkAppID)
	}

	return resList, err
}

// 解析xx.csv文件,文件中只有一列,返回该列数据
func ParseCsvFile(fileName string) ([]int64, error) {
	//fileName = "./" + fileName
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// 读取文件数据
	read := csv.NewReader(strings.NewReader(string(bs)))
	all, _ := read.ReadAll()
	count := len(all)
	var resList []int64
	for i := 0; i < count; i++ {
		//fmt.Println(all[i])    // [123]
		//fmt.Println(all[i][0]) // 123 key的数据  可以作为map的数据的值
		id, err := strconv.Atoi(all[i][0])
		if err != nil {
			if i == 0 {
				resList = append(resList, 1400358110)
				continue
			}
			fmt.Println(fmt.Sprintf("转换%v出错,err：%v", all[i][0], err))
		}
		resList = append(resList, int64(id))
	}

	return resList, err
}
