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
* `default_license_tier` - Specifies the default license tier which would be used by new clouds.
* `dns_configuration` - Dict settings for systemconfiguration.
* `dns_virtualservice_refs` - Dns virtualservices hosting fqdn records for applications across avi vantage.
* `docker_mode` - Boolean flag to set docker_mode.
* `email_configuration` - Dict settings for systemconfiguration.
* `fips_mode` - Enable fips mode.
* `global_tenant_config` - Dict settings for systemconfiguration.
* `linux_configuration` - Dict settings for systemconfiguration.
* `mgmt_ip_access_control` - Configure ip access control for controller to restrict open access.
* `ntp_configuration` - Dict settings for systemconfiguration.
* `portal_configuration` - Dict settings for systemconfiguration.
* `proxy_configuration` - Dict settings for systemconfiguration.
* `secure_channel_configuration` - Configure secure channel properties.
* `snmp_configuration` - Dict settings for systemconfiguration.
* `ssh_ciphers` - Allowed ciphers list for ssh to the management interface on the controller and service engines.
* `ssh_hmacs` - Allowed hmac list for ssh to the management interface on the controller and service engines.
* `uuid` - Unique object identifier of the object.
* `welcome_workflow_complete` - This flag is set once the initial controller setup workflow is complete.

