variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
}

//This can be get after app registration from azure portal
variable "client_id" {
  description = "Enter Client ID for Application created in Azure AD"
}

//This can be get after app registration from azure portal
variable "client_secret" {
  description = "Enter Client secret for Application in Azure AD"
}

variable "tenant_id" {
  description = "Enter Tenant ID / Directory ID of your Azure AD. Run Get-AzureSubscription to know your Tenant ID"
}

variable "location" {
  description = "The default Azure region for the resource provisioning"
  default     = "westus2"
}


variable "azure_vip_subnet_ip" {
  default = "10.10.0.200"
}

variable "azure_vip_subnet_mask" {
  default = 24
}

variable "cloud_name" {
  default = "azure"
}
variable "vs_name" {
  default = "avi_vs_azure"
}

variable "avi_username" {
  default = "admin"
}
variable "floating_ip" {
  default = true
}

variable "avi_password" {
}

variable "avi_version" {
  default = "21.1.1"
}

variable "project_name" {
  default = "remoctl"
}
variable "pip_name" {
  default = "pip-ctl1"
}
variable "azure_cloud" {
  default = "azure"
}
variable "waf_name" {
  default = "custom_waf_policy"
}
# this is the new SEG that was created in step 4
variable "se_group_name" {
  default = "avidemoseg"
}
