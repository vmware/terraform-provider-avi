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
resource "SystemConfiguration" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `admin_auth_configuration` - (Optional ) argument_description.
        * `default_license_tier` - (Optional ) argument_description.
        * `dns_configuration` - (Optional ) argument_description.
        * `dns_virtualservice_refs` - (Optional ) argument_description.
        * `docker_mode` - (Optional ) argument_description.
        * `email_configuration` - (Optional ) argument_description.
        * `global_tenant_config` - (Optional ) argument_description.
        * `linux_configuration` - (Optional ) argument_description.
        * `mgmt_ip_access_control` - (Optional ) argument_description.
        * `ntp_configuration` - (Optional ) argument_description.
        * `portal_configuration` - (Optional ) argument_description.
        * `proxy_configuration` - (Optional ) argument_description.
        * `secure_channel_configuration` - (Optional ) argument_description.
        * `snmp_configuration` - (Optional ) argument_description.
        * `ssh_ciphers` - (Optional ) argument_description.
        * `ssh_hmacs` - (Optional ) argument_description.
            * `welcome_workflow_complete` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                    * `uuid` - argument_description.
        
