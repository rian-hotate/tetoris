package main

import (
    "time"
    "models/tetoris"
)


//timer初期化
func initTime() models.Time {
    t := models.Time{}
    t.Status = true
    t.Span = 700
    return t
}

//タイマーイベント
func timerLoop(tch chan models.Time, pch chan models.Piece) {
    t := initTime()
    timer := time.NewTicker(time.Duration(t.Span) * time.Millisecond)
    tch <- t
    for {
        select{
        case <-timer.C:
            t.Status = true
            tch <- t
        }
    }
}
