
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
page_title: "AVI: avi_analyticsprofile"
sidebar_current: "docs-avi-datasource-analyticsprofile"
description: |-
  Get information of Avi AnalyticsProfile.
---

# avi_analyticsprofile

This data source is used to to get avi_analyticsprofile objects.

## Example Usage

```hcl
data "AnalyticsProfile" "foo_AnalyticsProfile" {
    uuid = "AnalyticsProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AnalyticsProfile by name.
* `uuid` - (Optional) Search AnalyticsProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `apdex_response_threshold` - If a client receives an http response in less than the satisfactory latency threshold, the request is considered satisfied.
* `apdex_response_tolerated_factor` - Client tolerated response latency factor.
* `apdex_rtt_threshold` - Satisfactory client to avi round trip time(rtt).
* `apdex_rtt_tolerated_factor` - Tolerated client to avi round trip time(rtt) factor.
* `apdex_rum_threshold` - If a client is able to load a page in less than the satisfactory latency threshold, the pageload is considered satisfied.
* `apdex_rum_tolerated_factor` - Virtual service threshold factor for tolerated page load time (plt) as multiple of apdex_rum_threshold.
* `apdex_server_response_threshold` - A server http response is considered satisfied if latency is less than the satisfactory latency threshold.
* `apdex_server_response_tolerated_factor` - Server tolerated response latency factor.
* `apdex_server_rtt_threshold` - Satisfactory client to avi round trip time(rtt).
* `apdex_server_rtt_tolerated_factor` - Tolerated client to avi round trip time(rtt) factor.
* `client_log_config` - Configure which logs are sent to the avi controller from ses and how they are processed.
* `client_log_streaming_config` - Configure to stream logs to an external server.
* `conn_lossy_ooo_threshold` - A connection between client and avi is considered lossy when more than this percentage of out of order packets are received.
* `conn_lossy_timeo_rexmt_threshold` - A connection between client and avi is considered lossy when more than this percentage of packets are retransmitted due to timeout.
* `conn_lossy_total_rexmt_threshold` - A connection between client and avi is considered lossy when more than this percentage of packets are retransmitted.
* `conn_lossy_zero_win_size_event_threshold` - A client connection is considered lossy when percentage of times a packet could not be trasmitted due to tcp zero window is above this threshold.
* `conn_server_lossy_ooo_threshold` - A connection between avi and server is considered lossy when more than this percentage of out of order packets are received.
* `conn_server_lossy_timeo_rexmt_threshold` - A connection between avi and server is considered lossy when more than this percentage of packets are retransmitted due to timeout.
* `conn_server_lossy_total_rexmt_threshold` - A connection between avi and server is considered lossy when more than this percentage of packets are retransmitted.
* `conn_server_lossy_zero_win_size_event_threshold` - A server connection is considered lossy when percentage of times a packet could not be trasmitted due to tcp zero window is above this threshold.
* `description` - General description.
* `disable_ondemand_metrics` - Virtual service (vs) metrics are processed only when there is live data traffic on the vs.
* `disable_se_analytics` - Disable node (service engine) level analytics forvs metrics.
* `disable_server_analytics` - Disable analytics on backend servers.
* `disable_vs_analytics` - Disable virtualservice (frontend) analytics.
* `enable_advanced_analytics` - Enables advanced analytics features like anomaly detection.
* `exclude_client_close_before_request_as_error` - Exclude client closed connection before an http request could be completed from being classified as an error.
* `exclude_dns_policy_drop_as_significant` - Exclude dns policy drops from the list of errors.
* `exclude_gs_down_as_error` - Exclude queries to gslb services that are operationally down from the list of errors.
* `exclude_http_error_codes` - List of http status codes to be excluded from being classified as an error.
* `exclude_invalid_dns_domain_as_error` - Exclude dns queries to domains outside the domains configured in the dns application profile from the list of errors.
* `exclude_invalid_dns_query_as_error` - Exclude invalid dns queries from the list of errors.
* `exclude_no_dns_record_as_error` - Exclude queries to domains that did not have configured services/records from the list of errors.
* `exclude_no_valid_gs_member_as_error` - Exclude queries to gslb services that have no available members from the list of errors.
* `exclude_persistence_change_as_error` - Exclude persistence server changed while load balancing' from the list of errors.
* `exclude_server_dns_error_as_error` - Exclude server dns error response from the list of errors.
* `exclude_server_tcp_reset_as_error` - Exclude server tcp reset from errors.
* `exclude_sip_error_codes` - List of sip status codes to be excluded from being classified as an error.
* `exclude_syn_retransmit_as_error` - Exclude 'server unanswered syns' from the list of errors.
* `exclude_tcp_reset_as_error` - Exclude tcp resets by client from the list of potential errors.
* `exclude_unsupported_dns_query_as_error` - Exclude unsupported dns queries from the list of errors.
* `healthscore_max_server_limit` - Skips health score computation of pool servers when number of servers in a pool is more than this setting.
* `hs_event_throttle_window` - Time window (in secs) within which only unique health change events should occur.
* `hs_max_anomaly_penalty` - Maximum penalty that may be deducted from health score for anomalies.
* `hs_max_resources_penalty` - Maximum penalty that may be deducted from health score for high resource utilization.
* `hs_max_security_penalty` - Maximum penalty that may be deducted from health score based on security assessment.
* `hs_min_dos_rate` - Dos connection rate below which the dos security assessment will not kick in.
* `hs_performance_boost` - Adds free performance score credits to health score.
* `hs_pscore_traffic_threshold_l4_client` - Threshold number of connections in 5min, below which apdexr, apdexc, rum_apdex, and other network quality metrics are not computed.
* `hs_pscore_traffic_threshold_l4_server` - Threshold number of connections in 5min, below which apdexr, apdexc, rum_apdex, and other network quality metrics are not computed.
* `hs_security_certscore_expired` - Score assigned when the certificate has expired.
* `hs_security_certscore_gt30d` - Score assigned when the certificate expires in more than 30 days.
* `hs_security_certscore_le07d` - Score assigned when the certificate expires in less than or equal to 7 days.
* `hs_security_certscore_le30d` - Score assigned when the certificate expires in less than or equal to 30 days.
* `hs_security_chain_invalidity_penalty` - Penalty for allowing certificates with invalid chain.
* `hs_security_cipherscore_eq000b` - Score assigned when the minimum cipher strength is 0 bits.
* `hs_security_cipherscore_ge128b` - Score assigned when the minimum cipher strength is greater than equal to 128 bits.
* `hs_security_cipherscore_lt128b` - Score assigned when the minimum cipher strength is less than 128 bits.
* `hs_security_encalgo_score_none` - Score assigned when no algorithm is used for encryption.
* `hs_security_encalgo_score_rc4` - Score assigned when rc4 algorithm is used for encryption.
* `hs_security_hsts_penalty` - Penalty for not enabling hsts.
* `hs_security_nonpfs_penalty` - Penalty for allowing non-pfs handshakes.
* `hs_security_selfsignedcert_penalty` - Deprecated.
* `hs_security_ssl30_score` - Score assigned when supporting ssl3.0 encryption protocol.
* `hs_security_tls10_score` - Score assigned when supporting tls1.0 encryption protocol.
* `hs_security_tls11_score` - Score assigned when supporting tls1.1 encryption protocol.
* `hs_security_tls12_score` - Score assigned when supporting tls1.2 encryption protocol.
* `hs_security_weak_signature_algo_penalty` - Penalty for allowing weak signature algorithm(s).
* `name` - The name of the analytics profile.
* `ondemand_metrics_idle_timeout` - This flag sets the time duration of no live data traffic after which virtual service metrics processing is suspended.
* `ranges` - List of http status code ranges to be excluded from being classified as an error.
* `resp_code_block` - Block of http response codes to be excluded from being classified as an error.
* `sensitive_log_profile` - Rules applied to the http application log for filtering sensitive information.
* `sip_log_depth` - Maximum number of sip messages added in logs for a sip transaction.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the analytics profile.

