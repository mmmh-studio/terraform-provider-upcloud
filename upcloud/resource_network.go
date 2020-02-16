package upcloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	upcloud "github.com/mmmh-studio/upcloud-go"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Delete: resourceNetworkDelete,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "Unique network identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Names the network.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone in which the network is configured.",
			},
			"family": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IP family of the network.",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Sets address space for the network.",
			},
			"dhcp": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Toggles DHCP service for the network.",
			},
			"dhcp_default_route": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Defines if the gateway should be given as default route by DHCP. Defaults to yes on public networks, and no on other ones.",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	service := meta.(*upcloud.Service)

	network, err := service.CreateNetwork(upcloud.CreateNetworkRequest{
		Name: d.Get("name").(string),
		Zone: d.Get("zone").(string),
		IPNetworks: []upcloud.CreateIPNetwork{
			{
				Address: d.Get("address").(string),
				DHCP:    d.Get("dhcp").(bool),
				Family:  d.Get("family").(string),
			},
		},
	})
	if err != nil {
		return err
	}

	d.SetId(network.UUID)
	d.Set("name", network.Name)
	d.Set("zone", network.Zone)

	log.Printf("[INFO] Network %s with UUID %s created", network.Name, network.UUID)

	return resourceNetworkRead(d, meta)
}

func resourceNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	service := meta.(*upcloud.Service)

	return service.DeleteNetwork(upcloud.DeleteNetworkRequest{
		UUID: d.Id(),
	})
}

func resourceNetworkRead(d *schema.ResourceData, meta interface{}) error {
	service := meta.(*upcloud.Service)

	network, err := service.GetNetworkDetails(upcloud.GetNetworkDetailsRequest{
		UUID: d.Id(),
	})
	if err != nil {
		return err
	}

	d.Set("name", network.Name)
	d.Set("zone", network.Zone)

	return nil
}

func resourceNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("resourceNetworkUpdate() not implemented")
}
