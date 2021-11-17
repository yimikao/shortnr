package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"shotnr/shortener"
	"shotnr/store"

	"github.com/gin-gonic/gin"
)

type Url struct {
	Path string `json:"url"`
}

func CreateShortUrl(c *gin.Context) {
	u := Url{}
	jon, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jon, &u)
	if err != nil {
		log.Fatal(err)
	}
	if u.Path == "" {
		fmt.Println("no fotm val")
		return
	}
	id, exists := c.Get("user_id")
	if !exists {
		log.Println("needs user id")
		return
	}
	l := shortener.GenerateShortLink(u.Path, id.(string))

	store.SaveUrlMapping(l, u.Path, id.(string))
}

func RedirectUrl(c *gin.Context) {
	su := c.Param("short_url")
	lu := store.RetrieveInitialUrl(su)
	fmt.Println(lu + "man")
	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("http://%v", lu))
}
