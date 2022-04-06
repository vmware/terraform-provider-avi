<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_trafficcloneprofile"
sidebar_current: "docs-avi-datasource-trafficcloneprofile"
description: |-
  Get information of Avi TrafficCloneProfile.
---

# avi_trafficcloneprofile

This data source is used to to get avi_trafficcloneprofile objects.

## Example Usage

```hcl
data "avi_trafficcloneprofile" "foo_trafficcloneprofile" {
    uuid = "trafficcloneprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search TrafficCloneProfile by name.
* `uuid` - (Optional) Search TrafficCloneProfile by uuid.
* `cloud_ref` - (Optional) Search TrafficCloneProfile by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `clone_servers` - Field introduced in 17.1.1. Maximum of 10 items allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cloud_ref` - It is a reference to an object of type cloud. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name for the traffic clone profile. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `preserve_client_ip` - Specifies if client ip needs to be preserved to clone destination. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid of the traffic clone profile. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

