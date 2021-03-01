############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `admin_auth_configuration` - (Optional) Dict settings for systemconfiguration.
* `common_criteria_mode` - (Optional) Enable common criteria compliance mode (disabled by default). Warn  toggling this field is disruptive and will result in reduced behavior with ssh and tls protocols. Expect possible warm start of control and data planes. Field introduced in 20.1.3.
* `default_license_tier` - (Optional) Specifies the default license tier which would be used by new clouds. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS. Field introduced in 17.2.5. Allowed in basic edition, essentials edition, enterprise edition. Special default for basic edition is basic, essentials edition is essentials, enterprise is enterprise.
* `dns_configuration` - (Optional) Dict settings for systemconfiguration.
* `dns_virtualservice_refs` - (Optional) Dns virtualservices hosting fqdn records for applications across avi vantage. If no virtualservices are provided, avi vantage will provide dns services for configured applications. Switching back to avi vantage from dns virtualservices is not allowed. It is a reference to an object of type virtualservice.
* `docker_mode` - (Optional) Boolean flag to set docker_mode.
* `email_configuration` - (Optional) Dict settings for systemconfiguration.
* `enable_cors` - (Optional) Enable cors header. Field introduced in 20.1.3. Allowed in basic edition, essentials edition, enterprise edition.
* `fips_mode` - (Optional) Enable fips mode. Field introduced in 20.1.1.
* `global_tenant_config` - (Optional) Dict settings for systemconfiguration.
* `linux_configuration` - (Optional) Dict settings for systemconfiguration.
* `mgmt_ip_access_control` - (Optional) Configure ip access control for controller to restrict open access.
* `ntp_configuration` - (Optional) Dict settings for systemconfiguration.
* `portal_configuration` - (Optional) Dict settings for systemconfiguration.
* `proxy_configuration` - (Optional) Dict settings for systemconfiguration.
* `secure_channel_configuration` - (Optional) Configure secure channel properties. Field introduced in 18.1.4, 18.2.1.
* `snmp_configuration` - (Optional) Dict settings for systemconfiguration.
* `ssh_ciphers` - (Optional) Allowed ciphers list for ssh to the management interface on the controller and service engines. If this is not specified, all the default ciphers are allowed.
* `ssh_hmacs` - (Optional) Allowed hmac list for ssh to the management interface on the controller and service engines. If this is not specified, all the default hmacs are allowed.
* `welcome_workflow_complete` - (Optional) This flag is set once the initial controller setup workflow is complete. Field introduced in 18.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

