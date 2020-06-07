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

* `datetime` - This field represents the creation time of the federateddiff.
* `name` - Name of the site.
* `prev_target_version` - Previous targer version for a site.
* `replication_state` - Replication state for a site.
* `site_id` - Cluster uuid of the site.
* `target_timeline` - Target timeline of the site.
* `target_version` - Target version of the site.
* `tenant_ref` - Tenant that this object belongs to.
* `timeline` - Timeline of the site.
* `uuid` - Uuid of the siteversion object.
* `version` - Version of the site.
* `version_type` - Type of message for which version is maintained.

