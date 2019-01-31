package main

import (
    "models/tetoris"
    "github.com/nsf/termbox-go"
    "math/rand"
    "time"
)

//初期化
func initGame() models.Piece {
    p := models.Piece{End: true}
    p.Vector.X, p.Vector.Y = 0, 1
    p.TargetOccupancy = initPiece(p)
    p.Occupancy = []models.Point{}
    p.Span = 700
    p.Score = 0
    return p
}

//テトリミノ初期化
func initPiece(p models.Piece) []models.Point {
    rand.Seed(time.Now().UnixNano())
    p.TargetType = rand.Intn(7)
    piece := []models.Point{}

    if p.TargetType == 0 {
        for i := 0; i < 2; i++ {
            for j:= 0; j < 2; j++ {
                piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
            }
        }
    } else if p.TargetType == 1 {
        for i := 0; i < 4; i++ {
            piece = append(piece, models.Point{X: (models.WIDTH/2), Y: 5+i})
        }
    } else if p.TargetType == 2 {
        for i := 0; i < 2; i++ {
            for j:= 0; j < 3; j++ {
                if (i == 0) {
                    if (j == 1) {
                      piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                } else {
                    piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                }
            }
        }
    } else if p.TargetType == 3 {
        for i := 0; i < 2; i++ {
            for j:= 0; j < 3; j++ {
                if i == 0 {
                    if j != 2 {
                        piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                } else {
                    if j != 0 {
                        piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                }
            }
        }
     } else if p.TargetType == 4 {
        for i := 0; i < 2; i++ {
            for j:= 0; j < 3; j++ {
                if i == 0 {
                    if j != 0 {
                        piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                } else {
                    if j != 2 {
                        piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                }
            }
        }
     } else if p.TargetType == 5 {
        for i := 0; i < 3; i++ {
            for j:= 0; j < 2; j++ {
                if i == 0 {
                    piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                } else {
                    if j == 1 {
                      piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                }
            }
        }
     } else if p.TargetType == 6 {
        for i := 0; i < 3; i++ {
            for j:= 0; j < 2; j++ {
                if i == 0 {
                    piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                } else {
                    if j == 0 {
                      piece = append(piece, models.Point{X: (models.WIDTH/2)+j, Y: 5+i})
                    }
                }
            }
        }
    }

    return piece
}

//テトリミノ回転
func rotationPiece(p models.Piece) (models.Piece, bool) {
    cos := 0
    sin := 1

    for i:= range p.TargetOccupancy {
        if (i != 0) {
            x := p.TargetOccupancy[i].X - p.TargetOccupancy[0].X
            y := p.TargetOccupancy[i].Y - p.TargetOccupancy[0].Y
            p.TargetOccupancy[i].X = cos*x - sin*y + p.TargetOccupancy[0].X
            p.TargetOccupancy[i].Y = sin*x + cos*y + p.TargetOccupancy[0].Y
        }
    }
    f := false
    for j := range p.TargetOccupancy {
        for k := range p.Occupancy {
            if p.TargetOccupancy[j].X == p.Occupancy[k].X && p.TargetOccupancy[j].Y == p.Occupancy[k].Y {
                f = true
            }
        }
        if p.TargetOccupancy[j].X == 0 || p.TargetOccupancy[j].X == models.WIDTH || p.TargetOccupancy[j].Y == models.HEIGHT || p.TargetOccupancy[j].Y == 4 {
            f = true
        }
    }
    if f {
        for i:= range p.TargetOccupancy {
            if i != 0 {
                x := p.TargetOccupancy[i].X - p.TargetOccupancy[0].X
                y := p.TargetOccupancy[i].Y - p.TargetOccupancy[0].Y
                p.TargetOccupancy[i].X = cos*x + sin*y + p.TargetOccupancy[0].X
                p.TargetOccupancy[i].Y = -sin*x + cos*y + p.TargetOccupancy[0].Y
            }
        }
    }

    return p, f
}

//当たり判定
func checkCollision(p models.Piece) models.Piece {
    occupantion := false
    for i := range p.TargetOccupancy {
        if p.TargetOccupancy[i].Y >= models.HEIGHT-1 {
            occupantion = true
            break
        } else {
            for j := range p.Occupancy {
                if p.TargetOccupancy[i].X == p.Occupancy[j].X && p.TargetOccupancy[i].Y+1 == p.Occupancy[j].Y {
                    occupantion = true
                    break
                }
            }
        }
    }
    if occupantion {
        for i := range p.TargetOccupancy {
            p.Occupancy = append(p.Occupancy, p.TargetOccupancy[i])
        }
        p.TargetOccupancy = initPiece(p)
    }

    p = checkRow(p)

    for j := range p.TargetOccupancy {
        for k := range p.Occupancy {
            if p.TargetOccupancy[j].X == p.Occupancy[k].X && p.TargetOccupancy[j].Y == p.Occupancy[k].Y {
                p.End = true
            }
        }
    }

    return p
}

//行占有判定
func checkRow(p models.Piece) models.Piece {
    row := map[int]int{}
    for i := range p.Occupancy {
        row[p.Occupancy[i].Y]++
    }
    deleteTarget := []models.Point{}
    for key, value := range row {
        if value == models.WIDTH-1 {
            for j := range p.Occupancy {
                if p.Occupancy[j].Y == key {
                    deleteTarget = append(deleteTarget, p.Occupancy[j])
                }
            }
        }
    }

    for k := range deleteTarget {
        p.Occupancy = deleteElement(p.Occupancy, deleteTarget[k])
        p.Score += 5
        if (p.HighScore < p.Score) {
            p.HighScore = p.Score
        }
        p.Span = 700 - p.Score/1
        if p.Span < 100 {
            p.Span = 100
        }
    }

    for key, value := range row {
        if value == models.WIDTH-1 {
            for j := range p.Occupancy {
                if (p.Occupancy[j].Y <= key) {
                    p.Occupancy[j].Y++
                }
            }
        }
    }

    return p
}

//slice要素の削除
func deleteElement(target []models.Point, element models.Point) []models.Point {
    ret := []models.Point{}
    for i := range target {
        if target[i] != element {
            ret = append(ret, target[i])
        }
    }
    return ret
}

func controller(kch chan termbox.Key) {
    p := initGame()
    drawInit(p)
    var timer time.Ticker
    for {
        select {
        case k := <-kch: //キーイベント
            switch k {
            case termbox.KeyEsc, termbox.KeyCtrlC: //ゲーム終了
                timer.Stop()
                p.End = true
                return
            case termbox.KeySpace, termbox.KeyEnter: //ゲームスタート
                if p.End {
                    p.End = false
                    timer = *time.NewTicker(time.Duration(700) * time.Millisecond)
                    drawPiece(p)
                }
                break
            case termbox.KeyArrowLeft: //ひだり
                f := true
                for i := range p.TargetOccupancy {
                    if f {
                        if p.TargetOccupancy[i].X <= 1 {
                            f = false
                        } else {
                            for j := range p.Occupancy {
                                if p.TargetOccupancy[i].X == p.Occupancy[j].X+1 && p.TargetOccupancy[i].Y == p.Occupancy[j].Y {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range p.TargetOccupancy {
                        p.TargetOccupancy[i].X--
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(p.Span) * time.Millisecond)
                }
                drawPiece(p)
                break
            case termbox.KeyArrowRight: //みぎ
                f := true
                for i := range p.TargetOccupancy {
                    if f {
                        if p.TargetOccupancy[i].X >= models.WIDTH-1 {
                            f = false
                        } else {
                            for j := range p.Occupancy {
                                if p.TargetOccupancy[i].X == p.Occupancy[j].X-1 && p.TargetOccupancy[i].Y == p.Occupancy[j].Y {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range p.TargetOccupancy {
                        p.TargetOccupancy[i].X++
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(p.Span) * time.Millisecond)
                }
                drawPiece(p)
                break
            case termbox.KeyArrowDown: //した
                f := true
                for i := range p.TargetOccupancy {
                    if f {
                        if p.TargetOccupancy[i].Y >= models.HEIGHT-1 {
                            f = false
                        } else {
                            for j := range p.Occupancy {
                                if p.TargetOccupancy[i].X == p.Occupancy[j].X && p.TargetOccupancy[i].Y == p.Occupancy[j].Y-1 {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range p.TargetOccupancy {
                        p.TargetOccupancy[i].Y++
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(p.Span) * time.Millisecond)
                }
                drawPiece(p)
                break
            case termbox.KeyArrowUp: //うえ
                var f bool
                p, f = rotationPiece(p)
                if (!f) {
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(p.Span) * time.Millisecond)
                }
                drawPiece(p)
                break
            }
            break
        case <-timer.C: //タイマーイベント
            timer.Stop()
            p = checkCollision(p)
            if p.End == false {
                f := true
                for i := range p.TargetOccupancy {
                    if f {
                        if p.TargetOccupancy[i].Y >= models.HEIGHT-1 {
                            f = false
                        } else {
                            for j := range p.Occupancy {
                                if p.TargetOccupancy[i].X == p.Occupancy[j].X && p.TargetOccupancy[i].Y == p.Occupancy[j].Y-1 {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range p.TargetOccupancy {
                        p.TargetOccupancy[i].Y += p.Vector.Y
                    }
                }

                drawPiece(p)
                timer = *time.NewTicker(time.Duration(p.Span) * time.Millisecond)
            } else if p.End == true {
                p = initGame()
                drawInit(p)
            }
            break
        default:
            break
        }

    }
}
