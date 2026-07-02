<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_apipath"
sidebar_current: "docs-avi-datasource-apipath"
description: |-
  Get information of Avi ApiPath.
---

# avi_apipath

This data source is used to to get avi_apipath objects.

## Example Usage

```hcl
data "avi_apipath" "foo_apipath" {
    uuid = "apipath-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApiPath by name.
* `uuid` - (Optional) Search ApiPath by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description of this api path. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `endpoints` - List of api endpoints for this path. Field introduced in 32.2.1. Maximum of 10 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of this object, unique per tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `path_template` - The uri path template for the object. Parameters can be defined in curly braces, for example /pet/{pet_id}. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `source` - Indicates whether this path was user-defined or imported from an openapi specification file. Enum options - SOURCE_USER_DEFINED, SOURCE_API_SPEC, SOURCE_DISCOVERED. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `unknown_http_method_action` - Action to take when a request matches this path but uses an http method not defined for this path. Overrides the policy-level unknown_http_method_action when not inherit. Enum options - API_ACTION_INHERIT_FROM_API_POLICY, API_ACTION_PASS, API_ACTION_LEARN, API_ACTION_FLAG, API_ACTION_REJECT. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - The object uuid. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

