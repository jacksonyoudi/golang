package main

//import "image/draw"

type Shape interface {
	Draw()
}


type ShapeFactory struct {

}

func (s ShapeFactory) gen() Shape {


}