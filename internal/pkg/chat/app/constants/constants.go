package constants

const (
	ErrSessionExpired      = "Срок действия сессии пользователя истек"
	LogBadRequest          = "Bad Request"
	LogInternalServerError = "Internal Server Error"
	LogSuccess             = "Success"
	LogGetMethod           = "GET"
	LogPostMethod          = "POST"
	LogOnMessageHandler    = "Сообщение в WebSocket"
	LogConnectionHandler   = "Подключение к WebSocket"
	LogUnknownFlag         = "Неизвестный flag от клиента"
	LogCloseHandler        = "Закрытие WebSocket"
	FormatData             = "15:04 02.01.2006"
	LogWSUnexpectedClose   = "WebSocket is unexpected close"
	LogWSClose             = "WebSocket is closed"
)
