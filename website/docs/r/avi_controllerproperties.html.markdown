---
layout: "avi"
page_title: "Avi: avi_controllerproperties"
sidebar_current: "docs-avi-resource-controllerproperties"
description: |-
  Creates and manages Avi ControllerProperties.
---

# avi_controllerproperties

The ControllerProperties resource allows the creation and management of Avi ControllerProperties

## Example Usage

```hcl
resource "ControllerProperties" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `allow_ip_forwarding` - (Optional ) argument_description.
        * `allow_unauthenticated_apis` - (Optional ) argument_description.
        * `allow_unauthenticated_nodes` - (Optional ) argument_description.
        * `api_idle_timeout` - (Optional ) argument_description.
        * `api_perf_logging_threshold` - (Optional ) argument_description.
        * `appviewx_compat_mode` - (Optional ) argument_description.
        * `attach_ip_retry_interval` - (Optional ) argument_description.
        * `attach_ip_retry_limit` - (Optional ) argument_description.
        * `bm_use_ansible` - (Optional ) argument_description.
        * `cleanup_expired_authtoken_timeout_period` - (Optional ) argument_description.
        * `cleanup_sessions_timeout_period` - (Optional ) argument_description.
        * `cloud_reconcile` - (Optional ) argument_description.
        * `cluster_ip_gratuitous_arp_period` - (Optional ) argument_description.
        * `consistency_check_timeout_period` - (Optional ) argument_description.
        * `crashed_se_reboot` - (Optional ) argument_description.
        * `dead_se_detection_timer` - (Optional ) argument_description.
        * `dns_refresh_period` - (Optional ) argument_description.
        * `dummy` - (Optional ) argument_description.
        * `enable_api_sharding` - (Optional ) argument_description.
        * `enable_memory_balancer` - (Optional ) argument_description.
        * `fatal_error_lease_time` - (Optional ) argument_description.
        * `max_dead_se_in_grp` - (Optional ) argument_description.
        * `max_pcap_per_tenant` - (Optional ) argument_description.
        * `max_seq_attach_ip_failures` - (Optional ) argument_description.
        * `max_seq_vnic_failures` - (Optional ) argument_description.
        * `persistence_key_rotate_period` - (Optional ) argument_description.
        * `portal_token` - (Optional ) argument_description.
        * `process_locked_useraccounts_timeout_period` - (Optional ) argument_description.
        * `process_pki_profile_timeout_period` - (Optional ) argument_description.
        * `query_host_fail` - (Optional ) argument_description.
        * `safenet_hsm_version` - (Optional ) argument_description.
        * `se_create_timeout` - (Optional ) argument_description.
        * `se_failover_attempt_interval` - (Optional ) argument_description.
        * `se_from_marketplace` - (Optional ) argument_description.
        * `se_offline_del` - (Optional ) argument_description.
        * `se_vnic_cooldown` - (Optional ) argument_description.
        * `secure_channel_cleanup_timeout` - (Optional ) argument_description.
        * `secure_channel_controller_token_timeout` - (Optional ) argument_description.
        * `secure_channel_se_token_timeout` - (Optional ) argument_description.
        * `seupgrade_fabric_pool_size` - (Optional ) argument_description.
        * `seupgrade_segroup_min_dead_timeout` - (Optional ) argument_description.
        * `ssl_certificate_expiry_warning_days` - (Optional ) argument_description.
        * `unresponsive_se_reboot` - (Optional ) argument_description.
        * `upgrade_dns_ttl` - (Optional ) argument_description.
        * `upgrade_lease_time` - (Optional ) argument_description.
            * `vnic_op_fail_time` - (Optional ) argument_description.
        * `vs_apic_scaleout_timeout` - (Optional ) argument_description.
        * `vs_awaiting_se_timeout` - (Optional ) argument_description.
        * `vs_key_rotate_period` - (Optional ) argument_description.
        * `vs_scaleout_ready_check_interval` - (Optional ) argument_description.
        * `vs_se_attach_ip_fail` - (Optional ) argument_description.
        * `vs_se_bootup_fail` - (Optional ) argument_description.
        * `vs_se_create_fail` - (Optional ) argument_description.
        * `vs_se_ping_fail` - (Optional ) argument_description.
        * `vs_se_vnic_fail` - (Optional ) argument_description.
        * `vs_se_vnic_ip_fail` - (Optional ) argument_description.
        * `warmstart_se_reconnect_wait_time` - (Optional ) argument_description.
        * `warmstart_vs_resync_wait_time` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                                                                                        * `uuid` - argument_description.
                                                        
