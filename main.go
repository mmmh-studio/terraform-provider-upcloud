package main // import "github.com/mmmh-studio/terraform-provider-upcloud"

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	"github.com/mmmh-studio/terraform-provider-upcloud/upcloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return upcloud.Provider()
		},
	})
}
