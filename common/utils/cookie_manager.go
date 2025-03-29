package utils

import (
	"net/http"
	"time"
)

type CookieManager struct {
	writer   http.ResponseWriter
	request  *http.Request
	InSecure bool
}

func NewCookieManager(w http.ResponseWriter, r *http.Request) *CookieManager {
	return &CookieManager{
		writer:   w,
		request:  r,
		InSecure: true,
	}
}

func (c *CookieManager) SetCookie(name, value string, opts CookieOptions) {
	if c.InSecure {
		opts.Secure = false
		opts.SameSite = http.SameSiteLaxMode
	}

	http.SetCookie(c.writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     opts.Path,
		Domain:   opts.Domain,
		MaxAge:   opts.MaxAge,
		Secure:   opts.Secure,
		HttpOnly: opts.HttpOnly,
		SameSite: opts.SameSite,
	})
}

func (c *CookieManager) GetCookie(name string) (*http.Cookie, error) {
	return c.request.Cookie(name)
}

func (c *CookieManager) GetCookieValue(name string) (string, error) {
	cookie, err := c.GetCookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func (c *CookieManager) DeleteCookie(name string, opts CookieOptions) {
	if c.InSecure {
		opts.Secure = false
		opts.SameSite = http.SameSiteLaxMode
	}

	http.SetCookie(c.writer, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     opts.Path,
		Domain:   opts.Domain,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   opts.Secure,
		HttpOnly: opts.HttpOnly,
		SameSite: opts.SameSite,
	})
}

type CookieOptions struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}
