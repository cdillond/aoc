package client

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const BaseURL = "https://adventofcode.com/"

type Client struct {
	Day, Year string
	Token     string
}

func New(day, year string) (Client, error) {
	token := os.Getenv("AOC_SESSION")
	if token == "" {
		return Client{}, errors.New("unable to obtain session cookie")
	}

	return Client{
		Day:   day,
		Year:  year,
		Token: token,
	}, nil
}

func (c Client) GetInput(w io.Writer) error {
	req, err := http.NewRequest(http.MethodGet, BaseURL+c.Year+"/day/"+c.Day+"/input", nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.Token,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	return err
}

func (c Client) Submit(part, answer string, w io.Writer) error {
	form := url.Values{}
	form.Add("level", part)
	form.Add("answer", answer)

	req, err := http.NewRequest(http.MethodPost, BaseURL+c.Year+"/day/"+c.Day+"/answer", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.Token,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	return err
}
