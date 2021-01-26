package lacia

import (
	"fmt"
	"testing"
)

func TestCheckTimeStrListIsCorrect(t *testing.T) {
	r, err := CheckTimeStrListIsCorrect([]string{"20201001", "20210110"})
	fmt.Println(r, err) // true <nil>
	r, err = CheckTimeStrListIsCorrect([]string{"20201001", "20200931"})
	fmt.Println(r, err) // false time day_str [31] area error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"20201000", "20210110"})
	fmt.Println(r, err) // false time day_str [00] area error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"202010001"})
	fmt.Println(r, err) // false time str [202010001] length error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"20201a01", "20210110"})
	fmt.Println(r, err) // false time month_str [1a] format error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"202a1001", "20210110"})
	fmt.Println(r, err) // false time year_str [202a] format error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"aaaabbbb"})
	fmt.Println(r, err) // false time year_str [aaaa] format error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"20202020"})
	fmt.Println(r, err) // false time month_str [20] area error, please check
	r, err = CheckTimeStrListIsCorrect([]string{"19691009"})
	fmt.Println(r, err) // false time year_str [1969] area error, please check
}

func TestGetNMonthAgoOrAfterThisDayByN1(t *testing.T) {
	// 指定月数之前
	timeT, timeStr, err := GetNMonthAgoOrAfterThisDayByN(nil, "20201231", -1)
	fmt.Println("20201231的1月前：", timeT, err, timeStr) // 2020-11-30 00:00:00 +0000 UTC <nil> 20201130

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201130", -1)
	fmt.Println("20201130的1月前：", timeT, err, timeStr) // 2020-10-30 00:00:00 +0000 UTC <nil> 20201030

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201130", -9)
	fmt.Println("20201130的9月前：", timeT, err, timeStr) // 20201130的9月前： 2020-02-29 00:00:00 +0000 UTC <nil> 20200229

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20210101", -1)
	fmt.Println("20210101的1月前：", timeT, err, timeStr) // 20210101的1月前： 2020-12-01 00:00:00 +0000UTC <nil> 20201201

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20210201", -1)
	fmt.Println("20210201的1月前：", timeT, err, timeStr) // 20210201的1月前： 2021-01-01 00:00:00 +0000UTC <nil> 20210101

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20210202", -1)
	fmt.Println("20210202的1月前：", timeT, err, timeStr) // 20210202的1月前： 2021-01-02 00:00:00 +0000UTC <nil> 20210102

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20210131", -2)
	fmt.Println("20210131的2月前：", timeT, err, timeStr) // 20210131的2月前： 2020-11-30 00:00:00 +0000 UTC <nil> 20201130

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20210228", -12)
	fmt.Println("20210228的12月前：", timeT, err, timeStr) // 20210228的12月前： 2020-02-28 00:00:00 +0000UTC <nil> 20200228

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200229", -13)
	fmt.Println("20200229的13月前：", timeT, err, timeStr) // 20200229的13月前： 2019-01-29 00:00:00 +0000UTC <nil> 20190129

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200229", -23)
	fmt.Println("20200229的23月前：", timeT, err, timeStr) // 20200229的23月前： 2018-03-29 00:00:00 +0000 UTC <nil> 20180329

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200229", -25)
	fmt.Println("20200229的25月前：", timeT, err, timeStr) // 20200229的25月前： 2018-01-29 00:00:00 +0000 UTC <nil> 20180129

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201129", -15)
	fmt.Println("20201129的15月前：", timeT, err, timeStr) // 20201129的15月前： 2019-08-29 00:00:00 +0000 UTC <nil> 20190829

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20211231", -13)
	fmt.Println("20211231的13月前：", timeT, err, timeStr) // 20211231的13月前： 2020-11-30 00:00:00 +0000 UTC <nil> 20201130

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20211231", -26)
	fmt.Println("20211231的26月前：", timeT, err, timeStr) // 20211231的26月前： 2019-10-31 00:00:00 +0000 UTC <nil> 20191031

}

