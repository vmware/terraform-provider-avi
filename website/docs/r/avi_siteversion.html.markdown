############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `datetime` - (Optional) This field represents the creation time of the federateddiff.
* `name` - (Optional) Name of the site.
* `prev_target_version` - (Optional) Previous targer version for a site.
* `replication_state` - (Optional) Replication state for a site.
* `site_id` - (Optional) Cluster uuid of the site.
* `target_timeline` - (Optional) Target timeline of the site.
* `target_version` - (Optional) Target version of the site.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `timeline` - (Optional) Timeline of the site.
* `version` - (Optional) Version of the site.
* `version_type` - (Optional) Type of message for which version is maintained.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the siteversion object.

