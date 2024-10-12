package form

// Для каждого публичного метода отдельный файл, поскольку так легче искать и теститься

type formManager interface {
	// Здесь все методы слоя менеджеров
	// Если слой менеджера для функции содержит чисто вызов следующего слоя, то не надо
	// для него делать слой менеджера, создай отдельный интерфейс для таких методов и вызывай
	// сразу слой репы
}

type formStorer interface {
	// Слой репы, о котором написано выше
}

type Handler struct {
	formManager formManager
}

func NewFormHandler(formManager formManager) *Handler {
	return &Handler{formManager: formManager}
}
