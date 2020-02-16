package upcloud

import (
	"github.com/hashicorp/terraform/helper/schema"

	upcloud "github.com/mmmh-studio/upcloud-go"
	"github.com/mmmh-studio/upcloud-go/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: configure,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPCLOUD_USERNAME", nil),
				Description: "Upcloud API username",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPCLOUD_PASSWORD", nil),
				Description: "Upcloud API password",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"upcloud_network": resourceNetwork(),
			"upcloud_server":  resourceServer(),
		},
	}
}

type Config struct {
	Username string
	Password string
}

func configure(d *schema.ResourceData) (interface{}, error) {
	var (
		config = Config{
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
		}
		client  = client.New(config.Username, config.Password)
		service = upcloud.NewService(client)
	)

	_, err := service.GetAccount()
	if err != nil {
		return nil, err
	}

	return service, nil
}
