provider "avi" {
  avi_username   = var.avi_username
  avi_tenant     = "admin"
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_version    = var.avi_version
  avi_api_timeout    = 50
}


data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_pool" "lb_pool" {
  for_each     = var.pools
  name         = each.key
  lb_algorithm = lookup(each.value, "lb_algorithm", null)
  lb_algorithm_consistent_hash_hdr = lookup(each.value, "lb_algorithm_consistent_hash_hdr", null)
  dynamic "servers" {
      for_each = [for server in lookup(each.value, "serversList", []): {
        server_ip = server.ip
        server_port = server.port
      }]
      content {
        ip {
          type = "V4"
          addr = servers.value.server_ip
        }
        port = servers.value.server_port
      }
  }
  tenant_ref = data.avi_tenant.default_tenant.id
}
