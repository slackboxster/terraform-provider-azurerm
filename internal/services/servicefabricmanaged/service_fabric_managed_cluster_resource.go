package servicefabricmanaged

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/sdk/2021-05-01/managedcluster"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type CustomFabricSetting struct {
	Parameter string `tfschema:"parameter"`
	Section   string `tfschema:"section"`
	Value     string `tfschema:"value"`
}

type LBRule struct {
	BackendPort      int64                        `tfschema:"backend_port"`
	FrontendPort     int64                        `tfschema:"frontend_port"`
	ProbeProtocol    managedcluster.ProbeProtocol `tfschema:"probe_protocol"`
	ProbeRequestPath string                       `tfschema:"probe_request_path"`
	Protocol         managedcluster.Protocol      `tfschema:"protocol"`
}

type Networking struct {
	ClientConnectionPort int64    `tfschema:"client_connection_port"`
	HTTPGatewayPort      int64    `tfschema:"http_gateway_port"`
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

type CertType string

const (
	CertTypeAdmin    CertType = "AdminClient"
	CertTypeReadOnly CertType = "ReadOnlyClient"
)

type ClusterResource struct{}

var _ sdk.ResourceWithUpdate = ClusterResource{}

type ClusterResourceModel struct {
	BackupRestoreService bool   `tfschema:"backup_restore_service"`
	DNSService           bool   `tfschema:"dns_service"`
	Location             string `tfschema:"location"`
	Name                 string `tfschema:"name"`
	Username             string `tfschema:"username"`
	Password             string `tfschema:"password"`
	ResourceGroup        string `tfschema:"resource_group_name"`

	Authentication       Authentication                       `tfschema:"authentication"`
	CustomFabricSettings []CustomFabricSetting                `tfschema:"custom_fabric_settings"`
	Networking           Networking                           `tfschema:"networking"`
	NodeTypes            []NodeType                           `tfschema:"node_type"`
	Sku                  managedcluster.SkuName               `tfschema:"sku"`
	Tags                 map[string]interface{}               `tfschema:"tags"`
	UpgradeWave          managedcluster.ClusterUpgradeCadence `tfschema:"upgrade_wave"`
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
		"location": azure.SchemaLocation(),
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
		"resource_group_name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotWhiteSpace,
		},
		"node_type": {
			Type:     pluginsdk.TypeSet,
			Optional: true,
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
										string(CertTypeAdmin),
										string(CertTypeReadOnly),
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
			Required: true,
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
										string(managedcluster.ProbeProtocolHttp),
										string(managedcluster.ProbeProtocolHttps),
										string(managedcluster.ProbeProtocolTcp),
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
										string(managedcluster.ProtocolTcp),
										string(managedcluster.ProtocolUdp),
									}, false),
								},
							},
						},
					},
				},
			},
		},
		"sku": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Default:  "free",
			ValidateFunc: validation.StringInSlice([]string{
				string(managedcluster.SkuNameBasic),
				string(managedcluster.SkuNameStandard),
			}, false),
		},
		"tags": tags.Schema(),
		"upgrade_wave": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Default:  managedcluster.ClusterUpgradeCadenceWaveZero,
			ValidateFunc: validation.StringInSlice([]string{
				string(managedcluster.ClusterUpgradeCadenceWaveZero),
				string(managedcluster.ClusterUpgradeCadenceWaveOne),
				string(managedcluster.ClusterUpgradeCadenceWaveTwo)}, false),
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
			var model ClusterResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding %+v", err)
			}
			ctx, cancel := timeouts.ForCreate(metadata.Client.StopContext, metadata.ResourceData)
			defer cancel()

			clusterClient := metadata.Client.ServiceFabricManaged.ManagedClusterClient
			subscriptionId := metadata.Client.Account.SubscriptionId

			managedClusterId := managedcluster.NewManagedClusterID(subscriptionId, model.ResourceGroup, model.Name)

			cluster := managedcluster.ManagedCluster{
				Location:   model.Location,
				Name:       utils.String(model.Name),
				Properties: expandClusterProperties(&model),
				Sku:        &managedcluster.Sku{Name: model.Sku},
				//Tags: tags.Expand(model.Tags),
			}

			err := clusterClient.CreateOrUpdateThenPoll(ctx, managedClusterId, cluster)
			if err != nil {
				return fmt.Errorf("while creating cluster %q: %+v", model.Name, err)
			}

			metadata.SetID(managedClusterId)

			return nil
		},
		Timeout: 30 * time.Minute,
	}
}

func (k ClusterResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			resourceId, err := managedcluster.ParseManagedClusterID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("while parsing resourceID: %+v", err)
			}
			clusterClient := metadata.Client.ServiceFabricManaged.ManagedClusterClient
			cluster, err := clusterClient.Get(ctx, *resourceId)

			// Most of the resource model's data lives in the Properties
			model := flattenClusterProperties(cluster.Model.Properties)

			// fill in the rest
			model.Name = utils.NormalizeNilableString(cluster.Model.Name)
			model.Location = cluster.Model.Location
			model.ResourceGroup = resourceId.ResourceGroup

			if sku := cluster.Model.Sku; sku != nil {
				model.Sku = sku.Name
			}

			return metadata.Encode(&model)
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

