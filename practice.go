package main

import (
	"bytes"
	"encoding/json"
	"text/template"
)

var tmpl, err = template.New("UserResourceT").Funcs(getFuncMap()).ParseFiles("../templates/UserResourceT")

var user = &User{&NewInventory{"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit",
	&SecondInventory{"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit",
		&ThirdInventory{"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit"}}},
	"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit"}

/*var userResource=&NewUserResource{"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit",
	"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit",
	"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit",
	"Sumit",1,"Gupta","9868901497","Gupta",1,"Sumit","Gupta","Sumit","Sumit","Sumit"}*/

var userResource UserResource

func GoTemplate(){
	buf := &bytes.Buffer{}
	if err != nil { panic(err) }
	err = tmpl.Execute(buf, user)
	//fmt.Print(buf.String())
	json.Unmarshal([]byte(buf.String()), &userResource)
	//fmt.Print(userResource)
	if err != nil { panic(err) }
}
func conversionDirect(){
	userResource.Id=user.Id
	userResource.Surname=user.Surname
	userResource.Contact=user.Contact
	userResource.SecondName=user.SecondName
	userResource.SecondId=user.SecondId
	userResource.SecondSurname=user.SecondSurname
	userResource.SecondContact=user.SecondContact
	userResource.TName=user.TName
	userResource.TSurname=user.TSurname
	userResource.TContact=user.TContact

	userResource.SecId=user.First.Id
	userResource.SecSurname=user.First.Surname
	userResource.SecContact=user.First.Contact
	userResource.SecSecondName=user.First.SecondName
	userResource.SecSecondId=user.First.SecondId
	userResource.SecSecondSurname=user.First.SecondSurname
	userResource.SecSecondContact=user.First.SecondContact
	userResource.SecTName=user.First.TName
	userResource.SecTSurname=user.First.TSurname
	userResource.SecTContact=user.First.TContact

	userResource.ThId=user.First.Second.Id
	userResource.ThSurname=user.First.Second.Surname
	userResource.ThContact=user.First.Second.Contact
	userResource.ThSecondName=user.First.Second.SecondName
	userResource.ThSecondId=user.First.Second.SecondId
	userResource.ThSecondSurname=user.First.Second.SecondSurname
	userResource.ThSecondContact=user.First.Second.SecondContact
	userResource.ThTName=user.First.Second.TName
	userResource.ThTSurname=user.First.Second.TSurname
	userResource.ThTContact=user.First.Second.TContact

	userResource.FoId=user.First.Second.Third.Id
	userResource.FoSurname=user.First.Second.Third.Surname
	userResource.FoContact=user.First.Second.Third.Contact
	userResource.FoSecondName=user.First.Second.Third.SecondName
	userResource.FoSecondId=user.First.Second.Third.SecondId
	userResource.FoSecondSurname=user.First.Second.Third.SecondSurname
	userResource.FoSecondContact=user.First.Second.Third.SecondContact
	userResource.FoTName=user.First.Second.Third.TName
	userResource.FoTSurname=user.First.Second.Third.TSurname
	userResource.FoTContact=user.First.Second.Third.TContact
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//fmt.Print(userResource)
	
}
func main() {
	GoTemplate()
	conversionDirect()
}

