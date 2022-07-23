output "PrivateIPs" {
  value = azurerm_network_interface.terraform_network_interface[*].private_ip_address
}

# output "NIC-2" {
#   value = azurerm_network_interface.terraform_network_interface_2.private_ip_address
# }

# output "NIC-3" {
#   value = azurerm_network_interface.terraform_network_interface_3.private_ip_address
# }
# # output "subnet_id" {
#   value = data.azurerm_subnet.terraform_subnet.id
# }

output "publicIPs" {
  value = azurerm_public_ip.avi_pip[*].ip_address
}
output "avi_main_controller_name" {
  value = azurerm_linux_virtual_machine.virtualmachine[0].name
}
output avi_fqdn {
  value = azurerm_public_ip.avi_pip.*.fqdn
}
