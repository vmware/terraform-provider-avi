resource "azurerm_virtual_network" "terraform_virtual_network" {
  name                = "${var.project_name}-terraform-virtual-network"
  address_space       = [var.virtual_network_cidr]
  location            = var.location
  resource_group_name = azurerm_resource_group.terraform_resource_group.name
  tags = {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

resource "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  count                = var.controllers
  address_prefixes     = [var.subnet_cidr]
  virtual_network_name = azurerm_virtual_network.terraform_virtual_network.name
  resource_group_name  = azurerm_resource_group.terraform_resource_group.name
}



resource "azurerm_network_interface" "terraform_network_interface" {
  count               = var.controllers
  name                = "${var.project_name}-terraform-network-interface-${count.index +1}"
  location            = var.location
  resource_group_name = azurerm_resource_group.terraform_resource_group.name
  enable_ip_forwarding = true
  enable_accelerated_networking = true
  ip_configuration {
    name      = "${var.project_name}-terraform-private-ip-configuration-${count.index +1}"
    subnet_id = azurerm_subnet.terraform_subnet[count.index % length(local.azs)].id

    public_ip_address_id          = count.index == 0 ? azurerm_public_ip.avi_pip[0].id : null
    private_ip_address_allocation = "Dynamic"

  }
  tags = {
    environment = "${var.project_name}-terraform-${var.project_environment}"
  }
}

resource "azurerm_public_ip" "avi_pip" {
  count               = var.controllers == 3 ? var.controllers : 1
  depends_on          = [azurerm_virtual_network.terraform_virtual_network]
  #name               = var.pip_name
  name                = "${var.pip_name}-${count.index}"
  location            = var.location
  resource_group_name = azurerm_resource_group.terraform_resource_group.name
  allocation_method   = "Static"   # Static is required due to the use of the Standard sku
  sku                 = "Standard" # the Standard sku is required due to the use of availability zones
  domain_name_label   = var.fqdn
  zones               = local.azs
}
