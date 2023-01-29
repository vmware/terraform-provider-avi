variable "avi_version" {
  default = "22.1.2"
}
variable "avi_controller_ips" {
}

variable "avi_license" {
  default = "ENTERPRISE"
}
variable "avi_default_license_tier" {
  default = "ENTERPRISE"
# default = "ENTERPRISE_WITH_CLOUD_SERVICES"
}

variable "avi_tenant" {
  default = "admin"
}
variable "avi_password" {}
variable "avi_dns_server_ips" {
  default = "10.79.16.132, 10.79.16.133"
}
#max 3 NTP servers
variable "avi_ntp_server_ips" {
  default = "10.10.10.32, 10.10.10.33"
}
variable "ntp_servers" {
  type    = list(any)
  default = ["0.us.pool.ntp.org", "1.us.pool.ntp.org", "2.us.pool.ntp.org", "3.us.pool.ntp.org"]
}
variable "dns_servers" {
  default = ["10.79.16.132", "10.79.16.133", "192.168.100.10"]
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
variable "banner" {
  default = "Avi Demo with Terraform"
}
variable "search_domain" {
  default = "io.local"
}
