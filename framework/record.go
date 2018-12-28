package framework

/*
// Record is a type that holds all the information
// relating to the current record for a given animal
*/
type Record struct {
	Animal		string
	Holder		string
	Score		float64
	Scoresheet	string
}

/*
// NewRecord is a constructor-like factory function that
// initializes an instance of a Record and returns a pointer
// to the variable
*/
func NewRecord(animal string) *Record {
	rec := new(Record)
	rec.Animal = animal
	rec.Holder = "<nobody>"
	rec.Score = 0.0
	rec.Scoresheet = "0"
	return rec
}
