############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_siteversion"
sidebar_current: "docs-avi-resource-siteversion"
description: |-
  Creates and manages Avi SiteVersion.
---

# avi_siteversion

The SiteVersion resource allows the creation and management of Avi SiteVersion

## Example Usage

```hcl
resource "avi_siteversion" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the site. Field introduced in 20.1.1.
* `datetime` - (Optional) This field represents the creation time of the federateddiff. Field introduced in 20.1.1.
* `prev_target_version` - (Optional) Previous targer version for a site. Field introduced in 20.1.1.
* `replication_state` - (Optional) Replication state for a site. Enum options - REPLICATION_STATE_FASTFORWARD, REPLICATION_STATE_FORCESYNC, REPLICATION_STATE_STREAMING, REPLICATION_STATE_SUSPENDED, REPLICATION_STATE_INIT, REPLICATION_STATE_WAIT, REPLICATION_STATE_NOT_APPLICABLE. Field introduced in 20.1.1.
* `site_id` - (Optional) Cluster uuid of the site. Field introduced in 20.1.1.
* `target_timeline` - (Optional) Target timeline of the site. Field introduced in 20.1.1.
* `target_version` - (Optional) Target version of the site. Field introduced in 20.1.1.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `timeline` - (Optional) Timeline of the site. Field introduced in 20.1.1.
* `version` - (Optional) Version of the site. Field introduced in 20.1.1.
* `version_type` - (Optional) Type of message for which version is maintained. Enum options - CONFIG_VERSION, HEALTH_STATUS_VERSION. Field introduced in 20.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the siteversion object. Field introduced in 20.1.1.

