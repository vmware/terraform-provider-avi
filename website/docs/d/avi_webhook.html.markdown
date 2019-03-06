
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
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
data "Webhook" "foo_Webhook" {
    uuid = "Webhook-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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

