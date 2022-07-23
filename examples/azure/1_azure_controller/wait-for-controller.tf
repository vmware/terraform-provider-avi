resource "null_resource" "wait_for_controller" {
  count = length(azurerm_linux_virtual_machine.virtualmachine)
  provisioner "local-exec" {
    command = "./wait-for-controller.sh"
    environment = {
      CONTROLLER_ADDRESS = azurerm_public_ip.avi_pip[*].ip_address
      POLL_INTERVAL      = 45
    }
  }
  depends_on = [
    azurerm_linux_virtual_machine.virtualmachine
  ]
}
