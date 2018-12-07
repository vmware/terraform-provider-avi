---
layout: "avi"
page_title: "AVI: avi_pool"
sidebar_current: "docs-avi-datasource-pool"
description: |-
Get information of Avi Pool.
---

# avi_pool

This data source is used to to get avi_pool objects.

## Example Usage

```hcl
data "Pool" "foo_Pool" {
    uuid = "Pool-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search Pool by name.
* `uuid` - (Optional) Search Pool by uuid.
* `cloud_ref` - (Optional) Search Pool by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `apic_epg_name` - Synchronize cisco apic epg members with pool servers.
* `application_persistence_profile_ref` - Persistence will ensure the same user sticks to the same server for a desired duration of time.
* `autoscale_launch_config_ref` - If configured then avi will trigger orchestration of pool server creation and deletion.
* `autoscale_networks` - Network ids for the launch configuration.
* `autoscale_policy_ref` - Reference to server autoscale policy.
* `capacity_estimation` - Inline estimation of capacity of servers.
* `capacity_estimation_ttfb_thresh` - The maximum time-to-first-byte of a server.
* `cloud_config_cksum` - Checksum of cloud configuration for pool.
* `cloud_ref` - It is a reference to an object of type cloud.
* `connection_ramp_duration` - Duration for which new connections will be gradually ramped up to a server recently brought online.
* `created_by` - Creator name.
* `default_server_port` - Traffic sent to servers will use this destination server port unless overridden by the server's specific port attribute.
* `description` - A description of the pool.
* `domain_name` - Comma separated list of domain names which will be used to verify the common names or subject alternative names presented by server certificates.
* `east_west` - Inherited config from virtualservice.
* `enabled` - Enable or disable the pool.
* `external_autoscale_groups` - Names of external auto-scale groups for pool servers.
* `fail_action` - Enable an action - close connection, http redirect or local http response - when a pool failure happens.
* `fewest_tasks_feedback_delay` - Periodicity of feedback for fewest tasks server selection algorithm.
* `graceful_disable_timeout` - Used to gracefully disable a server.
* `health_monitor_refs` - Verify server health by applying one or more health monitors.
* `host_check_enabled` - Enable common name check for server certificate.
* `inline_health_monitor` - The passive monitor will monitor client to server connections and requests and adjust traffic load to servers based on successful responses.
* `ipaddrgroup_ref` - Use list of servers from ip address group.
* `lb_algorithm` - The load balancing algorithm will pick a server within the pool's list of available servers.
* `lb_algorithm_consistent_hash_hdr` - Http header name to be used for the hash key.
* `lb_algorithm_core_nonaffinity` - Degree of non-affinity for core afffinity based server selection.
* `lb_algorithm_hash` - Criteria used as a key for determining the hash between the client and  server.
* `lookup_server_by_name` - Allow server lookup by name.
* `max_concurrent_connections_per_server` - The maximum number of concurrent connections allowed to each server within the pool.
* `max_conn_rate_per_server` - Rate limit connections to each server.
* `name` - The name of the pool.
* `networks` - (internal-use) networks designated as containing servers for this pool.
* `nsx_securitygroup` - A list of nsx service groups where the servers for the pool are created.
* `pki_profile_ref` - Avi will validate the ssl certificate present by a server against the selected pki profile.
* `placement_networks` - Manually select the networks and subnets used to provide reachability to the pool's servers.
* `request_queue_depth` - Minimum number of requests to be queued when pool is full.
* `request_queue_enabled` - Enable request queue when pool is full.
* `rewrite_host_header_to_server_name` - Rewrite incoming host header to server name of the server to which the request is proxied.
* `rewrite_host_header_to_sni` - If sni server name is specified, rewrite incoming host header to the sni server name.
* `server_count` - General description.
* `server_name` - Fully qualified dns hostname which will be used in the tls sni extension in server connections if sni is enabled.
* `server_reselect` - Server reselect configuration for http requests.
* `servers` - The pool directs load balanced traffic to this list of destination servers.
* `sni_enabled` - Enable tls sni for server connections.
* `ssl_key_and_certificate_ref` - Service engines will present a client ssl certificate to the server.
* `ssl_profile_ref` - When enabled, avi re-encrypts traffic to the backend servers.
* `tenant_ref` - It is a reference to an object of type tenant.
* `use_service_port` - Do not translate the client's destination port when sending the connection to the server.
* `uuid` - Uuid of the pool.
* `vrf_ref` - Virtual routing context that the pool is bound to.

