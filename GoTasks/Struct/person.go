package Struct

type Person struct {
	emri    string `json:"emri"`
	mbiemri string `json:"mbiemri"`
	mosha   int    `json:"mosha"`
}

func (p *Person) Tedhenatperson(emri string, mbiemri string, mosha int) interface{} {

	p.emri = emri
	p.mbiemri = mbiemri
	p.mosha = mosha

	return p
}

func (p *Person) Ndryshomoshen() int {

	p.mosha = p.mosha + 10

	return p.mosha
}
