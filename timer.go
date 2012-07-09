package ticked

import "time"

type Timer struct {
    ticker *time.Ticker
    active func() bool
    tick func()
}

func NewTimer(d time.Duration, active func() bool, tick func()) *Timer {
    return &Timer{ticker: time.NewTicker(d), active: active, tick: tick}
}

func (timer *Timer) Run() {
    for {
        select {
        case <-timer.ticker.C:
            if timer.active() {
                timer.tick()
            }
        }
    }
}

type Timers []*Timer

func (timers Timers) Run() {
    for _, timer := range timers {
        go timer.Run()
    }
}
