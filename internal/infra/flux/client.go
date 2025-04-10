package flux

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/eznd-go/flux/internal/domain"
)

const (
	loginURL       = "/?module=account&action=login&return_url="
	authCookieName = "fluxSessionData"
)

type client struct {
	baseURL    string
	serverName string
	username   string
	password   string
	cookie     http.Cookie
	loggedInAt time.Time
	client     *http.Client
}

func NewClient(serverName, serverURL, username, password string) *client {
	return &client{
		username:   username,
		password:   password,
		serverName: serverName,
		baseURL:    serverURL,
		client: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (c *client) Get(path string) ([]byte, error) {
	err := c.checkLogin()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&c.cookie)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	return content, nil
}

func (c *client) login() error {
	payload := url.Values{}
	payload.Add("server", c.serverName)
	payload.Add("username", c.username)
	payload.Add("password", c.password)
	pl := payload.Encode()

	req, err := http.NewRequest(http.MethodPost, c.baseURL+loginURL, bytes.NewBufferString(pl))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	for _, cc := range resp.Cookies() {
		if cc.Name == authCookieName {
			if cc.Value != "" {
				c.cookie = *cc
				c.loggedInAt = time.Now()
				return nil
			}
		}
	}

	return domain.ErrNoSessionCookie
}

func (c *client) checkLogin() error {
	if time.Now().Sub(c.loggedInAt) <= 60*time.Minute {
		return nil
	}

	err := c.login()
	if err != nil {
		return err
	}

	return nil
}
