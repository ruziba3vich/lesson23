package models

type Person struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}
