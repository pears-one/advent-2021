package advent

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strconv"
)

type client struct {
	http.Client
	host   string
	year   int
	cookie string
}

type Input []string

func (in *Input) ToInt() ([]int, error) {
	var intArr []int
	for _, s := range *in {
		i, err := strconv.Atoi(s)
		if err != nil {
			return intArr, err
		}
		intArr = append(intArr, i)
	}
	return intArr, nil
}

type Client interface {
	GetInput(day int) (*Input, error)
}

func NewClient(cookie string) (Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &client{
		Client: http.Client{
			Jar: jar,
		},
		host:   "adventofcode.com",
		year:   2021,
		cookie: cookie,
	}, nil
}

func (c *client) GetInput(day int) (*Input, error) {
	var lines Input
	method := "GET"
	address := fmt.Sprintf("https://%s/%d/day/%d/input", c.host, c.year, day)
	req, err := c.initRequest(method, address, nil)
	if err != nil {
		return &lines, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return &lines, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &lines, nil
}

func (c *client) initRequest(method, address string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, address, body)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.cookie,
	})
	return req, nil
}
