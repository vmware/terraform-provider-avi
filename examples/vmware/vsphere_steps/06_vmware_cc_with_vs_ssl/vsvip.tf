resource "avi_vsvip" "avi_vsvip" {
  name = "terraform-vip"
  vip {
    vip_id = "0"
    ip_address {
      type = "V4"
      addr = var.vip_ip
    }
  }
  cloud_ref    = avi_cloud.vmware_cloud_tf.id
  tenant_ref   = var.tenant
}
