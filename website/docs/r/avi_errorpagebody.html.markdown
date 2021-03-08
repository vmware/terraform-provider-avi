############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_errorpagebody"
sidebar_current: "docs-avi-resource-errorpagebody"
description: |-
  Creates and manages Avi ErrorPageBody.
---

# avi_errorpagebody

The ErrorPageBody resource allows the creation and management of Avi ErrorPageBody

## Example Usage

```hcl
resource "avi_errorpagebody" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `error_page_body` - (Required) Error page body sent to client when match. Field introduced in 17.2.4.
* `name` - (Required) Field introduced in 17.2.4.
* `format` - (Optional) Format of an error page body html or json. Enum options - ERROR_PAGE_FORMAT_HTML, ERROR_PAGE_FORMAT_JSON. Field introduced in 18.2.3.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.2.4.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.4.

