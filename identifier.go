package webapi

type Identifier struct {
	Name  string
	Value int
}

func Bib(bib int) Identifier {
	return Identifier{
		Name:  "bib",
		Value: bib,
	}
}
func PID(pid int) Identifier {
	return Identifier{
		Name:  "pid",
		Value: pid,
	}
}
