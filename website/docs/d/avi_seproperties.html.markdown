---
layout: "avi"
page_title: "AVI: avi_seproperties"
sidebar_current: "docs-avi-datasource-seproperties"
description: |-
  Get information of Avi SeProperties.
---

# avi_seproperties

This data source is used to to get avi_seproperties objects.

## Example Usage

```hcl
data "avi_seproperties" "foo_seproperties" {
    uuid = "seproperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SeProperties by name.
* `uuid` - (Optional) Search SeProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `se_agent_properties` - Dict settings for seproperties.
* `se_bootup_properties` - Dict settings for seproperties.
* `se_runtime_properties` - Dict settings for seproperties.
* `uuid` - Unique object identifier of the object.

