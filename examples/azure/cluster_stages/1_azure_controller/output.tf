output "PrivateIPs" {
  value = azurerm_network_interface.terraform_network_interface[*].private_ip_address
}
output "publicIPs" {
  value = azurerm_public_ip.avi_pip[*].ip_address
}
output "avi_main_controller_name" {
  value = azurerm_linux_virtual_machine.virtualmachine[0].name
}
output "avi_fqdn" {
  value = azurerm_public_ip.avi_pip.*.fqdn
}
