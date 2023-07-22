package main

type Category struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

func GetCategories() ([]Category, error) {
	rows, err := DB.Query(`
        SELECT *
        FROM categories c
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    categories := []Category{}

	for rows.Next() {
        var c Category

		err := rows.Scan(
            &c.ID,
            &c.Name,
		)
		if err != nil {
			return nil, err
		}
        categories = append(categories, c)
	}
	return categories, nil
}
