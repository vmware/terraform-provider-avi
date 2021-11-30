<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `format` - (Optional) Format of an error page body html or json. Enum options - ERROR_PAGE_FORMAT_HTML, ERROR_PAGE_FORMAT_JSON. Field introduced in 18.2.3.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.2.4.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.4.

