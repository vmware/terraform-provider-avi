<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `name` - (Required) Name of the icap profile. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pool_group_ref` - (Required) The pool group which is used to connect to icap servers. It is a reference to an object of type poolgroup. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `service_uri` - (Required) The path and query component of the icap url. Host name and port will be taken from the pool. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `allow_204` - (Optional) Allow icap server to send 204 response as described in rfc 3507 section 4.5. Service engine will buffer the complete request if alllow_204 is enabled. If disabled, preview_size request body will be buffered if enable_preview is set to true, and rest of the request body will be streamed to the icap server. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `buffer_size` - (Optional) The maximum buffer size for the http request body. If the request body exceeds this size, the request will not be checked by the icap server. In this case, the configured action will be executed and a significant log entry will be generated. Allowed values are 1-51200. Field introduced in 20.1.1. Unit is kb. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `buffer_size_exceed_action` - (Optional) Decide what should happen if the request body size exceeds the configured buffer size. If this is set to fail open, the request will not be checked by the icap server. If this is set to fail closed, the request will be rejected with 413 status code. Enum options - ICAP_FAIL_OPEN, ICAP_FAIL_CLOSED. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cloud_ref` - (Optional) The cloud where this object belongs to. This must match the cloud referenced in the pool group below. It is a reference to an object of type cloud. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - (Optional) A description for this icap profile. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_preview` - (Optional) Use the icap preview feature as described in rfc 3507 section 4.5. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `fail_action` - (Optional) Decide what should happen if there is a problem with the icap server like communication timeout, protocol error, pool error, etc. If the icap server responds with 4xx-5xx error code the configured fail action is performed. If this is set to fail open, the request will continue, but will create a significant log entry. If this is set to fail closed, the request will be rejected with a 500 status code. Enum options - ICAP_FAIL_OPEN, ICAP_FAIL_CLOSED. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nsx_defender_config` - (Optional) Nsxdefender specific icap configurations. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `preview_size` - (Optional) The icap preview size as described in rfc 3507 section 4.5. This should not exceed the size supported by the icap server. If this is set to 0, only the http header will be sent to the icap server as a preview. To disable preview completely, set the enable-preview option to false.if vendor is lastline, recommended preview size is 1000 bytes,minimum preview size is 10 bytes. Allowed values are 0-5000. Field introduced in 20.1.1. Unit is bytes. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `response_timeout` - (Optional) Maximum time, client's request will be paused for icap processing. If this timeout is exceeded, the request to the icap server will be aborted and the configured fail action is executed. Allowed values are 50-3600000. Field introduced in 20.1.1. Unit is milliseconds. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `slow_response_warning_threshold` - (Optional) If the icap request takes longer than this value, this request will generate a significant log entry. Allowed values are 50-3600000. Field introduced in 20.1.1. Unit is milliseconds. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant which this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vendor` - (Optional) The vendor of the icap server. Enum options - ICAP_VENDOR_GENERIC, ICAP_VENDOR_OPSWAT, ICAP_VENDOR_LASTLINE. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the icap profile. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

