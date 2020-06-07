---
layout: "avi"
page_title: "Avi: avi_fileobject"
sidebar_current: "docs-avi-resource-fileobject"
description: |-
  Creates and manages Avi FileObject.
---

# avi_fileobject

The FileObject resource allows the creation and management of Avi FileObject

## Example Usage

```hcl
resource "avi_fileobject" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `checksum` - (Optional) Sha1 checksum of the file.
* `compressed` - (Optional) This field indicates whether the file is gzip-compressed.
* `created` - (Optional) Timestamp of creation for the file.
* `description` - (Optional) Description of the file.
* `expires_at` - (Optional) Timestamp when the file will be no longer needed and can be removed by the system.
* `is_federated` - (Optional) This field describes the object's replication scope.
* `name` - (Optional) Name of the file object.
* `path` - (Optional) Path to the file.
* `read_only` - (Optional) Enforce read-only on the file.
* `restrict_download` - (Optional) Flag to allow/restrict download of the file.
* `size` - (Optional) Size of the file.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `type` - (Optional) Type of the file.
* `version` - (Optional) Version of the file.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the file.

