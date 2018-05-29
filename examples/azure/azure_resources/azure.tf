provider "azurerm" {
  subscription_id = "${var.subscription_id}"
  client_id 		  = "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		  = "${var.tenant_id}"
}

resource "azurerm_resource_group" "terraform_resource_group" {
  name 		 = "${var.project_name}-terraform-resource-group"
  location = "${var.location}"
  tags {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

resource "azurerm_virtual_network" "terraform_virtual_network" {
  name 			          = "${var.project_name}-terraform-virtual-network"
  address_space 	    = ["${var.virtual_network_cidr}"]
  location 		        = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_resource_group.name}"
  tags {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

resource "azurerm_subnet" "terraform_subnet" {
  name 			           = "${var.project_name}-terraform-subnet"
  address_prefix 	     = "${var.subnet_cidr}"
  virtual_network_name = "${azurerm_virtual_network.terraform_virtual_network.name}"
  resource_group_name  = "${azurerm_resource_group.terraform_resource_group.name}"
}

/*
data "azurerm_subnet" "terraform_subnet" {
  name                 = ""
  virtual_network_name = ""
  resource_group_name  = "${var.resource_group_name}"
}

output "subnet_id" {
  value = "${data.azurerm_subnet.terraform_subnet.id}"
}
*/

resource "azurerm_network_interface" "terraform_network_interface" {
  name 		            = "${var.project_name}-terraform-network-interface"
  location 	          = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_resource_group.name}"
  //resource_group_name = "${var.resource_group_name}"
  ip_configuration {
    name 			                    = "${var.project_name}-terraform-private-ip-configuration"
    subnet_id 			              = "${azurerm_subnet.terraform_subnet.id}"
    //subnet_id 			              = "${data.azurerm_subnet.terraform_subnet.id}"
    private_ip_address_allocation = "dynamic"
  }
  tags {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

output "NIC" {
  value = "${azurerm_network_interface.terraform_network_interface.private_ip_address}"
}

resource "azurerm_storage_account" "terraform_storage_account" {
  name 			               = "${var.project_name}tfstorageaccount"
  resource_group_name 	   = "${azurerm_resource_group.terraform_resource_group.name}"
  //resource_group_name = "${var.resource_group_name}"
  location 		             = "${var.location}"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  tags {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

resource "azurerm_storage_container" "terraform_storage_container" {
  name 			            = "${var.project_name}-terraform-storage-container"
  resource_group_name 	= "${azurerm_resource_group.terraform_resource_group.name}"
  //resource_group_name = "${var.resource_group_name}"
  storage_account_name 	= "${azurerm_storage_account.terraform_storage_account.name}"
  container_access_type = "private"
}

resource "azurerm_virtual_machine" "terraform_controller" {
  name                          = "${var.project_name}-terraform-controller"
  location                      = "${var.location}"
  resource_group_name           = "${azurerm_resource_group.terraform_resource_group.name}"
  //resource_group_name = "${var.resource_group_name}"
  network_interface_ids         = ["${azurerm_network_interface.terraform_network_interface.id}"]
  vm_size                       = "Standard_F4s"
  delete_os_disk_on_termination = true
  storage_image_reference {
    publisher = "avi-networks"
    offer     = "avi-vantage-adc"
    sku       = "avi-vantage-adc-byol"
    version   = "latest"
  }
  plan {
    name      = "avi-vantage-adc-byol"
    publisher = "avi-networks"
    product   = "avi-vantage-adc"
  }
  storage_os_disk {
    name          = "${var.project_name}-osdisk"
    vhd_uri       = "${azurerm_storage_account.terraform_storage_account.primary_blob_endpoint}${azurerm_storage_container.terraform_storage_container.name}/${var.project_name}-osdisk.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }
  os_profile {
    computer_name  = "${var.project_name}-ubuntu"
    admin_username = "${var.vm_username}"
    admin_password = "${var.vm_password}"
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
  tags {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

