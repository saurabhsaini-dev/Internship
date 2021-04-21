package zoom

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Userinfo struct {
	Email     string `json:"email"`
	Type      int    `json:"type"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserRequest struct {
	Action   string   `json:"action"`
	UserInfo Userinfo `json:"user_info"`
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceOrderRead,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"firstname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lastname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func handleRequest(url string, httpMethod string, body []byte) (response []byte, err error) {
	httpClient := &http.Client{}

	var req *http.Request
	req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTg5NDM0NDksImlhdCI6MTYxODkzODA1MH0.Bq5i2gbzKKayF3Xj91UqhKtjzQUpngutYCFF_y2jNH8"

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	email := d.Get("email").(string)
	fn := d.Get("firstname").(string)
	ln := d.Get("lastname").(string)

	createUserRequest := CreateUserRequest{
		Action: "create",
		UserInfo: Userinfo{
			Email:     email,
			Type:      1,
			FirstName: fn,
			LastName:  ln,
		},
	}

	url := "https://api.zoom.us/v2/users/"
	httpMethod := http.MethodPost

	reqBody, err := json.Marshal(createUserRequest)

	if err != nil {
		return diag.FromErr(err)
	}

	_, err = handleRequest(url, httpMethod, reqBody)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(1))

	return diags
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
