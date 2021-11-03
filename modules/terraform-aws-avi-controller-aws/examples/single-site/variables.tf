variable "avi_version" {
  description = "The AVI Controller version that will be deployed"
  type        = string
}
variable "key_pair_name" {
  description = "The name of the existing EC2 Key pair that will be used to authenticate to the Avi Controller"
  type        = string
}
variable "private_key_path" {
  description = "The local private key path for the EC2 Key pair used for authenticating to the Avi Controller"
  type        = string
  sensitive   = true
}
variable "controller_password" {
  description = "The password that will be used authenticating with the Avi Controller. This password be a minimum of 8 characters and contain at least one each of uppercase, lowercase, numbers, and special characters"
  type        = string
  sensitive   = true
  validation {
    condition     = length(var.controller_password) > 7
    error_message = "The controller_password value must be more than 8 characters and contain at least one each of uppercase, lowercase, numbers, and special characters."
  }
}
variable "name_prefix" {
  description = "This prefix is appended to the names of the Controller and SEs"
  type        = string
}
variable "create_networking" {
  description = "This variable controls the VPC and subnet creation for the Avi Controller. When set to false the custom_vpc_name and custom_subnetwork_name must be set."
  type        = bool
  default     = false
}
variable "create_iam" {
  description = "Create IAM Roles and Role Bindings necessary for the Avi GCP Full Access Cloud. If not set the Roles and permissions in this document must be associated with the controller service account - https://Avinetworks.com/docs/latest/gcp-full-access-roles-and-permissions/"
  type        = bool
  default     = false
}
variable "controller_ha" {
  description = "If true a HA controller cluster is deployed and configured"
  type        = bool
  default     = false
}
variable "controller_public_address" {
  description = "This variable controls if the Controller has a Public IP Address. When set to false the Ansible provisioner will connect to the private IP of the Controller."
  type        = bool
  default     = true
}
variable "configure_dns_profile" {
  description = "Configure Avi DNS Profile for DNS Record Creation for Virtual Services. If set to true the dns_service_domain variable must also be set"
  type        = bool
  default     = false
}
variable "dns_service_domain" {
  description = "The DNS Domain that will be available for Virtual Services. Avi will be the Authorative Nameserver for this domain and NS records may need to be created pointing to the Avi Service Engine addresses. An example is demo.Avi.com"
  type        = string
  default     = ""
}
variable "configure_dns_vs" {
  description = "Create DNS Virtual Service. The configure_dns_profile and configure_ipam_profile variables must be set to true and their associated configuration variables must also be set"
  type        = bool
  default     = false
}
variable "aws_secret_key" {
  description = "The AWS Secret Key that will be used by Terraform and the Avi Controller resources"
  type        = string
  sensitive   = true
  default     = null
}
variable "aws_access_key" {
  description = "The AWS Access Key that will be used by Terraform and the Avi Controller resources"
  type        = string
  sensitive   = true
  default     = null
}
variable "avi_cidr_block" {
  description = "The CIDR that will be used for creating a subnet in the AVI VPC - a /16 should be provided "
  type        = string
  default     = "10.255.0.0/16"
}
variable "custom_vpc_id" {
  description = "This field can be used to specify an existing VPC for the controller and SEs. The create-networking variable must also be set to false for this network to be used."
  type        = string
  default     = null
}
variable "custom_subnet_ids" {
  description = "This field can be used to specify a list of existing VPC Subnets for the controller and SEs. The create-networking variable must also be set to false for this network to be used."
  type        = list(string)
  default     = null
}
variable "dns_vs_settings" {
  description = "Settings for the DNS Virtual Service. The subnet_name must be an existing AWS Subnet. If the allocate_public_ip option is set to true a EIP will be allocated for the VS. The VS IP address will automatically be allocated via the AWS IPAM. Example:{ subnet_name = \"subnet-dns\", allocate_public_ip = \"true\" }"
  type        = object({ subnet_name = string, allocate_public_ip = bool })
  default     = null
}