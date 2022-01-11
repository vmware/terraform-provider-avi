terraform {
  required_providers {
    avi = {
      source = "vmware/avi"
      version = "21.1.1"
    }
  }
}

provider "avi" {
  avi_username   = var.admin_username
  avi_password   = var.admin_old_password
  avi_controller = var.controller_ip
  avi_tenant     = "admin"
  avi_version = var.avi_version
}
data "avi_tenant" "default_tenant"{
    name= "admin"
	depends_on = [avi_useraccount.avi_user]
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
	depends_on = [avi_useraccount.avi_user]
}
data "avi_healthmonitor" "default_monitor" {
    name= "System-HTTP"
	depends_on = [avi_useraccount.avi_user]
}
resource "avi_useraccount" "avi_user" {
  username     = "admin"
  old_password = var.admin_old_password
  password     = var.admin_new_password
}

resource "avi_pool" "testPool2" {
	depends_on = [avi_useraccount.avi_user]
    name = "test-Pool-3"
    cloud_ref = data.avi_cloud.default_cloud.id
    tenant_ref = data.avi_tenant.default_tenant.id
    enabled = true
    health_monitor_refs = [data.avi_healthmonitor.default_monitor.id]
    description = "Test Terraform"
    fail_action {
        type = "FAIL_ACTION_CLOSE_CONN"
    }
    servers {
        ip {
            type = "V4"
            addr = "10.23.45.67"
         }
        //port = 8082
        ratio = 2
    }
    servers {
        ip {
            type = "V4"
            addr = "10.23.56.79"
         }
        ratio = 1
    }
}
