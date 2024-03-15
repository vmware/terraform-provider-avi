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

* `allow_pki_errors` - Exempt errors during certificate verification. Enum options - ALLOW_EXPIRED_CRL, ALLOW_ALL_ERRORS. Field introduced in 30.1.1. Maximum of 1 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ca_certs` - List of certificate authorities (root and intermediate) trusted that is used for certificate validation. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - Creator name. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `crl_check` - When enabled, avi will verify via crl checks that certificates in the trust chain have not been revoked. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `crl_file_refs` - Refers to fileobject containing crl body. It is a reference to an object of type fileobject. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ignore_peer_chain` - When enabled, avi will not trust intermediate and root certs presented by a client. Instead, only the chain certs configured in the certificate authority section will be used to verify trust of the client's cert. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition. Special default for essentials edition is true, basic edition is true, enterprise is false.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Name of the pki profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `validate_only_leaf_crl` - When enabled, avi will only validate the revocation status of the leaf certificate using crl. To enable validation for the entire chain, disable this option and provide all the relevant crls. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition.

