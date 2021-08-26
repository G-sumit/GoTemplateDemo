package main

type NewInventory struct {
	Name string
	Id int
	Surname string
	Contact string
	SecondName string
	SecondId int
	SecondSurname string
	SecondContact string
	TName string
	TSurname string
	TContact string
	Second *SecondInventory
}

type SecondInventory struct{
	Name string
	Id int
	Surname string
	Contact string
	SecondName string
	SecondId int
	SecondSurname string
	SecondContact string
	TName string
	TSurname string
	TContact string
	Third *ThirdInventory

}

type ThirdInventory struct{
	Name string
	Id int
	Surname string
	Contact string
	SecondName string
	SecondId int
	SecondSurname string
	SecondContact string
	TName string
	TSurname string
	TContact string

}

type User struct {
	First *NewInventory
	Name string
	Id int
	Surname string
	Contact string
	SecondName string
	SecondId int
	SecondSurname string
	SecondContact string
	TName string
	TSurname string
	TContact string
}

func (u User)GetId( id int) int {
	return 2+id
}
func (u User) HasPermission(deviceId string) bool {
	if deviceId=="" {
		return true
	} else{
		return false
	}
}
