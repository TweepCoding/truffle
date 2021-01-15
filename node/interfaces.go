package node

type Drawer interface {
	Draw() error
	OnDraw(func() error)
}

type Updater interface {
	Update(float64) error
	OnUpdate(func(float64) error)
}

type DrawUpdater interface {
	Drawer
	Updater
}

type Positioner interface {
	GetX() float64
	GetY() float64
	SetX(float64)
	SetY(float64)
}

type Measurer interface {
	GetW() int32
	GetH() int32
	SetW(int32)
	SetH(int32)
}

type Collisioner interface {
	Measurer
	OnCollision(func(Collisioner))
}
