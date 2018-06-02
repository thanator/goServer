package model

type Visitor interface {
	visitBoss()
	visitManager()
}
