---
layout: "avi"
page_title: "AVI: avi_clusterclouddetails"
sidebar_current: "docs-avi-datasource-clusterclouddetails"
description: |-
  Get information of Avi ClusterCloudDetails.
---

# avi_clusterclouddetails

This data source is used to to get avi_clusterclouddetails objects.

## Example Usage

```hcl
data "ClusterCloudDetails" "foo_ClusterCloudDetails" {
    uuid = "ClusterCloudDetails-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ClusterCloudDetails by name.
* `uuid` - (Optional) Search ClusterCloudDetails by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `azure_info` - Azure info to configure cluster_vip on the controller.
* `name` - Field introduced in 17.2.5.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.2.5.
