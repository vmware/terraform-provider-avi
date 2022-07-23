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
  default     = "WEST US 2"
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
variable "csr_version" {
  #default  = "CRS-2021-2" # 21.1.4
  default  = "CRS-2021-1" # 21.1.3
  #default = "CRS-2020-3" #20.1.6
}

variable "waf_name" {
  default = "System-WAF-Policy"
}
variable "cloud_name" {
  default = "Default-Cloud"
}
# variable "avi_controller" {
# }
