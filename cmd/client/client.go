package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const BaseURL = "https://adventofcode.com/"

func New(day, year int, token string) Client {
	return Client{
		Day:   strconv.Itoa(day),
		Year:  strconv.Itoa(year),
		Token: token,
	}
}

type Client struct {
	Day, Year string
	Token     string
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

func (c Client) Submit(part int, answer string, w io.Writer) error {
	if part != 1 && part != 2 {
		return fmt.Errorf("invalid part (%d); must be 1 or 2", part)
	}

	form := url.Values{}
	form.Add("level", strconv.Itoa(part))
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
