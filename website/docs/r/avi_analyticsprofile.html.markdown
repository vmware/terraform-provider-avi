
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

---
layout: "avi"
page_title: "Avi: avi_analyticsprofile"
sidebar_current: "docs-avi-resource-analyticsprofile"
description: |-
  Creates and manages Avi AnalyticsProfile.
---

# avi_analyticsprofile

The AnalyticsProfile resource allows the creation and management of Avi AnalyticsProfile

## Example Usage

```hcl
resource "AnalyticsProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `apdex_response_threshold` - (Optional ) argument_description.
        * `apdex_response_tolerated_factor` - (Optional ) argument_description.
        * `apdex_rtt_threshold` - (Optional ) argument_description.
        * `apdex_rtt_tolerated_factor` - (Optional ) argument_description.
        * `apdex_rum_threshold` - (Optional ) argument_description.
        * `apdex_rum_tolerated_factor` - (Optional ) argument_description.
        * `apdex_server_response_threshold` - (Optional ) argument_description.
        * `apdex_server_response_tolerated_factor` - (Optional ) argument_description.
        * `apdex_server_rtt_threshold` - (Optional ) argument_description.
        * `apdex_server_rtt_tolerated_factor` - (Optional ) argument_description.
        * `client_log_config` - (Optional ) argument_description.
        * `client_log_streaming_config` - (Optional ) argument_description.
        * `conn_lossy_ooo_threshold` - (Optional ) argument_description.
        * `conn_lossy_timeo_rexmt_threshold` - (Optional ) argument_description.
        * `conn_lossy_total_rexmt_threshold` - (Optional ) argument_description.
        * `conn_lossy_zero_win_size_event_threshold` - (Optional ) argument_description.
        * `conn_server_lossy_ooo_threshold` - (Optional ) argument_description.
        * `conn_server_lossy_timeo_rexmt_threshold` - (Optional ) argument_description.
        * `conn_server_lossy_total_rexmt_threshold` - (Optional ) argument_description.
        * `conn_server_lossy_zero_win_size_event_threshold` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `disable_ondemand_metrics` - (Optional ) argument_description.
        * `disable_se_analytics` - (Optional ) argument_description.
        * `disable_server_analytics` - (Optional ) argument_description.
        * `disable_vs_analytics` - (Optional ) argument_description.
        * `enable_advanced_analytics` - (Optional ) argument_description.
        * `exclude_client_close_before_request_as_error` - (Optional ) argument_description.
        * `exclude_dns_policy_drop_as_significant` - (Optional ) argument_description.
        * `exclude_gs_down_as_error` - (Optional ) argument_description.
        * `exclude_http_error_codes` - (Optional ) argument_description.
        * `exclude_invalid_dns_domain_as_error` - (Optional ) argument_description.
        * `exclude_invalid_dns_query_as_error` - (Optional ) argument_description.
        * `exclude_no_dns_record_as_error` - (Optional ) argument_description.
        * `exclude_no_valid_gs_member_as_error` - (Optional ) argument_description.
        * `exclude_persistence_change_as_error` - (Optional ) argument_description.
        * `exclude_server_dns_error_as_error` - (Optional ) argument_description.
        * `exclude_server_tcp_reset_as_error` - (Optional ) argument_description.
        * `exclude_sip_error_codes` - (Optional ) argument_description.
        * `exclude_syn_retransmit_as_error` - (Optional ) argument_description.
        * `exclude_tcp_reset_as_error` - (Optional ) argument_description.
        * `exclude_unsupported_dns_query_as_error` - (Optional ) argument_description.
        * `healthscore_max_server_limit` - (Optional ) argument_description.
        * `hs_event_throttle_window` - (Optional ) argument_description.
        * `hs_max_anomaly_penalty` - (Optional ) argument_description.
        * `hs_max_resources_penalty` - (Optional ) argument_description.
        * `hs_max_security_penalty` - (Optional ) argument_description.
        * `hs_min_dos_rate` - (Optional ) argument_description.
        * `hs_performance_boost` - (Optional ) argument_description.
        * `hs_pscore_traffic_threshold_l4_client` - (Optional ) argument_description.
        * `hs_pscore_traffic_threshold_l4_server` - (Optional ) argument_description.
        * `hs_security_certscore_expired` - (Optional ) argument_description.
        * `hs_security_certscore_gt30d` - (Optional ) argument_description.
        * `hs_security_certscore_le07d` - (Optional ) argument_description.
        * `hs_security_certscore_le30d` - (Optional ) argument_description.
        * `hs_security_chain_invalidity_penalty` - (Optional ) argument_description.
        * `hs_security_cipherscore_eq000b` - (Optional ) argument_description.
        * `hs_security_cipherscore_ge128b` - (Optional ) argument_description.
        * `hs_security_cipherscore_lt128b` - (Optional ) argument_description.
        * `hs_security_encalgo_score_none` - (Optional ) argument_description.
        * `hs_security_encalgo_score_rc4` - (Optional ) argument_description.
        * `hs_security_hsts_penalty` - (Optional ) argument_description.
        * `hs_security_nonpfs_penalty` - (Optional ) argument_description.
        * `hs_security_selfsignedcert_penalty` - (Optional ) argument_description.
        * `hs_security_ssl30_score` - (Optional ) argument_description.
        * `hs_security_tls10_score` - (Optional ) argument_description.
        * `hs_security_tls11_score` - (Optional ) argument_description.
        * `hs_security_tls12_score` - (Optional ) argument_description.
        * `hs_security_weak_signature_algo_penalty` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `ondemand_metrics_idle_timeout` - (Optional ) argument_description.
        * `ranges` - (Optional ) argument_description.
        * `resp_code_block` - (Optional ) argument_description.
        * `sensitive_log_profile` - (Optional ) argument_description.
        * `sip_log_depth` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                                                                                                                                                                                                                * `uuid` - argument_description.
    
