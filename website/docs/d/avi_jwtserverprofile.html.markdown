############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_jwtserverprofile"
sidebar_current: "docs-avi-datasource-jwtserverprofile"
description: |-
  Get information of Avi JWTServerProfile.
---

# avi_jwtserverprofile

This data source is used to to get avi_jwtserverprofile objects.

## Example Usage

```hcl
data "avi_jwtserverprofile" "foo_jwtserverprofile" {
    uuid = "jwtserverprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search JWTServerProfile by name.
* `uuid` - (Optional) Search JWTServerProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `issuer` - Uniquely identifiable name of the token issuer. Field introduced in 20.1.3.
* `jwks_keys` - Jwks key set used for validating the jwt. Field introduced in 20.1.3.
* `name` - Name of the jwt profile. Field introduced in 20.1.3.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 20.1.3.
* `uuid` - Uuid of the jwtprofile. Field introduced in 20.1.3.

