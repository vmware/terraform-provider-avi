<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
resource "avi_analyticsprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the analytics profile.
* `apdex_response_threshold` - (Optional) If a client receives an http response in less than the satisfactory latency threshold, the request is considered satisfied. It is considered tolerated if it is not satisfied and less than tolerated latency factor multiplied by the satisfactory latency threshold. Greater than this number and the client's request is considered frustrated. Allowed values are 1-30000. Unit is milliseconds. Allowed in basic(allowed values- 500) edition, essentials(allowed values- 500) edition, enterprise edition.
* `apdex_response_tolerated_factor` - (Optional) Client tolerated response latency factor. Client must receive a response within this factor times the satisfactory threshold (apdex_response_threshold) to be considered tolerated. Allowed values are 1-1000. Allowed in basic(allowed values- 4) edition, essentials(allowed values- 4) edition, enterprise edition.
* `apdex_rtt_threshold` - (Optional) Satisfactory client to avi round trip time(rtt). Allowed values are 1-2000. Unit is milliseconds. Allowed in basic(allowed values- 250) edition, essentials(allowed values- 250) edition, enterprise edition.
* `apdex_rtt_tolerated_factor` - (Optional) Tolerated client to avi round trip time(rtt) factor. It is a multiple of apdex_rtt_tolerated_factor. Allowed values are 1-1000. Allowed in basic(allowed values- 4) edition, essentials(allowed values- 4) edition, enterprise edition.
* `apdex_rum_threshold` - (Optional) If a client is able to load a page in less than the satisfactory latency threshold, the pageload is considered satisfied. It is considered tolerated if it is greater than satisfied but less than the tolerated latency multiplied by satisifed latency. Greater than this number and the client's request is considered frustrated. A pageload includes the time for dns lookup, download of all http objects, and page render time. Allowed values are 1-30000. Unit is milliseconds. Allowed in basic(allowed values- 5000) edition, essentials(allowed values- 5000) edition, enterprise edition.
* `apdex_rum_tolerated_factor` - (Optional) Virtual service threshold factor for tolerated page load time (plt) as multiple of apdex_rum_threshold. Allowed values are 1-1000. Allowed in basic(allowed values- 4) edition, essentials(allowed values- 4) edition, enterprise edition.
* `apdex_server_response_threshold` - (Optional) A server http response is considered satisfied if latency is less than the satisfactory latency threshold. The response is considered tolerated when it is greater than satisfied but less than the tolerated latency factor * s_latency. Greater than this number and the server response is considered frustrated. Allowed values are 1-30000. Unit is milliseconds. Allowed in basic(allowed values- 400) edition, essentials(allowed values- 400) edition, enterprise edition.
* `apdex_server_response_tolerated_factor` - (Optional) Server tolerated response latency factor. Servermust response within this factor times the satisfactory threshold (apdex_server_response_threshold) to be considered tolerated. Allowed values are 1-1000. Allowed in basic(allowed values- 4) edition, essentials(allowed values- 4) edition, enterprise edition.
* `apdex_server_rtt_threshold` - (Optional) Satisfactory client to avi round trip time(rtt). Allowed values are 1-2000. Unit is milliseconds. Allowed in basic(allowed values- 125) edition, essentials(allowed values- 125) edition, enterprise edition.
* `apdex_server_rtt_tolerated_factor` - (Optional) Tolerated client to avi round trip time(rtt) factor. It is a multiple of apdex_rtt_tolerated_factor. Allowed values are 1-1000. Allowed in basic(allowed values- 4) edition, essentials(allowed values- 4) edition, enterprise edition.
* `client_log_config` - (Optional) Configure which logs are sent to the avi controller from ses and how they are processed.
* `client_log_streaming_config` - (Optional) Configure to stream logs to an external server. Field introduced in 17.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `conn_lossy_ooo_threshold` - (Optional) A connection between client and avi is considered lossy when more than this percentage of out of order packets are received. Allowed values are 1-100. Unit is percent. Allowed in basic(allowed values- 50) edition, essentials(allowed values- 50) edition, enterprise edition.
* `conn_lossy_timeo_rexmt_threshold` - (Optional) A connection between client and avi is considered lossy when more than this percentage of packets are retransmitted due to timeout. Allowed values are 1-100. Unit is percent. Allowed in basic(allowed values- 20) edition, essentials(allowed values- 20) edition, enterprise edition.
* `conn_lossy_total_rexmt_threshold` - (Optional) A connection between client and avi is considered lossy when more than this percentage of packets are retransmitted. Allowed values are 1-100. Unit is percent. Allowed in basic(allowed values- 50) edition, essentials(allowed values- 50) edition, enterprise edition.
* `conn_lossy_zero_win_size_event_threshold` - (Optional) A client connection is considered lossy when percentage of times a packet could not be trasmitted due to tcp zero window is above this threshold. Allowed values are 0-100. Unit is percent. Allowed in basic(allowed values- 2) edition, essentials(allowed values- 2) edition, enterprise edition.
* `conn_server_lossy_ooo_threshold` - (Optional) A connection between avi and server is considered lossy when more than this percentage of out of order packets are received. Allowed values are 1-100. Unit is percent. Allowed in basic(allowed values- 50) edition, essentials(allowed values- 50) edition, enterprise edition.
* `conn_server_lossy_timeo_rexmt_threshold` - (Optional) A connection between avi and server is considered lossy when more than this percentage of packets are retransmitted due to timeout. Allowed values are 1-100. Unit is percent. Allowed in basic(allowed values- 20) edition, essentials(allowed values- 20) edition, enterprise edition.
* `conn_server_lossy_total_rexmt_threshold` - (Optional) A connection between avi and server is considered lossy when more than this percentage of packets are retransmitted. Allowed values are 1-100. Unit is percent. Allowed in basic(allowed values- 50) edition, essentials(allowed values- 50) edition, enterprise edition.
* `conn_server_lossy_zero_win_size_event_threshold` - (Optional) A server connection is considered lossy when percentage of times a packet could not be trasmitted due to tcp zero window is above this threshold. Allowed values are 0-100. Unit is percent. Allowed in basic(allowed values- 2) edition, essentials(allowed values- 2) edition, enterprise edition.
* `description` - (Optional) User defined description for the object.
* `enable_adaptive_config` - (Optional) Enable adaptive configuration for optimizing resource usage. Field introduced in 20.1.1.
* `enable_advanced_analytics` - (Optional) Enables advanced analytics features like anomaly detection. If set to false, anomaly computation (and associated rules/events) for vs, pool and server metrics will be deactivated. However, setting it to false reduces cpu and memory requirements for analytics subsystem. Field introduced in 17.2.13, 18.1.5, 18.2.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition. Special default for basic edition is false, essentials edition is false, enterprise is true.
* `enable_ondemand_metrics` - (Optional) Virtual service (vs) metrics are processed only when there is live data traffic on the vs. In case, vs is idle for a period of time as specified by ondemand_metrics_idle_timeout then metrics processing is suspended for that vs. Field introduced in 20.1.3.
* `enable_se_analytics` - (Optional) Enable node (service engine) level analytics forvs metrics. Field introduced in 20.1.3.
* `enable_server_analytics` - (Optional) Enables analytics on backend servers. This may be desired in container environment when there are large number of ephemeral servers. Additionally, no healthscore of servers is computed when server analytics is enabled. Field introduced in 20.1.3.
* `enable_vs_analytics` - (Optional) Enable virtualservice (frontend) analytics. This flag enables metrics and healthscore for virtualservice. Field introduced in 20.1.3.
* `exclude_client_close_before_request_as_error` - (Optional) Exclude client closed connection before an http request could be completed from being classified as an error. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_dns_policy_drop_as_significant` - (Optional) Exclude dns policy drops from the list of errors. Field introduced in 17.2.2. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_gs_down_as_error` - (Optional) Exclude queries to gslb services that are operationally down from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_http_error_codes` - (Optional) List of http status codes to be excluded from being classified as an error. Error connections or responses impacts health score, are included as significant logs, and may be classified as part of a dos attack.
* `exclude_invalid_dns_domain_as_error` - (Optional) Exclude dns queries to domains outside the domains configured in the dns application profile from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_invalid_dns_query_as_error` - (Optional) Exclude invalid dns queries from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_issuer_revoked_ocsp_responses_as_error` - (Optional) Exclude the issuer-revoked ocsp responses from the list of errors. Field introduced in 20.1.1. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `exclude_no_dns_record_as_error` - (Optional) Exclude queries to domains that did not have configured services/records from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_no_valid_gs_member_as_error` - (Optional) Exclude queries to gslb services that have no available members from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_persistence_change_as_error` - (Optional) Exclude persistence server changed while load balancing' from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_revoked_ocsp_responses_as_error` - (Optional) Exclude the revoked ocsp certificate status responses from the list of errors. Field introduced in 20.1.1. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `exclude_server_dns_error_as_error` - (Optional) Exclude server dns error response from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_server_tcp_reset_as_error` - (Optional) Exclude server tcp reset from errors. It is common for applications like ms exchange. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_sip_error_codes` - (Optional) List of sip status codes to be excluded from being classified as an error. Field introduced in 17.2.13, 18.1.5, 18.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `exclude_stale_ocsp_responses_as_error` - (Optional) Exclude the stale ocsp certificate status responses from the list of errors. Field introduced in 20.1.1. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `exclude_syn_retransmit_as_error` - (Optional) Exclude 'server unanswered syns' from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_tcp_reset_as_error` - (Optional) Exclude tcp resets by client from the list of potential errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `exclude_unavailable_ocsp_responses_as_error` - (Optional) Exclude the unavailable ocsp responses from the list of errors. Field introduced in 20.1.1. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `exclude_unsupported_dns_query_as_error` - (Optional) Exclude unsupported dns queries from the list of errors. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `healthscore_max_server_limit` - (Optional) Skips health score computation of pool servers when number of servers in a pool is more than this setting. Allowed values are 0-5000. Special values are 0- 'server health score is deactivated'. Field introduced in 17.2.13, 18.1.4. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition. Special default for basic edition is 0, essentials edition is 0, enterprise is 20.
* `hs_event_throttle_window` - (Optional) Time window (in secs) within which only unique health change events should occur. Allowed in basic(allowed values- 1209600) edition, essentials(allowed values- 1209600) edition, enterprise edition.
* `hs_max_anomaly_penalty` - (Optional) Maximum penalty that may be deducted from health score for anomalies. Allowed values are 0-100. Allowed in basic(allowed values- 10) edition, essentials(allowed values- 10) edition, enterprise edition.
* `hs_max_resources_penalty` - (Optional) Maximum penalty that may be deducted from health score for high resource utilization. Allowed values are 0-100. Allowed in basic(allowed values- 25) edition, essentials(allowed values- 25) edition, enterprise edition.
* `hs_max_security_penalty` - (Optional) Maximum penalty that may be deducted from health score based on security assessment. Allowed values are 0-100. Allowed in basic(allowed values- 100) edition, essentials(allowed values- 100) edition, enterprise edition.
* `hs_min_dos_rate` - (Optional) Dos connection rate below which the dos security assessment will not kick in. Allowed in basic(allowed values- 1000) edition, essentials(allowed values- 1000) edition, enterprise edition.
* `hs_performance_boost` - (Optional) Adds free performance score credits to health score. It can be used for compensating health score for known slow applications. Allowed values are 0-100. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `hs_pscore_traffic_threshold_l4_client` - (Optional) Threshold number of connections in 5min, below which apdexr, apdexc, rum_apdex, and other network quality metrics are not computed. Allowed in basic(allowed values- 10) edition, essentials(allowed values- 10) edition, enterprise edition.
* `hs_pscore_traffic_threshold_l4_server` - (Optional) Threshold number of connections in 5min, below which apdexr, apdexc, rum_apdex, and other network quality metrics are not computed. Allowed in basic(allowed values- 10) edition, essentials(allowed values- 10) edition, enterprise edition.
* `hs_security_certscore_expired` - (Optional) Score assigned when the certificate has expired. Allowed values are 0-5. Allowed in basic(allowed values- 0.0) edition, essentials(allowed values- 0.0) edition, enterprise edition.
* `hs_security_certscore_gt30d` - (Optional) Score assigned when the certificate expires in more than 30 days. Allowed values are 0-5. Allowed in basic(allowed values- 5.0) edition, essentials(allowed values- 5.0) edition, enterprise edition.
* `hs_security_certscore_le07d` - (Optional) Score assigned when the certificate expires in less than or equal to 7 days. Allowed values are 0-5. Allowed in basic(allowed values- 2.0) edition, essentials(allowed values- 2.0) edition, enterprise edition.
* `hs_security_certscore_le30d` - (Optional) Score assigned when the certificate expires in less than or equal to 30 days. Allowed values are 0-5. Allowed in basic(allowed values- 4.0) edition, essentials(allowed values- 4.0) edition, enterprise edition.
* `hs_security_chain_invalidity_penalty` - (Optional) Penalty for allowing certificates with invalid chain. Allowed values are 0-5. Allowed in basic(allowed values- 1.0) edition, essentials(allowed values- 1.0) edition, enterprise edition.
* `hs_security_cipherscore_eq000b` - (Optional) Score assigned when the minimum cipher strength is 0 bits. Allowed values are 0-5. Allowed in basic(allowed values- 0.0) edition, essentials(allowed values- 0.0) edition, enterprise edition.
* `hs_security_cipherscore_ge128b` - (Optional) Score assigned when the minimum cipher strength is greater than equal to 128 bits. Allowed values are 0-5. Allowed in basic(allowed values- 5.0) edition, essentials(allowed values- 5.0) edition, enterprise edition.
* `hs_security_cipherscore_lt128b` - (Optional) Score assigned when the minimum cipher strength is less than 128 bits. Allowed values are 0-5. Allowed in basic(allowed values- 3.5) edition, essentials(allowed values- 3.5) edition, enterprise edition.
* `hs_security_encalgo_score_none` - (Optional) Score assigned when no algorithm is used for encryption. Allowed values are 0-5. Allowed in basic(allowed values- 0.0) edition, essentials(allowed values- 0.0) edition, enterprise edition.
* `hs_security_encalgo_score_rc4` - (Optional) Score assigned when rc4 algorithm is used for encryption. Allowed values are 0-5. Allowed in basic(allowed values- 2.5) edition, essentials(allowed values- 2.5) edition, enterprise edition.
* `hs_security_hsts_penalty` - (Optional) Penalty for not enabling hsts. Allowed values are 0-5. Allowed in basic(allowed values- 1.0) edition, essentials(allowed values- 1.0) edition, enterprise edition.
* `hs_security_nonpfs_penalty` - (Optional) Penalty for allowing non-pfs handshakes. Allowed values are 0-5. Allowed in basic(allowed values- 1.0) edition, essentials(allowed values- 1.0) edition, enterprise edition.
* `hs_security_ocsp_revoked_score` - (Optional) Score assigned when ocsp certificate status is set to revoked or issuer revoked. Allowed values are 0.0-5.0. Field introduced in 20.1.1. Allowed in basic(allowed values- 0.0) edition, essentials(allowed values- 0.0) edition, enterprise edition.
* `hs_security_selfsignedcert_penalty` - (Optional) Deprecated. Allowed values are 0-5. Allowed in basic(allowed values- 1.0) edition, essentials(allowed values- 1.0) edition, enterprise edition.
* `hs_security_ssl30_score` - (Optional) Score assigned when supporting ssl3.0 encryption protocol. Allowed values are 0-5. Allowed in basic(allowed values- 3.5) edition, essentials(allowed values- 3.5) edition, enterprise edition.
* `hs_security_tls10_score` - (Optional) Score assigned when supporting tls1.0 encryption protocol. Allowed values are 0-5. Allowed in basic(allowed values- 5.0) edition, essentials(allowed values- 5.0) edition, enterprise edition.
* `hs_security_tls11_score` - (Optional) Score assigned when supporting tls1.1 encryption protocol. Allowed values are 0-5. Allowed in basic(allowed values- 5.0) edition, essentials(allowed values- 5.0) edition, enterprise edition.
* `hs_security_tls12_score` - (Optional) Score assigned when supporting tls1.2 encryption protocol. Allowed values are 0-5. Allowed in basic(allowed values- 5.0) edition, essentials(allowed values- 5.0) edition, enterprise edition.
* `hs_security_weak_signature_algo_penalty` - (Optional) Penalty for allowing weak signature algorithm(s). Allowed values are 0-5. Allowed in basic(allowed values- 1.0) edition, essentials(allowed values- 1.0) edition, enterprise edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `ondemand_metrics_idle_timeout` - (Optional) This flag sets the time duration of no live data traffic after which virtual service metrics processing is suspended. It is applicable only when enable_ondemand_metrics is set to false. Field introduced in 18.1.1. Unit is seconds.
* `ranges` - (Optional) List of http status code ranges to be excluded from being classified as an error.
* `resp_code_block` - (Optional) Block of http response codes to be excluded from being classified as an error. Enum options - AP_HTTP_RSP_4XX, AP_HTTP_RSP_5XX.
* `sensitive_log_profile` - (Optional) Rules applied to the http application log for filtering sensitive information. Field introduced in 17.2.10, 18.1.2. Allowed in basic edition, essentials edition, enterprise edition.
* `sip_log_depth` - (Optional) Maximum number of sip messages added in logs for a sip transaction. By default, this value is 20. Allowed values are 1-1000. Field introduced in 17.2.13, 18.1.5, 18.2.1. Allowed in basic(allowed values- 20) edition, essentials(allowed values- 20) edition, enterprise edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the analytics profile.

