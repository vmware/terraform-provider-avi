package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceWebhookBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSWebhookConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_webhook.testWebhook", "name", "test-wb-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSWebhookConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_webhook" "testWebhook" {
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-wb-test-abc"
}

data "avi_webhook" "testWebhook" {
    name= "${avi_webhook.testWebhook.name}"
}
`
