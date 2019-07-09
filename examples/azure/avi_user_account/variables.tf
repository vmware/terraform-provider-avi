variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
  default     = ""
}

variable "client_id" {
  description = "Enter Client ID for Application created in Azure AD"
  default     = ""
}

variable "client_secret" {
  description = "Enter Client secret for Application in Azure AD"
  default     = ""
}

variable "tenant_id" {
  description = "Enter Tenant ID / Directory ID of your Azure AD. Run Get-AzureSubscription to know your Tenant ID"
  default     = ""
}

variable "avi_username" {
  default = "admin"
}

variable "avi_version" {
}

variable "avi_current_password" {
}

variable "avi_new_password" {
}

variable "project_name" {
}

