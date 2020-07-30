// Configure the Google Cloud provider
provider "google" {
  credentials = file(var.credential_file)
  project     = var.project_id
}

// Getting our VM's (Controller's)
data "google_compute_instance" "avi_controller" {
  count = 3
  name  = "${var.controller_name}-${count.index}"
  zone  = var.zone
}

// Configure the Avi provider
provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_current_password
  avi_controller = data.google_compute_instance.avi_controller.*.network_interface.0.network_ip[0]
  avi_tenant     = var.tenant
  avi_version    = var.avi_version
}

// Creating our cluster
resource "avi_cluster" "gcp_cluster" {
  name = var.cluster_name

  nodes {
    ip {
      type = "V4"
      addr = data.google_compute_instance.avi_controller.*.network_interface.0.network_ip[0]
    }

    name = "${var.controller_name}-0"
  }

  nodes {
    ip {
      type = "V4"
      addr = data.google_compute_instance.avi_controller.*.network_interface.0.network_ip[1]
    }

    name = "${var.controller_name}-1"
  }

  nodes {
    ip {
      type = "V4"
      addr = data.google_compute_instance.avi_controller.*.network_interface.0.network_ip[2]
    }

    name = "${var.controller_name}-2"
  }
  depends_on = [avi_useraccount.avi_user]
}

// In case you would like to update your password
resource "avi_useraccount" "avi_user" {
  username     = var.avi_username
  old_password = var.avi_current_password
  password     = var.avi_new_password
}

