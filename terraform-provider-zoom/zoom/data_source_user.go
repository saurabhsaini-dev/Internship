package zoom

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ListUsers struct {
	PageCount     int    `json:"page_count"`
	PageNumber    int    `json:"page_number"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	NextPageToken string `json:"next_page_token"`
	Users         []User `json:"users"`
}

type User struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Type          int    `json:"type"`
	PMI           int    `json:"pmi"`
	Verified      int    `json:"verified"`
	CreatedAt     string `json:"created_at"`
	LastLoginTime string `json:"last_login_time"`
	PicUrl        string `json:"pic_url"`
	Status        string `json:"status"`
	RoleId        string `json:"role_id"`
}

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{
			"page_count": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"page_number": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"total_records": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"next_page_token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pmi": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"verified": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"created_at": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_login_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pic_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenData(users *[]User) []interface{} {
	if users != nil {
		ois := make([]interface{}, len(*users))

		for i, user := range *users {
			oi := make(map[string]interface{})

			oi["id"] = user.Id
			oi["first_name"] = user.FirstName
			oi["last_name"] = user.LastName
			oi["email"] = user.Email
			oi["type"] = user.Type
			oi["pmi"] = user.PMI
			oi["verified"] = user.Verified
			oi["created_at"] = user.CreatedAt
			oi["last_login_time"] = user.LastLoginTime
			oi["pic_url"] = user.PicUrl
			oi["status"] = user.Status
			oi["role_id"] = user.RoleId

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

// Implement Read
func dataSourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{}
	url := "https://api.zoom.us/v2/users/"
	httpMethod := http.MethodGet
	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTg4MTA2NzQsImlhdCI6MTYxODgwNTI3Nn0.Oty8DXzssVaCcP2jdqjFTH2hQxst5CsBEyp95bkCsKI"
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	var apiResponse ListUsers
	err = json.Unmarshal(resp, &apiResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	users := flattenData(&apiResponse.Users)
	if err := d.Set("users", users); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
