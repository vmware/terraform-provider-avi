<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `accepted_ciphers` - Ciphers suites represented as defined by https //www.openssl.org/docs/apps/ciphers.html. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `accepted_versions` - Set of versions accepted by the server. Minimum of 1 items required. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cipher_enums` - Enum options - tls_ecdhe_ecdsa_with_aes_128_gcm_sha256, tls_ecdhe_ecdsa_with_aes_256_gcm_sha384, tls_ecdhe_rsa_with_aes_128_gcm_sha256, tls_ecdhe_rsa_with_aes_256_gcm_sha384, tls_ecdhe_ecdsa_with_aes_128_cbc_sha256, tls_ecdhe_ecdsa_with_aes_256_cbc_sha384, tls_ecdhe_rsa_with_aes_128_cbc_sha256, tls_ecdhe_rsa_with_aes_256_cbc_sha384, tls_rsa_with_aes_128_gcm_sha256, tls_rsa_with_aes_256_gcm_sha384, tls_rsa_with_aes_128_cbc_sha256, tls_rsa_with_aes_256_cbc_sha256, tls_ecdhe_ecdsa_with_aes_128_cbc_sha, tls_ecdhe_ecdsa_with_aes_256_cbc_sha, tls_ecdhe_rsa_with_aes_128_cbc_sha, tls_ecdhe_rsa_with_aes_256_cbc_sha, tls_rsa_with_aes_128_cbc_sha, tls_rsa_with_aes_256_cbc_sha, tls_rsa_with_3des_ede_cbc_sha, tls_aes_256_gcm_sha384... Allowed in enterprise with any value edition, essentials(allowed values- tls_ecdhe_ecdsa_with_aes_128_gcm_sha256,tls_ecdhe_ecdsa_with_aes_256_gcm_sha384,tls_ecdhe_rsa_with_aes_128_gcm_sha256,tls_ecdhe_rsa_with_aes_256_gcm_sha384,tls_ecdhe_ecdsa_with_aes_128_cbc_sha256,tls_ecdhe_ecdsa_with_aes_256_cbc_sha384,tls_ecdhe_rsa_with_aes_128_cbc_sha256,tls_ecdhe_rsa_with_aes_256_cbc_sha384,tls_rsa_with_aes_128_gcm_sha256,tls_rsa_with_aes_256_gcm_sha384,tls_rsa_with_aes_128_cbc_sha256,tls_rsa_with_aes_256_cbc_sha256,tls_ecdhe_ecdsa_with_aes_128_cbc_sha,tls_ecdhe_ecdsa_with_aes_256_cbc_sha,tls_ecdhe_rsa_with_aes_128_cbc_sha,tls_ecdhe_rsa_with_aes_256_cbc_sha,tls_rsa_with_aes_128_cbc_sha,tls_rsa_with_aes_256_cbc_sha,tls_rsa_with_3des_ede_cbc_sha) edition, basic(allowed values- tls_ecdhe_ecdsa_with_aes_128_gcm_sha256,tls_ecdhe_ecdsa_with_aes_256_gcm_sha384,tls_ecdhe_rsa_with_aes_128_gcm_sha256,tls_ecdhe_rsa_with_aes_256_gcm_sha384,tls_ecdhe_ecdsa_with_aes_128_cbc_sha256,tls_ecdhe_ecdsa_with_aes_256_cbc_sha384,tls_ecdhe_rsa_with_aes_128_cbc_sha256,tls_ecdhe_rsa_with_aes_256_cbc_sha384,tls_rsa_with_aes_128_gcm_sha256,tls_rsa_with_aes_256_gcm_sha384,tls_rsa_with_aes_128_cbc_sha256,tls_rsa_with_aes_256_cbc_sha256,tls_ecdhe_ecdsa_with_aes_128_cbc_sha,tls_ecdhe_ecdsa_with_aes_256_cbc_sha,tls_ecdhe_rsa_with_aes_128_cbc_sha,tls_ecdhe_rsa_with_aes_256_cbc_sha,tls_rsa_with_aes_128_cbc_sha,tls_rsa_with_aes_256_cbc_sha,tls_rsa_with_3des_ede_cbc_sha) edition, enterprise with cloud services edition.
* `ciphersuites` - Tls 1.3 ciphers suites represented as defined by u(https //www.openssl.org/docs/manmaster/man1/ciphers.html). Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition. Special default for essentials edition is tls_aes_256_gcm_sha384-tls_aes_128_gcm_sha256, basic edition is tls_aes_256_gcm_sha384-tls_aes_128_gcm_sha256, enterprise is tls_aes_256_gcm_sha384-tls_chacha20_poly1305_sha256-tls_aes_128_gcm_sha256.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dhparam` - Dh parameters used in ssl. At this time, it is not configurable and is set to 2048 bits. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ec_named_curve` - Elliptic curve cryptography namedcurves (tls supported groups)represented as defined by rfc 8422-section 5.1.1 andhttps //www.openssl.org/docs/man1.1.0/man3/ssl_ctx_set1_curves.html. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `enable_early_data` - Enable early data processing for tls1.3 connections. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials(allowed values- false) edition, basic(allowed values- false) edition, enterprise with cloud services edition.
* `enable_ssl_session_reuse` - Enable ssl session re-use. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `prefer_client_cipher_ordering` - Prefer the ssl cipher ordering presented by the client during the ssl handshake over the one specified in the ssl profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `send_close_notify` - Send 'close notify' alert message for a clean shutdown of the ssl connection. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `signature_algorithm` - Signature algorithms represented as defined by rfc5246-section 7.4.1.4.1 andhttps //www.openssl.org/docs/man1.1.0/man3/ssl_ctx_set1_client_sigalgs_list.html. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ssl_rating` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ssl_session_timeout` - The amount of time in seconds before an ssl session expires. Unit is sec. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tags` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `type` - Ssl profile type. Enum options - SSL_PROFILE_TYPE_APPLICATION, SSL_PROFILE_TYPE_SYSTEM. Field introduced in 17.2.8. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

