---
- hosts: localhost
  connection: local
  gather_facts: no
  roles:
    - role: avinetworks.avisdk
  vars:
    avi_credentials:
        controller: "{{ controller_ip[0] }}"
        username: "admin"
        password: "{{ password }}"
        api_version: ${avi_version}
    controller: "{{ ansible_host }}"
    username: admin
    password: "{{ password }}"
    cloud_name: "Default-Cloud"
    controller_ip:
      ${ indent(6, yamlencode(controller_ip))}
    controller_names:
      ${ indent(6, yamlencode(controller_names))}
    ansible_become: yes
    ansible_become_password: "{{ password }}"
    aws_vpc_id: ${vpc_id}
    aws_region: ${aws_region}
    se_name_prefix: ${se_name_prefix}
    se_ha_mode: ${se_ha_mode}
    mgmt_security_group: ${mgmt_security_group}
    data_security_group: ${data_security_group}
    controller_ha: ${controller_ha}
%{ if dns_servers != null ~}
    dns_servers:
%{ for item in dns_servers ~}
      - addr: "${item}"
        type: "V4"
%{ endfor ~}
    dns_search_domain: ${dns_search_domain}
%{ endif ~}
    ntp_servers:
%{ for item in ntp_servers ~}
      - server:
          addr: "${item.addr}"
          type: "${item.type}"
%{ endfor ~}
    configure_dns_route_53: ${configure_dns_route_53}
%{ if configure_dns_profile ~}
    dns_service_domain: ${dns_service_domain}
%{ endif ~}
%{ if configure_dns_vs ~}
    dns_vs_settings: 
      ${ indent(6, yamlencode(dns_vs_settings))}
%{ endif ~}
%{ if configure_gslb ~}
    gslb_site_name: ${gslb_site_name}
    additional_gslb_sites:
      ${ indent(6, yamlencode(additional_gslb_sites))}
%{ endif ~}
  tasks:
    - name: Wait for Controller to become ready
      wait_for:
        port: 443
        timeout: 600
        sleep: 30
    - name: Configure System Configurations
      avi_systemconfiguration:
        avi_credentials: "{{ avi_credentials }}"
        state: present
        email_configuration:
          smtp_type: "SMTP_LOCAL_HOST"
          from_email: admin@avicontroller.net
        global_tenant_config:
          se_in_provider_context: true
          tenant_access_to_provider_se: true
          tenant_vrf: false
%{ if dns_servers != null ~}
        dns_configuration:
          server_list: "{{ dns_servers }}"
          search_domain: "{{ dns_search_domain }}"
%{ endif ~}
        ntp_configuration:
          ntp_servers: "{{ ntp_servers }}"        
        portal_configuration:
          allow_basic_authentication: false
          disable_remote_cli_shell: false
          enable_clickjacking_protection: true
          enable_http: true
          enable_https: true
          password_strength_check: true
          redirect_to_https: true
          use_uuid_from_input: false
        welcome_workflow_complete: true

    - name: Configure Cloud
      avi_cloud:
        avi_credentials: "{{ avi_credentials }}"
        state: present
        name: "{{ cloud_name }}"
        vtype: CLOUD_AWS
        dhcp_enabled: true
        license_type: "LIC_CORES"
        aws_configuration:
          access_key_id: "{{ aws_access_key_id }}"
          secret_access_key: "{{ aws_secret_access_key }}"
          region: "{{ aws_region }}"
          asg_poll_interval: 60
          vpc_id: "{{ aws_vpc_id }}"
          route53_integration: "{{ configure_dns_route_53 }}"
          zones: 
%{ for zone, mgmt_subnet in se_mgmt_subnets ~}
            - availability_zone: "${zone}"
              mgmt_network_name: "${mgmt_subnet["mgmt_network_name"]}"
              mgmt_network_uuid: "${mgmt_subnet["mgmt_network_uuid"]}"
%{ endfor ~}
      register: avi_cloud
    - name: Set Backup Passphrase
      avi_backupconfiguration:
        avi_credentials: "{{ avi_credentials }}"
        state: present
        name: Backup-Configuration
        backup_passphrase: "{{ password }}"
        upload_to_remote_host: false
