output "east2_controller_info" {
  description = "IP address for the East region controller"
  value       = module.avi_controller_aws_east2.controllers
}
output "westus2_controller_info" {
  description = "IP address for the West region controller"
  value       = module.avi_controller_aws_west2.controllers
}