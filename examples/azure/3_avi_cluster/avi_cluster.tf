data "azurerm_public_ip" "avi_pip" {
  name                = "${var.pip_name}-0"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_network_interface" "controller_nic_1" {
  name                = "${var.project_name}-terraform-network-interface-1"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_network_interface" "controller_nic_2" {
  name                = "${var.project_name}-terraform-network-interface-2"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_network_interface" "controller_nic_3" {
  name                = "${var.project_name}-terraform-network-interface-3"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

resource "avi_cluster" "azure_cluster" {
  name = "cluster-0-1-tf"
  nodes {
    ip {
      type = "V4"
      addr = data.azurerm_network_interface.controller_nic_1.private_ip_address
    }
    name = "node01"
  }
  nodes {
    ip {
      type = "V4"
      addr = data.azurerm_network_interface.controller_nic_2.private_ip_address
    }
    name = "node02"
  }
  nodes {
    ip {
      type = "V4"
      addr = data.azurerm_network_interface.controller_nic_3.private_ip_address
    }
    name = "node03"
  }
}

resource "avi_systemconfiguration" "avi_system" {
  common_criteria_mode      = false
  uuid                      = "default-uuid"
  welcome_workflow_complete = true
  dns_configuration {
    server_list {
      addr = var.dns_servers[0]
      type = "V4"
    }
    server_list {
      addr = var.dns_servers[1]
      type = "V4"
    }
    server_list {
      addr = var.dns_servers[2]
      type = "V4"
    }
    search_domain = var.search_domain #"remo.local"
  }
  email_configuration {
    disable_tls      = var.mail_server_tls #false
    from_email       = var.email
    mail_server_name = var.mail_server      #"localhost"
    mail_server_port = var.mail_server_port #25
    smtp_type        = var.mail_type        #"SMTP_LOCAL_HOST"
  }
  linux_configuration {
    motd   = ""
    banner = var.banner
  }
  global_tenant_config {
    se_in_provider_context       = true
    tenant_access_to_provider_se = true
    tenant_vrf                   = false
  }

  ntp_configuration {
    ntp_servers {
      key_number = 1
      server {
        addr = var.ntp_servers[0]
        type = "DNS"
      }
    }
    ntp_servers {
      key_number = 1
      server {
        addr = var.ntp_servers[1]
        type = "DNS"
      }
    }
    ntp_servers {
      key_number = 1
      server {
        addr = var.ntp_servers[2]
        type = "DNS"
      }
    }
    ntp_servers {
      key_number = 1
      server {
        addr = var.ntp_servers[3]
        type = "DNS"
      }
    }
  }
}