func TestGetNMonthAgoOrAfterThisDayByN2(t *testing.T) {

	// 指定月数之后
	timeT, timeStr, err := GetNMonthAgoOrAfterThisDayByN(nil, "20200731", 2)
	fmt.Println("20200731的2月后：", timeT, err, timeStr) // 20200731的2月后： 2020-09-30 00:00:00 +0000 UTC <nil> 20200930

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201231", 1)
	fmt.Println("20201231的1月后：", timeT, err, timeStr) // 20201231的1月后： 2021-01-31 00:00:00 +0000 UTC <nil> 20210131

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200831", 1)
	fmt.Println("20200831的1月后：", timeT, err, timeStr) // 20200831的1月后： 2020-09-30 00:00:00 +0000 UTC <nil> 20200930

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200531", 9)
	fmt.Println("20200531的9月后：", timeT, err, timeStr) // 20200531的9月后： 2021-02-28 00:00:00 +0000 UTC <nil> 20210228

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201130", 3)
	fmt.Println("20201130的3月后：", timeT, err, timeStr) // 20201130的3月后： 2021-02-28 00:00:00 +0000 UTC <nil> 20210228

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201130", 15)
	fmt.Println("20201130的15月后：", timeT, err, timeStr) // 20201130的15月后： 2022-02-28 00:00:00 +0000UTC <nil> 20220228

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201231", 11)
	fmt.Println("20201231的11月后：", timeT, err, timeStr) // 20201231的11月后： 2021-11-30 00:00:00 +0000UTC <nil> 20211130

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201231", 12)
	fmt.Println("20201231的12月后：", timeT, err, timeStr) // 20201231的12月后： 2021-12-31 00:00:00 +0000 UTC <nil> 20211231

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200201", 11)
	fmt.Println("20200201的11月后：", timeT, err, timeStr) // 20200201的11月后： 2021-01-01 00:00:00 +0000UTC <nil> 20210101

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200201", 12)
	fmt.Println("20200201的12月后：", timeT, err, timeStr) // 20200201的12月后： 2021-02-01 00:00:00 +0000UTC <nil> 20210201

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20200201", 13)
	fmt.Println("20200201的13月后：", timeT, err, timeStr) // 20200201的13月后： 2021-03-01 00:00:00 +0000UTC <nil> 20210301

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201231", 13)
	fmt.Println("20201231的13月后：", timeT, err, timeStr) // 20201231的13月后： 2022-01-31 00:00:00 +0000UTC <nil> 20220131

	timeT, timeStr, err = GetNMonthAgoOrAfterThisDayByN(nil, "20201231", 23)
	fmt.Println("20201231的23月后：", timeT, err, timeStr) // 20201231的23月后： 2022-11-30 00:00:00 +0000UTC <nil> 20221130
}

func TestGetPreferTimeByTimeStrAndDifference(t *testing.T) {
	// 得到20201001的一天后的时间
	t1, err := GetPreferTimeByTimeStrAndDifference("20201001", 0, 0, 1)
	fmt.Println("20201001的下一天", t1, err, t1.String(), t1.Day(), t1.Unix()) // 2020-10-02 00:00:00 +0000 UTC <nil> 2020-10-02 00:00:00 +0000 UTC 2 1601596800
	t1, err = GetPreferTimeByTimeStrAndDifference("2020-10-01", 0, 0, 1)
	fmt.Println("2020-10-01的下一天", t1, err, t1.String(), t1.Day(), t1.Unix())
	// 得到20201231的一天后的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20201231", 0, 0, 1)
	fmt.Println("20201231的下一天", t1, err, t1.String(), t1.Day(), t1.Unix())
	t1, err = GetPreferTimeByTimeStrAndDifference("2020-12-31", 0, 0, 1)
	fmt.Println("2020-12-31的下一天", t1, err, t1.String(), t1.Day(), t1.Unix())

	// 得到20210101的一天前的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20210101", 0, 0, -1)
	fmt.Println("20210101的前一天", t1, err, t1.String(), t1.Day(), t1.Unix())
	t1, err = GetPreferTimeByTimeStrAndDifference("2021-01-01", 0, 0, -1)
	fmt.Println("2021-01-01的前一天", t1, err, t1.String(), t1.Day(), t1.Unix())

	// 得到20210102的30天前的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20210102", 0, 0, 30)
	fmt.Println("20210102的30天前的时间", t1, err, t1.String(), t1.Day(), t1.Unix())
	t1, err = GetPreferTimeByTimeStrAndDifference("2021-01-02", 0, 0, 30)
	fmt.Println("2021-01-02的30天前的时间", t1, err, t1.String(), t1.Day(), t1.Unix())

	// 得到20210102的1月前的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20210102", 0, -1, 0)
	fmt.Println("20210102的1月前的时间：", t1, err, t1.String(), t1.Day(), t1.Unix())
	t1, err = GetPreferTimeByTimeStrAndDifference("2021-01-02", 0, -1, 0)
	fmt.Println("2021-01-02的1月前的时间：", t1, err, t1.String(), t1.Day(), t1.Unix())

	// 得到20201231的1月前的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20201231", 0, -1, 0)
	fmt.Println("20201231的1月前的时间：", t1, err, t1.String(), t1.Day(), t1.Unix())
	t1, err = GetPreferTimeByTimeStrAndDifference("2020-12-31", 0, -1, 0)
	fmt.Println("2020-12-31的1月前的时间：", t1, err, t1.String(), t1.Day(), t1.Unix())

	// 月同比参考
	// 得到20201231的1月前的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20201231", 0, 0, -31)
	fmt.Println("20201231当天在上月的时间：", t1, err, t1.String(), t1.Day(), t1.Unix())
	// 得到20201130的1月前的时间
	t1, err = GetPreferTimeByTimeStrAndDifference("20201130", 0, 0, -31)
	fmt.Println("20201130当天在上月的时间：", t1, err, t1.String(), t1.Day(), t1.Unix())
}

