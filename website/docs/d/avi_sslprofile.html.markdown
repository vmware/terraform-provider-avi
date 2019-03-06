
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
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
data "SSLProfile" "foo_SSLProfile" {
    uuid = "SSLProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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
* `cipher_enums` - Enum options - tls_ecdhe_ecdsa_with_aes_128_gcm_sha256, tls_ecdhe_ecdsa_with_aes_256_gcm_sha384, tls_ecdhe_rsa_with_aes_128_gcm_sha256, tls_ecdhe_rsa_with_aes_256_gcm_sha384, tls_ecdhe_ecdsa_with_aes_128_cbc_sha256, tls_ecdhe_ecdsa_with_aes_256_cbc_sha384, tls_ecdhe_rsa_with_aes_128_cbc_sha256, tls_ecdhe_rsa_with_aes_256_cbc_sha384, tls_rsa_with_aes_128_gcm_sha256, tls_rsa_with_aes_256_gcm_sha384, tls_rsa_with_aes_128_cbc_sha256, tls_rsa_with_aes_256_cbc_sha256, tls_ecdhe_ecdsa_with_aes_128_cbc_sha, tls_ecdhe_ecdsa_with_aes_256_cbc_sha, tls_ecdhe_rsa_with_aes_128_cbc_sha, tls_ecdhe_rsa_with_aes_256_cbc_sha, tls_rsa_with_aes_128_cbc_sha, tls_rsa_with_aes_256_cbc_sha, tls_rsa_with_3des_ede_cbc_sha, tls_rsa_with_rc4_128_sha.
* `description` - General description.
* `dhparam` - Dh parameters used in ssl.
* `enable_ssl_session_reuse` - Enable ssl session re-use.
* `name` - General description.
* `prefer_client_cipher_ordering` - Prefer the ssl cipher ordering presented by the client during the ssl handshake over the one specified in the ssl profile.
* `send_close_notify` - Send 'close notify' alert message for a clean shutdown of the ssl connection.
* `ssl_rating` - General description.
* `ssl_session_timeout` - The amount of time in seconds before an ssl session expires.
* `tags` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Ssl profile type.
* `uuid` - General description.

