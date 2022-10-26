data "azurerm_public_ip" "avi_pip" {
  name                = "${var.pip_name}-0"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}


data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface-1"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

data "avi_applicationprofile" "system_http_profile" {
  name = "System-HTTP"
}

data "avi_applicationprofile" "system_https_profile" {
  name = "System-Secure-HTTP"
}

data "avi_tenant" "default_tenant" {
  name = var.avi_username
}

data "avi_cloud" "azure_cloud_cfg" {
  name = var.cloud_name
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_networkprofile" "system_tcp_profile" {
  name = "System-TCP-Proxy"
}

data "avi_analyticsprofile" "system_analytics_profile" {
  name = "System-Analytics-Profile"
}

data "avi_sslkeyandcertificate" "system_default_cert" {
  name = "System-Default-Cert"
}

data "avi_sslprofile" "system_standard_sslprofile" {
  name = "System-Standard"
}

data "avi_serviceenginegroup" "se_group" {
  name      = var.se_group_name
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_poolgroup" "azure-poolgroup" {
  name       = "azure_poolgroup"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_wafpolicy" "vs_wafpolicy" {
   name = var.waf_name
}
data "avi_wafpolicy" "template_wafpolicy" {
  name = "Template-WAF-Policy-Standard"
}