%{ if se_ha_mode == "active/active" }
    - name: Configure SE-Group
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: "serviceenginegroup"
        tenant: "admin"
        data:
          name: "Default-Group" 
          cloud_ref: "/api/cloud?name={{ cloud_name }}"
          ha_mode: HA_MODE_SHARED_PAIR
          min_scaleout_per_vs: 2
          algo: PLACEMENT_ALGO_DISTRIBUTED
          buffer_se: "0"
          max_se: "10"
          se_name_prefix: "{{ se_name_prefix }}"
          accelerated_networking: true
          disable_avi_securitygroups: true
          custom_securitygroups_mgmt:
            - "{{ mgmt_security_group }}"
          custom_securitygroups_data:
            - "{{ data_security_group }}"
          realtime_se_metrics:
            duration: "10080"
            enabled: true
%{ endif }
%{ if se_ha_mode == "n+m" }
    - name: Configure SE-Group
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: "serviceenginegroup"
        tenant: "admin"
        data:
          name: "Default-Group" 
          state: present
          cloud_ref: "{{ avi_cloud.obj.url }}"
          ha_mode: HA_MODE_SHARED
          min_scaleout_per_vs: 1
          algo: PLACEMENT_ALGO_PACKED
          buffer_se: "1"
          max_se: "10"
          se_name_prefix: "{{ se_name_prefix }}"
          accelerated_networking: true
          disable_avi_securitygroups: true
          custom_securitygroups_mgmt:
            - "{{ mgmt_security_group }}"
          custom_securitygroups_data:
            - "{{ data_security_group }}"
          realtime_se_metrics:
            duration: "10080"
            enabled: true
%{ endif }
%{ if se_ha_mode == "active/standby" }
    - name: Configure SE-Group
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: "serviceenginegroup"
        tenant: "admin"
        data:
          name: "Default-Group" 
          cloud_ref: "{{ avi_cloud.obj.url }}"
          ha_mode: HA_MODE_LEGACY_ACTIVE_STANDBY
          min_scaleout_per_vs: 1
          buffer_se: "0"
          max_se: "2"
          se_name_prefix: "{{ se_name_prefix }}_se"
          accelerated_networking: true
          disable_avi_securitygroups: true
          custom_securitygroups_mgmt:
            - "{{ mgmt_security_group }}"
          custom_securitygroups_data:
            - "{{ data_security_group }}"
          realtime_se_metrics:
            duration: "10080"
            enabled: true
%{ endif }
%{ if configure_dns_profile }
    - name: Create Avi DNS Profile
      avi_ipamdnsproviderprofile:
        avi_credentials: "{{ avi_credentials }}"
        state: present
        name: Avi_DNS
        type: IPAMDNS_TYPE_INTERNAL_DNS
        internal_profile:
          dns_service_domain:
          - domain_name: "{{ dns_service_domain }}"
            pass_through: true
          ttl: 30
      register: create_dns
    - name: Update Cloud Configuration with DNS profile 
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: patch
        path: "cloud/{{ avi_cloud.obj.uuid }}"
        data:
          add:
            dns_provider_ref: "{{ create_dns.obj.url }}"
%{ endif }
%{ if configure_gslb }
    - name: Configure GSLB SE-Group
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: "serviceenginegroup"
        tenant: "admin"
        data:
          name: "g-dns" 
          cloud_ref: "{{ avi_cloud.obj.url }}"
          ha_mode: HA_MODE_SHARED
          algo: PLACEMENT_ALGO_PACKED
          buffer_se: "1"
          max_se: "4"
          max_vs_per_se: "2"
          extra_shared_config_memory: 2000
          se_name_prefix: "{{ se_name_prefix }}"
          realtime_se_metrics:
            duration: "10080"
            enabled: true
      register: gslb_se_group
