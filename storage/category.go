package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func InsertCategory(db *sql.DB, category models.CreateCategory) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO category (
			id,
			name,
			updated_at
		) VALUES ($1, $2, now())
	`

	_, err := db.Exec(query,
		id,
		category.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdCategory(db *sql.DB, req models.CategoryPrimeryKey) (models.Category2, error) {

	var (
		category models.Category2
	)
	query := `
	select
		c.id,
		c.name,
		c.created_at,
		c.updated_at
	from bookandcategory as bac
	join category as c on c.id = bac.category_id
	where bac.category_id = $1`

	query1 := `
	select
		b.id,
		b.name,
		b.price,
		b.description
	from bookandcategory as bac
	join books as b on b.id = bac.book_id
	where bac.category_id = $1
	group by b.name, b.id`

	rows, err := db.Query(query, req.Id)
	if err != nil {
		return models.Category2{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		rows, err = db.Query(query1, req.Id)

		for rows.Next() {
			var book models.BookCat
			err = rows.Scan(
				&book.Id,
				&book.Name,
				&book.Price,
				&book.Description,
			)
			category.Book = append(category.Book, book)
		}
	}
	if err != nil {
		return models.Category2{}, err
	}

	return category, nil
}

func GetListCategory(db *sql.DB, req models.GetListCategoryRequest) (models.GetListCategoryResponse, error) {

	var (
		resp   models.GetListCategoryResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			created_at,
			updated_at
		FROM category
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := db.Query(query)
	if err != nil {
		return models.GetListCategoryResponse{}, err
	}

	for rows.Next() {
		var category models.Category

		err = rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return models.GetListCategoryResponse{}, err
		}

		resp.Categories = append(resp.Categories, category)
	}

	return resp, nil
}

func UpdateCategory(db *sql.DB, category models.UpdateCategory) error {

	query := `
		UPDATE 
			category 
		SET 
			name = $2,
			updated_at = now()
		WHERE id = $1
	`

	_, err := db.Exec(query,
		category.Id,
		category.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sql.DB, req models.CategoryPrimeryKey) error {
	_, err := db.Exec("DELETE FROM category WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
