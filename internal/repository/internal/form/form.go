package form

import "context"

// Здесь описываются чисто sql запросы и походы в бд,
// я бы хотел пользоваться pgx сразу, но если не успеваешь, то используй любой
// Тут чисто функции, без структуры, так будет проще, когда надо будет комбинировать

func CreateForm(ctx context.Context) error {
	const query = `
		SELECT 1`

	return nil
}
