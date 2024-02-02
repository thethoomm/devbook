package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func (p Posts) Create(post models.Post) (uint64, error) {
	stmt, err := p.db.Prepare("INSERT INTO posts(title, content, authorId) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	response, err := stmt.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, nil
	}

	lastID, err := response.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (p Posts) FindById(postID uint64) (*models.Post, error) {
	rows, err := p.db.Query("SELECT p.*, u.username FROM posts p INNER JOIN users u ON u.id = p.authorId WHERE p.id = ?", postID)
	if err != nil {
		return &models.Post{}, err
	}
	defer rows.Close()

	var post models.Post

	if rows.Next() {
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername); err != nil {
			return &models.Post{}, err
		}
	}

	return &post, nil
}

func (p Posts) Find(userID uint64) (*[]models.Post, error) {
	rows, err := p.db.Query(`
	SELECT DISTINCT p.*, u.username 
	FROM posts P
	INNER JOIN users u ON u.id = p.authorId 
	INNER JOIN followers f ON p.authorId = f.userId
	WHERE u.id = ? or f.followerId = ?
	ORDER BY 1 DESC
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

func (p Posts) Update(postID uint64, post models.Post) error {
	stmt, err := p.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (p Posts) Delete(postID uint64) error {
	stmt, err := p.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (p Posts) FindPostByUser(userID uint64) (*[]models.Post, error) {
	rows, err := p.db.Query(`
	SELECT p.*, u.username 
	FROM posts p JOIN users u ON u.id = p.authorId 
	WHERE p.authorId = ?
	`, userID)
	if err != nil {
		return nil, err
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUsername); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

func (p Posts) LikePost(postID uint64) error {
	stmt, err := p.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (p Posts) DislikePost(postID uint64) error {
	stmt, err := p.db.Prepare(`
		UPDATE posts SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE likes 
		END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db: db}
}
