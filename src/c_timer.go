package main

import (
    "time"
    "models/tetoris"
)


//timer初期化
func initTime() models.Time {
    t := models.Time{}
    t.Status = true
    return t
}

//タイマーイベント
func timerLoop(tch chan models.Time, span int, stopCh chan bool, doneCh chan bool) {
    defer func() { doneCh <- true }()

    t := initTime()
    timer := time.NewTicker(time.Duration(span) * time.Millisecond)
    for {
        select {
        case <-timer.C:
            t.Status = true
            tch <- t
            break
        case <-stopCh:
            timer.Stop()
            return
        default:
        }
    }
    timer.Stop()
}
