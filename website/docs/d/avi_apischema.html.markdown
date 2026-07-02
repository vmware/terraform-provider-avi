<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_apischema"
sidebar_current: "docs-avi-datasource-apischema"
description: |-
  Get information of Avi ApiSchema.
---

# avi_apischema

This data source is used to to get avi_apischema objects.

## Example Usage

```hcl
data "avi_apischema" "foo_apischema" {
    uuid = "apischema-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApiSchema by name.
* `uuid` - (Optional) Search ApiSchema by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `additional_object_key_action` - Action to take on unspecified keys in an object. Enum options - API_ACTION_INHERIT_FROM_API_POLICY, API_ACTION_PASS, API_ACTION_LEARN, API_ACTION_FLAG, API_ACTION_REJECT. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `additional_properties_schema` - Type constraint for additional properties not defined in object_properties. When set, unknown keys must conform to this schema. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `allow_additional_properties` - When true, object keys not defined in object_properties are permitted. Corresponds to openapi additionalproperties  true. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `array_item_type` - If the type is array, this is the type of the array items. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `composite_types` - Sub-schemas for this composite type (oneof, anyof, or allof). Each entry must be a schema_type_reference pointing to an apischema. Field introduced in 32.2.1. Maximum of 64 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description of this api schema. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `discriminator` - Property used to distinguish between sub-schemas in oneof/anyof composite types. Maps a discriminator property value to the matching schema. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `max_items` - Maximum number of items allowed in an array. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `min_items` - Minimum number of items allowed in an array. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of this object, unique per tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `object_properties` - List of properties for this object schema. Field introduced in 32.2.1. Maximum of 512 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `source` - Indicates whether this schema was user-defined or imported from an openapi specification file. Enum options - SOURCE_USER_DEFINED, SOURCE_API_SPEC, SOURCE_DISCOVERED. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `type` - The data type of this schema. Can be object, array, or a composite type (oneof, anyof, allof). Enum options - SCHEMA_TYPE_UNDEFINED, SCHEMA_TYPE_STRING, SCHEMA_TYPE_INTEGER, SCHEMA_TYPE_NUMBER, SCHEMA_TYPE_BOOLEAN, SCHEMA_TYPE_NULL, SCHEMA_TYPE_ARRAY, SCHEMA_TYPE_OBJECT, SCHEMA_TYPE_REFERENCE, SCHEMA_TYPE_ONE_OF, SCHEMA_TYPE_ALL_OF, SCHEMA_TYPE_ANY_OF. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `unique_items` - If true, all items in the array must be unique. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - The object uuid. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