func TestGetTimeByStr(t *testing.T) {
	t1, errs := GetTimeByStr("2020-10-01")
	fmt.Println(t1)
	fmt.Println(errs)
	t2, errs := GetTimeByStr("20201031")
	fmt.Println(t2)
	fmt.Println(errs)
	t3, errs := GetTimeByStr("2020/10/31") // error
	fmt.Println(t3)
	fmt.Println(errs)
	t4, errs := GetTimeByStr("202110a1") // error
	fmt.Println(t4)
	fmt.Println(errs)
	t5, errs := GetTimeByStr("31/10/2020") // error
	fmt.Println(t5)
	fmt.Println(errs)
}

func TestGetTimeListByTwoDayStr(t *testing.T) {
	// 跨年
	t1, t2 := GetTimeListByTwoDayStr("20201101", "20210102", false)
	fmt.Println(len(t1), t1) //
	fmt.Println(len(t2), t2) //
	t1, t2 = GetTimeListByTwoDayStr("20201101", "20210102", true)
	fmt.Println(len(t1), t1) //
	fmt.Println(len(t2), t2) //

	// 跨月
	t1, t2 = GetTimeListByTwoDayStr("20201128", "20201201", false)
	fmt.Println(len(t1), t1) // 4 [2020-11-28 2020-11-29 2020-11-30 2020-12-01]
	fmt.Println(len(t2), t2) // 4 [20201128 20201129 20201130 20201201]
	t1, t2 = GetTimeListByTwoDayStr("20201128", "20201201", true)
	fmt.Println(len(t1), t1) // 4 [2020-11-29 2020-11-30 2020-12-01 2020-12-02]
	fmt.Println(len(t2), t2) // 4 [20201129 20201130 20201201 20201202]

	// 同月
	t1, t2 = GetTimeListByTwoDayStr("20201003", "20201005", false)
	fmt.Println(len(t1), t1) // 3 [2020-10-03 2020-10-04 2020-10-05]
	fmt.Println(len(t2), t2) // 3 [20201003 20201004 20201005]
	t1, t2 = GetTimeListByTwoDayStr("20201003", "20201005", true)
	fmt.Println(len(t1), t1) // 3 [2020-10-04 2020-10-05 2020-10-06]
	fmt.Println(len(t2), t2) // 3 [20201004 20201005 20201006]
}

