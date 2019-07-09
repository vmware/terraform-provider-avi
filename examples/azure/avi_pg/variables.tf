variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
  default     = ""
}

//This can be get after app registration from azure portal
variable "client_id" {
  description = "Enter Client ID for Application created in Azure AD"
  default     = ""
}

//This can be get after app registration from azure portal
variable "client_secret" {
  description = "Enter Client secret for Application in Azure AD"
  default     = ""
}

variable "tenant_id" {
  description = "Enter Tenant ID / Directory ID of your Azure AD. Run Get-AzureSubscription to know your Tenant ID"
  default     = ""
}

/*
// This can be removed if have authorization to create new rg
variable "resource_group_name" {
  description = "Resource group name that will contain various resources"
  default = ""
}
*/

variable "avi_username" {
  default = "admin"
}

variable "avi_password" {
}

variable "avi_version" {
}

variable "project_name" {
}

