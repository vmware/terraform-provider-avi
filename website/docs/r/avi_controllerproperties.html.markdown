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
resource "avi_controllerproperties" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `allow_admin_network_updates` - (Optional) Allow non-admin tenants to update admin vrfcontext and network objects.
* `allow_ip_forwarding` - (Optional) Field introduced in 17.1.1.
* `allow_unauthenticated_apis` - (Optional) Allow unauthenticated access for special apis.
* `allow_unauthenticated_nodes` - (Optional) Boolean flag to set allow_unauthenticated_nodes.
* `api_idle_timeout` - (Optional) Allowed values are 0-1440.
* `api_perf_logging_threshold` - (Optional) Threshold to log request timing in portal_performance.log and server-timing response header.
* `appviewx_compat_mode` - (Optional) Export configuration in appviewx compatibility mode.
* `attach_ip_retry_interval` - (Optional) Placeholder for description of property attach_ip_retry_interval of obj type controllerproperties field type integer  type int.
* `attach_ip_retry_limit` - (Optional) Placeholder for description of property attach_ip_retry_limit of obj type controllerproperties field type integer  type int.
* `bm_use_ansible` - (Optional) Use ansible for se creation in baremetal.
* `cleanup_expired_authtoken_timeout_period` - (Optional) Period for auth token cleanup job.
* `cleanup_sessions_timeout_period` - (Optional) Period for sessions cleanup job.
* `cloud_reconcile` - (Optional) Enable/disable periodic reconcile for all the clouds.
* `cluster_ip_gratuitous_arp_period` - (Optional) Period for cluster ip gratuitous arp job.
* `consistency_check_timeout_period` - (Optional) Period for consistency check job.
* `crashed_se_reboot` - (Optional) Placeholder for description of property crashed_se_reboot of obj type controllerproperties field type integer  type int.
* `dead_se_detection_timer` - (Optional) Placeholder for description of property dead_se_detection_timer of obj type controllerproperties field type integer  type int.
* `default_minimum_api_timeout` - (Optional) Minimum api timeout value.if this value is not 60, it will be the default timeout for all apis that do not have a specific timeout.if an api has a specific timeout but is less than this value, this value will become the new timeout.
* `dns_refresh_period` - (Optional) Period for refresh pool and gslb dns job.
* `dummy` - (Optional) Placeholder for description of property dummy of obj type controllerproperties field type integer  type int.
* `enable_api_sharding` - (Optional) This setting enables the controller leader to shard api requests to the followers (if any).
* `enable_memory_balancer` - (Optional) Enable/disable memory balancer.
* `fatal_error_lease_time` - (Optional) Placeholder for description of property fatal_error_lease_time of obj type controllerproperties field type integer  type int.
* `max_dead_se_in_grp` - (Optional) Placeholder for description of property max_dead_se_in_grp of obj type controllerproperties field type integer  type int.
* `max_pcap_per_tenant` - (Optional) Maximum number of pcap files stored per tenant.
* `max_seq_attach_ip_failures` - (Optional) Maximum number of consecutive attach ip failures that halts vs placement.
* `max_seq_vnic_failures` - (Optional) Placeholder for description of property max_seq_vnic_failures of obj type controllerproperties field type integer  type int.
* `permission_scoped_shared_admin_networks` - (Optional) Network and vrfcontext objects from the admin tenant will not be shared to non-admin tenants unless admin permissions are granted.
* `persistence_key_rotate_period` - (Optional) Period for rotate app persistence keys job.
* `portal_token` - (Optional) Token used for uploading tech-support to portal.
* `process_locked_useraccounts_timeout_period` - (Optional) Period for process locked user accounts job.
* `process_pki_profile_timeout_period` - (Optional) Period for process pki profile job.
* `query_host_fail` - (Optional) Placeholder for description of property query_host_fail of obj type controllerproperties field type integer  type int.
* `safenet_hsm_version` - (Optional) Version of the safenet package installed on the controller.
* `se_create_timeout` - (Optional) Placeholder for description of property se_create_timeout of obj type controllerproperties field type integer  type int.
* `se_failover_attempt_interval` - (Optional) Interval between attempting failovers to an se.
* `se_from_marketplace` - (Optional) This setting decides whether se is to be deployed from the cloud marketplace or to be created by the controller.
* `se_offline_del` - (Optional) Placeholder for description of property se_offline_del of obj type controllerproperties field type integer  type int.
* `se_vnic_cooldown` - (Optional) Placeholder for description of property se_vnic_cooldown of obj type controllerproperties field type integer  type int.
* `secure_channel_cleanup_timeout` - (Optional) Period for secure channel cleanup job.
* `secure_channel_controller_token_timeout` - (Optional) Placeholder for description of property secure_channel_controller_token_timeout of obj type controllerproperties field type integer  type int.
* `secure_channel_se_token_timeout` - (Optional) Placeholder for description of property secure_channel_se_token_timeout of obj type controllerproperties field type integer  type int.
* `seupgrade_copy_pool_size` - (Optional) This parameter defines the number of simultaneous se image downloads in a segroup.
* `seupgrade_fabric_pool_size` - (Optional) Pool size used for all fabric commands during se upgrade.
* `seupgrade_segroup_min_dead_timeout` - (Optional) Time to wait before marking segroup upgrade as stuck.
* `shared_ssl_certificates` - (Optional) Ssl certificates in the admin tenant can be used in non-admin tenants.
* `ssl_certificate_expiry_warning_days` - (Optional) Number of days for ssl certificate expiry warning.
* `unresponsive_se_reboot` - (Optional) Placeholder for description of property unresponsive_se_reboot of obj type controllerproperties field type integer  type int.
* `upgrade_dns_ttl` - (Optional) Time to account for dns ttl during upgrade.
* `upgrade_lease_time` - (Optional) Placeholder for description of property upgrade_lease_time of obj type controllerproperties field type integer  type int.
* `vnic_op_fail_time` - (Optional) Placeholder for description of property vnic_op_fail_time of obj type controllerproperties field type integer  type int.
* `vs_apic_scaleout_timeout` - (Optional) Time to wait for the scaled out se to become ready before marking the scaleout done, applies to apic configuration only.
* `vs_awaiting_se_timeout` - (Optional) Placeholder for description of property vs_awaiting_se_timeout of obj type controllerproperties field type integer  type int.
* `vs_key_rotate_period` - (Optional) Period for rotate vs keys job.
* `vs_scaleout_ready_check_interval` - (Optional) Interval for checking scaleout_ready status while controller is waiting for scaleoutready rpc from the service engine.
* `vs_se_attach_ip_fail` - (Optional) Time to wait before marking attach ip operation on an se as failed.
* `vs_se_bootup_fail` - (Optional) Placeholder for description of property vs_se_bootup_fail of obj type controllerproperties field type integer  type int.
* `vs_se_create_fail` - (Optional) Placeholder for description of property vs_se_create_fail of obj type controllerproperties field type integer  type int.
* `vs_se_ping_fail` - (Optional) Placeholder for description of property vs_se_ping_fail of obj type controllerproperties field type integer  type int.
* `vs_se_vnic_fail` - (Optional) Placeholder for description of property vs_se_vnic_fail of obj type controllerproperties field type integer  type int.
* `vs_se_vnic_ip_fail` - (Optional) Placeholder for description of property vs_se_vnic_ip_fail of obj type controllerproperties field type integer  type int.
* `warmstart_se_reconnect_wait_time` - (Optional) Placeholder for description of property warmstart_se_reconnect_wait_time of obj type controllerproperties field type integer  type int.
* `warmstart_vs_resync_wait_time` - (Optional) Timeout for warmstart vs resync.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

