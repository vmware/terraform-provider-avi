output "public_ip" {
 value = data.azurerm_public_ip.avi_pip.ip_address
}
