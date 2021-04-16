package zoomclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HostURL - Zoom URL
const HostURL string = "https://api.zoom.us/v2/users/"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// AuthStruct
type AuthStruct struct {
	ClientId     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}

// AuthResponse
type AuthResponse struct {
	Token string `json:"token"`
}

func NewClient(host, clientid, clientsecret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Zoom URL
		HostURL: HostURL,
	}

	/* if you are using JWT then directly assign the JWT token value to the
	   Token attribute
	c.Token = "Token Value"
	*/

	// For Outh
	if (clientid != nil) && (clientsecret != nil) {
		// form request body
		rb, err := json.Marshal(AuthStruct{
			ClientId:     *clientid,
			ClientSecret: *clientsecret,
		})
		if err != nil {
			return nil, err
		}

		// authenticate
		req, err := http.NewRequest("POST", "https://zoom.us/oauth/token?grant_type=client_credentials", strings.NewReader(string(rb)))
		if err != nil {
			return nil, err
		}

		body, err := c.doRequest(req)

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.Token = ar.Token
	}

	// return the Client
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
