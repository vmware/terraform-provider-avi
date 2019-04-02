
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

---
layout: "avi"
page_title: "AVI: avi_httppolicyset"
sidebar_current: "docs-avi-datasource-httppolicyset"
description: |-
  Get information of Avi HTTPPolicySet.
---

# avi_httppolicyset

This data source is used to to get avi_httppolicyset objects.

## Example Usage

```hcl
data "HTTPPolicySet" "foo_HTTPPolicySet" {
    uuid = "HTTPPolicySet-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search HTTPPolicySet by name.
* `uuid` - (Optional) Search HTTPPolicySet by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of cloud configuration for pool.
* `created_by` - Creator name.
* `description` - General description.
* `http_request_policy` - Http request policy for the virtual service.
* `http_response_policy` - Http response policy for the virtual service.
* `http_security_policy` - Http security policy for the virtual service.
* `is_internal_policy` - General description.
* `name` - Name of the http policy set.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the http policy set.
