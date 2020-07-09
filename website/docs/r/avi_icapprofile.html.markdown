---
layout: "avi"
page_title: "Avi: avi_icapprofile"
sidebar_current: "docs-avi-resource-icapprofile"
description: |-
  Creates and manages Avi IcapProfile.
---

# avi_icapprofile

The IcapProfile resource allows the creation and management of Avi IcapProfile

## Example Usage

```hcl
resource "avi_icapprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `buffer_size` - (Optional) The maximum buffer size for the http request body.
* `buffer_size_exceed_action` - (Optional) Decide what should happen if the request body size exceeds the configured buffer size.
* `cloud_ref` - (Optional) The cloud where this object belongs to.
* `description` - (Optional) A description for this icap profile.
* `enable_preview` - (Optional) Use the icap preview feature as described in rfc 3507 section 4.5.
* `fail_action` - (Optional) Decide what should happen if there is a problem with the icap server like communication timeout, protocol error, pool error, etc.
* `name` - (Optional) Name of the icap profile.
* `pool_group_ref` - (Optional) The pool group which is used to connect to icap servers.
* `preview_size` - (Optional) The icap preview size as described in rfc 3507 section 4.5.
* `response_timeout` - (Optional) How long do we wait for a request to the icap server to finish.
* `service_uri` - (Optional) The path and query component of the icap url.
* `slow_response_warning_threshold` - (Optional) If the icap request takes longer than this value, this request will generate a significant log entry.
* `tenant_ref` - (Optional) Tenant which this object belongs to.
* `vendor` - (Optional) The vendor of the icap server.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the icap profile.

