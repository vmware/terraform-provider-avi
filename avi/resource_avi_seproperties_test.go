package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
	"github.com/vmware/alb-sdk/go/models"
)

func TestAVISePropertiesBasic(t *testing.T) {
	var seProp models.SeProperties
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISePropertiesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISePropertiesExists("avi_seproperties.testSeProperties", &seProp),
					resource.TestCheckResourceAttr(
						"avi_seproperties.testSeProperties", "uuid", "global"),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.dp_enq_interval_msec", 20),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.controller_echo_rpc_aggressive_timeout", 2000),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.dp_max_wait_rsp_time_sec", 60),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.dp_batch_size", 100),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_runtime_properties.log_agent_max_active_adf_files_per_vs", 100),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_runtime_properties.dp_aggressive_hb_timeout_count", 10),
				),
			},
			{
				Config: testAccAVISePropertiesupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISePropertiesExists("avi_seproperties.testSeProperties", &seProp),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.dp_enq_interval_msec", 27),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.controller_echo_rpc_aggressive_timeout", 2010),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.dp_max_wait_rsp_time_sec", 100),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_agent_properties.dp_batch_size", 120),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_runtime_properties.log_agent_max_active_adf_files_per_vs", 150),
					testAccCheckAVISePropertiesValuesUpdated(&seProp, "se_runtime_properties.dp_aggressive_hb_timeout_count", 8),
				),
			},
			{
				ResourceName:      "avi_seproperties.testSeProperties",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVISePropertiesConfig,
			},
		},
	})

}

func testAccCheckAVISePropertiesExists(resourcename string, seProp *models.SeProperties) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj *models.SeProperties
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI SeProperties ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		*seProp = *obj
		return nil
	}
}

func testAccCheckAVISePropertiesValuesUpdated(seProp *models.SeProperties, key string, val interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if key == "se_agent_properties.dp_enq_interval_msec" && *seProp.SeAgentProperties.DpEnqIntervalMsec != int32(val.(int)) {
			return fmt.Errorf("bad dp_enq_interval_msec, expected \"%v\", got: %#v", val, *seProp.SeAgentProperties.DpEnqIntervalMsec)
		}
		if key == "se_agent_properties.controller_echo_rpc_aggressive_timeout" && *seProp.SeAgentProperties.ControllerEchoRPCAggressiveTimeout != int32(val.(int)) {
			return fmt.Errorf("bad controller_echo_rpc_aggressive_timeout, expected \"%v\", got: %#v", val, *seProp.SeAgentProperties.ControllerEchoRPCAggressiveTimeout)
		}
		if key == "se_agent_properties.dp_max_wait_rsp_time_sec" && *seProp.SeAgentProperties.DpMaxWaitRspTimeSec != int32(val.(int)) {
			return fmt.Errorf("bad dp_max_wait_rsp_time_sec, expected \"%v\", got: %#v", val, *seProp.SeAgentProperties.DpMaxWaitRspTimeSec)
		}
		if key == "se_agent_properties.dp_batch_size" && *seProp.SeAgentProperties.DpBatchSize != int32(val.(int)) {
			return fmt.Errorf("bad dp_batch_size, expected \"%v\", got: %#v", val, *seProp.SeAgentProperties.DpBatchSize)
		}
		if key == "se_runtime_properties.log_agent_max_active_adf_files_per_vs" && *seProp.SeRuntimeProperties.LogAgentMaxActiveAdfFilesPerVs != int32(val.(int)) {
			return fmt.Errorf("bad log_agent_max_active_adf_files_per_vs, expected \"%v\", got: %#v", val, *seProp.SeRuntimeProperties.LogAgentMaxActiveAdfFilesPerVs)
		}
		if key == "se_runtime_properties.dp_aggressive_hb_timeout_count" && *seProp.SeRuntimeProperties.DpAggressiveHbTimeoutCount != int32(val.(int)) {
			return fmt.Errorf("bad dp_aggressive_hb_timeout_count, expected \"%v\", got: %#v", val, *seProp.SeRuntimeProperties.DpAggressiveHbTimeoutCount)
		}
		return nil
	}
}

