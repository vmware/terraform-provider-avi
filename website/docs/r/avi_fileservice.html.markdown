---
layout: "avi"
page_title: "Avi: avi_fileservice"
sidebar_current: "docs-avi-resource-fileservice"
description: |-
  Upload and Download files.
---

# avi_fileservice

The Fileservice resource allows the download and upload of files

## Example Usage

```hcl
resource "avi_fileservice" "foo" {
    uri = "/uploads"
    local_file = "/file/path"
    upload = True
}
```

## Argument Reference

The following arguments are supported:

    * `uri` - (Required) argument_description.
    * `local_file` - (Required) argument_description.
    * `upload` - (Optional ) argument_description.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

    * `uuid` - argument_description.

                                                                                                                                                                                                                                                        * `uuid` - argument_description.
