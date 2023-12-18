<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_webapput"
sidebar_current: "docs-avi-resource-webapput"
description: |-
  Creates and manages Avi WebappUT.
---

# avi_webapput

The WebappUT resource allows the creation and management of Avi WebappUT

## Example Usage

```hcl
resource "avi_webapput" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the webapput object-level0. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `mandatory_test` - (Optional) Optional message for nested f_mandatory test cases defined at level1. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `mandatory_tests` - (Optional) Repeated message for nested f_mandatory test cases-level1. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `sensitive_test` - (Optional) Optional message for nested f_sensitive test cases defined at level1. Field introduced in 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `sensitive_tests` - (Optional) Repeated message for nested f_sensitive test cases-level1. Field introduced in 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `skip_optional_check_tests` - (Optional) Optional bool for nested skip_optional_check test cases-level1. Field introduced in 30.1.2. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `string_length_test` - (Optional) Optional message for nested  max string length test cases. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `string_length_tests` - (Optional) Repeated message for nested  max string length test cases. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant of the webapput object-level0. It is a reference to an object of type tenant. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `test_sensitive_string` - (Optional) The string for sensitive (secret) field. Object-level0. Field introduced in 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `test_string` - (Optional) The maximum string length. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the webapput object-level0. Field introduced in 21.1.5, 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

