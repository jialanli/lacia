package lacia

import (
	"fmt"
	lacia "github.com/jialanli/lacia/utils"
	"testing"
	"time"
)

func TestTimed1(t *testing.T) {
	now := time.Now().Unix() + 5
	fmt.Println("first exec=", now)
	r := NewTimed(now, 2)
	fmt.Println("--1-：", r.Snap.CurTime, r.Snap.NextTime)
	<-r.C
	fmt.Println("--2-：", r.Snap.CurTime, r.Snap.NextTime)
	r.StopTimedLan()
	fmt.Println("--3-：", r)
}

func TestTimed2(t *testing.T) {
	ts := time.Date(2021, 2, 2, 14, 18, 0, 0, time.Local).Unix()
	one := ts
	fmt.Println("now=", time.Now().Unix())
	fmt.Println("first=", one)
	r := NewTimed(one, 15)
	fmt.Println("--1-：", r.Snap.CurTime, r.Snap.NextTime)
	for {
		select {
		case <-r.C:
			fmt.Println("--3-：", r)
			fc()
		}
	}
}

func fc() {
	now := time.Now().Unix()
	fmt.Println("exec:", now, lacia.GetTimeStrOfDayTimeByTs(now))
}
