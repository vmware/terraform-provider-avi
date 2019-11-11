package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceAnalyticsProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSAnalyticsProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "name", "test-System-Analytics-Profile-abc"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_event_throttle_window", "1209600"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "disable_se_analytics", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_server_rtt_tolerated_factor", "4"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_nonpfs_penalty", "1"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_tls12_score", "5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_server_tcp_reset_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_min_dos_rate", "1000"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_no_dns_record_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_server_lossy_zero_win_size_event_threshold", "2"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_max_resources_penalty", "25"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_lossy_total_rexmt_threshold", "50"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_certscore_le07d", "2"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_pscore_traffic_threshold_l4_client", "10"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_no_valid_gs_member_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_encalgo_score_rc4", "2.5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_server_response_threshold", "400"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_cipherscore_ge128b", "5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_performance_boost", "0"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_rum_tolerated_factor", "4"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_invalid_dns_domain_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_max_anomaly_penalty", "10"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_gs_down_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_server_response_tolerated_factor", "4"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "disable_server_analytics", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_server_lossy_timeo_rexmt_threshold", "20"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_client_close_before_request_as_error", "true"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_weak_signature_algo_penalty", "1"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_persistence_change_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_selfsignedcert_penalty", "1"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_server_lossy_total_rexmt_threshold", "50"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_lossy_timeo_rexmt_threshold", "20"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_rtt_tolerated_factor", "4"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_certscore_gt30d", "5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_ssl30_score", "3.5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_server_lossy_ooo_threshold", "50"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_cipherscore_eq000b", "0"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_certscore_le30d", "4"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_syn_retransmit_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_dns_policy_drop_as_significant", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_server_rtt_threshold", "125"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_hsts_penalty", "1"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_server_dns_error_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_tls11_score", "5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_lossy_zero_win_size_event_threshold", "2"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_rum_threshold", "5000"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_pscore_traffic_threshold_l4_server", "10"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_max_security_penalty", "100"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_encalgo_score_none", "0"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_response_tolerated_factor", "4"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_tcp_reset_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_chain_invalidity_penalty", "1"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_invalid_dns_query_as_error", "false"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "conn_lossy_ooo_threshold", "50"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_cipherscore_lt128b", "3.5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_tls10_score", "5"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "hs_security_certscore_expired", "0"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_response_threshold", "500"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "apdex_rtt_threshold", "250"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "exclude_unsupported_dns_query_as_error", "false"),
				),
			},
		},
	})

}

const testAccAVIDSAnalyticsProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_analyticsprofile" "testAnalyticsProfile" {
	hs_event_throttle_window = "1209600"
	disable_se_analytics = false
	apdex_server_rtt_tolerated_factor = "4"
	hs_security_nonpfs_penalty = "1"
	hs_security_tls12_score = "5"
	exclude_server_tcp_reset_as_error = false
	hs_min_dos_rate = "1000"
	exclude_no_dns_record_as_error = false
	conn_server_lossy_zero_win_size_event_threshold = "2"
	hs_max_resources_penalty = "25"
	conn_lossy_total_rexmt_threshold = "50"
	hs_security_certscore_le07d = "2"
	hs_pscore_traffic_threshold_l4_client = "10"
	exclude_no_valid_gs_member_as_error = false
	hs_security_encalgo_score_rc4 = "2.5"
	apdex_server_response_threshold = "400"
	hs_security_cipherscore_ge128b = "5"
	hs_performance_boost = "0"
	apdex_rum_tolerated_factor = "4"
	exclude_invalid_dns_domain_as_error = false
	exclude_http_error_codes = ["475"]
	hs_max_anomaly_penalty = "10"
	exclude_gs_down_as_error = false
	apdex_server_response_tolerated_factor = "4"
	client_log_config {
		enable_significant_log_collection = true
		non_significant_log_processing = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
		significant_log_processing = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
		filtered_log_processing = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
	}
	disable_server_analytics = false
	conn_server_lossy_timeo_rexmt_threshold = "20"
	exclude_client_close_before_request_as_error = true
	hs_security_weak_signature_algo_penalty = "1"
	exclude_persistence_change_as_error = false
	hs_security_selfsignedcert_penalty = "1"
	conn_server_lossy_total_rexmt_threshold = "50"
	conn_lossy_timeo_rexmt_threshold = "20"
	apdex_rtt_tolerated_factor = "4"
	hs_security_certscore_gt30d = "5"
	hs_security_ssl30_score = "3.5"
	conn_server_lossy_ooo_threshold = "50"
	hs_security_cipherscore_eq000b = "0"
	hs_security_certscore_le30d = "4"
	exclude_syn_retransmit_as_error = false
	exclude_dns_policy_drop_as_significant = false
	apdex_server_rtt_threshold = "125"
	hs_security_hsts_penalty = "1"
	exclude_server_dns_error_as_error = false
	hs_security_tls11_score = "5"
	conn_lossy_zero_win_size_event_threshold = "2"
	apdex_rum_threshold = "5000"
	hs_pscore_traffic_threshold_l4_server = "10"
	name = "test-System-Analytics-Profile-abc"
	hs_max_security_penalty = "100"
	hs_security_encalgo_score_none = "0"
	tenant_ref = data.avi_tenant.default_tenant.id
	apdex_response_tolerated_factor = "4"
	exclude_tcp_reset_as_error = false
	hs_security_chain_invalidity_penalty = "1"
	exclude_invalid_dns_query_as_error = false
	conn_lossy_ooo_threshold = "50"
	hs_security_cipherscore_lt128b = "3.5"
	hs_security_tls10_score = "5"
	hs_security_certscore_expired = "0"
	apdex_response_threshold = "500"
	apdex_rtt_threshold = "250"
	exclude_unsupported_dns_query_as_error = false
}

data "avi_analyticsprofile" "testAnalyticsProfile" {
    name= "${avi_analyticsprofile.testAnalyticsProfile.name}"
}
`
