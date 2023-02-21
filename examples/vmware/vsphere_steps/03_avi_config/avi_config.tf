resource "avi_systemconfiguration" "avi_system" {
  common_criteria_mode      = false
  uuid                      = "default-uuid"
  default_license_tier = tonumber(split(".", var.avi_version)[0]) >= tonumber("21") && tonumber(split(".", var.avi_version)["1"]) >= tonumber("1") && tonumber(split(".", var.avi_version)["2"]) >= tonumber("3") && var.avi_default_license_tier != "" ? var.avi_default_license_tier : "ENTERPRISE"
  welcome_workflow_complete = true

  dns_configuration {
    dynamic server_list {
      for_each = flatten(split(",", replace(var.avi_dns_server_ips, " ", "")))
      content {
        addr = server_list.value
        type = "V4"
      }
    }
  }

  global_tenant_config {
    se_in_provider_context       = true
    tenant_access_to_provider_se = true
    tenant_vrf                   = false
  }

  ntp_configuration {
    dynamic ntp_servers {
      for_each = flatten(split(",", replace(var.avi_ntp_server_ips, " ", "")))
      content {
        key_number = 1
        server {
          addr = ntp_servers.value
          type = "V4"
        }
      }
    }
  }
}

resource "avi_backupconfiguration" "backup_config" {
  name       = "Backup-Configuration"
  tenant_ref = var.avi_tenant
  save_local             = true
  maximum_backups_stored = 4
  backup_passphrase      = var.avi_password
  configpb_attributes {
    version = 1
  }
}