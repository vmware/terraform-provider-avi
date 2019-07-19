// Configure the Google Cloud provider
provider "google" {
  credentials = "${file("${var.credential_file}")}"
  project     = "${var.project_id}"
  region      = "${var.region}"
}

// In case we have to create the bucket, its name should be unique
// reference: https://www.terraform.io/docs/providers/random/r/string.html
resource "random_string" "bucket_name" {
  length  = 8
  upper   = false
  lower   = true
  number  = false
  special = false
}

// If the bucket name isn't given, we need to generate bucket with random name, IMP NOTE: Adding life cycle rule here which will delete any object in the bucket after 1 day. You may want to delete this bucket after controller creation.
resource "google_storage_bucket" "image-store" {
  count         = "${var.bucket_name == "not-given" ? 1 : 0}"
  name          = "avi-networks-${random_string.bucket_name.result}"
  location      = "US"
  storage_class = "MULTI_REGIONAL"
  project       = "${var.bucket_project == "not-given" ? "${var.project_id}" : "${var.bucket_project}"}"
  force_destroy = true

  lifecycle_rule {
    action {
      type = "Delete"
    }

    condition {
      age        = 1
      with_state = "ANY"
    }
  }
}

// Upload our image to the bucket
resource "google_storage_bucket_object" "gcp_controller_dir" {
  name   = "gcp_controller.tar.gz"
  source = "${var.gcp_controller_dir}"
  bucket = "${var.bucket_name == "not-given" ? "avi-networks-${random_string.bucket_name.result}" : "${var.bucket_name}"}"
}

// Create our custom image from our uploaded raw disk
resource "google_compute_image" "gcp_controller_custom_image" {
  name = "gcp-controller-custom-image"

  raw_disk {
    source = "${google_storage_bucket_object.gcp_controller_dir.self_link}"
  }
}

locals {
  // How many controller should we make
  controllers = "${var.controller_count >= "3" ? 3 : 1}"
}

// Create our VM(s)
resource "google_compute_instance" "avi_controller" {
  count                     = "${local.controllers}"
  name                      = "${var.controller_name}-${count.index}"
  project                   = "${var.project_id}"
  machine_type              = "${var.machine_type}"
  allow_stopping_for_update = true
  zone                      = "${var.zone}"
  tags                      = ["avi-controller-firewall-rules"]
  boot_disk {
    initialize_params {
      size  = "${var.boot_disk_size}"
      type  = "pd-ssd"
      image = "${google_compute_image.gcp_controller_custom_image.self_link}"
    }
  }

  network_interface {
    network            = "${var.network}"
    subnetwork         = "${var.subnetwork}"
    subnetwork_project = "${var.subnetwork_project == "not-given" ? "${var.project_id}" : "${var.subnetwork_project}"}"
    network_ip         = "${length(var.network_ip) == local.controllers ? var.network_ip[count.index] : null}"
  }

  // have to add optional service account email
  dynamic "service_account" {
    for_each = "${var.email != "not-given" ? list("${var.email}") : []}"
    content {
      scopes = ["compute-rw", "storage-rw"]
      email  = "${var.email}"
    }
  }
}

// Add firewall rules
resource "google_compute_firewall" "avi_controller_firewall_rules" {
  name        = "${var.firewall_rules_name}"
  description = "Firewall Rules for our Avi Controller"

  network       = "${var.network}"
  priority      = 1000
  target_tags   = ["avi-controller-firewall-rules"]
  direction     = "INGRESS"
  source_ranges = ["0.0.0.0/0"]
  allow {
    protocol = "tcp"
    ports    = ["0-65534", "80", "443", "73", "74", "75", "63", "97"]
  }
  allow {
    protocol = "udp"
  }
}
