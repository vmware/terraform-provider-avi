---
layout: "avi"
page_title: "AVI: avi_icapprofile"
sidebar_current: "docs-avi-datasource-icapprofile"
description: |-
  Get information of Avi IcapProfile.
---

# avi_icapprofile

This data source is used to to get avi_icapprofile objects.

## Example Usage

```hcl
data "avi_icapprofile" "foo_icapprofile" {
    uuid = "icapprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search IcapProfile by name.
* `uuid` - (Optional) Search IcapProfile by uuid.
* `cloud_ref` - (Optional) Search IcapProfile by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `buffer_size` - The maximum buffer size for the http request body.
* `buffer_size_exceed_action` - Decide what should happen if the request body size exceeds the configured buffer size.
* `cloud_ref` - The cloud where this object belongs to.
* `description` - A description for this icap profile.
* `enable_preview` - Use the icap preview feature as described in rfc 3507 section 4.5.
* `fail_action` - Decide what should happen if there is a problem with the icap server like communication timeout, protocol error, pool error, etc.
* `name` - Name of the icap profile.
* `pool_group_ref` - The pool group which is used to connect to icap servers.
* `preview_size` - The icap preview size as described in rfc 3507 section 4.5.
* `response_timeout` - How long do we wait for a request to the icap server to finish.
* `service_uri` - The path and query component of the icap url.
* `slow_response_warning_threshold` - If the icap request takes longer than this value, this request will generate a significant log entry.
* `tenant_ref` - Tenant which this object belongs to.
* `uuid` - Uuid of the icap profile.
* `vendor` - The vendor of the icap server.

