package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Translates []string `json:"translates"`
}

func (s *Service) OssetianTranslate(c echo.Context) error {
	word := c.QueryParams().Get("value")
	var id int
	err := s.db.QueryRow(`SELECT id FROM words WHERE word = $1`, word).Scan(&id)
	if err != nil {
		return err
	}
	res, err := s.db.Query(`SELECT translate_word FROM translates WHERE word_id = $1`, id)
	if err != nil {
		return err
	}
	defer res.Close()
	var response Response
	for res.Next() {
		var tr string
		res.Scan(&tr)
		response.Translates = append(response.Translates, tr)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Service) RussianTranslate(c echo.Context) error {
	word := c.QueryParams().Get("value")
	var id int
	err := s.db.QueryRow(`SELECT word_id FROM translates WHERE translate_word = $1`, word).Scan(&id)
	if err != nil {
		return err
	}
	res, err := s.db.Query(`SELECT word FROM words WHERE id = $1`, id)
	if err != nil {
		return err
	}
	defer res.Close()

	var response Response
	for res.Next() {
		var tr string
		res.Scan(&tr)
		response.Translates = append(response.Translates, tr)
	}

	return c.JSON(http.StatusOK, response)
}
