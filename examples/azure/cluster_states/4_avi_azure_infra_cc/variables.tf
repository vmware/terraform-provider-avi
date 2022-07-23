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
  description = "Azure region for the resource provisioning, must be smallcaps without space inbetween for se creation"
  default     = "westus2"
}
# This just lables the user we will use in azure cc
variable "azure_user" {
  default = "matteir"
}
variable "avi_username" {
  default = "admin"
}
variable "avi_password" {
}
variable "avi_version" {
  default = "21.1.1"
}
# This is the project name we assigned on step 1 -- variables
variable "project_name" {
  default = "remoctl"
}
# This is the name we have assigned on step 1 -- variables
variable "pip_name" {
  default = "pip-ctl1"
}
variable "application_azure_id" {}
variable "auth_token" {}
variable "azure_cloud_name" {
  default = "azure"
}
variable "owner_key" {
  default = "Owner"
}
variable "owner_value" {
  default = "Your Name"
}
variable "policy_key" {
  default = "shutdown_policy"
}
variable "policy_value" {
  default = "noshut"
}
variable "dept_key" {
  default = "Dept"
}
variable "dept_value" {
  default = "My TEAM"
}

variable "connection_mem_percentage" {
  default     = 50
  description = "default"
}

variable "disk_per_se" {
  default     = 10
  description = "default"
}
variable "ha_mode" {
  default = "HA_MODE_SHARED"
}
variable "instance_flavor_se" {
  default = "Standard_F8s_v2"
}
variable "max_se" {
  default     = 2
  description = ""
}
variable "max_vs_per_se" {
  default     = 20
  description = "max se for this group"
}
variable "mem_per_se" {
  default     = 2048
  description = "for WAF you will need more ram"
}
variable "vcpus_per_se" {
  default     = 2
  description = "How many CPU per SE"
}
variable "deprovision_delay" {
  default = 5
}
variable "se_prefix_seg" {
  default = "Avi_SE_AA"
}
variable "se_group_name" {
  default = "avidemoseg"
}
