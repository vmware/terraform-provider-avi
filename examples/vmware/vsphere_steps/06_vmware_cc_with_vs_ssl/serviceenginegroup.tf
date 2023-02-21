resource "avi_serviceenginegroup" "vsphere_se_group" {
  name                         = "Default-Group"
  archive_shm_limit            = 8
  algo                         = "PLACEMENT_ALGO_PACKED"
  buffer_se                    = 0
  cloud_ref                    = avi_cloud.vmware_cloud_tf.id
  connection_memory_percentage = var.connection_mem_percentage #default 50
  disk_per_se                  = var.disk_per_se               # default 10
  ha_mode                      = var.ha_mode                   # default "HA_MODE_SHARED"
  license_tier                 = var.avi_license
  license_type                 = var.vcenter_license_type
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
  vcpus_per_se         = var.vcpus_per_se      # default 2
  se_deprovision_delay = var.deprovision_delay # default 5
  se_name_prefix       = var.se_prefix
  tenant_ref           = var.tenant
}
