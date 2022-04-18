package comment

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//tidak ada
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "DELETE FROM comments WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, comment.Id)
	if err != nil {
		return comment, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return comment, err
	}
	if rows == 0 {
		return comment, err
	}
	return comment, nil
}

func (repository *CommentRepositoryImpl) Update(ctx context.Context, id int32, commemnt entity.Comment) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return commemnt, err
	}

	if rows.Next() {
		script := "UPDATE list_calonmertua SET email = ?, comment = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, commemnt.Email, commemnt.Comment, id)
		if err != nil {
			return commemnt, err
		}
		commemnt.Id = id
		return commemnt, nil
	} else {
		return commemnt, errors.New(("Id" + strconv.Itoa(int(id)) + " Update Failed"))
	}
}
