package main

import (
    "models/tetoris"
    "github.com/nsf/termbox-go"
    "fmt"
)

func drawLoop(pch chan models.Piece) {
    for {
        p := <-pch
        models.Mu.Lock()
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
        drawLine(0, 0, "EXIT : ESC KEY")
        drawLine(0, 1, fmt.Sprintf("HighScore : %05d", p.HighScore))
        drawLine(20, 1, fmt.Sprintf("Score : %05d", p.Score))

        for i := 0; i <= models.WIDTH; i++ {
            for j := 4; j <= models.HEIGHT; j++ {
                if i == 0 || i == models.WIDTH {
                    drawLine(i, j, "/")
                } else if j == 4 || j == models.HEIGHT {
                    drawLine(i, j, "/")
                }
            }
        }

        if p.End == false {
            if len(p.TargetOccupancy) != 0 {
                for i := range p.TargetOccupancy {
                    drawLine(p.TargetOccupancy[i].X, p.TargetOccupancy[i].Y, "#")
                }
            }
            if len(p.Occupancy) != 0 {
                for i := range p.Occupancy {
                    drawLine(p.Occupancy[i].X, p.Occupancy[i].Y, "#")
                }
            }
        } else {
            drawLine(models.WIDTH/3, models.HEIGHT/2, "PUSH SPACE KEY")
        }

        termbox.Flush()
        models.Mu.Unlock()
    }
}

//行を描画
func drawLine(x, y int, str string) {
    runes := []rune(str)
    for i := 0; i < len(runes); i++ {
        termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
    }
}
