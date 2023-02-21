# Outputs for Terraform

output "controller_ips" {
  value = var.avi_controller_ips
}

output "avi_admin_password" {
  value = var.avi_password
}