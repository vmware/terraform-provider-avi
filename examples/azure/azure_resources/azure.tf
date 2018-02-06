provider "azurerm" {
  subscription_id 	= "${var.subscription_id}"
  client_id 		= "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		= "${var.tenant_id}"
}

resource "azurerm_resource_group" "terraform_rg" {
  name 		= "${var.resource_group_name}"
  location 	= "${var.location}"
}

resource "azurerm_virtual_network" "terraform_vnet" {
  name 			= "Terraform-VNet"
  address_space 	= ["${var.vnet_cidr}"]
  location 		= "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_subnet" "terraform_subnet" {
  name 			= "Terraform-Subnet"
  address_prefix 	= "${var.subnet_cidr}"
  virtual_network_name 	= "${azurerm_virtual_network.terraform_vnet.name}"
  resource_group_name 	= "${azurerm_resource_group.terraform_rg.name}"
}

resource "azurerm_storage_account" "terraform_storage" {
  name 			= "terraformstorage1account"
  resource_group_name 	= "${azurerm_resource_group.terraform_rg.name}"
  location 		= "${var.location}"
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_storage_container" "terraform_container" {
  name 			= "vhds"
  resource_group_name 	= "${azurerm_resource_group.terraform_rg.name}"
  storage_account_name 	= "${azurerm_storage_account.terraform_storage.name}"
  container_access_type = "private"
}

resource "azurerm_network_security_group" "terraform_web" {
  name 			= "Terraform-Web-NSG"
  location 		= "${var.location}"
  resource_group_name 	= "${azurerm_resource_group.terraform_rg.name}"

  security_rule {
	name 			= "AllowSSH"
	priority 		= 100
	direction 		= "Inbound"
	access 		        = "Allow"
	protocol 		= "Tcp"
	source_port_range       = "*"
    destination_port_range     	= "22"
    source_address_prefix      	= "*"
    destination_address_prefix 	= "*"
  }

  security_rule {
	name 			= "AllowHTTP"
	priority		= 200
	direction		= "Inbound"
	access 			= "Allow"
	protocol 		= "Tcp"
	source_port_range       = "*"
    destination_port_range     	= "80"
    source_address_prefix      	= "Internet"
    destination_address_prefix 	= "*"
  }

  security_rule {
	name 			= "AllowHTTPS"
	priority		= 201
	direction		= "Inbound"
	access 			= "Allow"
	protocol 		= "Tcp"
	source_port_range       = "*"
    destination_port_range     	= "443"
    source_address_prefix      	= "Internet"
    destination_address_prefix 	= "*"
  }

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_public_ip" "terraform_pip" {
  name 				= "Terraform-PIP"
  location 			= "${var.location}"
  resource_group_name 		= "${azurerm_resource_group.terraform_rg.name}"
  public_ip_address_allocation 	= "static"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_public_ip" "terraform_pip_web1" {
  name 				= "Terraform-PIP-WebVM1"
  location 			= "${var.location}"
  resource_group_name 		= "${azurerm_resource_group.terraform_rg.name}"
  public_ip_address_allocation 	= "static"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_public_ip" "terraform_pip_web2" {
  name 				= "Terraform-PIP-WebVM2"
  location 			= "${var.location}"
  resource_group_name 		= "${azurerm_resource_group.terraform_rg.name}"
  public_ip_address_allocation 	= "static"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_public_ip" "terraform_pip_web3" {
  name 				= "Terraform-PIP-WebVM3"
  location 			= "${var.location}"
  resource_group_name 		= "${azurerm_resource_group.terraform_rg.name}"
  public_ip_address_allocation 	= "static"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_public_ip" "terraform_pip_web4" {
  name 				= "Terraform-PIP-WebVM4"
  location 			= "${var.location}"
  resource_group_name 		= "${azurerm_resource_group.terraform_rg.name}"
  public_ip_address_allocation 	= "static"

  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_network_interface" "terraform_nic" {
  name 		      = "Terraform-NIC"
  location 	      = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_rg.name}"
  network_security_group_id = "${azurerm_network_security_group.terraform_web.id}"

  ip_configuration {
    name 			= "Terraform-ControllerPrivate"
    subnet_id 			= "${azurerm_subnet.terraform_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id	= "${azurerm_public_ip.terraform_pip.id}"
  }
  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_network_interface" "terraform_nic_web1" {
  name 		      = "Terraform-NIC-WebVM1"
  location 	      = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_rg.name}"
  network_security_group_id = "${azurerm_network_security_group.terraform_web.id}"

  ip_configuration {
    name 			= "Terraform-WebPrivateVM1"
    subnet_id 			= "${azurerm_subnet.terraform_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id	= "${azurerm_public_ip.terraform_pip_web1.id}"
  }
  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_network_interface" "terraform_nic_web2" {
  name 		      = "Terraform-NIC-WebVM2"
  location 	      = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_rg.name}"
  network_security_group_id = "${azurerm_network_security_group.terraform_web.id}"

  ip_configuration {
    name 			= "Terraform-WebPrivateVM2"
    subnet_id 			= "${azurerm_subnet.terraform_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id	= "${azurerm_public_ip.terraform_pip_web2.id}"
  }
  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_network_interface" "terraform_nic_web3" {
  name 		      = "Terraform-NIC-WebVM3"
  location 	      = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_rg.name}"
  network_security_group_id = "${azurerm_network_security_group.terraform_web.id}"

  ip_configuration {
    name 			= "Terraform-WebPrivateVM3"
    subnet_id 			= "${azurerm_subnet.terraform_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id	= "${azurerm_public_ip.terraform_pip_web3.id}"
  }
  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_network_interface" "terraform_nic_web4" {
  name 		      = "Terraform-NIC-WebVM4"
  location 	      = "${var.location}"
  resource_group_name = "${azurerm_resource_group.terraform_rg.name}"
  network_security_group_id = "${azurerm_network_security_group.terraform_web.id}"

  ip_configuration {
    name 			= "Terraform-WebPrivateVM4"
    subnet_id 			= "${azurerm_subnet.terraform_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id	= "${azurerm_public_ip.terraform_pip_web4.id}"
  }
  tags {
	group = "TerraformInternal"
  }
}

