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
variable "ntp_servers" {
  type    = list(any)
  default = ["0.us.pool.ntp.org", "1.us.pool.ntp.org", "2.us.pool.ntp.org", "3.us.pool.ntp.org"]
}
variable "mail_server_tls" {
  default = false
}
variable "email" {
  default = "admin@avinetworks.com"
}
variable "mail_server" {
  default = "localhost"
}
variable "mail_server_port" {
  default = 25
}
variable "mail_type" {
  default = "SMTP_LOCAL_HOST"
}
variable "dns_servers" {
  default = ["8.8.8.8", "8.8.4.4", "1.1.1.1"]
}
variable "search_domain" {
  default = "io.local"
}
variable "banner" {
  default = "Avi Demo with Terraform"
}
