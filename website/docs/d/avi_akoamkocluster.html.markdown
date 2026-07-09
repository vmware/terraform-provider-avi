<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_akoamkocluster"
sidebar_current: "docs-avi-datasource-akoamkocluster"
description: |-
  Get information of Avi AkoAmkoCluster.
---

# avi_akoamkocluster

This data source is used to to get avi_akoamkocluster objects.

## Example Usage

```hcl
data "avi_akoamkocluster" "foo_akoamkocluster" {
    uuid = "akoamkocluster-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search AkoAmkoCluster by name.
* `uuid` - (Optional) Search AkoAmkoCluster by uuid.
* `cloud_ref` - (Optional) Search AkoAmkoCluster by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of the cloud configuration for akoamkocluster object. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `cloud_ref` - Cloud reference uuid in avi controller. It is a reference to an object of type cloud. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `cluster_type` - Type of operator - ako or amko. Enum options - CLUSTER_TYPE_AKO, CLUSTER_TYPE_AMKO. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `created_by` - Ako/amko user identifier. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `deployment_info` - Deployment configuration information. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `metadata` - Additional cluster metadata. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the ako/amko cluster. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant that ako/amko cluster belongs to. It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the ako/amko cluster. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `version_info` - Version information including kubernetes and ako/amko versions. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

