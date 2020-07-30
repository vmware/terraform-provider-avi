############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `callback_url` - Callback url for the webhook.
* `description` - Field introduced in 17.1.1.
* `name` - The name of the webhook profile.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the webhook profile.
* `verification_token` - Verification token sent back with the callback asquery parameters.

