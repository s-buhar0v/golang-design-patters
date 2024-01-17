package duck

// Duck is base class.
type Duck struct {
	flyBehaviuor   FlyBehaviuor
	quackBehaviuor QuackBehaviuor
}

func (d *Duck) SetFlyBehaviuor(f FlyBehaviuor) {
	d.flyBehaviuor = f
}

func (d *Duck) PerformFly() {
	d.flyBehaviuor.Fly()
}

func (d *Duck) SetQuackBehaviuor(q QuackBehaviuor) {
	d.quackBehaviuor = q
}

func (d *Duck) PerformQuack() {
	d.quackBehaviuor.Quack()
}

// UsualDuck inherits Duck.
type UsualDuck struct{ *Duck }

func NewUsualDuck() *UsualDuck {
	return &UsualDuck{&Duck{}}
}

// RubberDuck inherits Duck.
type RubberDuck struct{ *Duck }

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{&Duck{}}
}
