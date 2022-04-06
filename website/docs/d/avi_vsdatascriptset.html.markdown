<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_vsdatascriptset"
sidebar_current: "docs-avi-datasource-vsdatascriptset"
description: |-
  Get information of Avi VSDataScriptSet.
---

# avi_vsdatascriptset

This data source is used to to get avi_vsdatascriptset objects.

## Example Usage

```hcl
data "avi_vsdatascriptset" "foo_vsdatascriptset" {
    uuid = "vsdatascriptset-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search VSDataScriptSet by name.
* `uuid` - (Optional) Search VSDataScriptSet by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `created_by` - Creator name. Field introduced in 17.1.11,17.2.4. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `datascript` - Datascripts to execute. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `geo_db_ref` - Geo location mapping database used by this datascriptset. It is a reference to an object of type geodb. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ip_reputation_db_ref` - Ip reputation database that can be used by datascript functions. It is a reference to an object of type ipreputationdb. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ipgroup_refs` - Uuid of ip groups that could be referred by vsdatascriptset objects. It is a reference to an object of type ipaddrgroup. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name for the virtual service datascript collection. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `pki_profile_refs` - Uuids of pkiprofile objects that could be referred by vsdatascriptset objects. It is a reference to an object of type pkiprofile. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `pool_group_refs` - Uuid of pool groups that could be referred by vsdatascriptset objects. It is a reference to an object of type poolgroup. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `pool_refs` - Uuid of pools that could be referred by vsdatascriptset objects. It is a reference to an object of type pool. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `protocol_parser_refs` - List of protocol parsers that could be referred by vsdatascriptset objects. It is a reference to an object of type protocolparser. Field introduced in 18.2.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `rate_limiters` - The rate limit definitions needed for this datascript. The name is composed of the virtual service name and the datascript name. Field introduced in 18.2.9. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ssl_key_certificate_refs` - Uuids of sslkeyandcertificate objects that could be referred by vsdatascriptset objects. It is a reference to an object of type sslkeyandcertificate. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ssl_profile_refs` - Uuids of sslprofile objects that could be referred by vsdatascriptset objects. It is a reference to an object of type sslprofile. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `string_group_refs` - Uuid of string groups that could be referred by vsdatascriptset objects. It is a reference to an object of type stringgroup. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid of the virtual service datascript collection. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

