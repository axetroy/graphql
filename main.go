package graphql

import (
  "net/http"
  "io/ioutil"
  "strings"
  "encoding/json"
)

type Variables map[string]interface{}

type Client struct {
  Url string
}

type Request struct {
  OperationName string
  Query         string
  Variables     Variables
}

type Response struct {
  Data interface{}
  Errors *[]struct {
    Message string
    Locations []struct {
      Line   int
      Column int
    }
  }
  raw []byte
}

func (r *Response) GetRaw() (body []byte) {
  return r.raw
}

func New(url string) *Client {
  return &Client{
    Url: url,
  }
}

func (c *Client) Run(req Request) (res *Response, err error) {
  var (
    httpRes         *http.Response
    queryStringByte []byte
    body            []byte
  )

  res = new(Response)

  if queryStringByte, err = json.Marshal(req); err != nil {
    return
  }

  if httpRes, err = http.Post(c.Url, "application/json", strings.NewReader(string(queryStringByte))); err != nil {
    return
  }

  if body, err = ioutil.ReadAll(httpRes.Body); err != nil {
    return
  }

  res.raw = body

  if err = json.Unmarshal(body, res); err != nil {
    return
  }

  return
}
