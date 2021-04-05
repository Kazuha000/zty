package main

import "fmt"

type Painter interface{
	Def(i,j float64) bool
	Range()(x1,x2,x3,x4 float64)
	Change()(x,y float64)
}
type Heart struct {

}
func (h Heart)Def(i,j float64) bool{
	d:=i*i+j*j-1
	return d*d*d-i*i*i*j*j<=0
}
func(h Heart) Range()(x1,x2,y1,y2 float64){
return 1.2,-1.2,-2.5,2.5
}
func(h Heart)Change()(x,y float64){
	return 0.08,0.05
}

type Round struct {

}
func (r Round)Def(x,y float64) bool{
	return x*x+y*y<=9
}
func(r Round) Range()(x1,x2,y1,y2 float64){
	return 5,-5,-5,5
}
func(r Round)Change()(x,y float64){
	return 0.5,0.2
}
func Paint(painter Painter){
	x1,x2,y1,y2:=painter.Range()
	x,y:=painter.Change()
	for i:=x1;i>x2;i-=x{
		for j:=y1 ;j<y2;j+=y{
			if painter.Def(i,j){
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func main()  {
	//var heart Heart;Paint(heart)
    var round Round;Paint(round)
}