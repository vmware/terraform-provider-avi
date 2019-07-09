// Initialize avi provider.
provider "avi" {
  avi_username   = var.avi_username
  avi_tenant     = "admin"
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_version    = "18.2.1"
}

// Upload license file to the controller.
resource "avi_fileservice" "uploadlicensefile" {
  uri        = "license"
  local_file = "/path/to/license_file"
  upload     = true
}

// Upload controller.pkg file to the controller.
resource "avi_fileservice" "uploadpkgfile" {
  uri        = "uploads"
  local_file = "/path/to/controller.pkg"
  upload     = true
}

// Upload safenet.tar file to the controller.
resource "avi_fileservice" "uploadsafenetfile" {
  uri        = "hsmpackages?hsmtype=safenet"
  local_file = "/path/to/safenet.tar"
  upload     = true
}

// Download se.ova file to local machine.
resource "avi_fileservice" "downloadseovafile" {
  uri        = "seova"
  local_file = "/path/to/se.ova"
  upload     = false
}

// Download scripts form the controller.
resource "avi_fileservice" "downloadscriptsfile" {
  uri        = "scripts"
  local_file = "/path/to/scripts"
  upload     = false
}

