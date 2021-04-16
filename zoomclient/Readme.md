## Client Library for Zoom Provider
This package contains two functions

1. ```func NewClient(host, clientid, clientsecret *string) (*Client, error)```
    <br>This function will create new Client and will get Token for that client and returns reference to the client.
2. ```func (c *Client) doRequest(req *http.Request) ([]byte, error)```
    <br>This function will send request to the endpoint for Client and return response body.
