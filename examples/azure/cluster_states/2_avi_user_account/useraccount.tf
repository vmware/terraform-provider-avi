data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface-1"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}
data "azurerm_public_ip" "avi_pip" {
  name                = "${var.pip_name}-0"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}
resource "avi_useraccount" "avi_user" {
  username     = var.avi_username
  old_password = var.avi_current_password
  password     = var.avi_new_password
}
