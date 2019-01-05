package framework

import "strings"

//
// Record is a type that holds all the information
// relating to the record for a given animal
//
type Record struct {
	Animal     string
	Holder     string
	Score      float64
	Scoresheet string
}

//
// NewRecord is a constructor-like factory function that
// initializes an instance of a Record and returns a pointer
// to the variable
//
func NewRecord(animal string) *Record {
	rec := new(Record)
	rec.Animal = animal
	rec.Holder = "<nobody>"
	rec.Score = 0.0
	rec.Scoresheet = "0"
	return rec
}

//
// GetRecordIndexByAnimal takes as input a slice of records and
// the animal name of which to get the record index of. It returns
// -1 when no record bound to the animal was found and a valid
// index when it is found
//
func GetRecordIndexByAnimal(records []*Record, animal string) int {
	for i, rec := range records {
		if strings.Contains(strings.ToLower(rec.Animal), strings.ToLower(animal)) {
			return i
		}
	}
	return -1
}
