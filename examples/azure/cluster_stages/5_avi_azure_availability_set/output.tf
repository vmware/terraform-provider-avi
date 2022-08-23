output "azure_scale_set1" {
  value = azurerm_linux_virtual_machine_scale_set.terraform_scale_set_v1.name
}

output "azure_scale_set2" {
  value = azurerm_linux_virtual_machine_scale_set.terraform_scale_set_v2.name
}

