package web

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	R *http.Request
	W http.ResponseWriter
}

func (c *Context) GetJson(req interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, req)
}

func (c *Context) SentJson(responseCode int, resp interface{}) error {
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)
	if err != nil {
		return err
	}

	c.W.WriteHeader(responseCode)
	return nil
}
func (c *Context) SuccessJson(resp interface{}) error {
	return c.SentJson(http.StatusOK, resp)
}

func (c *Context) ErrorRequestJson(resp interface{}) error {
	return c.SentJson(http.StatusBadRequest, resp)
}

func (c *Context) ServerErrorJson(resp interface{}) error {
	return c.SentJson(http.StatusInternalServerError, resp)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		R: r,
		W: w,
	}
}
