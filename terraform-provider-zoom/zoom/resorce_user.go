package zoom

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"firstname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"lastname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type

	var diags diag.Diagnostics

	httpClient := &http.Client{}
	url := "https://api.zoom.us/v2/users/"
	httpMethod := http.MethodPost
	authToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTg3NzA1NjYsImlhdCI6MTYxODc2NTE2OH0.FJcy1sv0ps1iT3eMAd7O9EIcf7adI0QjXSaYMIMkveY"

	user := d.Get("user")

	return diags
}

 
