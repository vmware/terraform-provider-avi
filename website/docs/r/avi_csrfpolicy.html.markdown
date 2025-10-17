<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_csrfpolicy"
sidebar_current: "docs-avi-resource-csrfpolicy"
description: |-
  Creates and manages Avi CSRFPolicy.
---

# avi_csrfpolicy

The CSRFPolicy resource allows the creation and management of Avi CSRFPolicy

## Example Usage

```hcl
resource "avi_csrfpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of this csrf protection policy. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `rules` - (Required) Rules to control which requests undergo csrf protection.if the client's request doesn't match with any rules matchtarget, bypass_csrf action is applied. Field introduced in 30.2.1. Minimum of 1 items required. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `cookie_name` - (Optional) Name of the cookie to be used for csrf token. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `csrf_file_ref` - (Optional) The file object that contains csrf javascript content. Must be of type 'csrf'. It is a reference to an object of type fileobject. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `description` - (Optional) Human-readable description of this csrf protection policy. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this policy belongs. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `token_validity_time_min` - (Optional) A csrf token is rotated when this amount of time has passed. Even after that, tokens will be accepted until twice this amount of time has passed. Note, however, that other timeouts from the underlying session layer also affect how long a given token can be used. A token will be invalidated (rotated or deleted) after one of 'token_validity_time_min' (this value), 'session_establishment_timeout', 'session_idle_timeout', 'session_maximum_timeout' is reached, whichever occurs first. Allowed values are 10-1440. Special values are 0- unlimited. Field introduced in 30.2.1. Unit is min. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier to this csrf protection policy. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

