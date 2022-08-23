variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
}

variable "client_id" {
  description = "Enter Client ID for Application created in Azure AD"
}

variable "client_secret" {
  description = "Enter Client secret for Application in Azure AD"
}

variable "tenant_id" {
  description = "Enter Tenant ID / Directory ID of your Azure AD. Run Get-AzureSubscription to know your Tenant ID"
}

variable "avi_username" {
  default = "admin"
}

variable "avi_version" {
  default = "21.1.1"
}

variable "avi_current_password" {
  sensitive = true
}

variable "avi_new_password" {
  sensitive = true
}

variable "project_name" {
  default = "remoctl"
}

variable "pip_name" {
  default = "pip-ctl1"
}
