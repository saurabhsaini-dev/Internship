## Client Library for Zoom Provider
This package contains two functions

1. ```func NewClient(host, clientid, clientsecret *string) (*Client, error)```
    This function will create new Client and will get Token for that client.
2. ```func (c *Client) doRequest(req *http.Request) ([]byte, error)```
    This function will send request to the endpoint for Client
