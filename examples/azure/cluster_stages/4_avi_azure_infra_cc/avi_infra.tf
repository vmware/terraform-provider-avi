data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface-1"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

data "azurerm_public_ip" "avi_pip" {
  name                = "${var.pip_name}-0"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_cloudconnectoruser" "azure_user" {
  name       = var.azure_user
  tenant_ref = data.avi_tenant.default_tenant.id
  azure_serviceprincipal  {
    application_id       = var.application_azure_id
    authentication_token = var.auth_token
    tenant_id            = var.tenant_id
  }

  # azure_userpass {
  #   username    = var.azure_username
  #   password    = var.azure_password
  #   tenant_name = var.tenant_id
  # }
 }

resource "avi_cloud" "azure_cloud_cfg" {
  name         = var.azure_cloud_name
  vtype        = "CLOUD_AZURE"
  dhcp_enabled = true
  azure_configuration {
    use_managed_disks     = true
    use_enhanced_ha       = false
    cloud_credentials_ref = avi_cloudconnectoruser.azure_user.id
    use_azure_dns         = false
    resource_group        = "${var.project_name}-terraform-resource-group"
    location              = var.location
    subscription_id       = var.subscription_id
    network_info {
      se_network_id      = data.azurerm_subnet.terraform_subnet.id
      virtual_network_id = "/subscriptions/${var.subscription_id}/resourceGroups/${var.project_name}-terraform-resource-group/providers/Microsoft.Network/virtualNetworks/${var.project_name}-terraform-virtual-network"
    }
  }
  license_tier = "ENTERPRISE_18"
  license_type = "LIC_CORES"
  tenant_ref   = data.avi_tenant.default_tenant.id
}

resource "avi_serviceenginegroup" "azure_se_group" {
  name                         = var.se_group_name
  archive_shm_limit            = 8
  algo                         = "PLACEMENT_ALGO_PACKED"
  buffer_se                    = 0
  cloud_ref                    = avi_cloud.azure_cloud_cfg.id
  connection_memory_percentage = var.connection_mem_percentage #default 50
  disk_per_se                  = var.disk_per_se               # default 10
  ha_mode                      = var.ha_mode                   # default "HA_MODE_SHARED"
  instance_flavor              = var.instance_flavor_se        #default "t2.large"
  license_tier                 = "ENTERPRISE_18"
  license_type                 = "LIC_CORES"
  se_bandwidth_type            = "SE_BANDWIDTH_UNLIMITED"
  max_se                       = var.max_se        #default 2
  max_vs_per_se                = var.max_vs_per_se # default 20
  se_dp_max_hb_version         = 1
  memory_per_se                = var.mem_per_se #default 2048
  min_scaleout_per_vs          = 1
  realtime_se_metrics {
    duration = 0
    enabled  = true
  }
     custom_tag {
       tag_key = var.owner_key
       tag_val = var.owner_value
   }
     custom_tag {
       tag_key = var.policy_key
       tag_val = var.policy_value
   }
     custom_tag {
       tag_key = var.dept_key
       tag_val = var.dept_value
   }
   
  vcpus_per_se         = var.vcpus_per_se      # default 2
  se_deprovision_delay = var.deprovision_delay # default 5
  se_name_prefix       = var.se_group_name
  tenant_ref           = data.avi_tenant.default_tenant.id
}

resource "avi_backupconfiguration" "backup_config" {
  name       = "Backup-Configuration"
  tenant_ref = "admin"
  #tenant_ref              = data.avi_tenant.default_tenant.id
  save_local             = true
  maximum_backups_stored = 4
  backup_passphrase      = var.avi_password
  configpb_attributes {
    version = 1
  }
}
  # vcpus_per_se         = 2
  # se_deprovision_delay = 5
  # se_name_prefix       = var.project_name
  # tenant_ref           = data.avi_tenant.default_tenant.id
# }
