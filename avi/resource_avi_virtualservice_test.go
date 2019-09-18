package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIVirtualServiceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIVirtualServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIVirtualServiceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVirtualServiceExists("avi_virtualservice.testvs"),
					resource.TestCheckResourceAttr(
						"avi_virtualservice.testvs", "name", "vs-test")),
			},
			{
				Config: testAccUpdatedAVIVirtualServiceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIVirtualServiceExists("avi_virtualservice.testvs"),
					resource.TestCheckResourceAttr(
						"avi_virtualservice.testvs", "name", "vs-abc")),
			},
		},
	})

}

func testAccCheckAVIVirtualServiceExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI VirtualService ID is set")
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

func testAccCheckAVIVirtualServiceDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_virtualservice" {
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
			return fmt.Errorf("AVI VirtualService still exists")
		}
	}
	return nil
}

const testAccAVIVirtualServiceConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

data "avi_serviceenginegroup" "se_group" {
	name = "Default-Group"
}
  
data "avi_applicationprofile" "system_https_profile" {
	name= "System-Secure-HTTP"
}

data "avi_networkprofile" "system_tcp_profile" {
	name= "System-TCP-Proxy"
}
  
data "avi_analyticsprofile" "system_analytics_profile" {
	name= "System-Analytics-Profile"
}
  
data "avi_sslkeyandcertificate" "system_default_cert" {
	name= "System-Default-Cert"
}
  
data "avi_sslprofile" "system_standard_sslprofile" {
	name= "System-Standard"
}
  
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}

resource "avi_vsvip" "test_vsvip" {
	name= "vip-42"
	vip {
	  vip_id= "0"
	  ip_address {
		type= "V4"
		addr= "10.90.64.88"
	  }
      enabled= true
	}
	tenant_ref= data.avi_tenant.default_tenant.id
	cloud_ref= data.avi_cloud.default_cloud.id
	vrf_context_ref= data.avi_vrfcontext.global_vrf.id
  }
  
resource "avi_virtualservice" "testvs" {
	name= "vs-test"
	tenant_ref= data.avi_tenant.default_tenant.id
	cloud_ref= data.avi_cloud.default_cloud.id
	application_profile_ref= data.avi_applicationprofile.system_https_profile.id
	network_profile_ref = data.avi_networkprofile.system_tcp_profile.id
	vsvip_ref = avi_vsvip.test_vsvip.id
	vip {
	  vip_id= "0"
	  ip_address {
		type= "V4"
		addr= "10.90.64.88"
	  }
      enabled= true
	}
	services {
	  port= 80
	  enable_ssl= true
	  port_range_end= 80
	}
	se_group_ref= data.avi_serviceenginegroup.se_group.id
	analytics_profile_ref= data.avi_analyticsprofile.system_analytics_profile.id
	ssl_key_and_certificate_refs= [data.avi_sslkeyandcertificate.system_default_cert.id]
	ssl_profile_ref= data.avi_sslprofile.system_standard_sslprofile.id
	vrf_context_ref= data.avi_vrfcontext.global_vrf.id
  }
`

const testAccUpdatedAVIVirtualServiceConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}

data "avi_serviceenginegroup" "se_group" {
	name = "Default-Group"
}

data "avi_applicationprofile" "system_https_profile" {
	name= "System-Secure-HTTP"
}

data "avi_networkprofile" "system_tcp_profile" {
	name= "System-TCP-Proxy"
}

data "avi_analyticsprofile" "system_analytics_profile" {
	name= "System-Analytics-Profile"
}

data "avi_sslkeyandcertificate" "system_default_cert" {
	name= "System-Default-Cert"
}

data "avi_sslprofile" "system_standard_sslprofile" {
	name= "System-Standard"
}

data "avi_vrfcontext" "global_vrf" {
	name= "global"
}

resource "avi_vsvip" "test_vsvip" {
	name= "vip-42"
	vip {
	  vip_id= "0"
	  ip_address {
		type= "V4"
		addr= "10.90.64.88"
	  }
      enabled= true
	}
	tenant_ref= data.avi_tenant.default_tenant.id
	cloud_ref= data.avi_cloud.default_cloud.id
	vrf_context_ref= data.avi_vrfcontext.global_vrf.id
  }

resource "avi_virtualservice" "testvs" {
	name= "vs-abc"
	tenant_ref= data.avi_tenant.default_tenant.id
	cloud_ref= data.avi_cloud.default_cloud.id
	application_profile_ref= data.avi_applicationprofile.system_https_profile.id
	network_profile_ref = data.avi_networkprofile.system_tcp_profile.id
	vsvip_ref = avi_vsvip.test_vsvip.id
	vip {
	  vip_id= "0"
	  ip_address {
		type= "V4"
		addr= "10.90.64.88"
	  }
      enabled= true
	}
	services {
	  port= 80
	  enable_ssl= true
	  port_range_end= 80
	}
	se_group_ref= data.avi_serviceenginegroup.se_group.id
	analytics_profile_ref= data.avi_analyticsprofile.system_analytics_profile.id
	ssl_key_and_certificate_refs= [data.avi_sslkeyandcertificate.system_default_cert.id]
	ssl_profile_ref= data.avi_sslprofile.system_standard_sslprofile.id
	vrf_context_ref= data.avi_vrfcontext.global_vrf.id
  }
`
