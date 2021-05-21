<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_dnspolicy"
sidebar_current: "docs-avi-resource-dnspolicy"
description: |-
  Creates and manages Avi DnsPolicy.
---

# avi_dnspolicy

The DnsPolicy resource allows the creation and management of Avi DnsPolicy

## Example Usage

```hcl
resource "avi_dnspolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the dns policy. Field introduced in 17.1.1.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `created_by` - (Optional) Creator name. Field introduced in 17.1.1.
* `description` - (Optional) Field introduced in 17.1.1.
* `internal` - (Optional) The dns policy is created and modified by internal modules only. This should not be modified by users. Field introduced in 21.1.1.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `rule` - (Optional) Dns rules. Field introduced in 17.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the dns policy. Field introduced in 17.1.1.

