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


variable "avi_username" {
  default = "admin"
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


