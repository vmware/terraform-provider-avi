<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_webhook"
sidebar_current: "docs-avi-datasource-webhook"
description: |-
  Get information of Avi Webhook.
---

# avi_webhook

This data source is used to to get avi_webhook objects.

## Example Usage

```hcl
data "avi_webhook" "foo_webhook" {
    uuid = "webhook-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Webhook by name.
* `uuid` - (Optional) Search Webhook by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `callback_url` - Callback url for the webhook. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - The name of the webhook profile. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the webhook profile. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `verification_token` - Verification token sent back with the callback asquery parameters. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

