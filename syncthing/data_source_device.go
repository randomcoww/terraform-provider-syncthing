package syncthingdevice

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceDevice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDeviceRead,
		Schema: map[string]*schema.Schema{
			"cert_pem": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "PEM-encoded certificate",
			},
			"private_key_pem": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "PEM formatted string to use as the private key",
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDeviceRead(d *schema.ResourceData, meta interface{}) error {
	certPEM := d.Get("cert_pem").(string)
	privateKeyPEM := d.Get("private_key_pem").(string)

	return readCertPair(d, []byte(certPEM), []byte(privateKeyPEM))
}
