provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_tenant     = var.avi_tenant
}

resource "avi_cloud" "avi_cloud_cfg" {
  name         = "OpenStack"
  vtype        = "CLOUD_OPENSTACK"
  dhcp_enabled = true
  openstack_configuration {
    nuage_virtualip         = false
    allowed_address_pairs   = true
    use_internal_endpoints  = false
    auth_url                = var.openstack_url_ip
    security_groups         = true
    admin_tenant            = "admin"
    name_owner              = true
    map_admin_to_cloudadmin = true
    privilege               = "WRITE_ACCESS"
    free_floatingips        = false
    use_nuagevip            = false
    config_drive            = true
    username                = "admin"
    insecure                = false
    mgmt_network_name       = var.network_name
    tenant_se               = true
    contrail_plugin         = false
    port_security           = false
    anti_affinity           = true
    password                = var.openstack_password
    use_keystone_auth       = true
    contrail_disable_policy = false
    import_keystone_tenants = true
    neutron_rbac            = true
    hypervisor              = "KVM"
    region                  = var.op_region
    nuage_port              = 8443
    use_admin_url           = true
    img_format              = "OS_IMG_FMT_AUTO"
    external_networks       = false
  }
}

