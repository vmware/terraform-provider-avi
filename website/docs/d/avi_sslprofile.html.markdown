############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_sslprofile"
sidebar_current: "docs-avi-datasource-sslprofile"
description: |-
  Get information of Avi SSLProfile.
---

# avi_sslprofile

This data source is used to to get avi_sslprofile objects.

## Example Usage

```hcl
data "avi_sslprofile" "foo_sslprofile" {
    uuid = "sslprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SSLProfile by name.
* `uuid` - (Optional) Search SSLProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `accepted_ciphers` - Ciphers suites represented as defined by u(http //www.openssl.org/docs/apps/ciphers.html).
* `accepted_versions` - Set of versions accepted by the server.
* `cipher_enums` - Enum options - tls_ecdhe_ecdsa_with_aes_128_gcm_sha256, tls_ecdhe_ecdsa_with_aes_256_gcm_sha384, tls_ecdhe_rsa_with_aes_128_gcm_sha256, tls_ecdhe_rsa_with_aes_256_gcm_sha384, tls_ecdhe_ecdsa_with_aes_128_cbc_sha256, tls_ecdhe_ecdsa_with_aes_256_cbc_sha384, tls_ecdhe_rsa_with_aes_128_cbc_sha256, tls_ecdhe_rsa_with_aes_256_cbc_sha384, tls_rsa_with_aes_128_gcm_sha256, tls_rsa_with_aes_256_gcm_sha384, tls_rsa_with_aes_128_cbc_sha256, tls_rsa_with_aes_256_cbc_sha256, tls_ecdhe_ecdsa_with_aes_128_cbc_sha, tls_ecdhe_ecdsa_with_aes_256_cbc_sha, tls_ecdhe_rsa_with_aes_128_cbc_sha, tls_ecdhe_rsa_with_aes_256_cbc_sha, tls_rsa_with_aes_128_cbc_sha, tls_rsa_with_aes_256_cbc_sha, tls_rsa_with_3des_ede_cbc_sha, tls_aes_256_gcm_sha384, tls_chacha20_poly1305_sha256, tls_aes_128_gcm_sha256.
* `ciphersuites` - Tls 1.3 ciphers suites represented as defined by u(https //www.openssl.org/docs/manmaster/man1/ciphers.html).
* `description` - User defined description for the object.
* `dhparam` - Dh parameters used in ssl.
* `enable_early_data` - Enable early data processing for tls1.3 connections.
* `enable_ssl_session_reuse` - Enable ssl session re-use.
* `labels` - Key value pairs for granular object access control.
* `name` - Name of the object.
* `prefer_client_cipher_ordering` - Prefer the ssl cipher ordering presented by the client during the ssl handshake over the one specified in the ssl profile.
* `send_close_notify` - Send 'close notify' alert message for a clean shutdown of the ssl connection.
* `ssl_rating` - Dict settings for sslprofile.
* `ssl_session_timeout` - The amount of time in seconds before an ssl session expires.
* `tags` - List of list.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Ssl profile type.
* `uuid` - Unique object identifier of the object.

