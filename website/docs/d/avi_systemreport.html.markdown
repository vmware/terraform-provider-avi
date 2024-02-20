<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_systemreport"
sidebar_current: "docs-avi-datasource-systemreport"
description: |-
  Get information of Avi SystemReport.
---

# avi_systemreport

This data source is used to to get avi_systemreport objects.

## Example Usage

```hcl
data "avi_systemreport" "foo_systemreport" {
    uuid = "systemreport-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SystemReport by name.
* `uuid` - (Optional) Search SystemReport by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `archive_ref` - Relative path to the report archive file on filesystem.the archive includes exported system configuration and current object as json. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `controller_patch_image_ref` - Controller patch image associated with the report. It is a reference to an object of type image. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `downloadable` - Indicates whether this report is downloadable as an archive. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `events` - List of events associated with the report. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `image_ref` - System image associated with the report. It is a reference to an object of type image. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of the report dervied from operation in a readable format. Ex  upgrade_system_1a5c. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `readiness_reports` - Readiness state of the system. Ex  upgrade pre-check results. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_patch_image_ref` - Se patch image associated with the report. It is a reference to an object of type image. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `state` - Report state combines all applicable states. Ex  readiness_reports.system_readiness.state. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `summary` - Summary of the report. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the report. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.

