package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVISSLProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVISSLProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISSLProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISSLProfileExists("avi_sslprofile.testSSLProfile"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "name", "test-System-Standard-abc"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "ssl_session_timeout", "86400"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "prefer_client_cipher_ordering", "false"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "enable_ssl_session_reuse", "true"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "send_close_notify", "true"),
				),
			},
			{
				Config: testAccAVISSLProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISSLProfileExists("avi_sslprofile.testSSLProfile"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "name", "test-System-Standard-updated"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "ssl_session_timeout", "86400"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "prefer_client_cipher_ordering", "false"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "enable_ssl_session_reuse", "true"),
					resource.TestCheckResourceAttr(
						"avi_sslprofile.testSSLProfile", "send_close_notify", "true"),
				),
			},
			{
				ResourceName:      "avi_sslprofile.testSSLProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVISSLProfileConfig,
			},
		},
	})

}

func testAccCheckAVISSLProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI SSLProfile ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVISSLProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_sslprofile" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI SSLProfile still exists")
		}
	}
	return nil
}

const testAccAVISSLProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_sslprofile" "testSSLProfile" {
	name = "test-System-Standard-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	ssl_session_timeout = "86400"
	accepted_ciphers = "ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-ECDSA-AES256-SHA384:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA"
	prefer_client_cipher_ordering = false
	accepted_versions {
	type = "SSL_VERSION_TLS1"
}
accepted_versions {
	type = "SSL_VERSION_TLS1_1"
}
accepted_versions {
	type = "SSL_VERSION_TLS1_2"
}
	enable_ssl_session_reuse = true
	cipher_enums = ["TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA","TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA","TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","TLS_RSA_WITH_AES_128_GCM_SHA256","TLS_RSA_WITH_AES_256_GCM_SHA384","TLS_RSA_WITH_AES_128_CBC_SHA256","TLS_RSA_WITH_AES_256_CBC_SHA256","TLS_RSA_WITH_AES_128_CBC_SHA","TLS_RSA_WITH_AES_256_CBC_SHA","TLS_RSA_WITH_3DES_EDE_CBC_SHA","TLS_AES_256_GCM_SHA384","TLS_CHACHA20_POLY1305_SHA256","TLS_AES_128_GCM_SHA256"]
	send_close_notify = true
	ssl_rating {
		performance_rating = "SSL_SCORE_EXCELLENT"
		security_score = "100.0"
		compatibility_rating = "SSL_SCORE_EXCELLENT"
	}
	dhparam = <<EOF
-----BEGIN DH PARAMETERS-----
MIIBCAKCAQEAohUmEGbnPo1dxqvGg7zslnKTZAPPNnE7l1SdTbuPbsYF83J+VDkE
pUorADcUydwdPM9nTLEk4qKGnNsbt0S+WXf6EcP0oa+rjRFXsvb4B+tD4VHGmtDA
/iivo51hKu93xaoS0xe9TjI9ZZBcirzyz3V55u/OICgNwRM6nL/Fxx3RXG3LGHP4
JF73p/kR5cNB9ebYuKYaEzkTOg6pmCyguBEBdg40br+I59rQLgzn2WMRb1bZRUzy
qSIIAMyok9/bsaxyCCsgVzkTPTtlYM9ooJzSGarlNaBhP3AerMdCV6rAQFYfP9vw
KvbiJcx+IEdHViJLl2LAFl/gYxOnIBHAswIBAg==
-----END DH PARAMETERS-----
EOF
	type = "SSL_PROFILE_TYPE_APPLICATION"
}
`

const testAccAVISSLProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_sslprofile" "testSSLProfile" {
	name = "test-System-Standard-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	ssl_session_timeout = "86400"
	accepted_ciphers = "ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-ECDSA-AES256-SHA384:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA"
	prefer_client_cipher_ordering = false
	accepted_versions {
	type = "SSL_VERSION_TLS1"
}
accepted_versions {
	type = "SSL_VERSION_TLS1_1"
}
accepted_versions {
	type = "SSL_VERSION_TLS1_2"
}
	enable_ssl_session_reuse = true
	cipher_enums = ["TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA","TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA","TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","TLS_RSA_WITH_AES_128_GCM_SHA256","TLS_RSA_WITH_AES_256_GCM_SHA384","TLS_RSA_WITH_AES_128_CBC_SHA256","TLS_RSA_WITH_AES_256_CBC_SHA256","TLS_RSA_WITH_AES_128_CBC_SHA","TLS_RSA_WITH_AES_256_CBC_SHA","TLS_RSA_WITH_3DES_EDE_CBC_SHA","TLS_AES_256_GCM_SHA384","TLS_CHACHA20_POLY1305_SHA256","TLS_AES_128_GCM_SHA256"]
	send_close_notify = true
	ssl_rating {
		performance_rating = "SSL_SCORE_EXCELLENT"
		security_score = "100.0"
		compatibility_rating = "SSL_SCORE_EXCELLENT"
	}
	dhparam = <<EOF
-----BEGIN DH PARAMETERS-----
MIIBCAKCAQEAohUmEGbnPo1dxqvGg7zslnKTZAPPNnE7l1SdTbuPbsYF83J+VDkE
pUorADcUydwdPM9nTLEk4qKGnNsbt0S+WXf6EcP0oa+rjRFXsvb4B+tD4VHGmtDA
/iivo51hKu93xaoS0xe9TjI9ZZBcirzyz3V55u/OICgNwRM6nL/Fxx3RXG3LGHP4
JF73p/kR5cNB9ebYuKYaEzkTOg6pmCyguBEBdg40br+I59rQLgzn2WMRb1bZRUzy
qSIIAMyok9/bsaxyCCsgVzkTPTtlYM9ooJzSGarlNaBhP3AerMdCV6rAQFYfP9vw
KvbiJcx+IEdHViJLl2LAFl/gYxOnIBHAswIBAg==
-----END DH PARAMETERS-----
EOF
	type = "SSL_PROFILE_TYPE_APPLICATION"
}
`
