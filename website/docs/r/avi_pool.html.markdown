---
layout: "avi"
page_title: "Avi: avi_pool"
sidebar_current: "docs-avi-resource-pool"
description: |-
  Creates and manages Avi Pool.
---

# avi_pool

The Pool resource allows the creation and management of Avi Pool

## Example Usage

```hcl
resource "Pool" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `apic_epg_name` - (Optional ) argument_description.
        * `application_persistence_profile_ref` - (Optional ) argument_description.
        * `autoscale_launch_config_ref` - (Optional ) argument_description.
        * `autoscale_networks` - (Optional ) argument_description.
        * `autoscale_policy_ref` - (Optional ) argument_description.
        * `capacity_estimation` - (Optional ) argument_description.
        * `capacity_estimation_ttfb_thresh` - (Optional ) argument_description.
        * `cloud_config_cksum` - (Optional ) argument_description.
        * `cloud_ref` - (Optional ) argument_description.
        * `connection_ramp_duration` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `default_server_port` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `domain_name` - (Optional ) argument_description.
        * `east_west` - (Optional ) argument_description.
        * `enabled` - (Optional ) argument_description.
        * `external_autoscale_groups` - (Optional ) argument_description.
        * `fail_action` - (Optional ) argument_description.
        * `fewest_tasks_feedback_delay` - (Optional ) argument_description.
        * `graceful_disable_timeout` - (Optional ) argument_description.
        * `health_monitor_refs` - (Optional ) argument_description.
        * `host_check_enabled` - (Optional ) argument_description.
        * `inline_health_monitor` - (Optional ) argument_description.
        * `ipaddrgroup_ref` - (Optional ) argument_description.
        * `lb_algorithm` - (Optional ) argument_description.
        * `lb_algorithm_consistent_hash_hdr` - (Optional ) argument_description.
        * `lb_algorithm_core_nonaffinity` - (Optional ) argument_description.
        * `lb_algorithm_hash` - (Optional ) argument_description.
        * `lookup_server_by_name` - (Optional ) argument_description.
        * `max_concurrent_connections_per_server` - (Optional ) argument_description.
        * `max_conn_rate_per_server` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `networks` - (Optional ) argument_description.
        * `nsx_securitygroup` - (Optional ) argument_description.
        * `pki_profile_ref` - (Optional ) argument_description.
        * `placement_networks` - (Optional ) argument_description.
        * `request_queue_depth` - (Optional ) argument_description.
        * `request_queue_enabled` - (Optional ) argument_description.
        * `rewrite_host_header_to_server_name` - (Optional ) argument_description.
        * `rewrite_host_header_to_sni` - (Optional ) argument_description.
        * `server_count` - (Optional ) argument_description.
        * `server_name` - (Optional ) argument_description.
        * `server_reselect` - (Optional ) argument_description.
        * `servers` - (Optional ) argument_description.
        * `sni_enabled` - (Optional ) argument_description.
        * `ssl_key_and_certificate_ref` - (Optional ) argument_description.
        * `ssl_profile_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `use_service_port` - (Optional ) argument_description.
            * `vrf_ref` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                                                                                                        * `uuid` - argument_description.
