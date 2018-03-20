package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIAnalyticsProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAnalyticsProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAnalyticsProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAnalyticsProfileExists("avi_analyticsprofile.testAnalyticsProfile"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "name", "testSystem-Analytics-Profile")),
			},
			{
				Config: testAccAVIAnalyticsProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAnalyticsProfileExists("avi_analyticsprofile.testAnalyticsProfile"),
					resource.TestCheckResourceAttr(
						"avi_analyticsprofile.testAnalyticsProfile", "name", "testSystem-Analytics-Profile-abc")),
			},
		},
	})

}

func testAccCheckAVIAnalyticsProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI AnalyticsProfile ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIAnalyticsProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_analyticsprofile" {
			continue
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI AnalyticsProfile still exists")
		}
	}
	return nil
}

const testAccAVIAnalyticsProfileConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
resource "avi_analyticsprofile" "testAnalyticsProfile" {
"hs_event_throttle_window" = "1209600"
"disable_se_analytics" = false
"apdex_server_rtt_tolerated_factor" = "4.0"
"hs_security_nonpfs_penalty" = "1.0"
"hs_security_tls12_score" = "5.0"
"exclude_server_tcp_reset_as_error" = false
"hs_min_dos_rate" = "1000"
"exclude_no_dns_record_as_error" = false
"conn_server_lossy_zero_win_size_event_threshold" = "2"
"hs_max_resources_penalty" = "25"
"conn_lossy_total_rexmt_threshold" = "50"
"hs_security_certscore_le07d" = "2.0"
"hs_pscore_traffic_threshold_l4_client" = "10.0"
"exclude_no_valid_gs_member_as_error" = false
"hs_security_encalgo_score_rc4" = "2.5"
"apdex_server_response_threshold" = "400"
"hs_security_cipherscore_ge128b" = "5.0"
"hs_performance_boost" = "0"
"exclude_invalid_dns_domain_as_error" = false
"exclude_http_error_codes" = ["475"]
"hs_max_anomaly_penalty" = "10"
"exclude_gs_down_as_error" = false
"apdex_server_response_tolerated_factor" = "4.0"
"client_log_config" {
"enable_significant_log_collection" = true
"non_significant_log_processing" = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
"significant_log_processing" = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
"filtered_log_processing" = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
}
"disable_server_analytics" = false
"conn_server_lossy_timeo_rexmt_threshold" = "20"
"exclude_client_close_before_request_as_error" = true
"hs_security_weak_signature_algo_penalty" = "1.0"
"exclude_persistence_change_as_error" = false
"hs_security_selfsignedcert_penalty" = "1.0"
"conn_server_lossy_total_rexmt_threshold" = "50"
"conn_lossy_timeo_rexmt_threshold" = "20"
"apdex_rtt_tolerated_factor" = "4.0"
"hs_security_certscore_gt30d" = "5.0"
"hs_security_ssl30_score" = "3.5"
"conn_server_lossy_ooo_threshold" = "50"
"apdex_rum_tolerated_factor" = "4.0"
"hs_security_cipherscore_eq000b" = "0.0"
"hs_security_certscore_le30d" = "4.0"
"exclude_syn_retransmit_as_error" = false
"exclude_dns_policy_drop_as_significant" = false
"apdex_server_rtt_threshold" = "125"
"hs_security_hsts_penalty" = "1.0"
"exclude_server_dns_error_as_error" = false
"hs_security_tls11_score" = "5.0"
"conn_lossy_zero_win_size_event_threshold" = "2"
"apdex_rum_threshold" = "5000"
"hs_pscore_traffic_threshold_l4_server" = "10.0"
"name" = "testSystem-Analytics-Profile"
"hs_security_encalgo_score_none" = "0.0"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"apdex_response_tolerated_factor" = "4.0"
"exclude_tcp_reset_as_error" = false
"hs_max_security_penalty" = "100"
"exclude_invalid_dns_query_as_error" = false
"conn_lossy_ooo_threshold" = "50"
"hs_security_cipherscore_lt128b" = "3.5"
"hs_security_chain_invalidity_penalty" = "1.0"
"hs_security_tls10_score" = "5.0"
"hs_security_certscore_expired" = "0.0"
"apdex_response_threshold" = "500"
"apdex_rtt_threshold" = "250"
"exclude_unsupported_dns_query_as_error" = false
}
`

const testAccAVIAnalyticsProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
resource "avi_analyticsprofile" "testAnalyticsProfile" {
"hs_event_throttle_window" = "1209600"
"disable_se_analytics" = false
"apdex_server_rtt_tolerated_factor" = "4.0"
"hs_security_nonpfs_penalty" = "1.0"
"hs_security_tls12_score" = "5.0"
"exclude_server_tcp_reset_as_error" = false
"hs_min_dos_rate" = "1000"
"exclude_no_dns_record_as_error" = false
"conn_server_lossy_zero_win_size_event_threshold" = "2"
"hs_max_resources_penalty" = "25"
"conn_lossy_total_rexmt_threshold" = "50"
"hs_security_certscore_le07d" = "2.0"
"hs_pscore_traffic_threshold_l4_client" = "10.0"
"exclude_no_valid_gs_member_as_error" = false
"hs_security_encalgo_score_rc4" = "2.5"
"apdex_server_response_threshold" = "400"
"hs_security_cipherscore_ge128b" = "5.0"
"hs_performance_boost" = "0"
"exclude_invalid_dns_domain_as_error" = false
"exclude_http_error_codes" = ["475"]
"hs_max_anomaly_penalty" = "10"
"exclude_gs_down_as_error" = false
"apdex_server_response_tolerated_factor" = "4.0"
"client_log_config" {
"enable_significant_log_collection" = true
"non_significant_log_processing" = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
"significant_log_processing" = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
"filtered_log_processing" = "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"
}
"disable_server_analytics" = false
"conn_server_lossy_timeo_rexmt_threshold" = "20"
"exclude_client_close_before_request_as_error" = true
"hs_security_weak_signature_algo_penalty" = "1.0"
"exclude_persistence_change_as_error" = false
"hs_security_selfsignedcert_penalty" = "1.0"
"conn_server_lossy_total_rexmt_threshold" = "50"
"conn_lossy_timeo_rexmt_threshold" = "20"
"apdex_rtt_tolerated_factor" = "4.0"
"hs_security_certscore_gt30d" = "5.0"
"hs_security_ssl30_score" = "3.5"
"conn_server_lossy_ooo_threshold" = "50"
"apdex_rum_tolerated_factor" = "4.0"
"hs_security_cipherscore_eq000b" = "0.0"
"hs_security_certscore_le30d" = "4.0"
"exclude_syn_retransmit_as_error" = false
"exclude_dns_policy_drop_as_significant" = false
"apdex_server_rtt_threshold" = "125"
"hs_security_hsts_penalty" = "1.0"
"exclude_server_dns_error_as_error" = false
"hs_security_tls11_score" = "5.0"
"conn_lossy_zero_win_size_event_threshold" = "2"
"apdex_rum_threshold" = "5000"
"hs_pscore_traffic_threshold_l4_server" = "10.0"
"name" = "testSystem-Analytics-Profile-abc"
"hs_security_encalgo_score_none" = "0.0"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"apdex_response_tolerated_factor" = "4.0"
"exclude_tcp_reset_as_error" = false
"hs_max_security_penalty" = "100"
"exclude_invalid_dns_query_as_error" = false
"conn_lossy_ooo_threshold" = "50"
"hs_security_cipherscore_lt128b" = "3.5"
"hs_security_chain_invalidity_penalty" = "1.0"
"hs_security_tls10_score" = "5.0"
"hs_security_certscore_expired" = "0.0"
"apdex_response_threshold" = "500"
"apdex_rtt_threshold" = "250"
"exclude_unsupported_dns_query_as_error" = false
}
`