%{ endif}
%{ if configure_dns_vs ~}
    - name: DNS VS Config | Get AWS Subnet Information
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: get
        path: "vimgrnwruntime?name={{ dns_vs_settings.subnet_name }}"
      register: dns_vs_subnet
    - name: Create DNS VSVIP
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: "vsvip"
        tenant: "admin"
        data:
          east_west_placement: false
          cloud_ref: "{{ avi_cloud.obj.url }}"
%{ if configure_gslb ~}
          se_group_ref: "{{ gslb_se_group.obj.url }}"
%{ endif ~}
          vip:
          - enabled: true
            vip_id: 0      
            auto_allocate_ip: "true"
            auto_allocate_floating_ip: "{{ dns_vs_settings.allocate_public_ip }}"
            avi_allocated_vip: true
            avi_allocated_fip: "{{ dns_vs_settings.allocate_public_ip }}"
            auto_allocate_ip_type: V4_ONLY
            prefix_length: 32
            placement_networks: []
            ipam_network_subnet:
              network_ref: "{{ dns_vs_subnet.obj.results.0.url }}"
              subnet:
                ip_addr:
                  addr: "{{ dns_vs_subnet.obj.results.0.ip_subnet.0.prefix.ip_addr.addr }}"
                  type: "{{ dns_vs_subnet.obj.results.0.ip_subnet.0.prefix.ip_addr.type }}"
                mask: "{{ dns_vs_subnet.obj.results.0.ip_subnet.0.prefix.mask }}"
          dns_info:
          - type: DNS_RECORD_A
            algorithm: DNS_RECORD_RESPONSE_CONSISTENT_HASH
            fqdn: "dns.{{ dns_service_domain }}"
          name: vsvip-DNS-VS-Default-Cloud
      register: vsvip_results

    - name: Create DNS Virtual Service
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: "virtualservice"
        tenant: "admin"
        data:
          name: DNS-VS
          enabled: true
          analytics_policy:
            full_client_logs:
              enabled: true
              duration: 30
              throttle: 10
            client_insights: NO_INSIGHTS
            all_headers: false
            metrics_realtime_update:
              enabled: true
              duration: 30
            udf_log_throttle: 10
            significant_log_throttle: 10
            learning_log_policy:
              enabled: false
            client_log_filters: []
            client_insights_sampling: {}
          enable_autogw: true
          weight: 1
          delay_fairness: false
          max_cps_per_client: 0
          limit_doser: false
          type: VS_TYPE_NORMAL
          use_bridge_ip_as_vip: false
          flow_dist: LOAD_AWARE
          ign_pool_net_reach: false
          ssl_sess_cache_avg_size: 1024
          remove_listening_port_on_vs_down: false
          close_client_conn_on_config_update: false
          bulk_sync_kvcache: false
          advertise_down_vs: false
          scaleout_ecmp: false
          active_standby_se_tag: ACTIVE_STANDBY_SE_1
          flow_label_type: NO_LABEL
          content_rewrite:
            request_rewrite_enabled: false
            response_rewrite_enabled: false
            rewritable_content_ref: /api/stringgroup?name=System-Rewritable-Content-Types
          sideband_profile:
            sideband_max_request_body_size: 1024
          use_vip_as_snat: false
          traffic_enabled: true
          allow_invalid_client_cert: false
          vh_type: VS_TYPE_VH_SNI
          application_profile_ref: /api/applicationprofile?name=System-DNS
          network_profile_ref: /api/networkprofile?name=System-UDP-Per-Pkt
          analytics_profile_ref: /api/analyticsprofile?name=System-Analytics-Profile
          %{ if configure_gslb }
          se_group_ref: "{{ gslb_se_group.obj.url }}"
          %{ endif}
          cloud_ref: "{{ avi_cloud.obj.url }}"
          services:
          - port: 53
            port_range_end: 53
            enable_ssl: false
            enable_http2: false
          - port: 53
            port_range_end: 53
            override_network_profile_ref: /api/networkprofile/?name=System-TCP-Proxy
            enable_ssl: false
            enable_http2: false
          vsvip_ref: "{{ vsvip_results.obj.url }}"
          vs_datascripts: []
      register: dns_vs

    - name: Add DNS-VS to System Configuration
      avi_systemconfiguration:
        avi_credentials: "{{ avi_credentials }}"
        avi_api_update_method: patch
        avi_api_patch_op: add
        tenant: admin
        dns_virtualservice_refs: "{{ dns_vs.obj.url }}"
