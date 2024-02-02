package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"athorId,omitempty"`
	AuthorUsername string    `json:"authorUsername,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

func (p *Post) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("título é obrigatório e não pode estar em branco")
	}
	if p.Content == "" {
		return errors.New("conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
