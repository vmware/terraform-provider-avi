<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_pkiprofile"
sidebar_current: "docs-avi-datasource-pkiprofile"
description: |-
  Get information of Avi PKIProfile.
---

# avi_pkiprofile

This data source is used to to get avi_pkiprofile objects.

## Example Usage

```hcl
data "avi_pkiprofile" "foo_pkiprofile" {
    uuid = "pkiprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PKIProfile by name.
* `uuid` - (Optional) Search PKIProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ca_certs` - List of certificate authorities (root and intermediate) trusted that is used for certificate validation.
* `created_by` - Creator name.
* `crl_check` - When enabled, avi will verify via crl checks that certificates in the trust chain have not been revoked.
* `crls` - Certificate revocation lists.
* `ignore_peer_chain` - When enabled, avi will not trust intermediate and root certs presented by a client. Instead, only the chain certs configured in the certificate authority section will be used to verify trust of the client's cert. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition. Special default for basic edition is true, essentials edition is true, enterprise is false.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the pki profile.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.
* `validate_only_leaf_crl` - When enabled, avi will only validate the revocation status of the leaf certificate using crl. To enable validation for the entire chain, disable this option and provide all the relevant crls. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.

