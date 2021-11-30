<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_siteversion"
sidebar_current: "docs-avi-datasource-siteversion"
description: |-
  Get information of Avi SiteVersion.
---

# avi_siteversion

This data source is used to to get avi_siteversion objects.

## Example Usage

```hcl
data "avi_siteversion" "foo_siteversion" {
    uuid = "siteversion-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SiteVersion by name.
* `uuid` - (Optional) Search SiteVersion by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `datetime` - This field represents the creation time of the federateddiff. Field introduced in 20.1.1.
* `name` - Name of the site. Field introduced in 20.1.1.
* `prev_target_version` - Previous targer version for a site. Field introduced in 20.1.1.
* `replication_state` - Replication state for a site. Enum options - REPLICATION_STATE_FASTFORWARD, REPLICATION_STATE_FORCESYNC, REPLICATION_STATE_STREAMING, REPLICATION_STATE_SUSPENDED, REPLICATION_STATE_INIT, REPLICATION_STATE_WAIT, REPLICATION_STATE_NOT_APPLICABLE. Field introduced in 20.1.1.
* `site_id` - Cluster uuid of the site. Field introduced in 20.1.1.
* `target_timeline` - Target timeline of the site. Field introduced in 20.1.1.
* `target_version` - Target version of the site. Field introduced in 20.1.1.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `timeline` - Timeline of the site. Field introduced in 20.1.1.
* `uuid` - Uuid of the siteversion object. Field introduced in 20.1.1.
* `version` - Version of the site. Field introduced in 20.1.1.
* `version_type` - Type of message for which version is maintained. Enum options - CONFIG_VERSION, HEALTH_STATUS_VERSION. Field introduced in 20.1.1.