//nolint
const testAccAVISePropertiesConfig = `
resource "avi_seproperties" "testSeProperties" {
	se_agent_properties {
		dp_enq_interval_msec = "20"
		sdb_scan_count = "1000"
		dp_reg_pending_max_wait_time = "75"
		controller_echo_rpc_aggressive_timeout = "2000"
		dp_aggressive_enq_interval_msec = "1"
		vnic_ip_delete_interval = "5"
		dp_max_wait_rsp_time_sec = "60"
		controller_rpc_timeout = "10"
		controller_heartbeat_miss_limit = "6"
		dp_aggressive_deq_interval_msec = "1"
		controller_echo_rpc_timeout = "2000"
		ignore_docker_mac_change = true
		sdb_pipeline_size = "100"
		cpustats_interval = "5"
		dp_batch_size = "100"
		vnic_probe_interval = "5"
		debug_mode = false
		vnic_dhcp_ip_max_retries = "10"
		dp_deq_interval_msec = "20"
		ctrl_reg_pending_max_wait_time = "150"
		headless_timeout_sec = "0"
		sdb_flush_interval = "100"
		controller_echo_miss_aggressive_limit = "2"
		controller_heartbeat_timeout_sec = "12"
		controller_registration_timeout_sec = "10"
		controller_echo_miss_limit = "4"
		vnic_dhcp_ip_check_interval = "6"
	}
	se_runtime_properties {
		log_agent_max_active_adf_files_per_vs = "100"
		se_auth_ldap_conns_per_server = "1"
		log_agent_file_sz_appl = "4"
		se_packet_buffer_max = "0"
		log_agent_max_logmessage_proto_sz = "65536"
		se_hb_persist_fudge_bits = "3"
		dp_aggressive_hb_timeout_count = "10"
		se_metrics_rt_interval = "1000"
		persistence_mem_max = "0"
		upstream_connpool_core_max_cache = "-1"
		log_agent_log_storage_min_sz = "1024"
		se_random_tcp_drops = false
		se_dp_if_state_poll_interval = "10"
		se_auth_ldap_cache_size = "100000"
		log_message_max_file_list_size = "64"
		services_accessible_all_interfaces = false
		dupip_timeout_count = "5"
		baremetal_dispatcher_handles_flows = false
		upstream_connpool_cache_thresh = "-1"
		connections_lossy_log_rate_limiter_threshold = "1000"
		log_agent_unknown_vs_timer = "1800"
		upstream_connpool_strategy = "-1"
		upstream_connpool_conn_idle_thresh_tmo = "-1"
		log_agent_min_storage_per_vs = "10"
		feproxy_vips_enable_proxy_arp = true
		service_port_ranges {
	start = "10000"
	end = "20000"
}
		scaleout_udp_per_pkt = true
		se_dp_log_nf_enqueue_percent = "70"
		http_rum_min_content_length = "64"
		flow_table_batch_push_frequency = "5"
		log_agent_file_sz_conn = "4"
		se_mac_error_threshold_to_disable_promiscious = "1000"
		disable_flow_probes = false
		lbaction_rq_per_request_max_retries = "22"
		lbaction_num_requests_to_dispatch = "4"
		log_agent_file_sz_debug = "4"
		connections_udfnf_log_rate_limiter_threshold = "1000"
		upstream_send_timeout = "3600000"
		log_agent_max_storage_ignore_percent = "20.0"
		log_agent_max_storage_excess_percent = "110"
		log_agent_max_concurrent_rsync = "1024"
		log_agent_compress_logs = true
		http_rum_console_log = false
		dp_hb_frequency = "100"
		dupip_frequency = "0"
		se_dp_compression {
			mobile_str = ["iPhone","iPod","Android","BB10","BlackBerry","webOS","IEMobile","iPad","PlayBook","Xoom","P160U","SCH-I800","Nexus 7","Touch"]
			min_length = "128"
			min_high_rtt = "200"
			max_low_rtt = "10"
		}
		dp_aggressive_hb_frequency = "100"
		log_agent_pause_interval = "0"
		user_defined_metric_age = "60"
		se_metrics_rt_enabled = true
		upstream_connpool_enable = true
		se_auth_ldap_connect_timeout = "10000"
		spdy_fwd_proxy_parse_enable = true
		se_auth_ldap_reconnect_timeout = "10000"
		enable_hsm_log = false
		se_auth_ldap_request_timeout = "10000"
		upstream_connect_timeout = "3600000"
		downstream_send_timeout = "3600000"
		log_agent_file_sz_event = "4"
		se_dp_log_udf_enqueue_percent = "90"
		log_agent_export_msg_buffer_size = "524288"
		se_dp_hm_drops = "0"
		se_memory_poison = true
		se_metrics_interval = "60000"
		log_agent_export_wait_time = "100"
		se_auth_ldap_bind_timeout = "5000"
		se_auth_ldap_servers_failover_only = false
		log_agent_conn_send_buffer_size = "16384"
		global_mtu = "0"
		upstream_read_timeout = "3600000"
		se_handle_interface_routes = false
		upstream_keepalive = false
		tcp_syncache_max_retransmit_default = "4"
		log_agent_sleep_interval = "10"
		ngx_free_connection_stack = false
		dp_hb_timeout_count = "10"
		app_headers {
	hdr_name = "Server"
	hdr_match_case = "SENSITIVE"
	hdr_string_op = "EQUALS"
}
app_headers {
	hdr_name = "MicrosoftSharePointTeamServices"
	hdr_match_case = "SENSITIVE"
	hdr_string_op = "EQUALS"
}
	}
	se_bootup_properties {
		docker_backend_portstart = "20480"
		se_log_buffer_chunk_count = "1024"
		se_dp_compression {
			window_size = "4096"
			buf_size = "4096"
			buf_num = "128"
			level_aggressive = "5"
			level_normal = "1"
			hash_size = "16384"
		}
		se_log_buffer_app_blocking_dequeue = false
		se_emulated_cores = "0"
		se_log_buffer_conn_blocking_dequeue = false
		l7_resvd_listen_conns_per_core = "256"
		se_log_buffer_connlog_size = "1024"
		docker_backend_portend = "30720"
		se_log_buffer_events_blocking_dequeue = true
		tcp_syncache_hashsize = "8192"
		geo_db_granularity = "1"
		fair_queueing_enabled = true
		log_agent_debug_enabled = false
		se_ip_encap_ipc = "0"
		se_log_buffer_applog_size = "4096"
		se_log_buffer_events_size = "512"
		ssl_sess_cache_timeout = "86400"
		l7_conns_per_core = "16384"
		ssl_sess_cache_per_vs = "4096"
		log_agent_trace_enabled = true
		se_l3_encap_ipc = "0"
	}
}
`

