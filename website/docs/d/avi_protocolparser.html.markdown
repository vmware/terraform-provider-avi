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
page_title: "AVI: avi_protocolparser"
sidebar_current: "docs-avi-datasource-protocolparser"
description: |-
  Get information of Avi ProtocolParser.
---

# avi_protocolparser

This data source is used to to get avi_protocolparser objects.

## Example Usage

```hcl
data "avi_protocolparser" "foo_protocolparser" {
    uuid = "protocolparser-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ProtocolParser by name.
* `uuid` - (Optional) Search ProtocolParser by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - Description of the protocol parser. Field introduced in 18.2.3.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the protocol parser. Field introduced in 18.2.3.
* `parser_code` - Command script provided inline. Field introduced in 18.2.3.
* `tenant_ref` - Tenant uuid of the protocol parser. It is a reference to an object of type tenant. Field introduced in 18.2.3.
* `uuid` - Uuid of the protocol parser. Field introduced in 18.2.3.

