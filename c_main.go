package main

import (
    "github.com/nsf/termbox-go"
    models "github.com/rian-hotate/tetoris/models/tetoris"
    "math/rand"
    "time"
)

//初期化
func (g *Game) initGame() {
    g.p = models.Piece{End: true}
    g.p.Vector.X, g.p.Vector.Y = 0, 1
    g.initPiece()
    g.p.Occupancy = []models.Point{}
    g.p.Span = 700
    g.p.Score = 0
}

//テトリミノ初期化
func (g *Game) initPiece() {
    rand.Seed(time.Now().UnixNano())
    g.p.TargetType = rand.Intn(7)
    piece := []models.Point{}

    if g.p.TargetType == 0 {
        for i := 0; i < 2; i++ {
            for j := 0; j < 2; j++ {
                piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
            }
        }
    } else if g.p.TargetType == 1 {
        for i := 0; i < 4; i++ {
            piece = append(piece, models.Point{X: (models.WIDTH / 2), Y: 5 + i})
        }
    } else if g.p.TargetType == 2 {
        for i := 0; i < 2; i++ {
            for j := 0; j < 3; j++ {
                if i == 0 {
                    if j == 1 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                } else {
                    piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                }
            }
        }
    } else if g.p.TargetType == 3 {
        for i := 0; i < 2; i++ {
            for j := 0; j < 3; j++ {
                if i == 0 {
                    if j != 2 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                } else {
                    if j != 0 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                }
            }
        }
    } else if g.p.TargetType == 4 {
        for i := 0; i < 2; i++ {
            for j := 0; j < 3; j++ {
                if i == 0 {
                    if j != 0 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                } else {
                    if j != 2 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                }
            }
        }
    } else if g.p.TargetType == 5 {
        for i := 0; i < 3; i++ {
            for j := 0; j < 2; j++ {
                if i == 0 {
                    piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                } else {
                    if j == 1 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                }
            }
        }
    } else if g.p.TargetType == 6 {
        for i := 0; i < 3; i++ {
            for j := 0; j < 2; j++ {
                if i == 0 {
                    piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                } else {
                    if j == 0 {
                        piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
                    }
                }
            }
        }
    }

    g.p.TargetOccupancy = piece
}

//テトリミノ回転
func (g *Game) rotationPiece() bool {
    cos := 0
    sin := 1

    for i := range g.p.TargetOccupancy {
        if i != 0 {
            x := g.p.TargetOccupancy[i].X - g.p.TargetOccupancy[0].X
            y := g.p.TargetOccupancy[i].Y - g.p.TargetOccupancy[0].Y
            g.p.TargetOccupancy[i].X = cos*x - sin*y + g.p.TargetOccupancy[0].X
            g.p.TargetOccupancy[i].Y = sin*x + cos*y + g.p.TargetOccupancy[0].Y
        }
    }
    f := false
    for j := range g.p.TargetOccupancy {
        for k := range g.p.Occupancy {
            if g.p.TargetOccupancy[j].X == g.p.Occupancy[k].X && g.p.TargetOccupancy[j].Y == g.p.Occupancy[k].Y {
                f = true
            }
        }
        if g.p.TargetOccupancy[j].X == 0 || g.p.TargetOccupancy[j].X == models.WIDTH || g.p.TargetOccupancy[j].Y == models.HEIGHT || g.p.TargetOccupancy[j].Y == 4 {
            f = true
        }
    }
    if f {
        for i := range g.p.TargetOccupancy {
            if i != 0 {
                x := g.p.TargetOccupancy[i].X - g.p.TargetOccupancy[0].X
                y := g.p.TargetOccupancy[i].Y - g.p.TargetOccupancy[0].Y
                g.p.TargetOccupancy[i].X = cos*x + sin*y + g.p.TargetOccupancy[0].X
                g.p.TargetOccupancy[i].Y = -sin*x + cos*y + g.p.TargetOccupancy[0].Y
            }
        }
    }

    return f
}

//当たり判定
func (g *Game) checkCollision() {
    occupantion := false
    for i := range g.p.TargetOccupancy {
        if g.p.TargetOccupancy[i].Y >= models.HEIGHT-1 {
            occupantion = true
            break
        } else {
            for j := range g.p.Occupancy {
                if g.p.TargetOccupancy[i].X == g.p.Occupancy[j].X && g.p.TargetOccupancy[i].Y+1 == g.p.Occupancy[j].Y {
                    occupantion = true
                    break
                }
            }
        }
    }
    if occupantion {
        for i := range g.p.TargetOccupancy {
            g.p.Occupancy = append(g.p.Occupancy, g.p.TargetOccupancy[i])
        }
        g.initPiece()
    }

    g.checkRow()

    for j := range g.p.TargetOccupancy {
        for k := range g.p.Occupancy {
            if g.p.TargetOccupancy[j].X == g.p.Occupancy[k].X && g.p.TargetOccupancy[j].Y == g.p.Occupancy[k].Y {
                g.p.End = true
            }
        }
    }
}

