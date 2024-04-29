package lacia

import (
	"log"
	"testing"
	"time"
)

func TestTimed1(t *testing.T) {
	now := time.Now().Unix() + 5 // 定时开始时间
	log.Println("first exec=", now)
	r := NewTimed(now, 2) // 每隔2秒一次
	log.Println("--1-：", r.Snap.CurTime, r.Snap.NextTime)
	<-r.C
	log.Println("--2-：", r.Snap.CurTime, r.Snap.NextTime)
	r.StopTimedLan()
	log.Println("--3-：", r)
}

func TestTimed2(t *testing.T) {
	ts := time.Date(2021, 2, 2, 14, 18, 0, 0, time.Local).Unix()
	one := ts // 定时开始时间
	log.Println("now=", time.Now().Unix())
	log.Println("first=", one)
	r := NewTimed(one, 15) // 每隔15秒一次
	log.Println("--1-：", r.Snap.CurTime, r.Snap.NextTime)
	for {
		select {
		case <-r.C:
			log.Println("--3-：", r)
			fc()
		}
	}
}

func fc() {
	now := time.Now().Unix()
	log.Println("exec:", now, GetTimeStrOfDayTimeByTs(now))
}
