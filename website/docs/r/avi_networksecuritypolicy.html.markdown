<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_networksecuritypolicy"
sidebar_current: "docs-avi-resource-networksecuritypolicy"
description: |-
  Creates and manages Avi NetworkSecurityPolicy.
---

# avi_networksecuritypolicy

The NetworkSecurityPolicy resource allows the creation and management of Avi NetworkSecurityPolicy

## Example Usage

```hcl
resource "avi_networksecuritypolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for network sec policy. Internally set by cloud connector. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `created_by` - (Optional) Creator name. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `geo_db_ref` - (Optional) Geo database. It is a reference to an object of type geodb. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `internal` - (Optional) Network security policy is created and modified by internal modules only. Should not be modified by users. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ip_reputation_db_ref` - (Optional) Ip reputation database. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `rules` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