resource "azurerm_virtual_machine" "terraform_controller" {
  name                  = "Terraform-Controller"
  location              = "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"
  network_interface_ids = ["${azurerm_network_interface.terraform_nic.id}"]
  vm_size               = "Standard_F4s"

  delete_os_disk_on_termination = true

  storage_image_reference {
    publisher = "avi-networks"
    offer     = "avi-vantage-adc"
    sku       = "avi-vantage-adc-byol"
    version   = "latest"
  }
  plan {
    name= "avi-vantage-adc-byol"
    publisher= "avi-networks"
    product= "avi-vantage-adc"
  }
  
  storage_os_disk {
    name          = "osdisk"
    vhd_uri       = "${azurerm_storage_account.terraform_storage.primary_blob_endpoint}${azurerm_storage_container.terraform_container.name}/osdisk.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "ubuntuweb"
    admin_username = "${var.vm_username}"
    admin_password = "${var.vm_password}"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags {
    group = "TerraformInternal"
  }
}

resource "azurerm_virtual_machine" "terraform_webvm1" {
  name                  = "Terraform-WebVM1"
  location              = "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"
  network_interface_ids = ["${azurerm_network_interface.terraform_nic_web1.id}"]
  vm_size               = "Standard_DS1_v2"

  delete_os_disk_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
  
  storage_os_disk {
    name          = "osdisk-1"
    vhd_uri       = "${azurerm_storage_account.terraform_storage.primary_blob_endpoint}${azurerm_storage_container.terraform_container.name}/osdisk-1.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "ubuntuweb"
    admin_username = "${var.vm_username}"
    admin_password = "${var.vm_password}"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags {
    group = "TerraformInternal"
  }
}

resource "azurerm_virtual_machine" "terraform_webvm2" {
  name                  = "Terraform-WebVM2"
  location              = "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"
  network_interface_ids = ["${azurerm_network_interface.terraform_nic_web2.id}"]
  vm_size               = "Standard_DS1_v2"

  delete_os_disk_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name          = "osdisk-2"
    vhd_uri       = "${azurerm_storage_account.terraform_storage.primary_blob_endpoint}${azurerm_storage_container.terraform_container.name}/osdisk-2.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "ubuntuweb"
    admin_username = "${var.vm_username}"
    admin_password = "${var.vm_password}"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags {
    group = "TerraformInternal"
  }
}

resource "azurerm_virtual_machine" "terraform_webvm3" {
  name                  = "Terraform-WebVM3"
  location              = "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"
  network_interface_ids = ["${azurerm_network_interface.terraform_nic_web3.id}"]
  vm_size               = "Standard_DS1_v2"

  delete_os_disk_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name          = "osdisk-3"
    vhd_uri       = "${azurerm_storage_account.terraform_storage.primary_blob_endpoint}${azurerm_storage_container.terraform_container.name}/osdisk-3.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "ubuntuweb"
    admin_username = "${var.vm_username}"
    admin_password = "${var.vm_password}"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags {
    group = "TerraformInternal"
  }
}

resource "azurerm_virtual_machine" "terraform_webvm4" {
  name                  = "Terraform-WebVM4"
  location              = "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"
  network_interface_ids = ["${azurerm_network_interface.terraform_nic_web4.id}"]
  vm_size               = "Standard_DS1_v2"

  delete_os_disk_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name          = "osdisk-4"
    vhd_uri       = "${azurerm_storage_account.terraform_storage.primary_blob_endpoint}${azurerm_storage_container.terraform_container.name}/osdisk-4.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "ubuntuweb"
    admin_username = "${var.vm_username}"
    admin_password = "${var.vm_password}"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags {
    group = "TerraformInternal"
  }
}
/*
resource "azurerm_virtual_machine_extension" "cs_apache" {
  name 			= "custom_apache"
  location 		= "${var.location}"
  resource_group_name   = "${azurerm_resource_group.terraform_rg.name}"
  virtual_machine_name  = "${azurerm_virtual_machine.terraform_webvm1.name}"
  publisher 		= "Microsoft.OSTCExtensions"
  type 			= "CustomScriptForLinux"
  type_handler_version  = "1.2"

  settings = <<SETTINGS
  {
	"commandToExecute": "sudo apt-get -y update && apt-get install -y apache2 && wget --no-check-certificate https://demoavi.com/ && cp index.html /var/www/html/index.html && apachectl restart"
  }
SETTINGS

  tags {
	group = "TerraformInternal"
  }
}
*/