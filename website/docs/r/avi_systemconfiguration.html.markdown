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
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `admin_auth_configuration` - (Optional) Dict settings for systemconfiguration.
* `default_license_tier` - (Optional) Specifies the default license tier which would be used by new clouds.
* `dns_configuration` - (Optional) Dict settings for systemconfiguration.
* `dns_virtualservice_refs` - (Optional) Dns virtualservices hosting fqdn records for applications across avi vantage.
* `docker_mode` - (Optional) Boolean flag to set docker_mode.
* `email_configuration` - (Optional) Dict settings for systemconfiguration.
* `global_tenant_config` - (Optional) Dict settings for systemconfiguration.
* `linux_configuration` - (Optional) Dict settings for systemconfiguration.
* `mgmt_ip_access_control` - (Optional) Configure ip access control for controller to restrict open access.
* `ntp_configuration` - (Optional) Dict settings for systemconfiguration.
* `portal_configuration` - (Optional) Dict settings for systemconfiguration.
* `proxy_configuration` - (Optional) Dict settings for systemconfiguration.
* `secure_channel_configuration` - (Optional) Configure secure channel properties.
* `snmp_configuration` - (Optional) Dict settings for systemconfiguration.
* `ssh_ciphers` - (Optional) Allowed ciphers list for ssh to the management interface on the controller and service engines.
* `ssh_hmacs` - (Optional) Allowed hmac list for ssh to the management interface on the controller and service engines.
* `welcome_workflow_complete` - (Optional) This flag is set once the initial controller setup workflow is complete.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

