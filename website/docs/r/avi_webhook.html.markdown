<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `name` - (Required) The name of the webhook profile. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `callback_url` - (Optional) Callback url for the webhook. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - (Optional) Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `verification_token` - (Optional) Verification token sent back with the callback asquery parameters. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the webhook profile. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

