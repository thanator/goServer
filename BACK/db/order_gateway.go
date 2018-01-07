package db

import (
	"time"
	"database/sql"
)

func CreateOrder(productId int, orderVolume int, date time.Time, customerId int) {
	//TODO прописывать ручками тип ордера - покупка
}

func ReadOrder(orderId int) (*sql.Rows, error) {
	return nil, nil;
}

func DeleteOrder(orderId int) {

}
