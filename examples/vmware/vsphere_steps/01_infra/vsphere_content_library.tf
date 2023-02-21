resource "null_resource" "download_avi" {
  provisioner "local-exec" {
    command = "curl -o /tmp/controller.ova \"${var.avi_controller_url}\""
  }
}

resource "vsphere_content_library" "library" {
  name            = "${var.content_library.basename}${random_string.id.result}"
  storage_backing = [data.vsphere_datastore.datastore.id]
}

resource "vsphere_content_library_item" "file" {
  depends_on = [null_resource.download_avi]
  name        = "controller.ova"
  library_id  = vsphere_content_library.library.id
  file_url = "/tmp/controller.ova"
}

resource "null_resource" "remove_download_avi" {
  depends_on = [vsphere_content_library_item.file]
  provisioner "local-exec" {
    command = "rm -f /tmp/controller.ova"
  }
}
