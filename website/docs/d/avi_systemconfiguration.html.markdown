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

* `admin_auth_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `common_criteria_mode` - Common criteria mode's current state. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `controller_analytics_policy` - Controller metrics event dynamic thresholds can be set here. Controller_cpu_high and controller_mem_high evets can take configured dynamic thresholds. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `default_license_tier` - Specifies the default license tier which would be used by new clouds. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS, ENTERPRISE_WITH_CLOUD_SERVICES. Field introduced in 17.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition. Special default for essentials edition is essentials, basic edition is basic, enterprise is enterprise_with_cloud_services.
* `dns_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dns_virtualservice_refs` - Dns virtualservices hosting fqdn records for applications across avi vantage. If no virtualservices are provided, avi vantage will provide dns services for configured applications. Switching back to avi vantage from dns virtualservices is not allowed. It is a reference to an object of type virtualservice. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `docker_mode` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `email_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `enable_cors` - Enable cors header. Field introduced in 20.1.3. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `fips_mode` - Fips mode current state. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `global_tenant_config` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `linux_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `mgmt_ip_access_control` - Configure ip access control for controller to restrict open access. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ntp_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `portal_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `proxy_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `secure_channel_configuration` - Configure secure channel properties. Field introduced in 18.1.4, 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `snmp_configuration` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ssh_ciphers` - Allowed ciphers list for ssh to the management interface on the controller and service engines. If this is not specified, all the default ciphers are allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ssh_hmacs` - Allowed hmac list for ssh to the management interface on the controller and service engines. If this is not specified, all the default hmacs are allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `welcome_workflow_complete` - This flag is set once the initial controller setup workflow is complete. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

