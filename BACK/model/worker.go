package model

type Worker interface {
	Accept(Visitor)
}
