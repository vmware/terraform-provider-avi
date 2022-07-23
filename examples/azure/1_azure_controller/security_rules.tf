locals {
  range_allowed_mgmt   = [var.allowed_ip]
  #range_allowed_mgmt   = [var.allowed_ip, var.virtual_network_cidr]
  public_avi_ips       = [for nic in azurerm_network_interface.terraform_network_interface : format("%s/32", nic.private_ip_address)]
}
resource "azurerm_network_security_group" "avinetworks_secgroup" {
  name                = "avi-nsg"
  location            = var.location
  resource_group_name = azurerm_resource_group.terraform_resource_group.name

  security_rule {
    name                       = "AllowHTTPS"
    description                = "allows port 443 inbound"
    priority                   = 1000
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "443"
    source_address_prefixes    = local.range_allowed_mgmt
    destination_address_prefix = "*"
  }
  security_rule {
    name                       = "AllowHTTPS8443"
    description                = "allows port 8443 inbound"
    priority                   = 1001
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "8443"
    source_address_prefixes    = local.range_allowed_mgmt
    destination_address_prefix = "*"
  }
  security_rule {
    name                       = "Allowssh"
    description                = "allows port ssh inbound"
    priority                   = 1002
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "22"
    source_address_prefixes    = local.range_allowed_mgmt
    destination_address_prefix = "*"
  }
  security_rule {
    name                       = "AllowNTP"
    description                = "allows port NTP inbound"
    priority                   = 1003
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "123"
    source_address_prefixes    = local.range_allowed_mgmt
    destination_address_prefix = "*"
  }
  security_rule {
    name                       = "Allow_ICMP"
    description                = "allows ICMP inbound"
    priority                   = 1900
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Icmp"
    source_port_range          = "*"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
    }
}
