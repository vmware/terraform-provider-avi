<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_cloudproperties"
sidebar_current: "docs-avi-datasource-cloudproperties"
description: |-
  Get information of Avi CloudProperties.
---

# avi_cloudproperties

This data source is used to to get avi_cloudproperties objects.

## Example Usage

```hcl
data "avi_cloudproperties" "foo_cloudproperties" {
    uuid = "cloudproperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CloudProperties by name.
* `uuid` - (Optional) Search CloudProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cc_props` - Cloudconnector properties. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cc_vtypes` - Cloud types supported by cloudconnector. Enum options - CLOUD_NONE, CLOUD_VCENTER, CLOUD_OPENSTACK, CLOUD_AWS, CLOUD_VCA, CLOUD_APIC, CLOUD_MESOS, CLOUD_LINUXSERVER, CLOUD_DOCKER_UCP, CLOUD_RANCHER, CLOUD_OSHIFT_K8S, CLOUD_AZURE, CLOUD_GCP, CLOUD_NSXT. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `hyp_props` - Hypervisor properties. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `info` - Properties specific to a cloud type. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

