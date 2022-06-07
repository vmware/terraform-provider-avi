<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
data "avi_gslbservice" "foo_gslbservice" {
    uuid = "gslbservice-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbService by name.
* `uuid` - (Optional) Search GslbService by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `application_persistence_profile_ref` - The federated application persistence associated with gslbservice site persistence functionality. It is a reference to an object of type applicationpersistenceprofile. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `controller_health_status_enabled` - Gs member's overall health status is derived based on a combination of controller and datapath health-status inputs. Note that the datapath status is determined by the association of health monitor profiles. Only the controller provided status is determined through this configuration. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `created_by` - Creator name. Field introduced in 17.1.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `domain_names` - Fully qualified domain name of the gslb service. Minimum of 1 items required. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `down_response` - Response to the client query when the gslb service is down. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enabled` - Enable or disable the gslb service. If the gslb service is enabled, then the vips are sent in the dns responses based on reachability and configured algorithm. If the gslb service is disabled, then the vips are no longer available in the dns response. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `groups` - Select list of pools belonging to this gslb service. Minimum of 1 items required. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `health_monitor_refs` - Verify vs health by applying one or more health monitors. Active monitors generate synthetic traffic from dns service engine and to mark a vs up or down based on the response. It is a reference to an object of type healthmonitor. Maximum of 6 items allowed. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `health_monitor_scope` - Health monitor probe can be executed for all the members or it can be executed only for third-party members. This operational mode is useful to reduce the number of health monitor probes in case of a hybrid scenario. In such a case, avi members can have controller derived status while non-avi members can be probed by via health monitor probes in dataplane. Enum options - GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS, GSLB_SERVICE_HEALTH_MONITOR_ONLY_NON_AVI_MEMBERS. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `hm_off` - This field is an internal field and is used in se. Field introduced in 18.2.2. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `is_federated` - This field indicates that this object is replicated across gslb federation. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `min_members` - The minimum number of members to distribute traffic to. Allowed values are 1-65535. Special values are 0 - disable. Field introduced in 17.2.4. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Name for the gslb service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `num_dns_ip` - Number of ip addresses of this gslb service to be returned by the dns service. Enter 0 to return all ip addresses. Allowed values are 1-20. Special values are 0- return all ip addresses. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pool_algorithm` - The load balancing algorithm will pick a gslb pool within the gslb service list of available pools. Enum options - GSLB_SERVICE_ALGORITHM_PRIORITY, GSLB_SERVICE_ALGORITHM_GEO. Field introduced in 17.2.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `resolve_cname` - This field indicates that for a cname query, respond with resolved cnames in the additional section with a records. Field introduced in 18.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `site_persistence_enabled` - Enable site-persistence for the gslbservice. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `topology_policy_enabled` - When enabled, topology policy rules are used for member selection first. If no valid member is found using the topology policy rules, configured gslb algorithms for pool selection and member selection are used. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ttl` - Ttl value (in seconds) for records served for this gslb service by the dns service. Allowed values are 0-86400. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `use_edns_client_subnet` - Use the client ip subnet from the edns option as source ipaddress for client geo-location and consistent hash algorithm. Default is true. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the gslb service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `wildcard_match` - Enable wild-card match of fqdn  if an exact match is not found in the dns table, the longest match is chosen by wild-carding the fqdn in the dns request. Default is false. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

