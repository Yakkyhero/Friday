package main
import (
	"fmt"
)
type Point struct{//ประกาศ struct ที่ชื่อว่า Point
	X float64
	y float64
}
func Distance(){
	fmt.Println("hello function Distance")
}
func (p Point) Distance(){
	fmt.Println("show method", p)
}
type Path []Point

func (p Path) Distance() {
	fmt.Println("hello path Distance", p)
}

func main(){
	path := Point{
		Point{X: 2,Y: 2},
	}
	p := Point{X: 1, Y: 1}  //ประกาศจุดแรก x1y1
	// q := Point{X: 1, Y: 4}
	p.Distance()
	Distance()
}