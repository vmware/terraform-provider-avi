############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_federationcheckpoint"
sidebar_current: "docs-avi-datasource-federationcheckpoint"
description: |-
  Get information of Avi FederationCheckpoint.
---

# avi_federationcheckpoint

This data source is used to to get avi_federationcheckpoint objects.

## Example Usage

```hcl
data "avi_federationcheckpoint" "foo_federationcheckpoint" {
    uuid = "federationcheckpoint-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search FederationCheckpoint by name.
* `uuid` - (Optional) Search FederationCheckpoint by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `date` - Date when the checkpoint was created. Field introduced in 20.1.1.
* `description` - Description for this checkpoint. Field introduced in 20.1.1.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.1.
* `name` - Name of the checkpoint. Field introduced in 20.1.1.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `uuid` - Uuid of the checkpoint. Field introduced in 20.1.1.

