package models

type NameAndDepartmentObject struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

type IdAndAge struct {
	Id  int `json:"id"`
	Age int `json:"age"`
}
