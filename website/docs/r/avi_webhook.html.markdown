---
layout: "avi"
page_title: "Avi: avi_webhook"
sidebar_current: "docs-avi-resource-webhook"
description: |-
  Creates and manages Avi Webhook.
---

# avi_webhook

The Webhook resource allows the creation and management of Avi Webhook

## Example Usage

```hcl
resource "avi_webhook" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the webhook profile.
* `callback_url` - (Optional) Callback url for the webhook.
* `description` - (Optional) Field introduced in 17.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `verification_token` - (Optional) Verification token sent back with the callback asquery parameters.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the webhook profile.

