package zoom

import (
	"fmt"
	"regexp"
	"terraform-provider-zoom/vendor/github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"terraform-provider-zoom/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUser_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoomUserExists("zoom_user.test_user"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_user", "email", "thsaurabhsaini@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_user", "firstname", "Saurabh"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_user", "lastname", "Saini"),
				),
			},
		},
	})
}

func testAccCheckZoomUserExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No User ID is set")
		}
		name := rs.Primary.ID
		apiClient := testAccProvider.Meta().(*client.Client)
		_, err := apiClient.GetUser(name)
		if err != nil {
			return fmt.Errorf("Error fetching user with resource %s. %s", resource, err)
		}
		return nil
	}
}

func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoomUserExists("zoom_user.test_update"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_update", "email", "thsaurabhsaini@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_update", "firstname", "Saurabh"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_update", "lastname", "Saini"),
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoomUserExists("example_item.test_update"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_update", "email", "thsaurabhsaini@gmail.com"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_update", "firstname", "Saurabh"),
					resource.TestCheckResourceAttr(
						"zoom_user.test_update", "lastname", "Kumar"),
				),
			},
		},
	})
}

func testAccCheckUserDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zoom_user" {
			continue
		}

		_, err := apiClient.GetUser(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Alert, User still exists")
		}
		notFoundErr := "User not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
	resource "zoom_user" "test_user" {
		email      = "thsaurabhsaini@gmail.com" 
  		firstname  = "Saurabh"
  		lastname   = "Saini"
}
`)
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
	resource "zoom_user" "test_update" {
		email      = "thsaurabhsaini@gmail.com" 
		firstname  = "Saurabh"
		lastname   = "Saini"
}
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
	resource "zoom_user" "test_update" {
		email      = "thsaurabhsaini@gmail.com" 
		firstname  = "Saurabh"
		lastname   = "Kumar"
}
`)
}
