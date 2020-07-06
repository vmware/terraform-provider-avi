---
layout: "avi"
page_title: "Avi: avi_albservicesconfig"
sidebar_current: "docs-avi-resource-albservicesconfig"
description: |-
  Creates and manages Avi ALBServicesConfig.
---

# avi_albservicesconfig

The ALBServicesConfig resource allows the creation and management of Avi ALBServicesConfig

## Example Usage

```hcl
resource "avi_albservicesconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `asset_contact` - (Optional) Information about the default contact for this controller cluster.
* `feature_opt_in_status` - (Optional) Information about the portal features opted in for controller.
* `ip_reputation_file_object_expiry_duration` - (Optional) Ip reputation db file object expiry duration in days.
* `ip_reputation_sync_interval` - (Optional) Ip reputation db sync interval in minutes.
* `polling_interval` - (Optional) Time interval in minutes.
* `portal_url` - (Optional) The fqdn or ip address of the customer portal.
* `proactive_support_defaults` - (Optional) Default values to be used during proactive case creation and techsupport attachment.
* `split_proxy_configuration` - (Optional) Split proxy configuration to connect external pulse services.
* `use_system_proxy` - (Optional) By default, use system proxy configurationif false, use split proxy configuration.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 18.2.6.

