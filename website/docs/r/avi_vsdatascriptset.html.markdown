<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_vsdatascriptset"
sidebar_current: "docs-avi-resource-vsdatascriptset"
description: |-
  Creates and manages Avi VSDataScriptSet.
---

# avi_vsdatascriptset

The VSDataScriptSet resource allows the creation and management of Avi VSDataScriptSet

## Example Usage

```hcl
resource "avi_vsdatascriptset" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the virtual service datascript collection.
* `created_by` - (Optional) Creator name. Field introduced in 17.1.11,17.2.4.
* `datascript` - (Optional) Datascripts to execute.
* `description` - (Optional) User defined description for the object.
* `ip_reputation_db_ref` - (Optional) Ip reputation database that can be used by datascript functions. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.3.
* `ipgroup_refs` - (Optional) Uuid of ip groups that could be referred by vsdatascriptset objects. It is a reference to an object of type ipaddrgroup.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `pool_group_refs` - (Optional) Uuid of pool groups that could be referred by vsdatascriptset objects. It is a reference to an object of type poolgroup.
* `pool_refs` - (Optional) Uuid of pools that could be referred by vsdatascriptset objects. It is a reference to an object of type pool.
* `protocol_parser_refs` - (Optional) List of protocol parsers that could be referred by vsdatascriptset objects. It is a reference to an object of type protocolparser. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `rate_limiters` - (Optional) The rate limit definitions needed for this datascript. The name is composed of the virtual service name and the datascript name. Field introduced in 18.2.9. Allowed in basic edition, essentials edition, enterprise edition.
* `string_group_refs` - (Optional) Uuid of string groups that could be referred by vsdatascriptset objects. It is a reference to an object of type stringgroup.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the virtual service datascript collection.

