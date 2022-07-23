data "azurerm_public_ip" "avi_pip" {
  name                = "${var.pip_name}-0"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface-1"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "avi_tenant" "default_tenant" {
  name = var.avi_username
}

data "avi_cloud" "azure_cloud_cfg" {
  name = var.azure_cloud
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_pool" "azure-pool-v1" {
  name       = "azure_poolv1"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_pool" "azure-pool-v2" {
  name       = "azure_poolv2"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.azure_cloud_cfg.id
}

resource "avi_poolgroup" "azure-poolgroup" {
  name       = "azure_poolgroup"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.azure_cloud_cfg.id

  members {
    pool_ref = data.avi_pool.azure-pool-v1.id
    ratio    = 100
  }

  members {
    pool_ref = data.avi_pool.azure-pool-v2.id
    ratio    = 100
  }
}

