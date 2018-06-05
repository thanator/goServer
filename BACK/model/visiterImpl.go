package model

import (
	"fmt"
	"io/ioutil"

	"../db"
)

type ExportXmlVisitor struct{}

func (v *ExportXmlVisitor) visitBoss() {
	prd := db.CreateXMLOtchet()
	println(prd)
	err := ioutil.WriteFile("datbos.xml", []byte(prd), 0644)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (v *ExportXmlVisitor) visitManager() {
	str := db.ExportOrder()
	println(str)
	err := ioutil.WriteFile("datmanag.xml", []byte(str), 0644)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
