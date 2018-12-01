## Syncthing Helper Terraform Provider 

Generate a Syncthing compatible device ID from pre-generated cert and key.

### Usage

Generate key pair as normal. Common name should be "syncthing" to avoid needing to specify it in Syncthing configuration.
```
resource "tls_private_key" "test" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P384"
}

resource "tls_cert_request" "test" {
  key_algorithm   = "${tls_private_key.test.algorithm}"
  private_key_pem = "${tls_private_key.test.private_key_pem}"

  subject {
    common_name = "syncthing"
  }
}

resource "tls_locally_signed_cert" "test" {
  cert_request_pem   = "${tls_cert_request.test.cert_request_pem}"
  ca_key_algorithm   = "${tls_private_key.root.algorithm}"
  ca_private_key_pem = "${tls_private_key.root.private_key_pem}"
  ca_cert_pem        = "${tls_self_signed_cert.root.cert_pem}"

  validity_period_hours = 8760

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
    "client_auth",
  ]
}
```

Call data source
```
data "syncthing_device" "test" {
  cert_pem        = "${tls_locally_signed_cert.test.cert_pem}"
  private_key_pem = "${tls_private_key.test.private_key_pem}"
}
```

Use device ID in variable
```
device_id = "${data.syncthing_device.test.device_id}"
```
