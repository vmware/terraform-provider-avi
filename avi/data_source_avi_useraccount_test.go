package avi

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	var newPassword = password + "123"
	var useraccount = `
resource "avi_useraccount" "avi_user" {
  username     = "admin"
  old_password = "` + password + `"
  password     = "` + newPassword + `"
}
resource "avi_useraccount" "avi_user1" {
  username     = "admin"
  old_password = "` + newPassword + `"
  password     = "` + password + `"
  depends_on = [avi_useraccount.avi_user]
}
`
	return useraccount
}
