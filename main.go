package main

import (
	"context"
	"terraform-provider-workos/workos"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name workos

func main() {
	providerserver.Serve(context.Background(), workos.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/vellanci/workos",
	})
}
