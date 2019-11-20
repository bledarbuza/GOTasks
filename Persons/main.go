package main
import(
	"fmt"
)

type Person struct  {
	Emri,  Mbiemri,  Gjinia  string
	Mosha 					int

}

func( p Person) m() int{
	return 10 + p.Mosha


}
//var p = Person {"Bledar","Buza","M",20}
func main(){


	p :=  Person{Mosha:5}
	fmt.Println("m: ", p.m())

	pts := &p
	fmt.Println(pts.m())


}