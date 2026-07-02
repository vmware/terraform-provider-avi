<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_apischema"
sidebar_current: "docs-avi-resource-apischema"
description: |-
  Creates and manages Avi ApiSchema.
---

# avi_apischema

The ApiSchema resource allows the creation and management of Avi ApiSchema

## Example Usage

```hcl
resource "avi_apischema" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this object, unique per tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `type` - (Required) The data type of this schema. Can be object, array, or a composite type (oneof, anyof, allof). Enum options - SCHEMA_TYPE_UNDEFINED, SCHEMA_TYPE_STRING, SCHEMA_TYPE_INTEGER, SCHEMA_TYPE_NUMBER, SCHEMA_TYPE_BOOLEAN, SCHEMA_TYPE_NULL, SCHEMA_TYPE_ARRAY, SCHEMA_TYPE_OBJECT, SCHEMA_TYPE_REFERENCE, SCHEMA_TYPE_ONE_OF, SCHEMA_TYPE_ALL_OF, SCHEMA_TYPE_ANY_OF. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `additional_object_key_action` - (Optional) Action to take on unspecified keys in an object. Enum options - API_ACTION_INHERIT_FROM_API_POLICY, API_ACTION_PASS, API_ACTION_LEARN, API_ACTION_FLAG, API_ACTION_REJECT. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `additional_properties_schema` - (Optional) Type constraint for additional properties not defined in object_properties. When set, unknown keys must conform to this schema. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `allow_additional_properties` - (Optional) When true, object keys not defined in object_properties are permitted. Corresponds to openapi additionalproperties  true. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `array_item_type` - (Optional) If the type is array, this is the type of the array items. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `composite_types` - (Optional) Sub-schemas for this composite type (oneof, anyof, or allof). Each entry must be a schema_type_reference pointing to an apischema. Field introduced in 32.2.1. Maximum of 64 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Description of this api schema. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `discriminator` - (Optional) Property used to distinguish between sub-schemas in oneof/anyof composite types. Maps a discriminator property value to the matching schema. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `max_items` - (Optional) Maximum number of items allowed in an array. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `min_items` - (Optional) Minimum number of items allowed in an array. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `object_properties` - (Optional) List of properties for this object schema. Field introduced in 32.2.1. Maximum of 512 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `source` - (Optional) Indicates whether this schema was user-defined or imported from an openapi specification file. Enum options - SOURCE_USER_DEFINED, SOURCE_API_SPEC, SOURCE_DISCOVERED. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `unique_items` - (Optional) If true, all items in the array must be unique. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  The object uuid. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

