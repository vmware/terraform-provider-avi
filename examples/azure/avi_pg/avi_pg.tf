provider "azurerm" {
  subscription_id 	= "${var.subscription_id}"
  client_id 		= "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		= "${var.tenant_id}"
}


data "azurerm_public_ip" "terraform_controller" {
  name = "${var.avi_controller_name}"
  resource_group_name = "${var.resource_group_name}"
}


output "controlller_ip" {
  value = "${data.azurerm_public_ip.terraform_controller.ip_address}"
}


provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.azurerm_public_ip.terraform_controller.ip_address}"
  avi_tenant     = "admin"
  avi_version    = "17.2.7"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_cloud" "azure_cloud_cfg" {
  name = "Azure"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"
}


data "avi_pool" "azure-pool-v1" {
  name = "azure_poolv1"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"
}

data "avi_pool" "azure-pool-v2" {
  name = "azure_poolv2"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"
}

resource "avi_poolgroup" "azure-poolgroup" {
  name       = "azure_poolgroup"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"

  members = {
    pool_ref = "${data.avi_pool.azure-pool-v1.id}"
    ratio    = 100
  }

  members = {
    pool_ref = "${data.avi_pool.azure-pool-v2.id}"
    ratio    = 100
  }

}
