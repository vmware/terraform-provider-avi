#variables required for configuration of Avi for Horizon deployment in a shared VIP with L7 and L4 Virtual Services.

variable "ipaddr_placement" {}
variable "mgmt_net" {}
variable "cloud_name" {}
variable "ip_vip" {}
variable "domain_name" {}
variable "avi_controller" {}
variable "avi_username" {}
variable "avi_password" {}
variable "avi_api_version" {}
variable "pool_server1" {}
variable "pool_server2" {}
variable "app_profile" {}
variable "horizon_cert" {}
variable "horizon_hm" {}
variable "udp_profile" {}
variable "tcp_profile" {}
variable "ip_group" {}
variable "ssl_profile" {}
variable "l4_pool" {}
variable "l4_app_profile" {}
variable "l7_pool" {}
variable "l7_vs" {}
variable "l4_vs" {}
variable "ip_vip_l4" {}
variable "l4_domain_name" {}
