resource "avi_pool" "azure-pool-v1" {
  name                = "azure_poolv1"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id
  vrf_ref             = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref           = data.avi_cloud.azure_cloud_cfg.id

  external_autoscale_groups = ["${azurerm_linux_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.project_name}-terraform-resource-group"]
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "azure-pool-v2" {
  name                = "azure_poolv2"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id
  vrf_ref             = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref           = data.avi_cloud.azure_cloud_cfg.id

  external_autoscale_groups = ["${azurerm_linux_virtual_machine_scale_set.terraform_scale_set_v2.name}@${var.project_name}-terraform-resource-group"]
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "azurerm_linux_virtual_machine_scale_set" "terraform_scale_set_v1" {
  computer_name_prefix = "ssv1"
  name                 = "${var.project_name}_scale_set_v1"
  location             = var.location
  resource_group_name  = "${var.project_name}-terraform-resource-group"
  sku                  = "Standard_F2"
  instances            = 1
  admin_username       = "adminuser"

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  network_interface {
    name    = "${var.project_name}terraformnetworkprofile"
    primary = true
    ip_configuration {
      name      = "${var.project_name}TestIPConfiguration"
      subnet_id = data.azurerm_subnet.terraform_subnet.id
      primary   = true
    }
  }

  admin_ssh_key {
    username   = "adminuser"
    public_key = file("~/.ssh/id_rsa.pub")
  }
  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  extension {
    name                 = "vmssextension"
    publisher            = "Microsoft.OSTCExtensions"
    type                 = "CustomScriptForLinux"
    type_handler_version = "1.2"
    settings             = <<SETTINGS
    {
        "commandToExecute": "apt -y install nginx"
    }
    
SETTINGS

  }

  tags = {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}


resource "azurerm_linux_virtual_machine_scale_set" "terraform_scale_set_v2" {
  computer_name_prefix = "ssv1"
  name                 = "${var.project_name}_scale_set_v2"
  location             = var.location
  resource_group_name  = "${var.project_name}-terraform-resource-group"
  sku                  = "Standard_F2"
  instances            = 1
  admin_username       = "adminuser"

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }


  network_interface {
    name    = "${var.project_name}terraformnetworkprofile"
    primary = true
    ip_configuration {
      name      = "${var.project_name}TestIPConfiguration"
      subnet_id = data.azurerm_subnet.terraform_subnet.id
      primary   = true
    }
  }

  admin_ssh_key {
    username   = "adminuser"
    public_key = file("~/.ssh/id_rsa.pub")

  }
  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  extension {
    name                 = "vmssextension"
    publisher            = "Microsoft.OSTCExtensions"
    type                 = "CustomScriptForLinux"
    type_handler_version = "1.2"
    settings             = <<SETTINGS
    {
        "commandToExecute": "apt -y install nginx"
    }
    
SETTINGS

  }

  tags = {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}
