package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

func TestAVIDataSourceUserAccount(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSUserAccountConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_useraccount.avi_user", "username", "admin"),
				),
			},
		},
	})

}

func testAccAVIDSUserAccountConfig() string {
	var password = os.Getenv("AVI_PASSWORD")
	var new_password = password + "123"
	var useraccount = `
resource "avi_useraccount" "avi_user" {
  username     = "admin"
  old_password = "` + password + `"
  password     = "` + new_password + `"
}
resource "avi_useraccount" "avi_user1" {
  username     = "admin"
  old_password = "` + new_password + `"
  password     = "` + password + `"
}
`
	return useraccount
}
