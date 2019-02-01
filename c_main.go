package main

import (
    "github.com/nsf/termbox-go"
    models "github.com/rian-hotate/tetoris/models/tetoris"
    "math/rand"
    "time"
)

//初期化
func (p *Piece) initGame() {
    p.piece.End = true
    p.piece.Vector.X, p.piece.Vector.Y = 0, 1
    p.initPiece()
    p.piece.Occupancy = []models.Point{}
    p.piece.Span = 700
    p.piece.Score = 0
}

//テトリミノ初期化
func (p *Piece) initPiece() {
    rand.Seed(time.Now().UnixNano())
    p.piece.TargetType = rand.Intn(7)
    piece := []models.Point{}

    if p.piece.TargetType == 0 {
        for i := 0; i < 2; i++ {
            for j := 0; j < 2; j++ {
                piece = append(piece, models.Point{X: (models.WIDTH / 2) + j, Y: 5 + i})
            }
        }
    } else if p.piece.TargetType == 1 {
        for i := 0; i < 4; i++ {
            piece = append(piece, models.Point{X: (models.WIDTH / 2), Y: 5 + i})
        }
    } else if p.piece.TargetType == 2 {
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
    } else if p.piece.TargetType == 3 {
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
    } else if p.piece.TargetType == 4 {
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
    } else if p.piece.TargetType == 5 {
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
    } else if p.piece.TargetType == 6 {
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

    p.piece.TargetOccupancy = piece
}

//テトリミノ回転
func (p *Piece) rotationPiece() bool {
    cos := 0
    sin := 1

    for i := range p.piece.TargetOccupancy {
        if i != 0 {
            x := p.piece.TargetOccupancy[i].X - p.piece.TargetOccupancy[0].X
            y := p.piece.TargetOccupancy[i].Y - p.piece.TargetOccupancy[0].Y
            p.piece.TargetOccupancy[i].X = cos*x - sin*y + p.piece.TargetOccupancy[0].X
            p.piece.TargetOccupancy[i].Y = sin*x + cos*y + p.piece.TargetOccupancy[0].Y
        }
    }
    f := false
    for j := range p.piece.TargetOccupancy {
        for k := range p.piece.Occupancy {
            if p.piece.TargetOccupancy[j].X == p.piece.Occupancy[k].X && p.piece.TargetOccupancy[j].Y == p.piece.Occupancy[k].Y {
                f = true
            }
        }
        if p.piece.TargetOccupancy[j].X == 0 || p.piece.TargetOccupancy[j].X == models.WIDTH || p.piece.TargetOccupancy[j].Y == models.HEIGHT || p.piece.TargetOccupancy[j].Y == 4 {
            f = true
        }
    }
    if f {
        for i := range p.piece.TargetOccupancy {
            if i != 0 {
                x := p.piece.TargetOccupancy[i].X - p.piece.TargetOccupancy[0].X
                y := p.piece.TargetOccupancy[i].Y - p.piece.TargetOccupancy[0].Y
                p.piece.TargetOccupancy[i].X = cos*x + sin*y + p.piece.TargetOccupancy[0].X
                p.piece.TargetOccupancy[i].Y = -sin*x + cos*y + p.piece.TargetOccupancy[0].Y
            }
        }
    }

    return f
}

//当たり判定
func (p *Piece) checkCollision() {
    occupantion := false
    for i := range p.piece.TargetOccupancy {
        if p.piece.TargetOccupancy[i].Y >= models.HEIGHT-1 {
            occupantion = true
            break
        } else {
            for j := range p.piece.Occupancy {
                if p.piece.TargetOccupancy[i].X == p.piece.Occupancy[j].X && p.piece.TargetOccupancy[i].Y+1 == p.piece.Occupancy[j].Y {
                    occupantion = true
                    break
                }
            }
        }
    }
    if occupantion {
        for i := range p.piece.TargetOccupancy {
            p.piece.Occupancy = append(p.piece.Occupancy, p.piece.TargetOccupancy[i])
        }
        p.initPiece()
    }

    p.checkRow()

    for j := range p.piece.TargetOccupancy {
        for k := range p.piece.Occupancy {
            if p.piece.TargetOccupancy[j].X == p.piece.Occupancy[k].X && p.piece.TargetOccupancy[j].Y == p.piece.Occupancy[k].Y {
                p.piece.End = true
            }
        }
    }
}

//行占有判定
func (p *Piece) checkRow() {
    row := map[int]int{}
    for i := range p.piece.Occupancy {
        row[p.piece.Occupancy[i].Y]++
    }
    deleteTarget := []models.Point{}
    for key, value := range row {
        if value == models.WIDTH-1 {
            for j := range p.piece.Occupancy {
                if p.piece.Occupancy[j].Y == key {
                    deleteTarget = append(deleteTarget, p.piece.Occupancy[j])
                }
            }
        }
    }

    for k := range deleteTarget {
        p.piece.Occupancy = deleteElement(p.piece.Occupancy, deleteTarget[k])
        p.piece.Score += 5
        p.piece.Span = 700 - p.piece.Score/1
        if p.piece.Span < 100 {
            p.piece.Span = 100
        }
    }

    for key, value := range row {
        if value == models.WIDTH-1 {
            for j := range p.piece.Occupancy {
                if p.piece.Occupancy[j].Y <= key {
                    p.piece.Occupancy[j].Y++
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
    p *Piece
}

type Piece struct {
    piece models.Piece
}

func NewGame() *Game {
    return &Game{}
}

func NewModel() *Piece {
    return &Piece{}
}

func (g *Game) init() error {
    g.p = NewModel()
    return termbox.Init()
}

func (g *Game) close() {
    termbox.Close()
}
func (g *Game) run() {
    kch := make(chan termbox.Key)

    go keyEventLoop(kch)
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    g.p.initGame()
    g.p.drawInit()
    var timer time.Ticker
    for {
        select {
        case k := <-kch: //キーイベント
            switch k {
            case termbox.KeyEsc, termbox.KeyCtrlC: //ゲーム終了
                timer.Stop()
                g.p.piece.End = true
                return
            case termbox.KeySpace, termbox.KeyEnter: //ゲームスタート
                if g.p.piece.End {
                    g.p.piece.End = false
                    timer = *time.NewTicker(time.Duration(g.p.piece.Span) * time.Millisecond)
                    g.p.drawPiece()
                }
                break
            case termbox.KeyArrowLeft: //ひだり
                if !g.p.piece.End {
                    f := true
                    for i := range g.p.piece.TargetOccupancy {
                        if f {
                            if g.p.piece.TargetOccupancy[i].X <= 1 {
                                f = false
                            } else {
                                for j := range g.p.piece.Occupancy {
                                    if g.p.piece.TargetOccupancy[i].X == g.p.piece.Occupancy[j].X+1 && g.p.piece.TargetOccupancy[i].Y == g.p.piece.Occupancy[j].Y {
                                        f = false
                                    }
                                }
                            }
                        }
                    }
                    if f {
                        for i := range g.p.piece.TargetOccupancy {
                            g.p.piece.TargetOccupancy[i].X--
                        }
                        timer.Stop()
                        timer = *time.NewTicker(time.Duration(g.p.piece.Span) * time.Millisecond)
                    }
                    g.p.drawPiece()
                }
                break
            case termbox.KeyArrowRight: //みぎ
                f := true
                for i := range g.p.piece.TargetOccupancy {
                    if f {
                        if g.p.piece.TargetOccupancy[i].X >= models.WIDTH-1 {
                            f = false
                        } else {
                            for j := range g.p.piece.Occupancy {
                                if g.p.piece.TargetOccupancy[i].X == g.p.piece.Occupancy[j].X-1 && g.p.piece.TargetOccupancy[i].Y == g.p.piece.Occupancy[j].Y {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range g.p.piece.TargetOccupancy {
                        g.p.piece.TargetOccupancy[i].X++
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(g.p.piece.Span) * time.Millisecond)
                }
                g.p.drawPiece()
                break
            case termbox.KeyArrowDown: //した
                f := true
                for i := range g.p.piece.TargetOccupancy {
                    if f {
                        if g.p.piece.TargetOccupancy[i].Y >= models.HEIGHT-1 {
                            f = false
                        } else {
                            for j := range g.p.piece.Occupancy {
                                if g.p.piece.TargetOccupancy[i].X == g.p.piece.Occupancy[j].X && g.p.piece.TargetOccupancy[i].Y == g.p.piece.Occupancy[j].Y-1 {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range g.p.piece.TargetOccupancy {
                        g.p.piece.TargetOccupancy[i].Y++
                    }
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(g.p.piece.Span) * time.Millisecond)
                }
                g.p.drawPiece()
                break
            case termbox.KeyArrowUp: //うえ
                var f bool
                f = g.p.rotationPiece()
                if !f {
                    timer.Stop()
                    timer = *time.NewTicker(time.Duration(g.p.piece.Span) * time.Millisecond)
                }
                g.p.drawPiece()
                break
            }
            break
        case <-timer.C: //タイマーイベント
            timer.Stop()
            g.p.checkCollision()
            if g.p.piece.End == false {
                f := true
                for i := range g.p.piece.TargetOccupancy {
                    if f {
                        if g.p.piece.TargetOccupancy[i].Y >= models.HEIGHT-1 {
                            f = false
                        } else {
                            for j := range g.p.piece.Occupancy {
                                if g.p.piece.TargetOccupancy[i].X == g.p.piece.Occupancy[j].X && g.p.piece.TargetOccupancy[i].Y == g.p.piece.Occupancy[j].Y-1 {
                                    f = false
                                }
                            }
                        }
                    }
                }
                if f {
                    for i := range g.p.piece.TargetOccupancy {
                        g.p.piece.TargetOccupancy[i].Y += g.p.piece.Vector.Y
                    }
                }

                g.p.drawPiece()
                timer = *time.NewTicker(time.Duration(g.p.piece.Span) * time.Millisecond)
            } else if g.p.piece.End == true {
                g.p.initGame()
                g.p.drawInit()
            }
            break
        default:
            break
        }

    }
}
