variable "avi_controller" {
  description = "Enter Controller IP"
}

variable "avi_username" {
  description = "Enter Controller Username"
}

variable "avi_password" {
  description = "Enter Controller Password"
}

variable "avi_tenant" {
  description = "Enter tenant name"
  default     = "admin"
}

variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
}

variable "resource_group" {
  description = "Enter Resource group for provisioning resources in Azure"
}

variable "virtual_network_id" {
  description = "Enter virtual network id for provisioning resources in Azure"
}

variable "azure_username" {
  description = "Enter azure username"
}

variable "azure_password" {
  description = "Enter azure password"
}

variable "azure_tenant_name" {
  description = "Enter azure tenant"
}

