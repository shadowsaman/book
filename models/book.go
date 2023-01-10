package models

type BookPrimeryKey struct {
	Id string `json:"id"`
}

type CreateBook struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type Book struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type Book2 struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Categories  []CategoryName
}

type CategoryName struct {
	Id   string `json:"id"`
	Name string `json:"Category"`
}

type UpdateBook struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type GetListBookRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListBookResponse struct {
	Count int64  `json:"count"`
	Books []Book `json:"books"`
}

type UpdateBookSwag struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type Empty struct{}
