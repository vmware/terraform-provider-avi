---
layout: "avi"
page_title: "AVI: avi_fileservice"
sidebar_current: "docs-avi-datasource-fileservice"
description: |-
  Get information of Avi Fileservice.
---

# avi_fileservice

This data source is used to to get fileservice objects.

## Example Usage

```hcl
data "avi_fileservice" "foo_Fileservice" {
    uuid = "filename"
  }
```

## Argument Reference

* `uuid` - (Optional) Search fileservice object by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` - Uuid of the fileservice.