//行占有判定
func (g *Game) checkRow() {
    row := map[int]int{}
    for i := range g.p.Occupancy {
        row[g.p.Occupancy[i].Y]++
    }
    deleteTarget := []models.Point{}
    for key, value := range row {
        if value == models.WIDTH-1 {
            for j := range g.p.Occupancy {
                if g.p.Occupancy[j].Y == key {
                    deleteTarget = append(deleteTarget, g.p.Occupancy[j])
                }
            }
        }
    }

    for k := range deleteTarget {
        g.p.Occupancy = deleteElement(g.p.Occupancy, deleteTarget[k])
        g.p.Score += 5
        g.p.Span = 700 - g.p.Score/1
        if g.p.Span < 100 {
            g.p.Span = 100
        }
    }

    for key, value := range row {
        if value == models.WIDTH-1 {
            for j := range g.p.Occupancy {
                if g.p.Occupancy[j].Y <= key {
                    g.p.Occupancy[j].Y++
                }
            }
        }
    }
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

type Game struct {
    p models.Piece
}

func NewGame() *Game {
    return &Game{}
}
func (g *Game) init() error {
    return termbox.Init()
}

func (g *Game) close() {
    termbox.Close()
}
func (g *Game) run() {
    kch := make(chan termbox.Key)

    go keyEventLoop(kch)
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    g.initGame()
    g.drawInit()
    var timer time.Ticker
    for {
        select {
        case k := <-kch: //キーイベント
            switch k {
            case termbox.KeyEsc, termbox.KeyCtrlC: //ゲーム終了
                timer.Stop()
                g.p.End = true
                return
            case termbox.KeySpace, termbox.KeyEnter: //ゲームスタート
                if g.p.End {
                    g.p.End = false
                    timer = *time.NewTicker(time.Duration(g.p.Span) * time.Millisecond)
                    drawPiece(g.p)
                }
                break
            case termbox.KeyArrowLeft: //ひだり
                if g.p.End {
                    f := true
                    for i := range g.p.TargetOccupancy {
                        if f {
                            if g.p.TargetOccupancy[i].X <= 1 {
                                f = false
                            } else {
                                for j := range g.p.Occupancy {
                                    if g.p.TargetOccupancy[i].X == g.p.Occupancy[j].X+1 && g.p.TargetOccupancy[i].Y == g.p.Occupancy[j].Y {
                                        f = false
                                    }
                                }
                            }
                        }
                    }
                    if f {
                        for i := range g.p.TargetOccupancy {
                            g.p.TargetOccupancy[i].X--
                        }
                        timer.Stop()
                        timer = *time.NewTicker(time.Duration(g.p.Span) * time.Millisecond)
                    }
                    drawPiece(g.p)
                }
                break
            case termbox.KeyArrowRight: //みぎ
                f := true
                for i := range g.p.TargetOccupancy {
                    if f {
                        if g.p.TargetOccupancy[i].X >= models.WIDTH-1 {
                            f = false
                        } else {
                            for j := range g.p.Occupancy {
                                if g.p.TargetOccupancy[i].X == g.p.Occupancy[j].X-1 && g.p.TargetOccupancy[i].Y == g.p.Occupancy[j].Y {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range g.p.TargetOccupancy {
                        g.p.TargetOccupancy[i].X++
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(g.p.Span) * time.Millisecond)
                }
                drawPiece(g.p)
                break
            case termbox.KeyArrowDown: //した
                f := true
                for i := range g.p.TargetOccupancy {
                    if f {
                        if g.p.TargetOccupancy[i].Y >= models.HEIGHT-1 {
                            f = false
                        } else {
                            for j := range g.p.Occupancy {
                                if g.p.TargetOccupancy[i].X == g.p.Occupancy[j].X && g.p.TargetOccupancy[i].Y == g.p.Occupancy[j].Y-1 {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range g.p.TargetOccupancy {
                        g.p.TargetOccupancy[i].Y++
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(g.p.Span) * time.Millisecond)
                }
                drawPiece(g.p)
                break
            case termbox.KeyArrowUp: //うえ
                var f bool
                f = g.rotationPiece()
                if !f {
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(g.p.Span) * time.Millisecond)
                }
                drawPiece(g.p)
                break
            }
            break
        case <-timer.C: //タイマーイベント
            timer.Stop()
            g.checkCollision()
            if g.p.End == false {
                f := true
                for i := range g.p.TargetOccupancy {
                    if f {
                        if g.p.TargetOccupancy[i].Y >= models.HEIGHT-1 {
                            f = false
                        } else {
                            for j := range g.p.Occupancy {
                                if g.p.TargetOccupancy[i].X == g.p.Occupancy[j].X && g.p.TargetOccupancy[i].Y == g.p.Occupancy[j].Y-1 {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range g.p.TargetOccupancy {
                        g.p.TargetOccupancy[i].Y += g.p.Vector.Y
                    }
                }

                drawPiece(g.p)
                timer = *time.NewTicker(time.Duration(g.p.Span) * time.Millisecond)
            } else if g.p.End == true {
                g.initGame()
                g.drawInit()
            }
            break
        default:
            break
        }

    }
}
