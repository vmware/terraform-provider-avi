<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_virtualservice"
sidebar_current: "docs-avi-datasource-virtualservice"
description: |-
  Get information of Avi VirtualService.
---

# avi_virtualservice

This data source is used to to get avi_virtualservice objects.

## Example Usage

```hcl
data "avi_virtualservice" "foo_virtualservice" {
    uuid = "virtualservice-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VirtualService by name.
* `uuid` - (Optional) Search VirtualService by uuid.
* `cloud_ref` - (Optional) Search VirtualService by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `active_standby_se_tag` - This configuration only applies if the virtualservice is in legacy active standby ha mode and load distribution among active standby is enabled. This field is used to tag the virtualservice so that virtualservices with the same tag will share the same active serviceengine. Virtualservices with different tags will have different active serviceengines. If one of the serviceengine's in the serviceenginegroup fails, all virtualservices will end up using the same active serviceengine. Redistribution of the virtualservices can be either manual or automated when the failed serviceengine recovers. Redistribution is based on the auto redistribute property of the serviceenginegroup. Enum options - ACTIVE_STANDBY_SE_1, ACTIVE_STANDBY_SE_2.
* `advertise_down_vs` - Keep advertising virtual service via bgp even if it is marked down by health monitor. This setting takes effect for future virtual service flaps. To advertise current vses that are down, please disable and re-enable the virtual service. Field introduced in 20.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `allow_invalid_client_cert` - Process request even if invalid client certificate is presented. Datascript apis need to be used for processing of such requests. Field introduced in 18.2.3. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `analytics_policy` - Determines analytics settings for the application.
* `analytics_profile_ref` - Specifies settings related to analytics. It is a reference to an object of type analyticsprofile.
* `application_profile_ref` - Enable application layer specific features for the virtual service. It is a reference to an object of type applicationprofile. Special default for essentials edition is system-l4-application.
* `bgp_peer_labels` - Select bgp peers, using peer label, for vsvip advertisement. Field introduced in 20.1.5. Maximum of 128 items allowed.
* `bot_policy_ref` - Bot detection policy for the virtual service. It is a reference to an object of type botdetectionpolicy. Field introduced in 21.1.1.
* `bulk_sync_kvcache` - (this is a beta feature). Sync key-value cache to the new ses when vs is scaled out. For ex  ssl sessions are stored using vs's key-value cache. When the vs is scaled out, the ssl session information is synced to the new se, allowing existing ssl sessions to be reused on the new se. Field introduced in 17.2.7, 18.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `close_client_conn_on_config_update` - Close client connection on vs config update. Field introduced in 17.2.4. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `cloud_config_cksum` - Checksum of cloud configuration for vs. Internally set by cloud connector.
* `cloud_ref` - It is a reference to an object of type cloud.
* `cloud_type` - Enum options - cloud_none, cloud_vcenter, cloud_openstack, cloud_aws, cloud_vca, cloud_apic, cloud_mesos, cloud_linuxserver, cloud_docker_ucp, cloud_rancher, cloud_oshift_k8s, cloud_azure, cloud_gcp, cloud_nsxt. Allowed in basic(allowed values- cloud_none,cloud_nsxt) edition, essentials(allowed values- cloud_none,cloud_vcenter) edition, enterprise edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `connections_rate_limit` - Rate limit the incoming connections to this virtual service.
* `content_rewrite` - Profile used to match and rewrite strings in request and/or response body.
* `created_by` - Creator name.
* `delay_fairness` - Select the algorithm for qos fairness. This determines how multiple virtual services sharing the same service engines will prioritize traffic over a congested network. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `description` - User defined description for the object.
* `dns_info` - Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record. Note that only one of fqdn and dns_info setting is allowed. Maximum of 1000 items allowed.
* `dns_policies` - Dns policies applied on the dns traffic of the virtual service. Field introduced in 17.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `east_west_placement` - Force placement on all se's in service group (mesos mode only). Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `enable_autogw` - Response traffic to clients will be sent back to the source mac address of the connection, rather than statically sent to a default gateway. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition. Special default for basic edition is false, essentials edition is false, enterprise is true.
* `enable_rhi` - Enable route health injection using the bgp config in the vrf context.
* `enable_rhi_snat` - Enable route health injection for source nat'ted floating ip address using the bgp config in the vrf context.
* `enabled` - Enable or disable the virtual service.
* `error_page_profile_ref` - Error page profile to be used for this virtualservice.this profile is used to send the custom error page to the client generated by the proxy. It is a reference to an object of type errorpageprofile. Field introduced in 17.2.4.
* `flow_dist` - Criteria for flow distribution among ses. Enum options - LOAD_AWARE, CONSISTENT_HASH_SOURCE_IP_ADDRESS, CONSISTENT_HASH_SOURCE_IP_ADDRESS_AND_PORT. Allowed in basic(allowed values- load_aware) edition, essentials(allowed values- load_aware) edition, enterprise edition.
* `flow_label_type` - Criteria for flow labelling. Enum options - NO_LABEL, APPLICATION_LABEL, SERVICE_LABEL.
* `fqdn` - Dns resolvable, fully qualified domain name of the virtualservice. Only one of 'fqdn' and 'dns_info' configuration is allowed.
* `host_name_xlate` - Translate the host name sent to the servers to this value. Translate the host name sent from servers back to the value used by the client.
* `http_policies` - Http policies applied on the data traffic of the virtual service.
* `icap_request_profile_refs` - The config settings for the icap server when checking the http request. It is a reference to an object of type icapprofile. Field introduced in 20.1.1. Maximum of 1 items allowed. Allowed in basic edition, essentials edition, enterprise edition.
* `ign_pool_net_reach` - Ignore pool servers network reachability constraints for virtual service placement.
* `jwt_config` - Application-specific config for jwt validation. Field introduced in 20.1.3.
* `l4_policies` - L4 policies applied to the data traffic of the virtual service. Field introduced in 17.2.7.
* `ldap_vs_config` - Application-specific ldap config. Field introduced in 21.1.1.
* `limit_doser` - Limit potential dos attackers who exceed max_cps_per_client significantly to a fraction of max_cps_per_client for a while.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `max_cps_per_client` - Maximum connections per second per client ip. Allowed values are 10-1000. Special values are 0- unlimited.
* `microservice_ref` - Microservice representing the virtual service. It is a reference to an object of type microservice.
* `min_pools_up` - Minimum number of up pools to mark vs up. Field introduced in 18.2.1, 17.2.12.
* `name` - Name for the virtual service.
* `network_profile_ref` - Determines network settings such as protocol, tcp or udp, and related options for the protocol. It is a reference to an object of type networkprofile. Special default for essentials edition is system-tcp-fast-path.
* `network_security_policy_ref` - Network security policies for the virtual service. It is a reference to an object of type networksecuritypolicy.
* `nsx_securitygroup` - A list of nsx groups representing the clients which can access the virtual ip of the virtual service. Field introduced in 17.1.1.
* `oauth_vs_config` - Virtualservice specific oauth config. Field introduced in 21.1.3.
* `performance_limits` - Optional settings that determine performance limits like max connections or bandwdith etc.
* `pool_group_ref` - The pool group is an object that contains pools. It is a reference to an object of type poolgroup.
* `pool_ref` - The pool is an object that contains destination servers and related attributes such as load-balancing and persistence. It is a reference to an object of type pool.
* `remove_listening_port_on_vs_down` - Remove listening port if virtualservice is down.
* `requests_rate_limit` - Rate limit the incoming requests to this virtual service.
* `saml_sp_config` - Application-specific saml config. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `scaleout_ecmp` - Disable re-distribution of flows across service engines for a virtual service. Enable if the network itself performs flow hashing with ecmp in environments such as gcp.
* `se_group_ref` - The service engine group to use for this virtual service. Moving to a new se group is disruptive to existing connections for this vs. It is a reference to an object of type serviceenginegroup.
* `security_policy_ref` - Security policy applied on the traffic of the virtual service. This policy is used to perform security actions such as distributed denial of service (ddos) attack mitigation, etc. It is a reference to an object of type securitypolicy. Field introduced in 18.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `server_network_profile_ref` - Determines the network settings profile for the server side of tcp proxied connections. Leave blank to use the same settings as the client to vs side of the connection. It is a reference to an object of type networkprofile.
* `service_metadata` - Metadata pertaining to the service provided by this virtual service. In openshift/kubernetes environments, egress pod info is stored. Any user input to this field will be overwritten by avi vantage.
* `service_pool_select` - Select pool based on destination port.
* `services` - List of services defined for this virtual service. Maximum of 2048 items allowed.
* `sideband_profile` - Sideband configuration to be used for this virtualservice.it can be used for sending traffic to sideband vips for external inspection etc.
* `snat_ip` - Nat'ted floating source ip address(es) for upstream connection to servers. Maximum of 32 items allowed.
* `ssl_key_and_certificate_refs` - Select or create one or two certificates, ec and/or rsa, that will be presented to ssl/tls terminated connections. It is a reference to an object of type sslkeyandcertificate.
* `ssl_profile_ref` - Determines the set of ssl versions and ciphers to accept for ssl/tls terminated connections. It is a reference to an object of type sslprofile.
* `ssl_profile_selectors` - Select ssl profile based on client ip address match. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `ssl_sess_cache_avg_size` - Expected number of ssl session cache entries (may be exceeded). Allowed values are 1024-16383.
* `sso_policy_ref` - The sso policy attached to the virtualservice. It is a reference to an object of type ssopolicy. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `static_dns_records` - List of static dns records applied to this virtual service. These are static entries and no health monitoring is performed against the ip addresses. Maximum of 1000 items allowed.
* `tenant_ref` - It is a reference to an object of type tenant.
* `test_se_datastore_level_1_ref` - Used for testing se datastore upgrade 2.0 functionality. It is a reference to an object of type testsedatastorelevel1. Field introduced in 18.2.6.
* `topology_policies` - Topology policies applied on the dns traffic of the virtual service based ongslb topology algorithm. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `traffic_clone_profile_ref` - Server network or list of servers for cloning traffic. It is a reference to an object of type trafficcloneprofile. Field introduced in 17.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `traffic_enabled` - Knob to enable the virtual service traffic on its assigned service engines. This setting is effective only when the enabled flag is set to true. Field introduced in 17.2.8.
* `type` - Specify if this is a normal virtual service, or if it is the parent or child of an sni-enabled virtual hosted virtual service. Enum options - VS_TYPE_NORMAL, VS_TYPE_VH_PARENT, VS_TYPE_VH_CHILD. Allowed in basic(allowed values- vs_type_normal,vs_type_vh_parent) edition, essentials(allowed values- vs_type_normal) edition, enterprise edition.
* `use_bridge_ip_as_vip` - Use bridge ip as vip on each host in mesos deployments. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `use_vip_as_snat` - Use the virtual ip as the snat ip for health monitoring and sending traffic to the backend servers instead of the service engine interface ip. The caveat of enabling this option is that the virtualservice cannot be configued in an active-active ha mode. Dns based multi vip solution has to be used for ha & non-disruptive upgrade purposes. Field introduced in 17.1.9,17.2.3. Allowed in essentials(allowed values- false) edition, enterprise edition.
* `uuid` - Uuid of the virtualservice.
* `vh_domain_name` - The exact name requested from the client's sni-enabled tls hello domain name field. If this is a match, the parent vs will forward the connection to this child vs.
* `vh_matches` - Host and path match criteria to select this child vs. Field introduced in 20.1.3.
* `vh_parent_vs_ref` - Specifies the virtual service acting as virtual hosting (sni) parent. It is a reference to an object of type virtualservice.
* `vh_type` - Specify if the virtual hosting vs is of type sni or enhanced. Enum options - VS_TYPE_VH_SNI, VS_TYPE_VH_ENHANCED. Field introduced in 20.1.3. Allowed in basic(allowed values- vs_type_vh_enhanced) edition, enterprise edition. Special default for basic edition is vs_type_vh_enhanced, enterprise is vs_type_vh_sni.
* `vip` - List of virtual service ips. While creating a 'shared vs',please use vsvip_ref to point to the shared entities. Field introduced in 17.1.1.
* `vrf_context_ref` - Virtual routing context that the virtual service is bound to. This is used to provide the isolation of the set of networks the application is attached to. It is a reference to an object of type vrfcontext.
* `vs_datascripts` - Datascripts applied on the data traffic of the virtual service.
* `vsvip_cloud_config_cksum` - Checksum of cloud configuration for vsvip. Internally set by cloud connector. Field introduced in 17.2.9, 18.1.2.
* `vsvip_ref` - Mostly used during the creation of shared vs, this field refers to entities that can be shared across virtual services. It is a reference to an object of type vsvip. Field introduced in 17.1.1.
* `waf_policy_ref` - Waf policy for the virtual service. It is a reference to an object of type wafpolicy. Field introduced in 17.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `weight` - The quality of service weight to assign to traffic transmitted from this virtual service. A higher weight will prioritize traffic versus other virtual services sharing the same service engines. Allowed values are 1-128. Allowed in basic(allowed values- 1) edition, essentials(allowed values- 1) edition, enterprise edition.

