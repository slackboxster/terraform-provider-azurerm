package servicefabricmanaged

import (
	"context"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

type CustomFabricSetting struct {
	Parameter string `tfschema:"parameter"`
	Section   string `tfschema:"section"`
	Value     string `tfschema:"value"`
}

type LBRule struct {
	BackendPort      int          `tfschema:"backend_port"`
	FrontendPort     int          `tfschema:"frontend_port"`
	ProbeProtocol    LBProbeProto `tfschema:"probe_protocol"`
	ProbeRequestPath string       `tfschema:"ProbeRequestPath"`
	Protocol         LBProto      `tfschema:"protocol"`
}

type Networking struct {
	ClientConnectionPort int      `tfschema:"client_connection_port"`
	HTTPGatewayPort      int      `tfschema:"http_gateway_port"`
	LBRules              []LBRule `tfschema:"lb_rules"`
}

type ThumbprintAuth struct {
	CertificateType CertType `tfschema:"type"`
	CommonName      string   `tfschema:"common_name"`
	Thumbprint      string   `tfschema:"thumbprint"`
}

type ADAuthentication struct {
	ClientApp  string `tfschema:"client_application_id"`
	ClusterApp string `tfschema:"cluster_application_id"`
	TenantId   string `tfschema:"tenant_id"`
}

type Authentication struct {
	ADAuth             ADAuthentication `tfscherma:"active_directory"`
	CertAuthentication []ThumbprintAuth `tfschema:"certificate"`
}

type NodeType struct {
	DataDiskSize int    `tfschema:"data_disk_size"`
	NodeCount    int    `tfschema:"node_count"`
	NodeSize     string `tfschema:"node_size"`
	OSImage      string `tfschema:"os_image"`
}
type UpgradeWave string

const (
	UpgradeWave0 UpgradeWave = "Wave0"
	UpgradeWave1 UpgradeWave = "Wave1"
	UpgradeWave2 UpgradeWave = "Wave2"
)

type LBProto string

const (
	LBProtoTCP LBProto = "TCP"
	LBProtoUDP LBProto = "UDP"
)

type LBProbeProto string

const (
	LBProbeProtoHTTP  = "HTTP"
	LBProbeProtoHTTPS = "HTTPS"
	LBProbeProtoTCP   = "TCP"
)

type CertType string

const (
	Admin    CertType = "AdminClient"
	ReadOnly CertType = "ReadOnlyClient"
)

type ClusterResource struct{}

var _ sdk.ResourceWithUpdate = ClusterResource{}

type ClusterResourceModel struct {
	BackupRestoreService bool   `tfschema:"backup_restore_service"`
	DNSService           bool   `tfschema:"dns_service"`
	Name                 string `tfschema:"name"`
	Username             string `tfschema:"username"`
	Password             string `tfschema:"password"`
	ResourceGroup        string `tfschema:"resource_group"`

	Authentication       Authentication         `tfschema:"authentication"`
	CustomFabricSettings []CustomFabricSetting  `tfschema:"custom_fabric_settings"`
	Networking           Networking             `tfschema:"networking"`
	NodeTypes            []NodeType             `tfschema:"node_type"`
	Tags                 map[string]interface{} `tfschema:"tags"`
	UpgradeWave          UpgradeWave            `tfschema:"upgrade_wave"`
}

func (k ClusterResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"backup_restore_service": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},
		"dns_service": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
		},
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ValidateFunc: validation.All(
				validation.StringLenBetween(4, 23),
				validation.StringMatch(regexp.MustCompile(`^[a-z0-9]+(-*[a-z0-9])*$`), "The name of the cluster must have lowercase letters, numbers and hyphens. The first character must be a letter and the last character a letter or number")),
		},
		"username": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ValidateFunc: validation.All(
				validation.StringLenBetween(1, 15),
				validation.StringMatch(regexp.MustCompile("^[^\\\\/\"\\[\\]:|<>+=;,?*$]{1,14}$"), "User names cannot contain special characters \\/\"\"[]:|<>+=;,$?*@")),
		},
		"password": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ValidateFunc: validation.All(
				validation.StringLenBetween(8, 123),
				validation.StringIsNotWhiteSpace),
		},
		"resource_group": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotWhiteSpace,
		},
		"node_type": {
			Type:     pluginsdk.TypeSet,
			Required: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"os_image": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"node_count": {
						Type:         pluginsdk.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(3, 100),
					},
					"node_size": {
						Type:     pluginsdk.TypeString,
						Required: true,
						// TODO: validate
					},
					"data_disk_size": {
						Type:     pluginsdk.TypeInt,
						Required: true,
						// TODO: validate
						// TODO: DataDiskSize is more complex than a string (maybe?)

					},
				},
			},
		},
		"authentication": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"active_directory": {
						Type:     pluginsdk.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"client_application_id": {
									Type:         pluginsdk.TypeString,
									Required:     true,
									ValidateFunc: validation.IsUUID,
								},
								"cluster_application_id": {
									Type:         pluginsdk.TypeString,
									Required:     true,
									ValidateFunc: validation.IsUUID,
								},
								"tenant_id": {
									Type:         pluginsdk.TypeString,
									Required:     true,
									ValidateFunc: validation.IsUUID,
								},
							},
						},
					},
					"certificate": {
						Type:     pluginsdk.TypeSet,
						Optional: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"common_name": {
									Type:         pluginsdk.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringIsNotWhiteSpace,
								},
								"thumbprint": {
									Type:         pluginsdk.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotWhiteSpace,
								},
								"type": {
									Type:     pluginsdk.TypeString,
									Required: true,
									ValidateFunc: validation.StringInSlice([]string{
										string(Admin),
										string(ReadOnly),
									}, false),
								},
							},
						},
					},
				},
			},
		},
		"custom_fabric_setting": {
			Type:     pluginsdk.TypeSet,
			Optional: true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"parameter": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotWhiteSpace,
					},
					"section": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotWhiteSpace,
					},
					"value": {
						Type:         pluginsdk.TypeString,
						Required:     true,
						ValidateFunc: validation.StringIsNotWhiteSpace,
					},
				},
			},
		},
		"networking": {
			Type:     pluginsdk.TypeList,
			MaxItems: 1,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"client_connection_port": {
						Type:         pluginsdk.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(1500, 65535),
					},
					"http_gateway_port": {
						Type:         pluginsdk.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(1500, 65535),
					},
					"lb_rules": {
						Type:     pluginsdk.TypeSet,
						Required: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"backend_port": {
									Type:         pluginsdk.TypeInt,
									Required:     true,
									ValidateFunc: validation.IntBetween(1, 65535),
								},
								"frontend_port": {
									Type:         pluginsdk.TypeInt,
									Required:     true,
									ValidateFunc: validation.IntBetween(1, 65535),
								},
								"probe_protocol": {
									Type:     pluginsdk.TypeString,
									Required: true,
									ValidateFunc: validation.StringInSlice([]string{
										string(LBProbeProtoHTTP),
										string(LBProbeProtoHTTPS),
										string(LBProbeProtoTCP),
									}, false),
								},
								"probe_request_path": {
									Type:         pluginsdk.TypeString,
									Required:     true,
									ValidateFunc: validation.StringIsNotWhiteSpace,
								},
								"protocol": {
									Type:     pluginsdk.TypeString,
									Required: true,
									ValidateFunc: validation.StringInSlice([]string{
										string(LBProtoTCP),
										string(LBProtoUDP),
									}, false),
								},
							},
						},
					},
				},
			},
		},
		"tags": tags.Schema(),
		"upgrade_wave": {
			Type:     pluginsdk.TypeString,
			Required: true,
			Default:  UpgradeWave0,
			ValidateFunc: validation.StringInSlice([]string{
				string(UpgradeWave0),
				string(UpgradeWave1),
				string(UpgradeWave2)}, false),
		},
	}
}

func (k ClusterResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (k ClusterResource) ModelObject() interface{} {
	return &ClusterResourceModel{}
}

func (k ClusterResource) ResourceType() string {
	return "azurerm_service_fabric_managed_cluster"
}

func (k ClusterResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {

			return nil
		},
		Timeout: 30 * time.Minute,
	}
}

func (k ClusterResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
		Timeout: 5 * time.Minute,
	}
}

func (k ClusterResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
		Timeout: 30 * time.Minute,
	}
}

func (k ClusterResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
		Timeout: 30 * time.Minute,
	}
}

func (k ClusterResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validate.ServiceFabricManagedClusterID
}

func flattenThing(d *pluginsdk.ResourceData) *ClusterResource {
	return nil
}
