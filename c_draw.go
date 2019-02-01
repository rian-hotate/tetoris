package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
    models "github.com/rian-hotate/tetoris/models/tetoris"
)

func (p *Piece) drawStatus() {
    drawLine(0, 0, "EXIT : ESC KEY")
    drawLine(20, 1, fmt.Sprintf("Score : %05d", p.piece.Score))

    for i := 0; i <= models.WIDTH; i++ {
        for j := 4; j <= models.HEIGHT; j++ {
            if i == 0 || i == models.WIDTH {
                drawLine(i, j, "/")
            } else if j == 4 || j == models.HEIGHT {
                drawLine(i, j, "/")
            }
        }
    }
}

func (p *Piece) drawInit() error {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    p.drawStatus()
    drawLine(models.WIDTH/3, models.HEIGHT/2, "PUSH SPACE KEY")
    return termbox.Flush()
}

func (p *Piece) drawPiece() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    p.drawStatus()
    if len(p.piece.TargetOccupancy) != 0 {
        for i := range p.piece.TargetOccupancy {
            drawLine(p.piece.TargetOccupancy[i].X, p.piece.TargetOccupancy[i].Y, "#")
        }
    }
    if len(p.piece.Occupancy) != 0 {
        for i := range p.piece.Occupancy {
            drawLine(p.piece.Occupancy[i].X, p.piece.Occupancy[i].Y, "#")
        }
    }
    termbox.Flush()
}

//行を描画
func drawLine(x, y int, str string) {
    runes := []rune(str)
    for i := 0; i < len(runes); i++ {
        termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
    }
}
