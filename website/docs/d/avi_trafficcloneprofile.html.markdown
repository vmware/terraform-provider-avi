############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `clone_servers` - Field introduced in 17.1.1. Maximum of 10 items allowed.
* `cloud_ref` - It is a reference to an object of type cloud. Field introduced in 17.1.1.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - Name for the traffic clone profile. Field introduced in 17.1.1.
* `preserve_client_ip` - Specifies if client ip needs to be preserved to clone destination. Field introduced in 17.1.1.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1.
* `uuid` - Uuid of the traffic clone profile. Field introduced in 17.1.1.

