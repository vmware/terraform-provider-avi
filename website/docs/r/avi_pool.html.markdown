<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
resource "avi_pool" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the pool. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `analytics_policy` - (Optional) Determines analytics settings for the pool. Field introduced in 18.1.5, 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `analytics_profile_ref` - (Optional) Specifies settings related to analytics. It is a reference to an object of type analyticsprofile. Field introduced in 18.1.4,18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `append_port` - (Optional) Allows the option to append port to hostname in the host header while sending a request to the server. By default, port is appended for non-default ports. This setting will apply for pool's 'rewrite host header to server name', 'rewrite host header to sni' features and server's 'rewrite host header' settings as well as http healthmonitors attached to pools. Enum options - NON_DEFAULT_80_443, NEVER, ALWAYS. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- never), basic edition(allowed values- never), enterprise with cloud services edition. Special default for essentials edition is never, basic edition is never, enterprise is non_default_80_443.
* `application_persistence_profile_ref` - (Optional) Persistence will ensure the same user sticks to the same server for a desired duration of time. It is a reference to an object of type applicationpersistenceprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `autoscale_launch_config_ref` - (Optional) If configured then avi will trigger orchestration of pool server creation and deletion. It is a reference to an object of type autoscalelaunchconfig. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `autoscale_networks` - (Optional) Network ids for the launch configuration. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `autoscale_policy_ref` - (Optional) Reference to server autoscale policy. It is a reference to an object of type serverautoscalepolicy. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `capacity_estimation` - (Optional) Inline estimation of capacity of servers. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `capacity_estimation_ttfb_thresh` - (Optional) The maximum time-to-first-byte of a server. Allowed values are 1-5000. Special values are 0 - automatic. Unit is milliseconds. Allowed in enterprise edition with any value, essentials edition(allowed values- 0), basic edition(allowed values- 0), enterprise with cloud services edition.
* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for pool. Internally set by cloud connector. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `conn_pool_properties` - (Optional) Connnection pool properties. Field introduced in 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `connection_ramp_duration` - (Optional) Duration for which new connections will be gradually ramped up to a server recently brought online. Useful for lb algorithms that are least connection based. Allowed values are 1-300. Special values are 0 - immediate. Unit is min. Allowed in enterprise edition with any value, essentials edition(allowed values- 0), basic edition(allowed values- 0), enterprise with cloud services edition. Special default for essentials edition is 0, basic edition is 0, enterprise is 10.
* `created_by` - (Optional) Creator name. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `default_server_port` - (Optional) Traffic sent to servers will use this destination server port unless overridden by the server's specific port attribute. The ssl checkbox enables avi to server encryption. Allowed values are 1-65535. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `delete_server_on_dns_refresh` - (Optional) Indicates whether existing ips are disabled(false) or deleted(true) on dns hostname refreshdetail -- on a dns refresh, some ips set on pool may no longer be returned by the resolver. These ips are deleted from the pool when this knob is set to true. They are disabled, if the knob is set to false. Field introduced in 18.2.3. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition.
* `description` - (Optional) A description of the pool. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `domain_name` - (Optional) Comma separated list of domain names which will be used to verify the common names or subject alternative names presented by server certificates. It is performed only when common name check host_check_enabled is enabled. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `east_west` - (Optional) Inherited config from virtualservice. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_http2` - (Optional) Enable http/2 for traffic from virtualservice to all backend servers in this pool. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `enabled` - (Optional) Enable or disable the pool. Disabling will terminate all open connections and pause health monitors. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `external_autoscale_groups` - (Optional) Names of external auto-scale groups for pool servers. Currently available only for aws and azure. Field introduced in 17.1.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `fail_action` - (Optional) Enable an action - close connection, http redirect or local http response - when a pool failure happens. By default, a connection will be closed, in case the pool experiences a failure. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `fewest_tasks_feedback_delay` - (Optional) Periodicity of feedback for fewest tasks server selection algorithm. Allowed values are 1-300. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `graceful_disable_timeout` - (Optional) Used to gracefully disable a server. Virtual service waits for the specified time before terminating the existing connections  to the servers that are disabled. Allowed values are 1-7200. Special values are 0 - immediate, -1 - infinite. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `gslb_sp_enabled` - (Optional) Indicates if the pool is a site-persistence pool. Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `health_monitor_refs` - (Optional) Verify server health by applying one or more health monitors. Active monitors generate synthetic traffic from each service engine and mark a server up or down based on the response. The passive monitor listens only to client to server communication. It raises or lowers the ratio of traffic destined to a server based on successful responses. It is a reference to an object of type healthmonitor. Maximum of 50 items allowed. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `horizon_profile` - (Optional) Horizon uag configuration. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `host_check_enabled` - (Optional) Enable common name check for server certificate. If enabled and no explicit domain name is specified, avi will use the incoming host header to do the match. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `http2_properties` - (Optional) Http2 pool properties. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `ignore_server_port` - (Optional) Ignore the server port in building the load balancing state.applicable only for consistent hash load balancing algorithm or disable port translation (use_service_port) use cases. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `inline_health_monitor` - (Optional) The passive monitor will monitor client to server connections and requests and adjust traffic load to servers based on successful responses. This may alter the expected behavior of the lb method, such as round robin. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ipaddrgroup_ref` - (Optional) Use list of servers from ip address group. It is a reference to an object of type ipaddrgroup. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `lb_algo_rr_per_se` - (Optional) Do round robin load load balancing at se level instead of the default per core load balancing. Field introduced in 21.1.5, 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `lb_algorithm` - (Optional) The load balancing algorithm will pick a server within the pool's list of available servers. Values lb_algorithm_nearest_server and lb_algorithm_topology are only allowed for gslb pool. Enum options - LB_ALGORITHM_LEAST_CONNECTIONS, LB_ALGORITHM_ROUND_ROBIN, LB_ALGORITHM_FASTEST_RESPONSE, LB_ALGORITHM_CONSISTENT_HASH, LB_ALGORITHM_LEAST_LOAD, LB_ALGORITHM_FEWEST_SERVERS, LB_ALGORITHM_RANDOM, LB_ALGORITHM_FEWEST_TASKS, LB_ALGORITHM_NEAREST_SERVER, LB_ALGORITHM_CORE_AFFINITY, LB_ALGORITHM_TOPOLOGY. Allowed in enterprise edition with any value, essentials edition(allowed values- lb_algorithm_least_connections,lb_algorithm_round_robin,lb_algorithm_consistent_hash), basic edition(allowed values- lb_algorithm_least_connections,lb_algorithm_round_robin,lb_algorithm_consistent_hash), enterprise with cloud services edition.
* `lb_algorithm_consistent_hash_hdr` - (Optional) Http header name to be used for the hash key. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `lb_algorithm_core_nonaffinity` - (Optional) Degree of non-affinity for core affinity based server selection. Allowed values are 1-65535. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials edition(allowed values- 2), basic edition(allowed values- 2), enterprise with cloud services edition.
* `lb_algorithm_hash` - (Optional) Criteria used as a key for determining the hash between the client and  server. Enum options - LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS, LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS_AND_PORT, LB_ALGORITHM_CONSISTENT_HASH_URI, LB_ALGORITHM_CONSISTENT_HASH_CUSTOM_HEADER, LB_ALGORITHM_CONSISTENT_HASH_CUSTOM_STRING, LB_ALGORITHM_CONSISTENT_HASH_CALLID. Allowed in enterprise edition with any value, essentials edition(allowed values- lb_algorithm_consistent_hash_source_ip_address), basic edition(allowed values- lb_algorithm_consistent_hash_source_ip_address), enterprise with cloud services edition.
* `lookup_server_by_name` - (Optional) Allow server lookup by name. Field introduced in 17.1.11,17.2.4. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `max_concurrent_connections_per_server` - (Optional) The maximum number of concurrent connections allowed to each server within the pool. Note  applied value will be no less than the number of service engines that the pool is placed on. If set to 0, no limit is applied. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_conn_rate_per_server` - (Optional) Rate limit connections to each server. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `min_health_monitors_up` - (Optional) Minimum number of health monitors in up state to mark server up. Field introduced in 18.2.1, 17.2.12. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `min_servers_up` - (Optional) Minimum number of servers in up state for marking the pool up. Field introduced in 18.2.1, 17.2.12. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `networks` - (Optional) (internal-use) networks designated as containing servers for this pool. The servers may be further narrowed down by a filter. This field is used internally by avi, not editable by the user. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nsx_securitygroup` - (Optional) A list of nsx groups where the servers for the pool are created. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pki_profile_ref` - (Optional) Avi will validate the ssl certificate present by a server against the selected pki profile. It is a reference to an object of type pkiprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `placement_networks` - (Optional) Manually select the networks and subnets used to provide reachability to the pool's servers. Specify the subnet using the following syntax  10-1-1-0/24. Use static routes in vrf configuration when pool servers are not directly connected but routable from the service engine. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pool_type` - (Optional) Type or purpose, the pool is to be used for. Enum options - POOL_TYPE_GENERIC_APP, POOL_TYPE_OAUTH. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `request_queue_depth` - (Optional) Minimum number of requests to be queued when pool is full. Allowed in enterprise edition with any value, essentials edition(allowed values- 128), basic edition(allowed values- 128), enterprise with cloud services edition.
* `request_queue_enabled` - (Optional) Enable request queue when pool is full. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `resolve_pool_by_dns` - (Optional) This field is used as a flag to create a job for jobmanager. Field introduced in 18.2.10,20.1.2. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `rewrite_host_header_to_server_name` - (Optional) Rewrite incoming host header to server name of the server to which the request is proxied. Enabling this feature rewrites host header for requests to all servers in the pool. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `rewrite_host_header_to_sni` - (Optional) If sni server name is specified, rewrite incoming host header to the sni server name. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `routing_pool` - (Optional) Enable to do routing when this pool is selected to send traffic. No servers present in routing pool. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `server_disable_type` - (Optional) Server graceful disable timeout behaviour. Enum options - DISALLOW_NEW_CONNECTION, ALLOW_NEW_CONNECTION_IF_PERSISTENCE_PRESENT. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `server_name` - (Optional) Fully qualified dns hostname which will be used in the tls sni extension in server connections if sni is enabled. If no value is specified, avi will use the incoming host header instead. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `server_reselect` - (Optional) Server reselect configuration for http requests. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `server_timeout` - (Optional) Server timeout value specifies the time within which a server connection needs to be established and a request-response exchange completes between avi and the server. Value of 0 results in using default timeout of 60 minutes. Allowed values are 0-21600000. Field introduced in 18.1.5,18.2.1. Unit is milliseconds. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `servers` - (Optional) The pool directs load balanced traffic to this list of destination servers. The servers can be configured by ip address, name, network or via ip address group. Maximum of 5000 items allowed. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `service_metadata` - (Optional) Metadata pertaining to the service provided by this pool. In openshift/kubernetes environments, app metadata info is stored. Any user input to this field will be overwritten by avi vantage. Field introduced in 17.2.14,18.1.5,18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `sni_enabled` - (Optional) Enable tls sni for server connections. If disabled, avi will not send the sni extension as part of the handshake. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `sp_gs_info` - (Optional) Gslb service associated with the site persistence pool. Field introduced in 22.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `ssl_key_and_certificate_ref` - (Optional) Service engines will present a client ssl certificate to the server. It is a reference to an object of type sslkeyandcertificate. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ssl_profile_ref` - (Optional) When enabled, avi re-encrypts traffic to the backend servers. The specific ssl profile defines which ciphers and ssl versions will be supported. It is a reference to an object of type sslprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tier1_lr` - (Optional) This tier1_lr field should be set same as virtualservice associated for nsx-t. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `use_service_port` - (Optional) Do not translate the client's destination port when sending the connection to the server. Monitor port needs to be specified for health monitors. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic, enterprise with cloud services edition.
* `use_service_ssl_mode` - (Optional) This applies only when use_service_port is set to true. If enabled, ssl mode of the connection to the server is decided by the ssl mode on the virtualservice service port, on which the request was received. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vrf_ref` - (Optional) Virtual routing context that the pool is bound to. This is used to provide the isolation of the set of networks the pool is attached to. The pool inherits the virtual routing conext of the virtual service, and this field is used only internally, and is set by pb-transform. It is a reference to an object of type vrfcontext. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the pool. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

