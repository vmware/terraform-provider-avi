<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_positivesecuritypolicy"
sidebar_current: "docs-avi-resource-positivesecuritypolicy"
description: |-
  Creates and manages Avi PositiveSecurityPolicy.
---

# avi_positivesecuritypolicy

The PositiveSecurityPolicy resource allows the creation and management of Avi PositiveSecurityPolicy

## Example Usage

```hcl
resource "avi_positivesecuritypolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Details of the positive security configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `enable_positive_security_rule_updates` - (Optional) Enable positive security rule generation using the application learning data rules will be programmed in a dedicated learning group. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `enable_regex_programming` - (Optional) Enable dynamic regex generation for positive security rules. This is an experimental feature and shouldn't be used in production. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `name` - (Optional) The name of the positivesecurity configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `positive_security_params` - (Optional) Parameters for generating positive security rules. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Details of the tenant for positive security policy. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the positive security configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