func TestGetAllDaysStrByMonths(t *testing.T) {
	fmt.Println(GetAllDaysStrByMonths("202008", "202008", true))  // [2020802 2020803 2020804 2020805 2020806 2020807 2020808 2020809 2020810 2020811 2020812 2020813 2020814 2020815 2020816 2020817 2020818 2020819 2020820 2020821 2020822 2020823 2020824 2020825 2020826 2020827 2020828 2020829 2020830 2020831 2020901]
	fmt.Println(GetAllDaysStrByMonths("202008", "202008", false)) // [2020801 2020802 2020803 2020804 2020805 2020806 2020807 2020808 2020809 2020810 2020811 2020812 2020813 2020814 2020815 2020816 2020817 2020818 2020819 2020820 2020821 2020822 2020823 2020824 2020825 2020826 2020827 2020828 2020829 2020830 2020831]
	fmt.Println(GetAllDaysStrByMonths("202007", "202008", true))  // [2020702 2020703 2020704 2020705 2020706 2020707 2020708 2020709 2020710 2020711 2020712 2020713 2020714 2020715 2020716 2020717 2020718 2020719 2020720 2020721 2020722 2020723 2020724 2020725 2020726 2020727 2020728 2020729 2020730 2020731 2020801 2020802 2020803 2020804 2020805 2020806 2020807 2020808 2020809 2020810 2020811 2020812 2020813 2020814 2020815 2020816 2020817 2020818 2020819 2020820 2020821 2020822 2020823 2020824 2020825 2020826 2020827 2020828 2020829 2020830 2020831 2020901]
	fmt.Println(GetAllDaysStrByMonths("202007", "202008", false)) // [2020701 2020702 2020703 2020704 2020705 2020706 2020707 2020708 2020709 2020710 2020711 2020712 2020713 2020714 2020715 2020716 2020717 2020718 2020719 2020720 2020721 2020722 2020723 2020724 2020725 2020726 2020727 2020728 2020729 2020730 2020731 2020801 2020802 2020803 2020804 2020805 2020806 2020807 2020808 2020809 2020810 2020811 2020812 2020813 2020814 2020815 2020816 2020817 2020818 2020819 2020820 2020821 2020822 2020823 2020824 2020825 2020826 2020827 2020828 2020829 2020830 2020831]
	fmt.Println(GetAllDaysStrByMonths("202007", "202009", true))  // [2020702 2020703 2020704 2020705 2020706 2020707 2020708 2020709 2020710 2020711 2020712 2020713 2020714 2020715 2020716 2020717 2020718 2020719 2020720 2020721 2020722 2020723 2020724 2020725 2020726 2020727 2020728 2020729 2020730 2020731 2020801 2020802 2020803 2020804 2020805 2020806 2020807 2020808 2020809 2020810 2020811 2020812 2020813 2020814 2020815 2020816 2020817 2020818 2020819 2020820 2020821 2020822 2020823 2020824 2020825 2020826 2020827 2020828 2020829 2020830 2020831 2020901 2020902 2020903 2020904 2020905 2020906 2020907 2020908 2020909 2020910 2020911 2020912 2020913 2020914 2020915 2020916 2020917 2020918 2020919 2020920 2020921 2020922 2020923 2020924 2020925 2020926 2020927 2020928 2020929 2020930 20201001]
	fmt.Println(GetAllDaysStrByMonths("202007", "202009", false)) // [2020701 2020702 2020703 2020704 2020705 2020706 2020707 2020708 2020709 2020710 2020711 2020712 2020713 2020714 2020715 2020716 2020717 2020718 2020719 2020720 2020721 2020722 2020723 2020724 2020725 2020726 2020727 2020728 2020729 2020730 2020731 2020801 2020802 2020803 2020804 2020805 2020806 2020807 2020808 2020809 2020810 2020811 2020812 2020813 2020814 2020815 2020816 2020817 2020818 2020819 2020820 2020821 2020822 2020823 2020824 2020825 2020826 2020827 2020828 2020829 2020830 2020831 2020901 2020902 2020903 2020904 2020905 2020906 2020907 2020908 2020909 2020910 2020911 2020912 2020913 2020914 2020915 2020916 2020917 2020918 2020919 2020920 2020921 2020922 2020923 2020924 2020925 2020926 2020927 2020928 2020929 2020930]
}

func TestGetYearMonthListByTwoDay(t *testing.T) {
	fmt.Println(GetYearMonthListByTwoDay("202008", "202008"))     // [202008]
	fmt.Println(GetYearMonthListByTwoDay("202007", "202008"))     // [202007 202008]
	fmt.Println(GetYearMonthListByTwoDay("202007", "202009"))     // [202007 202008 202009]
	fmt.Println(GetYearMonthListByTwoDay("20200703", "20200905")) // [202007 202008 202009]
	fmt.Println(GetYearMonthListByTwoDay("20200703", "20200805")) // [202007 202008]
}

