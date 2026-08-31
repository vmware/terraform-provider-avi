# AVI Controller Deployment on AWS Terraform module
This Terraform module creates and configures an AVI (NSX Advanced Load-Balancer) Controller on AWS

## Module Functions
The module is meant to be modular and can create all or none of the prerequiste resources needed for the AVI AWS Deployment including:
* VPC and Subnets for the Controller and SEs (configured with create_networking variable)
* IAM Roles, Policy, and Instance Profile (configured with create_iam variable)
* Security Groups for AVI Controller and SE communication
* AWS EC2 Instance using an official AVI AMI
* High Availability AVI Controller Deployment (configured with controller_ha variable)

During the creation of the Controller instance the following initialization steps are performed:
* Copy Ansible playbook to controller using the assigned public IP
* Run Ansible playbook to configure initial settings and AWS Full Access Cloud

The Ansible playbook can optionally add these configurations:
* Create Avi DNS Profile (configured with configure_dns_profile and dns_service_domain variables)
* Create Avi DNS Virtual Service (configured with configure_dns_vs and dns_vs_settings variables)
* Configure GSLB (configured with configure_gslb, gslb_site_name, gslb_domains, and configure_gslb_additional_sites variables)

## Usage
This is an example of a controller deployment that leverages an existing VPC (with a cidr_block of 10.154.0.0/16) and 3 subnets. The public key is already created in EC2 and the private key found in the "/home/<user>/.ssh/id_rsa" will be used to copy and run the Ansible playbook to configure the Controller.
```hcl
terraform {
  backend "local" {
  }
}
module "avi_controller_aws" {
  source  = "slarimore02/avi-controller-aws/aws"
  version = "1.0.x"

  region = "us-west-1"
  aws_access_key = "<access-key>"
  aws_secret_key = "<secret-key>"
  create_networking = "false"
  create_iam = "false"
  avi_version = "20.1.5"
  custom_vpc_id = "vpc-<id>"
  custom_subnet_ids = ["subnet-<id>","subnet-<id>","subnet-<id>"]
  avi_cidr_block = "10.154.0.0/16"
  controller_password = "<newpassword>"
  key_pair_name = "<key>"
  private_key_path = "/home/<user>/.ssh/id_rsa"
  name_prefix = "<name>"
  custom_tags = { "Role" : "Avi-Controller", "Owner" : "admin", "Department" : "IT", "shutdown_policy" : "noshut" }
}
output "controller_info" {
  value = module.avi_controller_aws.controllers
}
```
## GSLB Deployment Example
The example below shows a GSLB deployment with 2 regions utilized.
```hcl
terraform {
  backend "local" {
  }
}
module "avi_controller_aws_west2" {
  source                = "slarimore02/avi-controller-aws/aws"
  version               = "1.0.x"

  region                = "us-west-2"
  aws_access_key        = "<access-key>"
  aws_secret_key        = "<secret-key>"
  create_networking     = "false"
  create_iam            = "false"
  controller_ha         = true
  avi_version           = "20.1.6"
  custom_vpc_id         = "vpc-<id>"
  custom_subnet_ids     = ["subnet-<id>","subnet-<id>","subnet-<id>"]
  avi_cidr_block        = "10.154.0.0/16"
  controller_password   = "<newpassword>"
  key_pair_name         = "<key>"
  private_key_path      = "/home/<user>/.ssh/id_rsa"
  name_prefix           = "<name>"
  custom_tags           = { "Role" : "Avi-Controller", "Owner" : "admin", "Department" : "IT" }
  se_ha_mode            = "active/active"
  configure_dns_profile = "true"
  dns_service_domain    = "west1.avidemo.net"
  configure_dns_vs      = "true"
  dns_vs_settings       = { allocate_public_ip = "true", subnet_name           = "companyname-avi-subnet" }
}
module "avi_controller_aws_east1" {
  source  = "slarimore02/avi-controller-aws/aws"
  version = "1.0.x"

  region                          = "us-east-1"
  aws_access_key                  = "<access-key>"
  aws_secret_key                  = "<secret-key>"
  create_networking               = "false"
  create_iam                      = "false"
  controller_ha                   = true
  avi_version                     = "20.1.6"
  custom_vpc_id                   = "vpc-<id>"
  custom_subnet_ids               = ["subnet-<id>","subnet-<id>","subnet-<id>"]
  avi_cidr_block                  = "10.155.0.0/16"
  controller_password             = "<newpassword>"
  key_pair_name                   = "<key>"
  private_key_path                = "/home/<user>/.ssh/id_rsa"
  name_prefix                     = "<name>"
  custom_tags                     = { "Role" : "Avi-Controller", "Owner" : "admin", "Department" : "IT", "shutdown_policy" : "noshut" }
  se_ha_mode                      = "active/active"
  configure_dns_profile           = "true"
  dns_service_domain              = "east1.avidemo.net"
  configure_dns_vs                = "true"
  dns_vs_settings                 = { allocate_public_ip = "true", subnet_name = "companyname-avi-subnet" }
  configure_gslb                  = "true"
  gslb_site_name                  = "East1"
  gslb_domains                    = ["gslb.avidemo.net"]
  configure_gslb_additional_sites = "true"
  additional_gslb_sites           = [{name = "West2", ip_address_list = module.avi_controller_aws_west2.controllers[*].private_ip_address, dns_vs_name = "DNS-VS"}]
}
output "east1_controller_info" {
  value = module.avi_controller_aws_east1.controllers
}
output "westus2_controller_info" {
  value = module.avi_controller_aws_west2.controllers
}
```
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13.6 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 3.25.0 |
| <a name="requirement_null"></a> [null](#requirement\_null) | 3.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 3.25.0 |
| <a name="provider_null"></a> [null](#provider\_null) | 3.0.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_ec2_tag.custom_controller_1](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_tag) | resource |
| [aws_ec2_tag.custom_controller_2](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_tag) | resource |
| [aws_ec2_tag.custom_controller_3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ec2_tag) | resource |
| [aws_iam_instance_profile.avi](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_instance_profile) | resource |
| [aws_iam_role.avi](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy.avi_asg_notification](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_autoscaling](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_ec2](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_iam](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_kms](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_r53](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_sns](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_iam_role_policy.avi_sqs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_instance.avi_controller](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |
| [aws_internet_gateway.avi](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/internet_gateway) | resource |
| [aws_route.default_route](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route) | resource |
| [aws_security_group.avi_controller_sg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_security_group.avi_data_sg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_security_group.avi_se_mgmt_sg](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_subnet.avi](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet) | resource |
| [aws_vpc.avi](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [null_resource.ansible_provisioner](https://registry.terraform.io/providers/hashicorp/null/3.0.0/docs/resources/resource) | resource |
| [aws_availability_zones.azs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_subnet.custom](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/subnet) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_additional_gslb_sites"></a> [additional\_gslb\_sites](#input\_additional\_gslb\_sites) | The Names and IP addresses of the GSLB Sites that will be configured. | `list(object({ name = string, ip_address_list = list(string), dns_vs_name = string }))` | <pre>[<br>  {<br>    "dns_vs_name": "DNS-VS",<br>    "ip_address_list": [<br>      ""<br>    ],<br>    "name": ""<br>  }<br>]</pre> | no |
| <a name="input_avi_cidr_block"></a> [avi\_cidr\_block](#input\_avi\_cidr\_block) | The CIDR that will be used for creating a subnet in the AVI VPC - a /16 should be provided | `string` | `"10.255.0.0/16"` | no |
| <a name="input_avi_version"></a> [avi\_version](#input\_avi\_version) | The AVI Controller version that will be deployed | `string` | n/a | yes |
| <a name="input_aws_access_key"></a> [aws\_access\_key](#input\_aws\_access\_key) | The Access Key that will be used to deploy AWS resources | `string` | n/a | yes |
| <a name="input_aws_secret_key"></a> [aws\_secret\_key](#input\_aws\_secret\_key) | The Secret Key that will be used to deploy AWS resources | `string` | n/a | yes |
| <a name="input_boot_disk_size"></a> [boot\_disk\_size](#input\_boot\_disk\_size) | The boot disk size for the Avi controller | `number` | `128` | no |
| <a name="input_configure_dns_profile"></a> [configure\_dns\_profile](#input\_configure\_dns\_profile) | Configure Avi DNS Profile for DNS Record Creation for Virtual Services. If set to true the dns\_service\_domain variable must also be set | `bool` | `"false"` | no |
| <a name="input_configure_dns_route_53"></a> [configure\_dns\_route\_53](#input\_configure\_dns\_route\_53) | Configures Avi Cloud with Route53 DNS Provider. The following variables must be set to false if enabled: configure\_dns\_profile, configure\_dns\_vs, configure\_gslb | `bool` | `"false"` | no |
| <a name="input_configure_dns_vs"></a> [configure\_dns\_vs](#input\_configure\_dns\_vs) | Create Avi DNS Virtual Service. The configure\_dns\_profile variable must also be set to true | `bool` | `"false"` | no |
| <a name="input_configure_gslb"></a> [configure\_gslb](#input\_configure\_gslb) | Configure GSLB. The gslb\_site\_name, gslb\_domains, and configure\_dns\_vs variables must also be set. Optionally the additional\_gslb\_sites variable can be used to add active GSLB sites | `bool` | `"false"` | no |
| <a name="input_configure_gslb_additional_sites"></a> [configure\_gslb\_additional\_sites](#input\_configure\_gslb\_additional\_sites) | Configure Additional GSLB Sites. The additional\_gslb\_sites, gslb\_site\_name, gslb\_domains, and configure\_dns\_vs variables must also be set. Optionally the additional\_gslb\_sites variable can be used to add active GSLB sites | `bool` | `"false"` | no |
| <a name="input_controller_ha"></a> [controller\_ha](#input\_controller\_ha) | If true a HA controller cluster is deployed and configured | `bool` | `"false"` | no |
| <a name="input_controller_password"></a> [controller\_password](#input\_controller\_password) | The password that will be used authenticating with the AVI Controller. This password be a minimum of 8 characters and contain at least one each of uppercase, lowercase, numbers, and special characters | `string` | n/a | yes |
| <a name="input_controller_public_address"></a> [controller\_public\_address](#input\_controller\_public\_address) | This variable controls if the Controller has a Public IP Address. When set to false the Ansible provisioner will connect to the private IP of the Controller. | `bool` | `"false"` | no |
| <a name="input_create_iam"></a> [create\_iam](#input\_create\_iam) | Create IAM Service Account, Roles, and Role Bindings for Avi GCP Full Access Cloud | `bool` | `"false"` | no |
| <a name="input_create_networking"></a> [create\_networking](#input\_create\_networking) | This variable controls the VPC and subnet creation for the AVI Controller. When set to false the custom-vpc-name and custom-subnetwork-name must be set. | `bool` | `"true"` | no |
| <a name="input_custom_subnet_ids"></a> [custom\_subnet\_ids](#input\_custom\_subnet\_ids) | This field can be used to specify a list of existing VPC Subnets for the controller and SEs. The create-networking variable must also be set to false for this network to be used. | `list(string)` | `null` | no |
| <a name="input_custom_tags"></a> [custom\_tags](#input\_custom\_tags) | Custom tags added to AWS Resources created by the module | `map(string)` | `{}` | no |
| <a name="input_custom_vpc_id"></a> [custom\_vpc\_id](#input\_custom\_vpc\_id) | This field can be used to specify an existing VPC for the controller and SEs. The create-networking variable must also be set to false for this network to be used. | `string` | `null` | no |
| <a name="input_dns_search_domain"></a> [dns\_search\_domain](#input\_dns\_search\_domain) | The optional DNS search domain that will be used by the controller | `string` | `null` | no |
| <a name="input_dns_servers"></a> [dns\_servers](#input\_dns\_servers) | The optional DNS servers that will be used for local DNS resolution by the controller. Example ["8.8.4.4", "8.8.8.8"] | `list(string)` | `null` | no |
| <a name="input_dns_service_domain"></a> [dns\_service\_domain](#input\_dns\_service\_domain) | The DNS Domain that will be available for Virtual Services. Avi will be the Authorative Nameserver for this domain and NS records may need to be created pointing to the Avi Service Engine addresses. An example is demo.Avi.com | `string` | `""` | no |
| <a name="input_dns_vs_settings"></a> [dns\_vs\_settings](#input\_dns\_vs\_settings) | Settings for the DNS Virtual Service. The subnet\_name must be an existing AWS Subnet. If the allocate\_public\_ip option is set to true a EIP will be allocated for the VS. The VS IP address will automatically be allocated via the AWS IPAM. Example:{ subnet\_name = "subnet-dns", allocate\_public\_ip = "true" } | `object({ subnet_name = string, allocate_public_ip = bool })` | `null` | no |
| <a name="input_email_config"></a> [email\_config](#input\_email\_config) | The Email settings that will be used for sending password reset information or for trigged alerts. The default setting will send emails directly from the Avi Controller | `object({ smtp_type = string, from_email = string, mail_server_name = string, mail_server_port = string, auth_username = string, auth_password = string })` | <pre>{<br>  "auth_password": "",<br>  "auth_username": "",<br>  "from_email": "admin@avicontroller.net",<br>  "mail_server_name": "localhost",<br>  "mail_server_port": "25",<br>  "smtp_type": "SMTP_LOCAL_HOST"<br>}</pre> | no |
| <a name="input_gslb_domains"></a> [gslb\_domains](#input\_gslb\_domains) | A list of GSLB domains that will be configured | `list(string)` | <pre>[<br>  ""<br>]</pre> | no |
| <a name="input_gslb_site_name"></a> [gslb\_site\_name](#input\_gslb\_site\_name) | The name of the GSLB site the deployed Controller(s) will be a member of. | `string` | `""` | no |
| <a name="input_instance_type"></a> [instance\_type](#input\_instance\_type) | The EC2 instance type for the Avi Controller | `string` | `"m5.2xlarge"` | no |
| <a name="input_key_pair_name"></a> [key\_pair\_name](#input\_key\_pair\_name) | The name of the existing EC2 Key pair that will be used to authenticate to the Avi Controller | `string` | n/a | yes |
| <a name="input_name_prefix"></a> [name\_prefix](#input\_name\_prefix) | This prefix is appended to the names of the Controller and SEs | `string` | n/a | yes |
| <a name="input_ntp_servers"></a> [ntp\_servers](#input\_ntp\_servers) | The NTP Servers that the Avi Controllers will use. The server should be a valid IP address (v4 or v6) or a DNS name. Valid options for type are V4, DNS, or V6 | `list(object({ addr = string, type = string }))` | <pre>[<br>  {<br>    "addr": "0.us.pool.ntp.org",<br>    "type": "DNS"<br>  },<br>  {<br>    "addr": "1.us.pool.ntp.org",<br>    "type": "DNS"<br>  },<br>  {<br>    "addr": "2.us.pool.ntp.org",<br>    "type": "DNS"<br>  },<br>  {<br>    "addr": "3.us.pool.ntp.org",<br>    "type": "DNS"<br>  }<br>]</pre> | no |
| <a name="input_private_key_path"></a> [private\_key\_path](#input\_private\_key\_path) | The local private key path for the EC2 Key pair used for authenticating to the Avi Controller | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | The Region that the AVI controller and SEs will be deployed to | `string` | n/a | yes |
| <a name="input_se_ha_mode"></a> [se\_ha\_mode](#input\_se\_ha\_mode) | The HA mode of the Service Engine Group. Possible values active/active, n+m, or active/standby | `string` | `"active/active"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_controller_private_addresses"></a> [controller\_private\_addresses](#output\_controller\_private\_addresses) | The Private IP Addresses allocated for the Avi Controller(s) |
| <a name="output_controllers"></a> [controllers](#output\_controllers) | The AVI Controller(s) Information |
| <a name="output_public_address"></a> [public\_address](#output\_public\_address) | Public IP Addresses for the AVI Controller(s) |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
