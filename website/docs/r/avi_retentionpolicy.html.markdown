<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_retentionpolicy"
sidebar_current: "docs-avi-resource-retentionpolicy"
description: |-
  Creates and manages Avi RetentionPolicy.
---

# avi_retentionpolicy

The RetentionPolicy resource allows the creation and management of Avi RetentionPolicy

## Example Usage

```hcl
resource "avi_retentionpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `policy` - (Required) Policy specification. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enabled` - (Optional) Enables the policy. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `history` - (Optional) History of previous runs. Field introduced in 31.1.1. Maximum of 10 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - (Optional) Name of the policy. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `summary` - (Optional) Details of most recent run. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the policy. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