func flattenClusterProperties(properties *managedcluster.ManagedClusterProperties) *ClusterResourceModel {
	model := &ClusterResourceModel{}

	if features := properties.AddonFeatures; features != nil {
		for _, feature := range *features {
			if feature == managedcluster.AddonFeaturesDnsService {
				model.DNSService = true
			} else if feature == managedcluster.AddonFeaturesBackupRestoreService {
				model.BackupRestoreService = true
			}
		}
	}
	model.Username = properties.AdminUserName
	model.Password = utils.NormalizeNilableString(properties.AdminPassword)

	if aad := properties.AzureActiveDirectory; aad != nil {
		adModel := model.Authentication.ADAuth
		adModel.ClientApp = utils.NormalizeNilableString(aad.ClientApplication)
		adModel.ClusterApp = utils.NormalizeNilableString(aad.ClusterApplication)
		adModel.TenantId = utils.NormalizeNilableString(aad.TenantId)
	}

	if clients := properties.Clients; clients != nil {
		certs := make([]ThumbprintAuth, len(*clients))
		for idx, client := range *clients {
			t := CertTypeReadOnly
			if client.IsAdmin {
				t = CertTypeAdmin
			}
			certs[idx] = ThumbprintAuth{
				CertificateType: t,
				CommonName:      utils.NormalizeNilableString(client.CommonName),
				Thumbprint:      utils.NormalizeNilableString(client.Thumbprint),
			}
		}
		model.Authentication.CertAuthentication = certs
	}

	if fss := properties.FabricSettings; fss != nil {
		cfs := make([]CustomFabricSetting, 0)
		for _, fs := range *fss {
			for _, param := range fs.Parameters {
				cfs = append(cfs, CustomFabricSetting{
					Parameter: fs.Name,
					Section:   param.Name,
					Value:     param.Value,
				})
			}
		}
	}

	model.Networking.ClientConnectionPort = utils.NormaliseNilableInt64(properties.ClientConnectionPort)
	model.Networking.HTTPGatewayPort = utils.NormaliseNilableInt64(properties.HttpGatewayConnectionPort)

	if lbrules := properties.LoadBalancingRules; lbrules != nil {
		model.Networking.LBRules = make([]LBRule, len(*lbrules))
		for idx, rule := range *lbrules {
			model.Networking.LBRules[idx] = LBRule{
				BackendPort:      rule.BackendPort,
				FrontendPort:     rule.FrontendPort,
				ProbeProtocol:    rule.ProbeProtocol,
				ProbeRequestPath: utils.NormalizeNilableString(rule.ProbeRequestPath),
				Protocol:         rule.Protocol,
			}
		}
	}

	return nil
}

func expandClusterProperties(model *ClusterResourceModel) *managedcluster.ManagedClusterProperties {
	out := &managedcluster.ManagedClusterProperties{}

	addons := make([]managedcluster.AddonFeatures, 0)
	if model.DNSService {
		addons = append(addons, managedcluster.AddonFeaturesDnsService)
	}
	if model.BackupRestoreService {
		addons = append(addons, managedcluster.AddonFeaturesBackupRestoreService)
	}
	out.AddonFeatures = &addons

	out.AdminPassword = utils.String(model.Password)
	out.AdminUserName = model.Username

	adAuth := model.Authentication.ADAuth
	if adAuth.ClientApp != "" && adAuth.ClusterApp != "" && adAuth.TenantId != "" {
		out.AzureActiveDirectory = &managedcluster.AzureActiveDirectory{
			ClientApplication:  utils.String(adAuth.ClientApp),
			ClusterApplication: utils.String(adAuth.ClusterApp),
			TenantId:           utils.String(adAuth.TenantId),
		}
	}

	out.ClientConnectionPort = &model.Networking.ClientConnectionPort

	if certs := model.Authentication.CertAuthentication; len(certs) > 0 {
		clients := make([]managedcluster.ClientCertificate, len(certs))
		for idx, cert := range certs {
			clients[idx] = managedcluster.ClientCertificate{
				CommonName: utils.String(cert.CommonName),
				IsAdmin:    cert.CertificateType == CertTypeAdmin,
				Thumbprint: utils.String(cert.Thumbprint),
			}
		}
		out.Clients = &clients
	}

	out.ClusterUpgradeCadence = &model.UpgradeWave

	if customSettings := model.CustomFabricSettings; len(customSettings) > 0 {
		fs := make([]managedcluster.SettingsSectionDescription, len(customSettings))

		// First we build a map of all settings per section
		fsMap := make(map[string][]managedcluster.SettingsParameterDescription)
		for _, cs := range customSettings {
			spd := managedcluster.SettingsParameterDescription{
				Name:  cs.Parameter,
				Value: cs.Value,
			}
			if v, ok := fsMap[cs.Section]; ok {
				v = append(v, spd)
			} else {
				fsMap[cs.Section] = []managedcluster.SettingsParameterDescription{spd}
			}
		}

		// Then we update the properties struct
		for k, v := range fsMap {
			fs = append(fs, managedcluster.SettingsSectionDescription{
				Name:       k,
				Parameters: v,
			})
		}
		out.FabricSettings = &fs
	}

	out.HttpGatewayConnectionPort = &model.Networking.HTTPGatewayPort

	if rules := model.Networking.LBRules; len(rules) > 0 {
		lbRules := make([]managedcluster.LoadBalancingRule, 0)
		for idx, rule := range lbRules {
			lbRules[idx] = managedcluster.LoadBalancingRule{
				BackendPort:      rule.BackendPort,
				FrontendPort:     rule.FrontendPort,
				ProbePort:        rule.ProbePort,
				ProbeProtocol:    rule.ProbeProtocol,
				ProbeRequestPath: rule.ProbeRequestPath,
				Protocol:         rule.Protocol,
			}
		}
		out.LoadBalancingRules = &lbRules
	}

	return out
}
