package zoom

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	//"terraform-provider-zoom/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccZoomUserBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoomUserDestroy,
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
		userID := rs.Primary.ID
		//apiClient := testAccProvider.Meta().(*client.Client)
		_, err := handleReadRequest(userID)
		//_, err := apiClient.GetUser(name)
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
		CheckDestroy: testAccCheckZoomUserDestroy,
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
					testAccCheckZoomUserExists("zoom_user.test_update"),
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

func testAccCheckZoomUserDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zoom_user" {
			continue
		}

		userID := rs.Primary.ID

		err := deleteUser(userID)
		if err != nil {
			return fmt.Errorf("status: %v", err)
		}

		/*_, err = handleReadRequest(userID)
		if err == nil {
			return fmt.Errorf("Alert, User still exists")
		}*/
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
