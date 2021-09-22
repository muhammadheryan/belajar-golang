package repository

import (
	"belajar-golang/belajar-database/entity"
	"context"
	"database/sql"
	"errors"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(DB *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments (email, comment) VALUES (?,?)"

	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	defer repository.DB.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)

	return comment, err

}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT * FROM comments WHERE id = ?"
	comment := entity.Comment{}

	row, err := repository.DB.QueryContext(ctx, script, id)
	if err != nil {
		return comment, err
	}
	defer row.Close()

	if row.Next() {
		row.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("ID Not Found")
	}

}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT * FROM comments"
	var comments []entity.Comment

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
