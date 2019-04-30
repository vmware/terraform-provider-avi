---
layout: "avi"
page_title: "Avi: avi_virtualservice"
sidebar_current: "docs-avi-resource-virtualservice"
description: |-
  Creates and manages Avi VirtualService.
---

# avi_virtualservice

The VirtualService resource allows the creation and management of Avi VirtualService

## Example Usage

```hcl
resource "VirtualService" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `active_standby_se_tag` - (Optional ) argument_description.
        * `allow_invalid_client_cert` - (Optional ) argument_description.
        * `analytics_policy` - (Optional ) argument_description.
        * `analytics_profile_ref` - (Optional ) argument_description.
        * `apic_contract_graph` - (Optional ) argument_description.
        * `application_profile_ref` - (Optional ) argument_description.
        * `bulk_sync_kvcache` - (Optional ) argument_description.
        * `client_auth` - (Optional ) argument_description.
        * `close_client_conn_on_config_update` - (Optional ) argument_description.
        * `cloud_config_cksum` - (Optional ) argument_description.
        * `cloud_ref` - (Optional ) argument_description.
        * `cloud_type` - (Optional ) argument_description.
        * `connections_rate_limit` - (Optional ) argument_description.
        * `content_rewrite` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `delay_fairness` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `dns_info` - (Optional ) argument_description.
        * `dns_policies` - (Optional ) argument_description.
        * `east_west_placement` - (Optional ) argument_description.
        * `enable_autogw` - (Optional ) argument_description.
        * `enable_rhi` - (Optional ) argument_description.
        * `enable_rhi_snat` - (Optional ) argument_description.
        * `enabled` - (Optional ) argument_description.
        * `error_page_profile_ref` - (Optional ) argument_description.
        * `flow_dist` - (Optional ) argument_description.
        * `flow_label_type` - (Optional ) argument_description.
        * `fqdn` - (Optional ) argument_description.
        * `host_name_xlate` - (Optional ) argument_description.
        * `http_policies` - (Optional ) argument_description.
        * `ign_pool_net_reach` - (Optional ) argument_description.
        * `l4_policies` - (Optional ) argument_description.
        * `limit_doser` - (Optional ) argument_description.
        * `max_cps_per_client` - (Optional ) argument_description.
        * `microservice_ref` - (Optional ) argument_description.
        * `min_pools_up` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `network_profile_ref` - (Optional ) argument_description.
        * `network_security_policy_ref` - (Optional ) argument_description.
        * `nsx_securitygroup` - (Optional ) argument_description.
        * `performance_limits` - (Optional ) argument_description.
        * `pool_group_ref` - (Optional ) argument_description.
        * `pool_ref` - (Optional ) argument_description.
        * `remove_listening_port_on_vs_down` - (Optional ) argument_description.
        * `requests_rate_limit` - (Optional ) argument_description.
        * `saml_sp_config` - (Optional ) argument_description.
        * `scaleout_ecmp` - (Optional ) argument_description.
        * `se_group_ref` - (Optional ) argument_description.
        * `security_policy_ref` - (Optional ) argument_description.
        * `server_network_profile_ref` - (Optional ) argument_description.
        * `service_metadata` - (Optional ) argument_description.
        * `service_pool_select` - (Optional ) argument_description.
        * `services` - (Optional ) argument_description.
        * `sideband_profile` - (Optional ) argument_description.
        * `snat_ip` - (Optional ) argument_description.
        * `ssl_key_and_certificate_refs` - (Optional ) argument_description.
        * `ssl_profile_ref` - (Optional ) argument_description.
        * `ssl_profile_selectors` - (Optional ) argument_description.
        * `ssl_sess_cache_avg_size` - (Optional ) argument_description.
        * `sso_policy_ref` - (Optional ) argument_description.
        * `static_dns_records` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `topology_policies` - (Optional ) argument_description.
        * `traffic_clone_profile_ref` - (Optional ) argument_description.
        * `traffic_enabled` - (Optional ) argument_description.
        * `type` - (Optional ) argument_description.
        * `use_bridge_ip_as_vip` - (Optional ) argument_description.
        * `use_vip_as_snat` - (Optional ) argument_description.
            * `vh_domain_name` - (Optional ) argument_description.
        * `vh_parent_vs_uuid` - (Optional ) argument_description.
        * `vip` - (Optional ) argument_description.
        * `vrf_context_ref` - (Optional ) argument_description.
        * `vs_datascripts` - (Optional ) argument_description.
        * `vsvip_cloud_config_cksum` - (Optional ) argument_description.
        * `vsvip_ref` - (Optional ) argument_description.
        * `waf_policy_ref` - (Optional ) argument_description.
        * `weight` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                                                                                                                                                                                    * `uuid` - argument_description.
                                        
