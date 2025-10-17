<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_positivesecuritypolicy"
sidebar_current: "docs-avi-datasource-positivesecuritypolicy"
description: |-
  Get information of Avi PositiveSecurityPolicy.
---

# avi_positivesecuritypolicy

This data source is used to to get avi_positivesecuritypolicy objects.

## Example Usage

```hcl
data "avi_positivesecuritypolicy" "foo_positivesecuritypolicy" {
    uuid = "positivesecuritypolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PositiveSecurityPolicy by name.
* `uuid` - (Optional) Search PositiveSecurityPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Details of the positive security configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enable_positive_security_rule_updates` - Enable positive security rule generation using the application learning data rules will be programmed in a dedicated learning group. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enable_regex_programming` - Enable dynamic regex generation for positive security rules. This is an experimental feature and shouldn't be used in production. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - The name of the positivesecurity configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `positive_security_params` - Parameters for generating positive security rules. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Details of the tenant for positive security policy. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid of the positive security configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

