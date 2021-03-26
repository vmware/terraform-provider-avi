terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "~>20.1.4"
    }
  }
}

data "avi_tenant" "default_tenant" {
  name = var.tenant
}
data "avi_serverautoscalepolicy" "ServerAutoScalePolicy" {
  count = var.autoscale_policy != null ? 1 : 0
  name = var.autoscale_policy
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_applicationpersistenceprofile" "ApplicationPersistenceProfile" {
  count = var.application_persistence_profile != null ? 1 : 0
  name = var.application_persistence_profile
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_sslprofile" "SSLProfile" {
  count = var.ssl_profile != null ? 1 : 0
  name= var.ssl_profile
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_pkiprofile" "PKIProfile" {
  count = var.pki_profile != null ? 1 : 0
  name= var.pki_profile
  tenant_ref = data.avi_tenant.default_tenant.id   
}
data "avi_sslkeyandcertificate" "SSLKeyAndCertificate" {
  count = var.ssl_key_and_certificate != null ? 1 : 0
  name= var.ssl_key_and_certificate
  tenant_ref = data.avi_tenant.default_tenant.id   
}
data "avi_autoscalelaunchconfig" "AutoScaleLaunchConfig" {
  count = var.autoscale_launch_config != null ? 1 : 0
  name= var.autoscale_launch_config
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_vrfcontext" "VrfContext" {
  count = var.vrf != null ? 1 : 0
  name= var.vrf
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_ipaddrgroup" "IpAddrGroup" {
  count = var.ipaddrgroup != null ? 1 : 0
  name= var.ipaddrgroup
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_analyticsprofile" "AnalyticsProfile" {
  count = var.analytics_profile != null ? 1 : 0
  name= var.analytics_profile
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_cloud" "Cloud" {
  count = var.cloud != null ? 1 : 0
  name= var.cloud
  tenant_ref = data.avi_tenant.default_tenant.id
}
data "avi_healthmonitor" "HealthMonitor" {
  count = var.health_monitors != null ? length(var.health_monitors) : 0
  name = var.health_monitors[count.index]
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_pool" "lb_pool" {
  name         = var.name
  dynamic "servers" {
    for_each = var.services
    content {
      ip {
        type = lookup(servers.value["meta"], "ip_type", "V4")
        addr = servers.value["address"]
      }
      port = servers.value["port"]
      hostname = lookup(servers.value["meta"], "hostname", null)
      enabled = tobool(lookup(servers.value["meta"], "enabled", null))
      ratio = tonumber(lookup(servers.value["meta"], "ratio", null))
      external_uuid = lookup(servers.value["meta"], "external_uuid", null)
      resolve_server_by_dns = tobool(lookup(servers.value["meta"], "resolve_server_by_dns", null))
      prst_hdr_val = lookup(servers.value["meta"], "prst_hdr_val", null)
      mac_address = lookup(servers.value["meta"], "mac_address", null)
      static = tobool(lookup(servers.value["meta"], "static", null))
      server_node = lookup(servers.value["meta"], "server_node", null)
      availability_zone = lookup(servers.value["meta"], "availability_zone", null)
      rewrite_host_header = tobool(lookup(servers.value["meta"], "rewrite_host_header", null))
      external_orchestration_id = lookup(servers.value["meta"], "external_orchestration_id", null)
      description = lookup(servers.value["meta"], "description", null)
      autoscaling_group_name = lookup(servers.value["meta"], "autoscaling_group_name", null)
    }
  }
  default_server_port = var.default_server_port
  graceful_disable_timeout = var.graceful_disable_timeout
  connection_ramp_duration = var.connection_ramp_duration
  max_concurrent_connections_per_server = var.max_concurrent_connections_per_server
  lb_algorithm = var.lb_algorithm
  lb_algorithm_hash = var.lb_algorithm_hash
  lb_algorithm_consistent_hash_hdr = var.lb_algorithm_consistent_hash_hdr
  inline_health_monitor = var.inline_health_monitor
  use_service_port = var.use_service_port
  capacity_estimation = var.capacity_estimation
  capacity_estimation_ttfb_thresh = var.capacity_estimation_ttfb_thresh
  apic_epg_name = var.apic_epg_name
  autoscale_networks = var.autoscale_networks
  fewest_tasks_feedback_delay = var.fewest_tasks_feedback_delay
  enabled = var.enabled
  east_west = var.east_west
  cloud_config_cksum = var.cloud_config_cksum
  request_queue_enabled = var.request_queue_enabled
  request_queue_depth = var.request_queue_depth
  host_check_enabled = var.host_check_enabled
  domain_name = var.domain_name
  sni_enabled = var.sni_enabled
  server_name = var.server_name
  rewrite_host_header_to_sni = var.rewrite_host_header_to_sni
  rewrite_host_header_to_server_name = var.rewrite_host_header_to_server_name
  nsx_securitygroup = var.nsx_securitygroup
  external_autoscale_groups = var.external_autoscale_groups
  lb_algorithm_core_nonaffinity = var.lb_algorithm_core_nonaffinity
  lookup_server_by_name = var.lookup_server_by_name
  service_metadata = var.service_metadata
  description = var.description
  min_servers_up = var.min_servers_up
  min_health_monitors_up = var.min_health_monitors_up
  server_timeout = var.server_timeout
  delete_server_on_dns_refresh = var.delete_server_on_dns_refresh
  enable_http2 = var.enable_http2
  ignore_server_port = var.ignore_server_port
  routing_pool = var.routing_pool
  tier1_lr = var.tier1_lr
  resolve_pool_by_dns = var.resolve_pool_by_dns

  health_monitor_refs = var.health_monitors != null ? [
    for index in range(length(var.health_monitors)) : 
        data.avi_healthmonitor.HealthMonitor[index].id
    ] : null
  pki_profile_ref = var.pki_profile != null ? data.avi_pkiprofile.PKIProfile[0].id : null
  ssl_key_and_certificate_ref = var.ssl_key_and_certificate != null ? data.avi_sslkeyandcertificate.SSLKeyAndCertificate[0].id : null
  ssl_profile_ref = var.ssl_profile != null ? data.avi_sslprofile.SSLProfile[0].id : null
  application_persistence_profile_ref = var.application_persistence_profile != null ? data.avi_applicationpersistenceprofile.ApplicationPersistenceProfile[0].id : null
  autoscale_policy_ref = var.autoscale_policy != null ? data.avi_serverautoscalepolicy.ServerAutoScalePolicy[0].id : null
  autoscale_launch_config_ref = var.autoscale_launch_config != null ? data.avi_autoscalelaunchconfig.AutoScaleLaunchConfig[0].id : null
  vrf_ref = var.vrf != null ? data.avi_vrfcontext.VrfContext[0].id : null
  ipaddrgroup_ref = var.ipaddrgroup != null ? data.avi_ipaddrgroup.IpAddrGroup[0].id : null
  analytics_profile_ref= var.analytics_profile != null ? data.avi_analyticsprofile.AnalyticsProfile[0].id : null
  cloud_ref = var.cloud != null ? data.avi_cloud.Cloud[0].id : null
  tenant_ref = data.avi_tenant.default_tenant.id
}
