package model

import (
	"fmt"
	"io/ioutil"
	"strings"

	"../db"
)

type ExportXmlVisitor struct{}

func (v *ExportXmlVisitor) visitBoss() {
	prd := db.ReadAllProducts()
	println(prd)
	err := ioutil.WriteFile("datbos", []byte(prd), 0644)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func (v *ExportXmlVisitor) visitManager() {
	prd := db.ReadAllOrderIds()
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(prd)), ","), "[]")
	println(str)
	err := ioutil.WriteFile("datmanag", []byte(str), 0644)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
