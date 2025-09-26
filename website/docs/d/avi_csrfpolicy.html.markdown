<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_csrfpolicy"
sidebar_current: "docs-avi-datasource-csrfpolicy"
description: |-
  Get information of Avi CSRFPolicy.
---

# avi_csrfpolicy

This data source is used to to get avi_csrfpolicy objects.

## Example Usage

```hcl
data "avi_csrfpolicy" "foo_csrfpolicy" {
    uuid = "csrfpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CSRFPolicy by name.
* `uuid` - (Optional) Search CSRFPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `cookie_name` - Name of the cookie to be used for csrf token. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `csrf_file_ref` - The file object that contains csrf javascript content. Must be of type 'csrf'. It is a reference to an object of type fileobject. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `description` - Human-readable description of this csrf protection policy. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - The name of this csrf protection policy. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `rules` - Rules to control which requests undergo csrf protection.if the client's request doesn't match with any rules matchtarget, bypass_csrf action is applied. Field introduced in 30.2.1. Minimum of 1 items required. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - The unique identifier of the tenant to which this policy belongs. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `token_validity_time_min` - A csrf token is rotated when this amount of time has passed. Even after that, tokens will be accepted until twice this amount of time has passed. Note, however, that other timeouts from the underlying session layer also affect how long a given token can be used. A token will be invalidated (rotated or deleted) after one of 'token_validity_time_min' (this value), 'session_establishment_timeout', 'session_idle_timeout', 'session_maximum_timeout' is reached, whichever occurs first. Allowed values are 10-1440. Special values are 0- unlimited. Field introduced in 30.2.1. Unit is min. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - A unique identifier to this csrf protection policy. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

