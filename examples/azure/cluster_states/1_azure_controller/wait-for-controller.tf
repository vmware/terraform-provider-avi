locals  {
  clusters_or_not = var.controllers == 3 ? 1 : 1
}
resource "null_resource" "wait_for_controller" {
  count = local.clusters_or_not
  #count = length(azurerm_linux_virtual_machine.virtualmachine[0])
  provisioner "local-exec" {
    command = "./wait-for-controller.sh"
    environment = {
      CONTROLLER_ADDRESS = azurerm_public_ip.avi_pip[0].ip_address
      POLL_INTERVAL      = 45
    }
  }
  depends_on = [
    azurerm_linux_virtual_machine.virtualmachine
  ]
}
