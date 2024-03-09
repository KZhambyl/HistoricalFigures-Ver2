package data

import (
	"database/sql"
	"errors"
	"github.com/KZhambyl/HistoricalFigures/internal/validator"
	_ "github.com/lib/pq"
	"time"
)

// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Figure struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	YearsOfLife string    `json:"years_of_life"`
	Description string    `json:"description"`
	Version     int32     `json:"version"`
}

func ValidateFigure(v *validator.Validator, f *Figure) {
	v.Check(f.Name != "", "name", "must be provided")
	v.Check(len(f.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(len(f.Description) <= 1000, "desvription", "must not be more than 1000 bytes long")
	v.Check(len(f.YearsOfLife) <= 10, "years_of_life", "must not be more than 10 bytes long")
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type FigureModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the movies table.
func (m FigureModel) Insert(figure *Figure) error {
	query := `
	INSERT INTO figures (name, years_of_life, description)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, version`
	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []interface{}{figure.Name, figure.YearsOfLife, figure.Description}
	// Use the QueryRow() method to execute the SQL query on our connection pool,
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the movie struct.
	return m.DB.QueryRow(query, args...).Scan(&figure.ID, &figure.CreatedAt, &figure.Version)
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m FigureModel) Get(id int64) (*Figure, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
	SELECT id, created_at, name, years_of_life, description, version
	FROM figures WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var figure Figure

	err := m.DB.QueryRow(query, id).Scan(
		&figure.ID,
		&figure.CreatedAt,
		&figure.Name,
		&figure.YearsOfLife,
		&figure.Description,
		&figure.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &figure, nil
}

// Add a placeholder method for updating a specific record in the movies table.
func (m FigureModel) Update(figure *Figure) error {
	// Declare the SQL query for updating the record and returning the new version
	// number.
	query := `
	UPDATE figures
	SET name = $1, years_of_life = $2, description = $3, version = version + 1 WHERE id = $4 RETURNING version`
	// Create an args slice containing the values for the placeholder parameters.
	args := []interface{}{
		figure.Name,
		figure.YearsOfLife,
		figure.Description,
		figure.ID,
	}
	// Use the QueryRow() method to execute the query, passing in the args slice as a
	// variadic parameter and scanning the new version value into the movie struct.
	return m.DB.QueryRow(query, args...).Scan(&figure.Version)
}

// Add a placeholder method for deleting a specific record from the movies table.
func (m FigureModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
	DELETE FROM figures
	WHERE id = $1`
	
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
