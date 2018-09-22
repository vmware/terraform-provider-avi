provider "openstack" {
  user_name   = "${var.openstack_username}"
  tenant_name = "${var.openstack_tenant_name}"
  password    = "${var.openstack_password}"
  auth_url    = "${var.openstack_url}"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_current_password}"
  avi_controller = "${var.avi_controller}"
  avi_tenant     = "admin"
}

resource "avi_useraccount" "avi_user"{
  username = "${var.avi_username}"
  old_password = "${var.avi_current_password}"
  password = "${var.avi_new_password}"
}

