provider "azurerm" {
  subscription_id 	= "${var.subscription_id}"
  client_id 		= "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		= "${var.tenant_id}"
}

data "azurerm_public_ip" "terraform_pip" {
  name = "Terraform-PIP"
  resource_group_name = "${var.resource_group_name}"
}

output "ipaddress" {
  value = "${data.azurerm_public_ip.terraform_pip.ip_address}"
}

data "azurerm_virtual_network" "terraform_vnet" {
  name 			= "Terraform-VNet"
  resource_group_name   = "${var.resource_group_name}"
}

output "subnets" {
  value = "${data.azurerm_virtual_network.terraform_vnet.subnets}"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.azurerm_public_ip.terraform_pip.ip_address}"
  avi_tenant     = "admin"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_cloudconnectoruser" "azure_user" {
    name = "azureclouduser"
    tenant_ref =  "${data.avi_tenant.default_tenant.id}"
    azure_userpass {
        username = "${var.azure_username}"
        password = "${var.azure_password}"
        tenant_name = "${var.tenant_id}"
    }
}

resource "avi_cloud" "azure_cloud_cfg" {
  name         = "azure_cloud_cfg"
  vtype        = "CLOUD_AZURE"
  dhcp_enabled = true
  
  azure_configuration = {
    use_managed_disks = true
    use_enhanced_ha= false
    cloud_credentials_ref = "${avi_cloudconnectoruser.azure_user.id}"
    use_azure_dns = false
    resource_group = "${var.resource_group_name}"
    location = "${var.location}"
    network_info {
        se_network_id = "${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"
        virtual_network_id = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}/providers/Microsoft.Network/virtualNetworks/${data.azurerm_virtual_network.terraform_vnet.name}"
    }
    subscription_id = "${var.subscription_id}"
  }

  license_tier = "ENTERPRISE_18"
  license_type = "LIC_CORES"
  tenant_ref   = "${data.avi_tenant.default_tenant.id}"
}

