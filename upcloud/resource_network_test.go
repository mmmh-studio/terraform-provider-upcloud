package upcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mmmh-studio/upcloud-go"
)

const testNetworkName = "terraform-provider-upcloud-network-create"

var testNetworkCreateConfig = fmt.Sprintf(`
resource "upcloud_network" "test-network-create" {
	name               = "%s"
	zone               = "de-fra1"
	family             = "IPv4"
	address            = "172.16.0.0/22"
	dhcp               = true
	dhcp_default_route = false
}
`, testNetworkName)

func TestAccNetwork_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		CheckDestroy: testCheckNetworkDestroy,
		PreCheck:     func() { testPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testNetworkCreateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"upcloud_network.test-network-create",
						"dhcp",
						"true",
					),
				),
			},
		},
	})
}

func testCheckNetworkDestroy(s *terraform.State) error {
	service := testAccProvider.Meta().(*upcloud.Service)

	networks, err := service.ListNetworksInZone(upcloud.ListNetworksInZoneRequest{
		Zone: "de-fra1",
	})
	if err != nil {
		return err
	}

	exists := false

	for _, network := range networks {
		if network.Name == testNetworkName {
			exists = true
			break
		}
	}

	if exists {
		return fmt.Errorf("test network still exists")
	}

	return nil
}
