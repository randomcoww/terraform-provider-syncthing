package syncthing

import (
	 "crypto/tls"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/syncthing/syncthing/lib/protocol"
)

func readCertPair(d *schema.ResourceData, certPEM, privateKeyPEM []byte) error {
	cert, err := tls.X509KeyPair(certPEM, privateKeyPEM)
	if err != nil {
		return fmt.Errorf("Failed to load keypair: %s", err)
	}

	deviceID := protocol.NewDeviceID(cert.Certificate[0]).String()
	d.SetId(deviceID)
	d.Set("device_id", deviceID)
	return nil
}
