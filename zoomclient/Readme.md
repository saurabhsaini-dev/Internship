## Client Library for Zoom Provider
This package contains two functions

1. ```func NewClient(host, clientid, clientsecret *string) (*Client, error)```
2. ```func (c *Client) doRequest(req *http.Request) ([]byte, error)```