%{ endif}
%{ if configure_gslb }
    - name: GSLB Config | Verify Cluster UUID
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: get
        path: cluster
      register: cluster
    - name: Create GSLB Config
      avi_gslb:
        avi_credentials: "{{ avi_credentials }}"
        name: "GSLB"
        sites:
          - name: "{{ gslb_site_name }}"
            username: "{{ username }}"
            password: "{{ password }}"
            ip_addresses:
              - type: "V4"
                addr: "{{ controller_ip[0] }}"
%{ if controller_ha ~}
              - type: "V4"
                addr: "{{ controller_ip[1] }}"
              - type: "V4"
                addr: "{{ controller_ip[2] }}"
%{ endif ~}
            enabled: True
            member_type: "GSLB_ACTIVE_MEMBER"
            port: 443
            dns_vses:
              - dns_vs_uuid: "{{ dns_vs.obj.uuid }}"
            cluster_uuid: "{{ cluster.obj.uuid }}"
        dns_configs:
          %{ for domain in gslb_domains }
          - domain_name: "${domain}"
          %{ endfor }
        leader_cluster_uuid: "{{ cluster.obj.uuid }}"
      register: gslb_results
  %{ endif }
  %{ if configure_gslb_additional_sites }%{ for site in additional_gslb_sites }

    - name: GSLB Config | Verify DNS configuration
      avi_api_session:
        controller: "${site.ip_address}"
        username: "admin"
        password: "{{ password }}"
        api_version: ${avi_version}
        http_method: get
        path: virtualservice?name=DNS-VS
      register: dns_vs_verify

    - name: Display DNS VS Verify
      ansible.builtin.debug:
        var: dns_vs_verify

    - name: GSLB Config | Verify GSLB site configuration
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: post
        path: gslbsiteops/verify
        data:
          name: name
          username: admin
          password: "{{ password }}"
          port: 443
          ip_addresses:
            - type: "V4"
              addr: "${site.ip_address_list[0]}"
      register: gslb_verify
      
    - name: Display GSLB Siteops Verify
      ansible.builtin.debug:
        var: gslb_verify

    - name: Add GSLB Sites
      avi_api_session:
        avi_credentials: "{{ avi_credentials }}"
        http_method: patch
        path: "gslb/{{ gslb_results.obj.uuid }}"
        tenant: "admin"
        data:
          add:
            sites:
              - name: "${site.name}"
                member_type: "GSLB_ACTIVE_MEMBER"
                username: "{{ username }}"
                password: "{{ password }}"
                cluster_uuid: "{{ gslb_verify.obj.rx_uuid }}"
                ip_addresses:
%{ for address in site.ip_address_list ~}
                  - type: "V4"
                    addr: "${address}"
%{ endfor ~}
                dns_vses:
                  - dns_vs_uuid: "{{ dns_vs_verify.obj.results.0.uuid }}"
  %{ endfor }%{ endif }
%{ if controller_ha }
    - name: Controller Cluster Configuration
      avi_cluster:
        avi_credentials: "{{ avi_credentials }}"
        state: present
        #virtual_ip:
        #  type: V4
        #  addr: "{{ controller_cluster_vip }}"
        nodes:
            - name:  "{{ controller_names[0] }}" 
              password: "{{ password }}"
              ip:
                type: V4
                addr: "{{ controller_ip[0] }}"
            - name:  "{{ controller_names[1] }}" 
              password: "{{ password }}"
              ip:
                type: V4
                addr: "{{ controller_ip[1] }}"
            - name:  "{{ controller_names[2] }}" 
              password: "{{ password }}"
              ip:
                type: V4
                addr: "{{ controller_ip[2] }}"
        name: "cluster01"
        tenant_uuid: "admin"
%{ endif }