package game

type AppFactory interface {
	Init()
	Update()
	Draw()
	Close()
}
