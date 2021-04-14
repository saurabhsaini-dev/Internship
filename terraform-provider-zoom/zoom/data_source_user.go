package zoom

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource {
		ReadContext: dataSourceUsersRead, 
		Schema: map[string]*schema.Schema{},
	}
}

Schema: map[string]*schema.Schema {
	"users": &schema.Schema {
		Type:    schema.TypeList,
		Computed:true,
		Elem:  &schema.Resource {
			Schema: map[string]*schema.Schema {
				"page_count": &schema.Schema {
					Type:     schema.TypeInt,
					Computed: true
				},
				"page_number": &schema.Schema {
					Type:     schema.TypeInt,
					Computed: true
				},
				"page_size": &schema.Schema {
					Type:     schema.TypeInt,
					Computed: true
				},
				"total_records": &schema.Schema {
					Type:     schema.TypeInt,
					Computed: true
				},
				"next_page_token": &schema.Schema {
					Type:     schema.TypeString,
					Computed: true
				},
				"users": &schema.Schema {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource {
						Schema: map[string]*schema.Schema {
							"id"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"first_name"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"last_name"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"email"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"type"  : &schema.Schema {
								Type:     schema.TypeInt,
								Computed: true,
							},
							"pmi"  : &schema.Schema {
								Type:     schema.TypeInt,
								Computed: true,
							},
							"verified"  : &schema.Schema {
								Type:     schema.TypeInt,
								Computed: true,
							},
							"created_at"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"last_login_time"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"pic_url"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"status"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
							"role_id"  : &schema.Schema {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			},
		},
	},
},

// Implement Read
func dataSourceUsersRead( ctx context.Context, d *schema.ResourceData, m interface{} )  diag.Diagnostics {
	client := &http.Client { Timeout: 10 * time.Second } 

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", "https://api.zoom.us/v2/users/", nil) 
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	users := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("users", users); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}