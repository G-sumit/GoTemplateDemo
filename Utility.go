package main

import "text/template"

func deviceEnabled(deviceId string) bool {
	if deviceId=="" {
		return true
	} else{
		return false
	}
}

func getFuncMap() template.FuncMap{
	var funMap = make(map[string]interface{})
	funMap["deviceEnabled"]= deviceEnabled
	return funMap
}

