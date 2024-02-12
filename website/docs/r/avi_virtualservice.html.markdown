<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
resource "avi_virtualservice" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the virtual service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `active_standby_se_tag` - (Optional) This configuration only applies if the virtualservice is in legacy active standby ha mode and load distribution among active standby is enabled. This field is used to tag the virtualservice so that virtualservices with the same tag will share the same active serviceengine. Virtualservices with different tags will have different active serviceengines. If one of the serviceengine's in the serviceenginegroup fails, all virtualservices will end up using the same active serviceengine. Redistribution of the virtualservices can be either manual or automated when the failed serviceengine recovers. Redistribution is based on the auto redistribute property of the serviceenginegroup. Enum options - ACTIVE_STANDBY_SE_1, ACTIVE_STANDBY_SE_2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `advertise_down_vs` - (Optional) Keep advertising virtual service via bgp even if it is marked down by health monitor. This setting takes effect for future virtual service flaps. To advertise current vses that are down, please disable and re-enable the virtual service. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `allow_invalid_client_cert` - (Optional) Process request even if invalid client certificate is presented. Datascript apis need to be used for processing of such requests. Field introduced in 18.2.3. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `analytics_policy` - (Optional) Determines analytics settings for the application. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `analytics_profile_ref` - (Optional) Specifies settings related to analytics. It is a reference to an object of type analyticsprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `application_profile_ref` - (Optional) Enable application layer specific features for the virtual service. It is a reference to an object of type applicationprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition. Special default for essentials edition is system-l4-application.
* `azure_availability_set` - (Optional) (internal-use)applicable for azure only. Azure availability set to which this vs is associated. Internally set by the cloud connector. Field introduced in 17.2.12, 18.1.2. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `bgp_local_preference` - (Optional) Local_pref to be used for vsvip advertised. Applicable only over ibgp. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `bgp_num_as_path_prepend` - (Optional) Number of times the local as should be prepended additionally to vsvip. Applicable only over ebgp. Allowed values are 1-10. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `bgp_peer_labels` - (Optional) Select bgp peers, using peer label, for vsvip advertisement. Field introduced in 20.1.5. Maximum of 128 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `bot_policy_ref` - (Optional) Bot detection policy for the virtual service. It is a reference to an object of type botdetectionpolicy. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `bulk_sync_kvcache` - (Optional) (this is a beta feature). Sync key-value cache to the new ses when vs is scaled out. For ex  ssl sessions are stored using vs's key-value cache. When the vs is scaled out, the ssl session information is synced to the new se, allowing existing ssl sessions to be reused on the new se. Field introduced in 17.2.7, 18.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `close_client_conn_on_config_update` - (Optional) Close client connection on vs config update. Field introduced in 17.2.4. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for vs. Internally set by cloud connector. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cloud_type` - (Optional) Enum options - cloud_none, cloud_vcenter, cloud_openstack, cloud_aws, cloud_vca, cloud_apic, cloud_mesos, cloud_linuxserver, cloud_docker_ucp, cloud_rancher, cloud_oshift_k8s, cloud_azure, cloud_gcp, cloud_nsxt. Allowed in enterprise edition with any value, essentials edition(allowed values- cloud_none,cloud_vcenter), basic edition(allowed values- cloud_none,cloud_nsxt), enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `connections_rate_limit` - (Optional) Rate limit the incoming connections to this virtual service. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `content_rewrite` - (Optional) Profile used to match and rewrite strings in request and/or response body. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `created_by` - (Optional) Creator name. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `csrf_policy_ref` - (Optional) Csrf protection policy for the virtual service. It is a reference to an object of type csrfpolicy. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `delay_fairness` - (Optional) Select the algorithm for qos fairness. This determines how multiple virtual services sharing the same service engines will prioritize traffic over a congested network. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dns_info` - (Optional) Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record. Note that only one of fqdn and dns_info setting is allowed. Maximum of 1000 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `dns_policies` - (Optional) Dns policies applied on the dns traffic of the virtual service. Field introduced in 17.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `east_west_placement` - (Optional) Force placement on all se's in service group (mesos mode only). Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `enable_autogw` - (Optional) Response traffic to clients will be sent back to the source mac address of the connection, rather than statically sent to a default gateway. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition. Special default for essentials edition is false, basic edition is false, enterprise is true.
* `enable_rhi` - (Optional) Enable route health injection using the bgp config in the vrf context. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `enable_rhi_snat` - (Optional) Enable route health injection for source nat'ted floating ip address using the bgp config in the vrf context. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `enable_session` - (Optional) Enable http sessions for this virtual service. If enabled, a session cookie will be added to http responses and persistent key-value store will be activated. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `enabled` - (Optional) Enable or disable the virtual service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `error_page_profile_ref` - (Optional) Error page profile to be used for this virtualservice.this profile is used to send the custom error page to the client generated by the proxy. It is a reference to an object of type errorpageprofile. Field introduced in 17.2.4. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `flow_dist` - (Optional) Criteria for flow distribution among ses. Enum options - LOAD_AWARE, CONSISTENT_HASH_SOURCE_IP_ADDRESS, CONSISTENT_HASH_SOURCE_IP_ADDRESS_AND_PORT. Allowed in enterprise edition with any value, essentials edition(allowed values- load_aware), basic edition(allowed values- load_aware), enterprise with cloud services edition.
* `flow_label_type` - (Optional) Criteria for flow labelling. Enum options - NO_LABEL, APPLICATION_LABEL, SERVICE_LABEL. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `fqdn` - (Optional) Dns resolvable, fully qualified domain name of the virtualservice. Only one of 'fqdn' and 'dns_info' configuration is allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `host_name_xlate` - (Optional) Translate the host name sent to the servers to this value. Translate the host name sent from servers back to the value used by the client. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `http_policies` - (Optional) Http policies applied on the data traffic of the virtual service. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `icap_request_profile_refs` - (Optional) The config settings for the icap server when checking the http request. It is a reference to an object of type icapprofile. Field introduced in 20.1.1. Maximum of 1 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ign_pool_net_reach` - (Optional) Ignore pool servers network reachability constraints for virtual service placement. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `jwt_config` - (Optional) Application-specific config for jwt validation. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `l4_policies` - (Optional) L4 policies applied to the data traffic of the virtual service. Field introduced in 17.2.7. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ldap_vs_config` - (Optional) Application-specific ldap config. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `limit_doser` - (Optional) Limit potential dos attackers who exceed max_cps_per_client significantly to a fraction of max_cps_per_client for a while. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `max_cps_per_client` - (Optional) Maximum connections per second per client ip. Allowed values are 10-1000. Special values are 0- unlimited. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `microservice_ref` - (Optional) Microservice representing the virtual service. It is a reference to an object of type microservice. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `min_pools_up` - (Optional) Minimum number of up pools to mark vs up. Field introduced in 18.2.1, 17.2.12. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `network_profile_ref` - (Optional) Determines network settings such as protocol, tcp or udp, and related options for the protocol. It is a reference to an object of type networkprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition. Special default for essentials edition is system-tcp-fast-path.
* `network_security_policy_ref` - (Optional) Network security policies for the virtual service. It is a reference to an object of type networksecuritypolicy. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nsx_securitygroup` - (Optional) A list of nsx groups representing the clients which can access the virtual ip of the virtual service. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `oauth_vs_config` - (Optional) Virtualservice specific oauth config. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `performance_limits` - (Optional) Optional settings that determine performance limits like max connections or bandwdith etc. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pool_group_ref` - (Optional) The pool group is an object that contains pools. It is a reference to an object of type poolgroup. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `pool_ref` - (Optional) The pool is an object that contains destination servers and related attributes such as load-balancing and persistence. It is a reference to an object of type pool. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `remove_listening_port_on_vs_down` - (Optional) Remove listening port if virtualservice is down. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `requests_rate_limit` - (Optional) Rate limit the incoming requests to this virtual service. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `revoke_vip_route` - (Optional) Revoke the advertisement of virtual service via the cloud if it is marked down by health monitor. Supported for nsxt clouds only.this setting takes effect for future virtual service flaps. To advertise current vses that are down, please disable and re-enable the virtual service. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `saml_sp_config` - (Optional) Application-specific saml config. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `scaleout_ecmp` - (Optional) Disable re-distribution of flows across service engines for a virtual service. Enable if the network itself performs flow hashing with ecmp in environments such as gcp. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_group_ref` - (Optional) The service engine group to use for this virtual service. Moving to a new se group is disruptive to existing connections for this vs. It is a reference to an object of type serviceenginegroup. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `security_policy_ref` - (Optional) Security policy applied on the traffic of the virtual service. This policy is used to perform security actions such as distributed denial of service (ddos) attack mitigation, etc. It is a reference to an object of type securitypolicy. Field introduced in 18.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `server_network_profile_ref` - (Optional) Determines the network settings profile for the server side of tcp proxied connections. Leave blank to use the same settings as the client to vs side of the connection. It is a reference to an object of type networkprofile. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `service_metadata` - (Optional) Metadata pertaining to the service provided by this virtual service. In openshift/kubernetes environments, egress pod info is stored. Any user input to this field will be overwritten by avi vantage. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `service_pool_select` - (Optional) Select pool based on destination port. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `services` - (Optional) List of services defined for this virtual service. Maximum of 2048 items allowed. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `sideband_profile` - (Optional) Sideband configuration to be used for this virtualservice.it can be used for sending traffic to sideband vips for external inspection etc. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `snat_ip` - (Optional) Nat'ted floating source ip address(es) for upstream connection to servers. Maximum of 32 items allowed. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `snat_ip6_addresses` - (Optional) Ipv6 address for se snat. Field introduced in 30.2.1. Maximum of 32 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `sp_pool_refs` - (Optional) Gslb pools used to manage site-persistence functionality. Each site-persistence pool contains the virtualservices in all the other sites, that is auto-generated by the gslb manager. This is a read-only field for the user. It is a reference to an object of type pool. Field introduced in 17.2.2. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `ssl_key_and_certificate_refs` - (Optional) Select or create one or two certificates, ec and/or rsa, that will be presented to ssl/tls terminated connections. It is a reference to an object of type sslkeyandcertificate. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ssl_profile_ref` - (Optional) Determines the set of ssl versions and ciphers to accept for ssl/tls terminated connections. It is a reference to an object of type sslprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ssl_profile_selectors` - (Optional) Select ssl profile based on client ip address match. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ssl_sess_cache_avg_size` - (Optional) Expected number of ssl session cache entries (may be exceeded). Allowed values are 1024-16383. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `sso_policy_ref` - (Optional) The sso policy attached to the virtualservice. It is a reference to an object of type ssopolicy. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `static_dns_records` - (Optional) List of static dns records applied to this virtual service. These are static entries and no health monitoring is performed against the ip addresses. Maximum of 1000 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `test_se_datastore_level_1_ref` - (Optional) Used for testing se datastore upgrade 2.0 functionality. It is a reference to an object of type testsedatastorelevel1. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `topology_policies` - (Optional) Topology policies applied on the dns traffic of the virtual service based ongslb topology algorithm. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `traffic_clone_profile_ref` - (Optional) Server network or list of servers for cloning traffic. It is a reference to an object of type trafficcloneprofile. Field introduced in 17.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `traffic_enabled` - (Optional) Knob to enable the virtual service traffic on its assigned service engines. This setting is effective only when the enabled flag is set to true. Field introduced in 17.2.8. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - (Optional) Specify if this is a normal virtual service, or if it is the parent or child of an sni-enabled virtual hosted virtual service. Enum options - VS_TYPE_NORMAL, VS_TYPE_VH_PARENT, VS_TYPE_VH_CHILD. Allowed in enterprise edition with any value, essentials edition(allowed values- vs_type_normal), basic edition(allowed values- vs_type_normal,vs_type_vh_parent), enterprise with cloud services edition.
* `use_bridge_ip_as_vip` - (Optional) Use bridge ip as vip on each host in mesos deployments. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `use_vip_as_snat` - (Optional) Use the virtual ip as the snat ip for health monitoring and sending traffic to the backend servers instead of the service engine interface ip. The caveat of enabling this option is that the virtualservice cannot be configued in an active-active ha mode. Dns based multi vip solution has to be used for ha & non-disruptive upgrade purposes. Field introduced in 17.1.9,17.2.3. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic, enterprise with cloud services edition.
* `vh_domain_name` - (Optional) The exact name requested from the client's sni-enabled tls hello domain name field. If this is a match, the parent vs will forward the connection to this child vs. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vh_matches` - (Optional) Match criteria to select this child vs. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vh_parent_vs_ref` - (Optional) Specifies the virtual service acting as virtual hosting (sni) parent. It is a reference to an object of type virtualservice. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vh_type` - (Optional) Specify if the virtual hosting vs is of type sni or enhanced. Enum options - VS_TYPE_VH_SNI, VS_TYPE_VH_ENHANCED. Field introduced in 20.1.3. Allowed in enterprise edition with any value, basic edition(allowed values- vs_type_vh_sni,vs_type_vh_enhanced), enterprise with cloud services edition.
* `vip` - (Optional) List of virtual service ips. While creating a 'shared vs',please use vsvip_ref to point to the shared entities. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vrf_context_ref` - (Optional) Virtual routing context that the virtual service is bound to. This is used to provide the isolation of the set of networks the application is attached to. It is a reference to an object of type vrfcontext. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_datascripts` - (Optional) Datascripts applied on the data traffic of the virtual service. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `vsvip_cloud_config_cksum` - (Optional) Checksum of cloud configuration for vsvip. Internally set by cloud connector. Field introduced in 17.2.9, 18.1.2. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `vsvip_ref` - (Optional) Mostly used during the creation of shared vs, this field refers to entities that can be shared across virtual services. It is a reference to an object of type vsvip. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `waf_policy_ref` - (Optional) Waf policy for the virtual service. It is a reference to an object of type wafpolicy. Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `weight` - (Optional) The quality of service weight to assign to traffic transmitted from this virtual service. A higher weight will prioritize traffic versus other virtual services sharing the same service engines. Allowed values are 1-128. Allowed in enterprise edition with any value, essentials edition(allowed values- 1), basic edition(allowed values- 1), enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the virtualservice. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