func TestGetYearMonthListByTwoDayDelayOneDay(t *testing.T) {
	fmt.Println(GetYearMonthListByTwoDayDelayOneDay("202008", "202008"))     // [202009]
	fmt.Println(GetYearMonthListByTwoDayDelayOneDay("202007", "202008"))     // [202008 202009]
	fmt.Println(GetYearMonthListByTwoDayDelayOneDay("202007", "202009"))     // [202008 202009 202010]
	fmt.Println(GetYearMonthListByTwoDayDelayOneDay("20200705", "20200708")) // [202008]
	fmt.Println(GetYearMonthListByTwoDayDelayOneDay("20200705", "20200808")) // [202008 202009]
	fmt.Println(GetYearMonthListByTwoDayDelayOneDay("20200705", "20200908")) // [202008 202009 202010]
}

func TestGetDaysNumByTwoDays(t *testing.T) {
	days0 := GetDaysNumByTwoDays(2020, 11, 01, 12, 31) //跳月，返回33
	fmt.Println("总天数1=", days0)
	days1 := GetDaysNumByTwoDays(2020, 6, 30, 7, 2) //跳月，返回3
	fmt.Println("总天数2=", days1)
	days2 := GetDaysNumByTwoDays(2020, 6, 20, 6, 29) //当月，返回10
	fmt.Println("总天数3=", days2)
	days3 := GetDaysNumByTwoDays(2020, 5, 30, 8, 1) //跳月，返回64
	fmt.Println("总天数4=", days3)
}

func TestCompareTwoStrTime(t *testing.T) {
	fmt.Println("compare 20201001 and 20201001: ", CompareTwoStrTime("20201001", "20201001"))
	fmt.Println("compare 20201002 and 20201001: ", CompareTwoStrTime("20201002", "20201001"))
	fmt.Println("compare 20201102 and 20201001: ", CompareTwoStrTime("20201102", "20201001"))
	fmt.Println("compare 20201102 and 20201201: ", CompareTwoStrTime("20201102", "20201001"))
	fmt.Println("compare 202011 and 202010: ", CompareTwoStrTime("202011", "202010"))
	fmt.Println("compare 202101 and 202010: ", CompareTwoStrTime("202101", "202010"))
	fmt.Println("compare 2021 and 2020: ", CompareTwoStrTime("2021", "2020"))

	fmt.Println("compare 2021-01-01 and 2020-12-01: ", CompareTwoStrTime("2021-01-01", "2020-12-01"))
	fmt.Println("compare 2020-12-01 and 2020-12-01: ", CompareTwoStrTime("2021-01-01", "2020-12-01"))
	fmt.Println("compare 2020-12 and 2020-12: ", CompareTwoStrTime("2020-12", "2020-12"))
	fmt.Println("compare 2020-10-10 18:18:18 and 2020-10-10 18:18:18: ", CompareTwoStrTime("2020-10-10 18:18:18", "2020-10-10 18:18:18"))
	fmt.Println("compare 2020-10-10 18:18:19 and 2020-10-10 18:18:18: ", CompareTwoStrTime("2020-10-10 18:18:19", "2020-10-10 18:18:18"))
	fmt.Println("compare 2020-10-10 18:19:19 and 2020-10-10 18:18:19: ", CompareTwoStrTime("2020-10-10 18:19:19", "2020-10-10 18:18:19"))

	fmt.Println("compare 20201010 18:19:19 and 20201011 18:18:19: ", CompareTwoStrTime("20201010 18:19:19", "20201011 18:18:19"))
	fmt.Println("compare 20201010 18:19:19 and 20201010 18:18:19: ", CompareTwoStrTime("20201010 18:19:19", "20201010 18:18:19"))
	fmt.Println("compare 20201010 18:18:18 and 20201010 18:18:18: ", CompareTwoStrTime("20201010 18:18:18", "20201010 18:18:18"))

	fmt.Println("compare 2020/10/10 18:18:18 and 2020/10/10 18:18:18: ", CompareTwoStrTime("2020/10/10 18:18:18", "2020/10/10 18:18:18"))
	fmt.Println("compare 2020/10/10 18:18:19 and 2020/10/10 18:18:18: ", CompareTwoStrTime("2020/10/10 18:18:19", "2020/10/10 18:18:18"))
	fmt.Println("compare 2020/10/10 and 2020/10/10: ", CompareTwoStrTime("2020/10/10", "2020/10/10"))
	fmt.Println("compare 2020/10/11 and 2020/10/10: ", CompareTwoStrTime("2020/10/11", "2020/10/10"))
	fmt.Println("compare 2020/10 and 2020/10: ", CompareTwoStrTime("2020/10", "2020/10"))
	fmt.Println("compare 2020/10 and 2020/11: ", CompareTwoStrTime("2020/10", "2020/11"))
}
