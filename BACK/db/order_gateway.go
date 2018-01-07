package db

import (
	"time"
	"database/sql"
)

func CreateOrder(productId int, orderVolume int, date time.Time, customerId int) {
	//TODO прописывать ручками тип (покупка-продажа) ордера - покупка + статус - inProgress
}

func ReadOrder(orderId int) (*sql.Rows, error) {
	return nil, nil
}

func ReadorderWithParam(status int) (*sql.Rows, error) {
	return nil, nil
}

func UpdateOrder(orderId int, status int) {
}
