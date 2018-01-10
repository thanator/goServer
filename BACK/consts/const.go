package consts

const ORDER_WAITING int = 1
const ORDER_ACCEPTED int = 2
const ORDER_DECLINED int = 3
const ORDER_SELL = 1
const ORDER_BUY = 2
const (
	DB_USER     = "TanDS"
	DB_PASSWORD = "6364"
	DB_NAME     = "MilkDB"
)

var MILK_TYPE = map[int]string{
	1: "Молоко",
	2: "Кефир",
	3: "Сливки",
}

var MILK_CREATOR = map[int]string{
	1: "Домик в деревне",
	2: "36 копеек",
	3: "Простоквашино",
	4: "Яшкино",
}

var ORDER_STATUS = map[int]string{
	1: "Ожидание",
	2: "Одобрено",
	3: "Отклонено",
}

var TYPE_OF_ORDER = map[int]string{
	1: "Продажа",
	2: "Покупка",
}
