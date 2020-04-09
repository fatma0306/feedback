package models

//Review is a struct which represents review infos
type Review struct {
	ReviewID      int
	ReviewContent string
	ReviewState   bool
	CustomerID    int
}
