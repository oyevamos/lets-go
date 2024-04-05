package main

import (
	"github.com/oyevamos/lets-go.git/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
