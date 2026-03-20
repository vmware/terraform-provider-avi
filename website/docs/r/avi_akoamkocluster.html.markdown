<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_akoamkocluster"
sidebar_current: "docs-avi-resource-akoamkocluster"
description: |-
  Creates and manages Avi AkoAmkoCluster.
---

# avi_akoamkocluster

The AkoAmkoCluster resource allows the creation and management of Avi AkoAmkoCluster

## Example Usage

```hcl
resource "avi_akoamkocluster" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `created_by` - (Required) Ako/amko user identifier. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `name` - (Required) Name of the ako/amko cluster. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `cloud_ref` - (Optional) Cloud reference uuid in avi controller. It is a reference to an object of type cloud. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `cluster_type` - (Optional) Type of operator - ako or amko. Enum options - CLUSTER_TYPE_AKO, CLUSTER_TYPE_AMKO. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `deployment_info` - (Optional) Deployment configuration information. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `metadata` - (Optional) Additional cluster metadata. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `version_info` - (Optional) Version information including kubernetes and ako/amko versions. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the ako/amko cluster. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

