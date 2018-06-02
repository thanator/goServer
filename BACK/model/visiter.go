package model

type Visitor interface {
	visitBoss(*BossWorker)
	visitManager(*ManagerWorker)
}
