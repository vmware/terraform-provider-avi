
variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
  default = ""
}

//This can be get after app registration from azure portal
variable "client_id" {
  description = "Enter Client ID for Application created in Azure AD"
  default = ""
}

//This can be get after app registration from azure portal
variable "client_secret" {
  description = "Enter Client secret for Application in Azure AD"
  default = ""
}

variable "tenant_id" {
  description = "Enter Tenant ID / Directory ID of your Azure AD. Run Get-AzureSubscription to know your Tenant ID"
  default = "avinetworks.com"
}

variable "location" {
  description = "The default Azure region for the resource provisioning"
  default = "EAST US"
}

// This can be removed if have authorization to create new rg
variable "resource_group_name" {
  description = "Resource group name that will contain various resources"
  default = "abhinav-group-eu"
}

variable "azure_subnet_ip" {}

variable "azure_vip_subnet_ip" {}

variable "azure_vnet" {
  default = "abhinaveuvnetperf"
}

variable "azure_vm_password" {}

variable "azure_subnet_mask" {
    default = 24
}


variable "azure_vip_subnet_mask" {
  default = 24
}

variable "avi_username" {
  default = "admin"
}

variable "avi_password" {}

variable "avi_controller_name" {
  default = "abhinavcntrl1721"
}

variable "project_name" {}