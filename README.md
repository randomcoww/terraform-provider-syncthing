# Terraform Provider Syncthing (Terraform Plugin Framework)

_This template repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). The template repository built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) can be found at [terraform-provider-scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding). See [Which SDK Should I Use?](https://www.terraform.io/docs/plugin/which-sdk.html) in the Terraform documentation for additional information._

This repository is a *template* for a [Terraform](https://www.terraform.io) provider. It is intended as a starting point for creating Terraform providers, containing:

- A resource and a data source (`internal/provider/`),
- Examples (`examples/`) and generated documentation (`docs/`),
- Miscellaneous meta files.

These files contain boilerplate code that you will need to edit to create your own Terraform provider. Tutorials for creating Terraform providers can be found on the [HashiCorp Learn](https://learn.hashicorp.com/collections/terraform/providers-plugin-framework) platform. _Terraform Plugin Framework specific guides are titled accordingly._

Please see the [GitHub template repository documentation](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) for how to create a new repository from this template on GitHub.

Once you've written your provider, you'll want to [publish it on the Terraform Registry](https://www.terraform.io/docs/registry/providers/publishing.html) so that others can use it.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

#### Generate certs

Common name should be `syncthing` to avoid needing to specify it in Syncthing configuration

```shell
resource "tls_private_key" "test" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P384"
}

resource "tls_cert_request" "test" {
  key_algorithm   = tls_private_key.test.algorithm
  private_key_pem = tls_private_key.test.private_key_pem

  subject {
    common_name = "syncthing"
  }
}

resource "tls_locally_signed_cert" "test" {
  cert_request_pem   = tls_cert_request.test.cert_request_pem
  ca_key_algorithm   = tls_private_key.root.algorithm
  ca_private_key_pem = tls_private_key.root.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.root.cert_pem

  validity_period_hours = 8760

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
    "client_auth",
  ]
}
```

#### Generate Syncthing device ID from certs

```shell
data "syncthing_device" "test" {
  cert_pem        = tls_locally_signed_cert.test.cert_pem
  private_key_pem = tls_private_key.test.private_key_pem
}

device_id = data.syncthing_device.test.id
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
