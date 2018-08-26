
provider "azurerm" {
  subscription_id = "${var.subscription_id}"
  client_id 	  = "${var.client_id}"
  client_secret   = "${var.client_secret}"
  tenant_id 	  = "${var.tenant_id}"
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.azurerm_network_interface.controller_nic.private_ip_address}"
  avi_tenant     = "admin"
  avi_version    = "${var.avi_version}"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_authprofile" "samlauthprofile" {
	name       = "saml-test"
	type       = "AUTH_PROFILE_SAML"
	tenant_ref = "${data.avi_tenant.default_tenant.id}"
  saml = {
    /*
    #This needs to be added after creating/registering app at Service Provider.
    idp = {
      metadata = <<EOF
EOF
    }
    */
    sp = {
      org_display_name   = ""
      org_name           = ""
      org_url            = ""
      tech_contact_email = ""
      tech_contact_name  = ""
      saml_entity_type   = "AUTH_SAML_CLUSTER_VIP"
      #below can be used if the above is set to AUTH_SAML_DNS_FQDN
      #fqdn               = "controller.avinetworks.com"
    }
  }
}

/*
#This needs to be added after creating/registering app at Service Provider.
resource "avi_systemconfiguration" "sys_conf" {
  admin_auth_configuration {
    auth_profile_ref = "${avi_authprofile.samlauthprofile.id}"
    allow_local_user_login = true
    mapping_rules {
      index         = 0
      assign_tenant = "ASSIGN_ALL"
      assign_role   = "ASSIGN_ALL"
      is_superuser  = false
    }
  }
}
*/