package lacia

import (
	"time"
)

/*
sample timer
author: jialanli
*/

type TimeSnap struct {
	CurTime  int64
	NextTime int64
}

type TimedLan struct {
	Snap     TimeSnap
	C        chan *TimeSnap
	IsClosed bool
	Step     int64
	T        int64
}

type TaskInterface interface {
	NewTimed()
}

func NewTimed(ts, step int64) *TimedLan {
	if ts < time.Now().Unix() {
		ts = time.Now().Unix()
	}
	d := TimeSnap{CurTime: ts, NextTime: ts}
	r := new(TimedLan)
	r.C, r.IsClosed, r.Step, r.Snap = make(chan *TimeSnap, 1), false, step, d
	go r.StartTimedLan()
	return r
}

func (tl *TimedLan) StartTimedLan() {
	for {
		if tl.Snap.NextTime != tl.Snap.CurTime {
			tl.Snap.NextTime = tl.Snap.CurTime + tl.Step
		}
		if tl.Snap.NextTime <= time.Now().Unix() {
			//fmt.Println( tl.Snap.NextTime,  time.Now().Unix())
			tl.C <- &TimeSnap{CurTime: tl.Snap.CurTime, NextTime: tl.Snap.NextTime + tl.Step}
			tl.Snap.CurTime, tl.Snap.NextTime = tl.Snap.NextTime, tl.Snap.NextTime+tl.Step
		}
	}
}

func (tl *TimedLan) StopTimedLan() {
	close(tl.C)
	tl.IsClosed = true
}

func (tl *TimedLan) SetStopTimeLan(stopT int64) {
	tl.T = stopT
}
