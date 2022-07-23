data "azurerm_network_interface" "controller_nic_1" {
  name                = "${var.project_name}-terraform-network-interface-1"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

data "azurerm_public_ip" "avi_pip" {
  name                = "${var.pip_name}-0"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "avi_tenant" "default_tenant" {
  name = var.avi_tenant
}

data "avi_cloud" "azure_cloud_cfg" {
  name = var.azure_cloud
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_healthmonitor" "system_http_healthmonitor" {
  name = "System-HTTP"
}
