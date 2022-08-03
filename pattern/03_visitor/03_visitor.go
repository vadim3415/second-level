package main

import (
	"fmt"
)

/*
Посетитель — это поведенческий паттерн,
который позволяет добавить новую операцию для целой иерархии классов,
не изменяя код этих классов.
Плюсы:
-Нет необходимости изменять классы
-Похожие операции над разными объектами хранятся в одном месте
Минусы:
-Лишний код
-Может привести к нарушению инкапсуляции элементов.
*/

type VisitorFigureShapes interface {
	VisitQuadrate(rectangle *Quadrate)
	VisitCircle(circle *Circle)
}

type VisitorShapePrint struct{}


func (VisitorShapePrint) VisitQuadrate(quadrate *Quadrate) {
	fmt.Printf("прямоугольник: %+v\n", *quadrate)
}

func (VisitorShapePrint) VisitCircle(circle *Circle) {
	fmt.Printf("круг: %+v\n", *circle)
}

type figureShapes struct {
	quadrate Quadrate
	circle    Circle
}

type Quadrate struct{ 
	width, 
	height int
}

type Circle struct{ 
	radius int
}

func (s *figureShapes) Visit(v VisitorFigureShapes) {
	v.VisitQuadrate(&s.quadrate)
	v.VisitCircle(&s.circle)
}

func NewFigureSizes() *figureShapes {
	s := new(figureShapes)

	s.quadrate.width = 10
	s.quadrate.height = 10
	s.circle.radius = 10
	return s
}

func main() {
	shapes := NewFigureSizes()
	visitorShape := VisitorShapePrint{}
	shapes.Visit(visitorShape)
}