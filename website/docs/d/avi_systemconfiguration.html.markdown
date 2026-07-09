<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_systemconfiguration"
sidebar_current: "docs-avi-datasource-systemconfiguration"
description: |-
  Get information of Avi SystemConfiguration.
---

# avi_systemconfiguration

This data source is used to to get avi_systemconfiguration objects.

## Example Usage

```hcl
data "avi_systemconfiguration" "foo_systemconfiguration" {
    uuid = "systemconfiguration-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SystemConfiguration by name.
* `uuid` - (Optional) Search SystemConfiguration by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `admin_auth_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `allow_legacy_sha1_ntp_auth` - Allow ntp authentication using legacy md5 or sha1 algorithms. When enabled, configuring md5 or sha1 ntp keys is permitted but a warning event is generated in the controller ui. When disabled (default), only sha256 or stronger is accepted and configuring md5 or sha1 results in an api error. Field introduced in 32.1.3. Allowed with any value in enterprise, enterprise with cloud services edition.
* `avi_email_login_password` - Password for avi_email_login user. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `certificate_security_policy` - Certificate security policy for the system. Field introduced in 32.1.3. Allowed with any value in enterprise, enterprise with cloud services edition.
* `common_criteria_mode` - Common criteria mode's current state. Field introduced in 20.1.3. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `controller_analytics_policy` - Controller metrics event dynamic thresholds can be set here. Controller_cpu_high and controller_mem_high evets can take configured dynamic thresholds. Field introduced in 21.1.3. Allowed with any value in enterprise, enterprise with cloud services edition.
* `default_license_tier` - Specifies the default license tier which would be used by new clouds. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS, ENTERPRISE_WITH_CLOUD_SERVICES. Field introduced in 17.2.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Special default for essentials edition is essentials, basic edition is basic, enterprise edition is enterprise_with_cloud_services.
* `dns_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `dns_virtualservice_refs` - Dns virtualservices hosting fqdn records for applications across avi vantage. If no virtualservices are provided, avi vantage will provide dns services for configured applications. Switching back to avi vantage from dns virtualservices is not allowed. It is a reference to an object of type virtualservice. Allowed with any value in enterprise, enterprise with cloud services edition.
* `docker_mode` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `email_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `enable_cors` - Enable cors header. Field introduced in 20.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `enable_host_header_check` - Validates the host header against a list of trusted domains. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enable_license_quota` - Enable license quota for the system. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `fips_mode` - Fips mode current state. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `global_tenant_config` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `host_key_algorithm_exclude` - Users can specify comma separated list of deprecated host key algorithm.if nothing is specified, all known algorithms provided by openssh will be supported.this change could only apply on the controller node. Field introduced in 22.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `kex_algorithm_exclude` - Users can specify comma separated list of deprecated key exchange algorithm.if nothing is specified, all known algorithms provided by openssh will be supported.this change could only apply on the controller node. Field introduced in 22.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `legacy_ssl_support` - Allow outgoing connections from controller to servers using tls 1.0/1.1. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `license_quota` - License quota for the system. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `linux_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `mgmt_ip_access_control` - Configure ip access control for controller to restrict open access. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `ntp_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `password_policy_managed_at_ops` - Indicates whether password policy fields are managed by vcf-ops. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `portal_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `proxy_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `rekey_time_limit` - Users can specify and update the time limit of rekeylimit in sshd_config.if nothing is specified, the default setting will be none. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `rekey_volume_limit` - Users can specify and update the size/volume limit of rekeylimit in sshd_config.if nothing is specified, the default setting will be default. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `sddcmanager_fqdn` - Fqdn of sddc manager in vcf responsible for management of this alb controller cluster. Field introduced in 22.1.6,31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `secure_channel_configuration` - Configure secure channel properties. Field introduced in 18.1.4, 18.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `service_auth_configurations` - Service auth configurations. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `snmp_configuration` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `ssh_ciphers` - Allowed ciphers list for ssh to the management interface on the controller and service engines. If this is not specified, all the default ciphers are allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `ssh_hmacs` - Allowed hmac list for ssh to the management interface on the controller and service engines. If this is not specified, all the default hmacs are allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `sync_kex_host_to_se` - Ability to sync the kexalgorithms & hostkeyalgorithms to ses. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `sync_syslog_to_se` - Ability to sync the syslog server config to ses. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `syslog_servers` - The destination syslog server ip(v4/v6) address or fqdn. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `telemetry_configuration` - Telemetry configuration. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `trusted_host_profiles_refs` - Trusted host profiles for host header validation. Only works when host_header_check is set to true. It is a reference to an object of type trustedhostprofile. Field introduced in 31.1.1. Maximum of 20 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition.
* `truststore_pkiprofile_ref` - Reference to pkiprofile used for validating the ca certificates for external comminications from avi load balancer controller  this acts as trust store for avi load balancer controller. It is a reference to an object of type pkiprofile. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `welcome_workflow_complete` - This flag is set once the initial controller setup workflow is complete. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

