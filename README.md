# Terraform Provider WorkOS

This provider should be used to manage WorkOS.

Warning: this project has been made for educational purpose and is not guaranteed to be continued.

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-hashicups
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

## Developer Notes

### Documentation

Documentation is generated using `tfplugindocs`. See [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) for more details.

A `go:generate` comment is included in the `main.go` file to generate the documentation.
