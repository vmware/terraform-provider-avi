---
layout: "avi"
page_title: "AVI: avi_alertemailconfig"
sidebar_current: "docs-avi-datasource-alertemailconfig"
description: |-
Get information of Avi AlertEmailConfig.
---

# avi_alertemailconfig

This data source is used to to get avi_alertemailconfig objects.

## Example Usage

```hcl
data "AlertEmailConfig" "foo_AlertEmailConfig" {
    uuid = "AlertEmailConfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AlertEmailConfig by name.
* `uuid` - (Optional) Search AlertEmailConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cc_emails` - Alerts are copied to the comma separated list of  email recipients.
* `description` - General description.
* `name` - A user-friendly name of the email notification service.
* `tenant_ref` - It is a reference to an object of type tenant.
* `to_emails` - Alerts are sent to the comma separated list of  email recipients.
* `uuid` - General description.