//nolint
const testAccAVISePropertiesupdatedConfig = `
resource "avi_seproperties" "testSeProperties" {
	se_agent_properties {
		dp_enq_interval_msec = "27"
		sdb_scan_count = "1000"
		dp_reg_pending_max_wait_time = "75"
		controller_echo_rpc_aggressive_timeout = "2010"
		dp_aggressive_enq_interval_msec = "1"
		vnic_ip_delete_interval = "5"
		dp_max_wait_rsp_time_sec = "100"
		controller_rpc_timeout = "10"
		controller_heartbeat_miss_limit = "6"
		dp_aggressive_deq_interval_msec = "1"
		controller_echo_rpc_timeout = "2005"
		ignore_docker_mac_change = true
		sdb_pipeline_size = "100"
		cpustats_interval = "5"
		dp_batch_size = "120"
		vnic_probe_interval = "5"
		debug_mode = false
		vnic_dhcp_ip_max_retries = "10"
		dp_deq_interval_msec = "20"
		ctrl_reg_pending_max_wait_time = "180"
		headless_timeout_sec = "0"
		sdb_flush_interval = "100"
		controller_echo_miss_aggressive_limit = "2"
		controller_heartbeat_timeout_sec = "12"
		controller_registration_timeout_sec = "10"
		controller_echo_miss_limit = "4"
		vnic_dhcp_ip_check_interval = "6"
	}
	se_runtime_properties {
		log_agent_max_active_adf_files_per_vs = "150"
		se_auth_ldap_conns_per_server = "1"
		log_agent_file_sz_appl = "4"
		se_packet_buffer_max = "0"
		log_agent_max_logmessage_proto_sz = "65536"
		se_hb_persist_fudge_bits = "3"
		dp_aggressive_hb_timeout_count = "8"
		se_metrics_rt_interval = "1000"
		persistence_mem_max = "0"
		upstream_connpool_core_max_cache = "-1"
		log_agent_log_storage_min_sz = "1024"
		se_random_tcp_drops = false
		se_dp_if_state_poll_interval = "10"
		se_auth_ldap_cache_size = "100000"
		log_message_max_file_list_size = "64"
		services_accessible_all_interfaces = false
		dupip_timeout_count = "5"
		baremetal_dispatcher_handles_flows = false
		upstream_connpool_cache_thresh = "-1"
		connections_lossy_log_rate_limiter_threshold = "1000"
		log_agent_unknown_vs_timer = "1800"
		upstream_connpool_strategy = "-1"
		upstream_connpool_conn_idle_thresh_tmo = "-1"
		log_agent_min_storage_per_vs = "10"
		feproxy_vips_enable_proxy_arp = true
		service_port_ranges {
	start = "10000"
	end = "20000"
}
		scaleout_udp_per_pkt = true
		se_dp_log_nf_enqueue_percent = "70"
		http_rum_min_content_length = "64"
		flow_table_batch_push_frequency = "5"
		log_agent_file_sz_conn = "4"
		se_mac_error_threshold_to_disable_promiscious = "1000"
		disable_flow_probes = false
		lbaction_rq_per_request_max_retries = "22"
		lbaction_num_requests_to_dispatch = "4"
		log_agent_file_sz_debug = "4"
		connections_udfnf_log_rate_limiter_threshold = "1000"
		upstream_send_timeout = "3600000"
		log_agent_max_storage_ignore_percent = "20.0"
		log_agent_max_storage_excess_percent = "110"
		log_agent_max_concurrent_rsync = "1024"
		log_agent_compress_logs = true
		http_rum_console_log = false
		dp_hb_frequency = "100"
		dupip_frequency = "0"
		se_dp_compression {
			mobile_str = ["iPhone","iPod","Android","BB10","BlackBerry","webOS","IEMobile","iPad","PlayBook","Xoom","P160U","SCH-I800","Nexus 7","Touch"]
			min_length = "128"
			min_high_rtt = "200"
			max_low_rtt = "10"
		}
		dp_aggressive_hb_frequency = "100"
		log_agent_pause_interval = "0"
		user_defined_metric_age = "60"
		se_metrics_rt_enabled = true
		upstream_connpool_enable = true
		se_auth_ldap_connect_timeout = "10000"
		spdy_fwd_proxy_parse_enable = true
		se_auth_ldap_reconnect_timeout = "10000"
		enable_hsm_log = false
		se_auth_ldap_request_timeout = "10000"
		upstream_connect_timeout = "3600000"
		downstream_send_timeout = "3600000"
		log_agent_file_sz_event = "4"
		se_dp_log_udf_enqueue_percent = "90"
		log_agent_export_msg_buffer_size = "524288"
		se_dp_hm_drops = "0"
		se_memory_poison = true
		se_metrics_interval = "60000"
		log_agent_export_wait_time = "100"
		se_auth_ldap_bind_timeout = "5000"
		se_auth_ldap_servers_failover_only = false
		log_agent_conn_send_buffer_size = "16384"
		global_mtu = "0"
		upstream_read_timeout = "3600000"
		se_handle_interface_routes = false
		upstream_keepalive = false
		tcp_syncache_max_retransmit_default = "4"
		log_agent_sleep_interval = "10"
		ngx_free_connection_stack = false
		dp_hb_timeout_count = "10"
		app_headers {
	hdr_name = "Server"
	hdr_match_case = "SENSITIVE"
	hdr_string_op = "EQUALS"
}
app_headers {
	hdr_name = "MicrosoftSharePointTeamServices"
	hdr_match_case = "SENSITIVE"
	hdr_string_op = "EQUALS"
}
	}
	se_bootup_properties {
		docker_backend_portstart = "20480"
		se_log_buffer_chunk_count = "1024"
		se_dp_compression {
			window_size = "4096"
			buf_size = "4096"
			buf_num = "128"
			level_aggressive = "5"
			level_normal = "1"
			hash_size = "16384"
		}
		se_log_buffer_app_blocking_dequeue = false
		se_emulated_cores = "0"
		se_log_buffer_conn_blocking_dequeue = false
		l7_resvd_listen_conns_per_core = "256"
		se_log_buffer_connlog_size = "1024"
		docker_backend_portend = "30720"
		se_log_buffer_events_blocking_dequeue = true
		tcp_syncache_hashsize = "8192"
		geo_db_granularity = "1"
		fair_queueing_enabled = true
		log_agent_debug_enabled = false
		se_ip_encap_ipc = "0"
		se_log_buffer_applog_size = "4096"
		se_log_buffer_events_size = "512"
		ssl_sess_cache_timeout = "86400"
		l7_conns_per_core = "163843"
		ssl_sess_cache_per_vs = "4096"
		log_agent_trace_enabled = true
		se_l3_encap_ipc = "0"
	}
}
`
