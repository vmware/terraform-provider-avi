// Configure the Google Cloud provider
provider "google" {
  credentials = file(var.credential_file)
}

// Get our controllers
data "google_compute_instance" "avi_controller" {
  count   = var.controller_count
  name    = "${var.controller_name}-${count.index}"
  zone    = var.zone
  project = var.project_id
}

// Avi provider
provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_current_password
  avi_controller = data.google_compute_instance.avi_controller.*.network_interface.0.network_ip[0]
  avi_tenant     = var.tenant
  avi_version    = var.avi_version
}

// Avi tenant
data "avi_tenant" "default_tenant" {
  name = var.tenant
}

// CloudConnector user
resource "avi_cloudconnectoruser" "gcp_user" {
  name       = var.gcp_user
  tenant_ref = data.avi_tenant.default_tenant.id

  gcp_credentials {
    service_account_keyfile_data = file(var.credential_file)
  }
}

// Making our cloud
resource "avi_cloud" "gcp_cloud_cfg" {
  name  = var.gcp_cloud_name
  vtype = "CLOUD_GCP"

  gcp_configuration {
    vip_allocation_strategy {
        routes{
            match_se_group_subnet = false
        }
    }
    network_config {
      config = "INBAND_MANAGEMENT"

      inband {
        vpc_project_id   = var.vpc_project_id
        vpc_subnet_name  = var.vpc_subnetwork
        vpc_network_name = var.vpc_network
      }
    }
    region_name           = var.region_name
    se_project_id         = var.se_project_id
    firewall_target_tags  = var.firewall_target_tags
    zones                 = var.zones
    cloud_credentials_ref = avi_cloudconnectoruser.gcp_user.id
    # gcs_project_id        = "${var.gcs_project_id}"  // optional, read description for more info.
    # gcs_bucket_name       = "${var.gcs_bucket_name}" // optional, read description for more info.
  }

  license_tier = "ENTERPRISE"
  license_type = "LIC_CORES"
  tenant_ref   = data.avi_tenant.default_tenant.id
}

