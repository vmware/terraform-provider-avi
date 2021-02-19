############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `admin_auth_configuration` - Dict settings for systemconfiguration.
* `common_criteria_mode` - Enable common criteria compliance mode (disabled by default). Warn  toggling this field is disruptive and will result in reduced behavior with ssh and tls protocols. Expect possible warm start of control and data planes. Field introduced in 20.1.3.
* `default_license_tier` - Specifies the default license tier which would be used by new clouds. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS. Field introduced in 17.2.5. Allowed in basic edition, essentials edition, enterprise edition. Special default for basic edition is basic, essentials edition is essentials, enterprise is enterprise.
* `dns_configuration` - Dict settings for systemconfiguration.
* `dns_virtualservice_refs` - Dns virtualservices hosting fqdn records for applications across avi vantage. If no virtualservices are provided, avi vantage will provide dns services for configured applications. Switching back to avi vantage from dns virtualservices is not allowed. It is a reference to an object of type virtualservice.
* `docker_mode` - Boolean flag to set docker_mode.
* `email_configuration` - Dict settings for systemconfiguration.
* `enable_cors` - Enable cors header. Field introduced in 20.1.3. Allowed in basic edition, essentials edition, enterprise edition.
* `fips_mode` - Enable fips mode. Field introduced in 20.1.1.
* `global_tenant_config` - Dict settings for systemconfiguration.
* `linux_configuration` - Dict settings for systemconfiguration.
* `mgmt_ip_access_control` - Configure ip access control for controller to restrict open access.
* `ntp_configuration` - Dict settings for systemconfiguration.
* `portal_configuration` - Dict settings for systemconfiguration.
* `proxy_configuration` - Dict settings for systemconfiguration.
* `secure_channel_configuration` - Configure secure channel properties. Field introduced in 18.1.4, 18.2.1.
* `snmp_configuration` - Dict settings for systemconfiguration.
* `ssh_ciphers` - Allowed ciphers list for ssh to the management interface on the controller and service engines. If this is not specified, all the default ciphers are allowed.
* `ssh_hmacs` - Allowed hmac list for ssh to the management interface on the controller and service engines. If this is not specified, all the default hmacs are allowed.
* `uuid` - Unique object identifier of the object.
* `welcome_workflow_complete` - This flag is set once the initial controller setup workflow is complete. Field introduced in 18.2.3.

