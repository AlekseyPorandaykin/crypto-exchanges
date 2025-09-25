package client

type ExchangeConfig interface {
	Validate() error // Метод для проверки, что все обязательные поля заполнены
	Name() string    // Метод для получения названия биржи
}
