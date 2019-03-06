
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
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
data "TrafficCloneProfile" "foo_TrafficCloneProfile" {
    uuid = "TrafficCloneProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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

* `clone_servers` - Field introduced in 17.1.1.
* `cloud_ref` - It is a reference to an object of type cloud.
* `name` - Name for the traffic clone profile.
* `preserve_client_ip` - Specifies if client ip needs to be preserved to clone destination.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the traffic clone profile.

