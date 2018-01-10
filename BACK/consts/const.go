package consts

// константы - ну не писать же по сто раз эти одинаковые цифры,
// угадывая потом, что там было

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

var ORDER_STATUS = map[int]string{
	1: "Ожидание",
	2: "Одобрено",
	3: "Отклонено",
}

var TYPE_OF_ORDER = map[int]string{
	1: "Продажа",
	2: "Покупка",
}

var PRODUCT_STATUS = map[int]string{
	1: "Списано",
	2: "Не списано",
	4: "Просрочено",
}