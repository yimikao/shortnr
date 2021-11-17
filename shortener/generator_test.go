package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var userId = "moyinka"

func TestShortLinkGenerator(t *testing.T) {
	initLink := "wwww.futa.edu.ng/firars/student/dashboard"
	shortLink := GenerateShortLink(initLink, userId)
	assert.NotEmpty(t, shortLink)
}
