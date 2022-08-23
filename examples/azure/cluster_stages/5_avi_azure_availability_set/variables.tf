//This can be retreived with the az command
//to run the jq you will need to have that installed!!
//Run the command: az account list|jq  .'[]'.id
variable "subscription_id" {
  description = "Enter Subscription ID for provisioning resources in Azure"
}

//You must have an app registered to get this. You could use the az to create an app
//run the command: az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/SUBSCRIPTION_ID"
//
variable "client_id" {
  description = "Enter Client ID for Application created in Azure AD"
}

//This can be get after app registration from azure portal
variable "client_secret" {
  description = "Enter Client secret for Application in Azure AD"
}
//Run the command: az account tenant list
variable "tenant_id" {
  description = "Enter Tenant ID / Directory ID of your Azure AD. Run Get-AzureSubscription to know your Tenant ID"
}

variable "location" {
  description = "The default Azure region for the resource provisioning"
  default     = "westus2"
}

variable "azure_vm_password" {
}

variable "avi_username" {
  default = "admin"
}
variable "avi_tenant" {
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

variable "project_environment" {
  default = "remoctlenv"
}

variable "azure_cloud" {
  default = "azure"
}
