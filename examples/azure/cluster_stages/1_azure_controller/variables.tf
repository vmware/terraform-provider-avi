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


// This can be removed if have authorization to create new rg
variable "resource_group_name" {
  description = "Resource group name that will contain various resources"
  default     = "remoctl-rg"
}

variable "virtual_network_cidr" {
  description = "CIDR block for Virtual Network"
  default     = "10.0.0.0/8"
}

variable "subnet_cidr" {
  description = "CIDR block for Subnet within a Virtual Network"
  default     = "10.10.0.0/24"
}

variable "vm_username" {
  description = "Enter admin username to SSH into Linux VM"
  default     = "itlinux"
}

variable "vm_password" {
  description = "Enter admin password to SSH into VM"
}

variable "project_name" {
  default = "remoctl"
}

variable "project_environment" {
  default = "remoctlenv"
}

variable "login_name" {
  default = "itlinux"
}

variable "controller_ha" {
  default = "3"
}
variable "controllers" {
  default = "3"
}
variable "computer_name" {
  default = "avicontroller"
}
variable "disk_size" {
  description = "disk size in GB"
  default     = "150"
}
variable "sku_and_name" {
  default = "nsx-alb-controller-2101"
}
variable "avi_azure_version" {
  default = "latest"
  #default = "21.01.01"
}
variable "product_and_offer" {
  default = "avi-vantage-adc"
}
variable "avi_publisher" {
  default = "avi-networks"
}
variable "owner_tag" {
  default = "Remo Mattei"
}
variable "policy" {
  default = "noshut"
}
variable "storage_type" {
  default = "Premium_LRS"
}
variable "caching" {
  default = "ReadWrite"
}
# check image on az with the following command
# az vm image list --all --publisher avi-networks
variable "controller_size" {
  default = "Standard_D8s_v3"
}
variable "allowed_ip" {
  default = "172.10.163.251/32"
}
variable "dns_ip" {
  type    = list(string)
  default = ["8.8.8.8", "8.8.4.4"]
}
variable "ntp_servers" {
  type    = list(string)
  default = ["0.us.pool.ntp.org", "1.us.pool.ntp.org", "2.us.pool.ntp.org", "3.us.pool.ntp.org"]
}

variable "welcome_banner" {
  default = "Ciao and Welcome to my new Avi Controller!"
}
variable "search_default_domain" {
  default = "io.local"
}
variable "admin_password" {
  description = "Admin Password"
}
variable "azs" {
  default = ["1","2","3"]
}
variable "pip_name" {
  default = "pip-ctl1"
}
#If you are using east you will change this to east.cloudapp.azure.com if you use europe  westeurope.cloudapp.azure.com etc.. 
variable "fqdn" {
  default = "avictl"
}
