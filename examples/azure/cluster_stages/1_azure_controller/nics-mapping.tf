resource "azurerm_network_interface_security_group_association" "public" {
  count                     = var.controllers == 3 ? var.controllers : 1
  depends_on                = [azurerm_network_security_group.avinetworks_secgroup]
  network_interface_id      = azurerm_network_interface.terraform_network_interface[count.index].id
  network_security_group_id = azurerm_network_security_group.avinetworks_secgroup.id
}
