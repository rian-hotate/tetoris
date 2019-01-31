package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	models "github.com/rian-hotate/tetoris/models/tetoris"
)

func drawStatus(p models.Piece) {
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
}

func drawInit(p models.Piece) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawStatus(p)
	drawLine(models.WIDTH/3, models.HEIGHT/2, "PUSH SPACE KEY")
	termbox.Flush()
}

func drawPiece(p models.Piece) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawStatus(p)
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
	termbox.Flush()
}

//行を描画
func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}
