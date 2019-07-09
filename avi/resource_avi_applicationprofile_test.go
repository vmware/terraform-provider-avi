package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIApplicationProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIApplicationProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIApplicationProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIApplicationProfileExists("avi_applicationprofile.testApplicationProfile"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "name", "test-System-Secure-HTTP-abc"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_port", "false"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_ip", "false"),
				),
			},
			{
				Config: testAccAVIApplicationProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIApplicationProfileExists("avi_applicationprofile.testApplicationProfile"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "name", "test-System-Secure-HTTP-updated"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_port", "false"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_ip", "false"),
				),
			},
			{
				ResourceName:      "avi_applicationprofile.testApplicationProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIApplicationProfileConfig,
			},
		},
	})

}

func testAccCheckAVIApplicationProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ApplicationProfile ID is set")
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

func testAccCheckAVIApplicationProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_applicationprofile" {
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
			return fmt.Errorf("AVI ApplicationProfile still exists")
		}
	}
	return nil
}

const testAccAVIApplicationProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_stringgroup" "system_compressiblestringgroup" {
    name= "System-Compressible-Content-Types"
}
data "avi_stringgroup" "system_cacheablestringgroup" {
    name= "System-Cacheable-Resource-Types"
}
resource "avi_applicationprofile" "testApplicationProfile" {
	name = "test-System-Secure-HTTP-abc"
	type = "APPLICATION_PROFILE_TYPE_HTTP"
	tenant_ref = data.avi_tenant.default_tenant.id
	http_profile {
		max_rps_uri = "0"
		keepalive_header = false
		max_rps_cip_uri = "0"
		x_forwarded_proto_enabled = true
		connection_multiplexing_enabled = true
		websockets_enabled = true
		enable_request_body_buffering = false
		hsts_enabled = true
		compression_profile {
			compressible_content_ref = data.avi_stringgroup.system_compressiblestringgroup.id
			type = "AUTO_COMPRESSION"
			compression = false
			remove_accept_encoding_header = true
		}
		xff_enabled = true
		disable_keepalive_posts_msie6 = true
		keepalive_timeout = "30000"
		ssl_client_certificate_mode = "SSL_CLIENT_CERTIFICATE_NONE"
		http_to_https = true
		spdy_enabled = false
		respond_with_100_continue = true
		client_body_timeout = "30000"
		httponly_enabled = true
		hsts_max_age = "365"
		client_max_header_size = "12"
		server_side_redirect_to_https = true
		max_bad_rps_cip = "0"
		client_max_request_size = "48"
		cache_config {
			min_object_size = "100"
			query_cacheable = false
			xcache_header = true
			age_header = true
			enabled = false
			default_expire = "600"
			max_cache_size = "0"
			heuristic_expire = false
			date_header = true
			aggressive = false
			max_object_size = "4194304"
			mime_types_group_refs = [data.avi_stringgroup.system_cacheablestringgroup.id]
		}
		max_rps_unknown_uri = "0"
		ssl_everywhere_enabled = true
		spdy_fwd_proxy_mode = false
		allow_dots_in_header_name = false
		client_header_timeout = "10000"
		post_accept_timeout = "30000"
		secure_cookie_enabled = true
		max_response_headers_size = "48"
		xff_alternate_name = "X-Forwarded-For"
		max_rps_cip = "0"
		enable_fire_and_forget = false
		max_rps_unknown_cip = "0"
		max_bad_rps_cip_uri = "0"
		max_bad_rps_uri = "0"
		use_app_keepalive_timeout = false
	}
	preserve_client_port = false
	preserve_client_ip = false
}
`

const testAccAVIApplicationProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_stringgroup" "system_compressiblestringgroup" {
    name= "System-Compressible-Content-Types"
}
data "avi_stringgroup" "system_cacheablestringgroup" {
    name= "System-Cacheable-Resource-Types"
}
resource "avi_applicationprofile" "testApplicationProfile" {
	name = "test-System-Secure-HTTP-updated"
	type = "APPLICATION_PROFILE_TYPE_HTTP"
	tenant_ref = data.avi_tenant.default_tenant.id
	http_profile {
		max_rps_uri = "0"
		keepalive_header = false
		max_rps_cip_uri = "0"
		x_forwarded_proto_enabled = true
		connection_multiplexing_enabled = true
		websockets_enabled = true
		enable_request_body_buffering = false
		hsts_enabled = true
		compression_profile {
			compressible_content_ref = data.avi_stringgroup.system_compressiblestringgroup.id
			type = "AUTO_COMPRESSION"
			compression = false
			remove_accept_encoding_header = true
		}
		xff_enabled = true
		disable_keepalive_posts_msie6 = true
		keepalive_timeout = "30000"
		ssl_client_certificate_mode = "SSL_CLIENT_CERTIFICATE_NONE"
		http_to_https = true
		spdy_enabled = false
		respond_with_100_continue = true
		client_body_timeout = "30000"
		httponly_enabled = true
		hsts_max_age = "365"
		client_max_header_size = "12"
		server_side_redirect_to_https = true
		max_bad_rps_cip = "0"
		client_max_request_size = "48"
		cache_config {
			min_object_size = "100"
			query_cacheable = false
			xcache_header = true
			age_header = true
			enabled = false
			default_expire = "600"
			max_cache_size = "0"
			heuristic_expire = false
			date_header = true
			aggressive = false
			max_object_size = "4194304"
			mime_types_group_refs = [data.avi_stringgroup.system_cacheablestringgroup.id]
		}
		max_rps_unknown_uri = "0"
		ssl_everywhere_enabled = true
		spdy_fwd_proxy_mode = false
		allow_dots_in_header_name = false
		client_header_timeout = "10000"
		post_accept_timeout = "30000"
		secure_cookie_enabled = true
		max_response_headers_size = "48"
		xff_alternate_name = "X-Forwarded-For"
		max_rps_cip = "0"
		enable_fire_and_forget = false
		max_rps_unknown_cip = "0"
		max_bad_rps_cip_uri = "0"
		max_bad_rps_uri = "0"
		use_app_keepalive_timeout = false
	}
	preserve_client_port = false
	preserve_client_ip = false
}
`
