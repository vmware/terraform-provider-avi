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

* `allow_204` - Allow icap server to send 204 response as described in rfc 3507 section 4.5. Service engine will buffer the complete request if alllow_204 is enabled. If disabled, preview_size request body will be buffered if enable_preview is set to true, and rest of the request body will be streamed to the icap server. Field introduced in 20.1.3.
* `buffer_size` - The maximum buffer size for the http request body. If the request body exceeds this size, the request will not be checked by the icap server. In this case, the configured action will be executed and a significant log entry will be generated. Allowed values are 1-51200. Field introduced in 20.1.1. Unit is kb.
* `buffer_size_exceed_action` - Decide what should happen if the request body size exceeds the configured buffer size. If this is set to fail open, the request will not be checked by the icap server. If this is set to fail closed, the request will be rejected with 413 status code. Enum options - ICAP_FAIL_OPEN, ICAP_FAIL_CLOSED. Field introduced in 20.1.1.
* `cloud_ref` - The cloud where this object belongs to. This must match the cloud referenced in the pool group below. It is a reference to an object of type cloud. Field introduced in 20.1.1.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - A description for this icap profile. Field introduced in 20.1.1.
* `enable_preview` - Use the icap preview feature as described in rfc 3507 section 4.5. Field introduced in 20.1.1.
* `fail_action` - Decide what should happen if there is a problem with the icap server like communication timeout, protocol error, pool error, etc. If this is set to fail open, the request will continue, but will create a significant log entry. If this is set to fail closed, the request will be rejected with a 500 status code. Enum options - ICAP_FAIL_OPEN, ICAP_FAIL_CLOSED. Field introduced in 20.1.1.
* `name` - Name of the icap profile. Field introduced in 20.1.1.
* `nsx_defender_config` - Nsxdefender specific icap configurations. Field introduced in 21.1.1.
* `pool_group_ref` - The pool group which is used to connect to icap servers. It is a reference to an object of type poolgroup. Field introduced in 20.1.1.
* `preview_size` - The icap preview size as described in rfc 3507 section 4.5. This should not exceed the size supported by the icap server. If this is set to 0, only the http header will be sent to the icap server as a preview. To disable preview completely, set the enable-preview option to false. Allowed values are 0-5000. Field introduced in 20.1.1. Unit is bytes.
* `response_timeout` - Maximum time, client's request will be paused for icap processing. If this timeout is exceeded, the request to the icap server will be aborted and the configured fail action is executed. Allowed values are 50-3600000. Field introduced in 20.1.1. Unit is milliseconds.
* `service_uri` - The path and query component of the icap url. Host name and port will be taken from the pool. Field introduced in 20.1.1.
* `slow_response_warning_threshold` - If the icap request takes longer than this value, this request will generate a significant log entry. Allowed values are 50-3600000. Field introduced in 20.1.1. Unit is milliseconds.
* `tenant_ref` - Tenant which this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `uuid` - Uuid of the icap profile. Field introduced in 20.1.1.
* `vendor` - The vendor of the icap server. Enum options - ICAP_VENDOR_GENERIC, ICAP_VENDOR_OPSWAT, ICAP_VENDOR_LASTLINE. Field introduced in 20.1.1.

