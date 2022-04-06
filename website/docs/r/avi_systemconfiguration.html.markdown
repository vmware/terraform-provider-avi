<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_systemconfiguration"
sidebar_current: "docs-avi-resource-systemconfiguration"
description: |-
  Creates and manages Avi SystemConfiguration.
---

# avi_systemconfiguration

The SystemConfiguration resource allows the creation and management of Avi SystemConfiguration

## Example Usage

```hcl
resource "avi_systemconfiguration" "foo" {
    uuid = "default-uuid"
}
```

## Argument Reference

The following arguments are supported:

* `admin_auth_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `common_criteria_mode` - (Optional) Common criteria mode's current state. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `controller_analytics_policy` - (Optional) Controller metrics event dynamic thresholds can be set here. Controller_cpu_high and controller_mem_high evets can take configured dynamic thresholds. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `default_license_tier` - (Optional) Specifies the default license tier which would be used by new clouds. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS, ENTERPRISE_WITH_CLOUD_SERVICES. Field introduced in 17.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition. Special default for essentials edition is essentials, basic edition is basic, enterprise is enterprise_with_cloud_services.
* `dns_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dns_virtualservice_refs` - (Optional) Dns virtualservices hosting fqdn records for applications across avi vantage. If no virtualservices are provided, avi vantage will provide dns services for configured applications. Switching back to avi vantage from dns virtualservices is not allowed. It is a reference to an object of type virtualservice. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `docker_mode` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `email_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `enable_cors` - (Optional) Enable cors header. Field introduced in 20.1.3. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `fips_mode` - (Optional) Fips mode current state. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `global_tenant_config` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `linux_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `mgmt_ip_access_control` - (Optional) Configure ip access control for controller to restrict open access. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ntp_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `portal_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `proxy_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `secure_channel_configuration` - (Optional) Configure secure channel properties. Field introduced in 18.1.4, 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `snmp_configuration` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ssh_ciphers` - (Optional) Allowed ciphers list for ssh to the management interface on the controller and service engines. If this is not specified, all the default ciphers are allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ssh_hmacs` - (Optional) Allowed hmac list for ssh to the management interface on the controller and service engines. If this is not specified, all the default hmacs are allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `welcome_workflow_complete` - (Optional) This flag is set once the initial controller setup workflow is complete. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

