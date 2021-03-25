variable "avi_username" {
  type    = string
  default = ""
}
variable "avi_password" {
  type    = string
  default = ""
}
variable "avi_controller" {
  type    = string
  default = ""
}
variable "avi_version" {
  type    = string
  default = ""
}

# Below block is standard template for CTS. 
variable "services" {
  description = "Consul services monitored by Consul-Terraform-Sync"
  type = map(
    object({
      id        = string
      name      = string
      kind      = string
      address   = string
      port      = number
      meta      = map(string)
      tags      = list(string)
      namespace = string
      status    = string
      node                  = string
      node_id               = string
      node_address          = string
      node_datacenter       = string
      node_tagged_addresses = map(string)
      node_meta             = map(string)
      cts_user_defined_meta = map(string)
    })
  )
}

#Below block is for lb_pool related variables.

variable "pool_name" {
  type    = string
}
variable "default_server_port" {
  type    = number
  default = null
}
variable "graceful_disable_timeout" {
  type    = number
  default = null
}
variable "connection_ramp_duration" {
  type    = number
  default = null
}
variable "max_concurrent_connections_per_server" {
  type    = string
  default = null
}
variable "lb_algorithm" {
  type    = string
  default = null
}
variable "lb_algorithm_hash" {
  type    = string
  default = null
}
variable "lb_algorithm_consistent_hash_hdr" {
  type    = string
  default = null
}
variable "application_persistence_profile" {
  type    = string
  default = null
}
variable "ssl_profile" {
  type    = string
  default = null
}
variable "inline_health_monitor" {
  type    = bool
  default = null
}
variable "use_service_port" {
  type    = bool
  default = null
}
variable "capacity_estimation" {
  type    = bool
  default = null
}
variable "capacity_estimation_ttfb_thresh" {
  type    = number
  default = null
}
variable "pki_profile" {
  type    = string
  default = null
}
variable "ssl_key_and_certificate" {
  type    = string
  default = null
}
variable "autoscale_networks" {
  type    = list(string)
  default = null
}
variable "autoscale_policy" {
  type    = string
  default = null
}
variable "autoscale_launch_config" {
  type    = string
  default = null
}
variable "vrf" {
  type    = string
  default = null
}
variable "ipaddrgroup" {
  type    = string
  default = null
}
variable "fewest_tasks_feedback_delay" {
  type    = number
  default = null
}
variable "enabled" {
  type    = bool
  default = null
}
variable "east_west" {
  type    = bool
  default = null
}
variable "created_by" {
  type    = string
  default = null
}
variable "cloud_config_cksum" {
  type    = string
  default = null
}
variable "request_queue_enabled" {
  type    = bool
  default = null
}
variable "request_queue_depth" {
  type    = number
  default = null
}
variable "host_check_enabled" {
  type    = bool
  default = null
}
variable "sni_enabled" {
  type    = bool
  default = null
}
variable "server_name" {
  type    = string
  default = null
}
variable "rewrite_host_header_to_sni" {
  type    = bool
  default = null
}
variable "rewrite_host_header_to_server_name" {
  type    = bool
  default = null
}
variable "nsx_securitygroup" {
  type    = list(string)
  default = null
}
variable "external_autoscale_groups" {
  type    = list(string)
  default = null
}
variable "domain_name" {
  type    = list(string)
  default = null
}
variable "health_monitor_names" {
  type    = list(string)
  default = null
}
variable "lb_algorithm_core_nonaffinity" {
  type    = number
  default = null
}
variable "lookup_server_by_name" {
  type    = bool
  default = null
}
variable "analytics_profile" {
  type    = string
  default = null
}
variable "service_metadata" {
  type    = string
  default = null
}
variable "description" {
  type    = string
  default = null
}
variable "tenant" {
  type    = string
  default = "admin"
}
variable "cloud" {
  type    = number
  default = null
}
variable "min_servers_up" {
  type    = number
  default = null
}
variable "min_health_monitors_up" {
  type    = number
  default = null
}
variable "server_timeout" {
  type    = number
  default = null
}
variable "delete_server_on_dns_refresh" {
  type    = bool
  default = null
}
variable "enable_http2" {
  type    = bool
  default = null
}
variable "ignore_server_port" {
  type    = bool
  default = null
}
variable "routing_pool" {
  type    = bool
  default = null
}
variable "tier1_lr" {
  type    = string
  default = null
}
variable "resolve_pool_by_dns" {
  type    = bool
  default = null
}
variable "apic_epg_name" {
  type    = string
  default = null
}
variable "health_monitors" {
  type    = list(string)
  default = null
}
