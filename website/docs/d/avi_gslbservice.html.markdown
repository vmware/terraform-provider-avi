---
layout: "avi"
page_title: "AVI: avi_gslbservice"
sidebar_current: "docs-avi-datasource-gslbservice"
description: |-
  Get information of Avi GslbService.
---

# avi_gslbservice

This data source is used to to get avi_gslbservice objects.

## Example Usage

```hcl
data "GslbService" "foo_GslbService" {
    uuid = "GslbService-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbService by name.
* `uuid` - (Optional) Search GslbService by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `application_persistence_profile_ref` - The federated application persistence associated with gslbservice site persistence functionality.
* `controller_health_status_enabled` - Gs member's overall health status is derived based on a combination of controller and datapath health-status inputs.
* `created_by` - Creator name.
* `description` - General description.
* `domain_names` - Fully qualified domain name of the gslb service.
* `down_response` - Response to the client query when the gslb service is down.
* `enabled` - Enable or disable the gslb service.
* `groups` - Select list of pools belonging to this gslb service.
* `health_monitor_refs` - Verify vs health by applying one or more health monitors.
* `health_monitor_scope` - Health monitor probe can be executed for all the members or it can be executed only for third-party members.
* `is_federated` - This field indicates that this object is replicated across gslb federation.
* `min_members` - The minimum number of members to distribute traffic to.
* `name` - Name for the gslb service.
* `num_dns_ip` - Number of ip addresses of this gslb service to be returned by the dns service.
* `pool_algorithm` - The load balancing algorithm will pick a gslb pool within the gslb service list of available pools.
* `site_persistence_enabled` - Enable site-persistence for the gslbservice.
* `tenant_ref` - It is a reference to an object of type tenant.
* `ttl` - Ttl value (in seconds) for records served for this gslb service by the dns service.
* `use_edns_client_subnet` - Use the client ip subnet from the edns option as source ipaddress for client geo-location and consistent hash algorithm.
* `uuid` - Uuid of the gslb service.
* `wildcard_match` - Enable wild-card match of fqdn  if an exact match is not found in the dns table, the longest match is chosen by wild-carding the fqdn in the dns request.
