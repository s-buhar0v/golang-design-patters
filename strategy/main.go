package main

import (
	"github.com/s-buhar0v/golang-design-patters/strategy/duck"
)

func main() {
	usualDuck := duck.NewUsualDuck()
	usualDuck.SetFlyBehaviuor(&duck.WingFly{})
	usualDuck.SetQuackBehaviuor(&duck.DefaultQuack{})

	rubberDuck := duck.NewRubberDuck()
	rubberDuck.SetFlyBehaviuor(&duck.NoFly{})
	rubberDuck.SetQuackBehaviuor(&duck.DefaultQuack{})

	usualDuck.PerformFly()   // print 'I'm flying with wings!'
	usualDuck.PerformQuack() // print 'Quack!'

	rubberDuck.PerformFly()                         // print 'I can't fly :('
	rubberDuck.PerformQuack()                       // print 'Quack'
	rubberDuck.SetQuackBehaviuor(&duck.LoudQuack{}) // dynamicly change DefaultQuack to LoudQuack
	rubberDuck.PerformQuack()                       // print 'QUACK!'
}
