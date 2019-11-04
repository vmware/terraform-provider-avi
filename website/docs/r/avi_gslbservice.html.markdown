---
layout: "avi"
page_title: "Avi: avi_gslbservice"
sidebar_current: "docs-avi-resource-gslbservice"
description: |-
  Creates and manages Avi GslbService.
---

# avi_gslbservice

The GslbService resource allows the creation and management of Avi GslbService

## Example Usage

```hcl
resource "avi_gslbservice" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the gslb service.
* `application_persistence_profile_ref` - (Optional) The federated application persistence associated with gslbservice site persistence functionality.
* `controller_health_status_enabled` - (Optional) Gs member's overall health status is derived based on a combination of controller and datapath health-status inputs.
* `created_by` - (Optional) Creator name.
* `description` - (Optional) User defined description for the object.
* `domain_names` - (Optional) Fully qualified domain name of the gslb service.
* `down_response` - (Optional) Response to the client query when the gslb service is down.
* `enabled` - (Optional) Enable or disable the gslb service.
* `groups` - (Optional) Select list of pools belonging to this gslb service.
* `health_monitor_refs` - (Optional) Verify vs health by applying one or more health monitors.
* `health_monitor_scope` - (Optional) Health monitor probe can be executed for all the members or it can be executed only for third-party members.
* `hm_off` - (Optional) This field is an internal field and is used in se.
* `is_federated` - (Optional) This field indicates that this object is replicated across gslb federation.
* `min_members` - (Optional) The minimum number of members to distribute traffic to.
* `num_dns_ip` - (Optional) Number of ip addresses of this gslb service to be returned by the dns service.
* `pool_algorithm` - (Optional) The load balancing algorithm will pick a gslb pool within the gslb service list of available pools.
* `resolve_cname` - (Optional) This field indicates that for a cname query, respond with resolved cnames in the additional section with a records.
* `site_persistence_enabled` - (Optional) Enable site-persistence for the gslbservice.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `ttl` - (Optional) Ttl value (in seconds) for records served for this gslb service by the dns service.
* `use_edns_client_subnet` - (Optional) Use the client ip subnet from the edns option as source ipaddress for client geo-location and consistent hash algorithm.
* `wildcard_match` - (Optional) Enable wild-card match of fqdn  if an exact match is not found in the dns table, the longest match is chosen by wild-carding the fqdn in the dns request.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the gslb service.

