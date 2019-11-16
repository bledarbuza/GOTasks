package main
import(
	"fmt"
)

type Person struct  {
	Emri,  Mbiemri,  Gjinia  string
	Mosha 					int

}
func( p Person) m() int{
	return p.Mosha
}
//var p = Person {"Bledar","Buza","M",20}
func main(){


	p1 :=  Person{Mosha:20}

	pts := &p1
	fmt.Println(pts)


}