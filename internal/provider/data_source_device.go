package provider

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	syncthing "github.com/syncthing/syncthing/lib/protocol"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DeviceDataSource{}

func NewDeviceDataSource() datasource.DataSource {
	return &DeviceDataSource{}
}

// DeviceDataSource defines the data source implementation.
type DeviceDataSource struct{}

// DeviceDataSourceModel describes the data source data model.
type DeviceDataSourceModel struct {
	CertPEM       types.String `tfsdk:"cert_pem"`
	PrivateKeyPEM types.String `tfsdk:"private_key_pem"`
	ID            types.String `tfsdk:"id"`
}

func (d *DeviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (d *DeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Syncthing device ID data source",

		Attributes: map[string]schema.Attribute{
			"cert_pem": schema.StringAttribute{
				MarkdownDescription: "PEM-encoded certificate",
				Required:            true,
			},
			"private_key_pem": schema.StringAttribute{
				MarkdownDescription: "PEM formatted string to use as the private key",
				Required:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Device ID",
				Computed:            true,
			},
		},
	}
}

func (d *DeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
}

func (d *DeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DeviceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := tls.X509KeyPair([]byte(data.CertPEM.ValueString()), []byte(data.PrivateKeyPEM.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device ID",
			fmt.Sprintf("Failed to load keypair: %s", err.Error()),
		)
		return
	}

	data.ID = types.StringValue(syncthing.NewDeviceID(cert.Certificate[0]).String())

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
