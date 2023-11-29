// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createPostStmt, err = db.PrepareContext(ctx, createPost); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePost: %w", err)
	}
	if q.deletePostStmt, err = db.PrepareContext(ctx, deletePost); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePost: %w", err)
	}
	if q.dislikePostStmt, err = db.PrepareContext(ctx, dislikePost); err != nil {
		return nil, fmt.Errorf("error preparing query DislikePost: %w", err)
	}
	if q.getDislikeStmt, err = db.PrepareContext(ctx, getDislike); err != nil {
		return nil, fmt.Errorf("error preparing query GetDislike: %w", err)
	}
	if q.getLikeStmt, err = db.PrepareContext(ctx, getLike); err != nil {
		return nil, fmt.Errorf("error preparing query GetLike: %w", err)
	}
	if q.getPostByIdStmt, err = db.PrepareContext(ctx, getPostById); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostById: %w", err)
	}
	if q.likePostStmt, err = db.PrepareContext(ctx, likePost); err != nil {
		return nil, fmt.Errorf("error preparing query LikePost: %w", err)
	}
	if q.listPostsStmt, err = db.PrepareContext(ctx, listPosts); err != nil {
		return nil, fmt.Errorf("error preparing query ListPosts: %w", err)
	}
	if q.updatePostStmt, err = db.PrepareContext(ctx, updatePost); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePost: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createPostStmt != nil {
		if cerr := q.createPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPostStmt: %w", cerr)
		}
	}
	if q.deletePostStmt != nil {
		if cerr := q.deletePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePostStmt: %w", cerr)
		}
	}
	if q.dislikePostStmt != nil {
		if cerr := q.dislikePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dislikePostStmt: %w", cerr)
		}
	}
	if q.getDislikeStmt != nil {
		if cerr := q.getDislikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDislikeStmt: %w", cerr)
		}
	}
	if q.getLikeStmt != nil {
		if cerr := q.getLikeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLikeStmt: %w", cerr)
		}
	}
	if q.getPostByIdStmt != nil {
		if cerr := q.getPostByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostByIdStmt: %w", cerr)
		}
	}
	if q.likePostStmt != nil {
		if cerr := q.likePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likePostStmt: %w", cerr)
		}
	}
	if q.listPostsStmt != nil {
		if cerr := q.listPostsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPostsStmt: %w", cerr)
		}
	}
	if q.updatePostStmt != nil {
		if cerr := q.updatePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePostStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db              DBTX
	tx              *sql.Tx
	createPostStmt  *sql.Stmt
	deletePostStmt  *sql.Stmt
	dislikePostStmt *sql.Stmt
	getDislikeStmt  *sql.Stmt
	getLikeStmt     *sql.Stmt
	getPostByIdStmt *sql.Stmt
	likePostStmt    *sql.Stmt
	listPostsStmt   *sql.Stmt
	updatePostStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:              tx,
		tx:              tx,
		createPostStmt:  q.createPostStmt,
		deletePostStmt:  q.deletePostStmt,
		dislikePostStmt: q.dislikePostStmt,
		getDislikeStmt:  q.getDislikeStmt,
		getLikeStmt:     q.getLikeStmt,
		getPostByIdStmt: q.getPostByIdStmt,
		likePostStmt:    q.likePostStmt,
		listPostsStmt:   q.listPostsStmt,
		updatePostStmt:  q.updatePostStmt,
	}
}
