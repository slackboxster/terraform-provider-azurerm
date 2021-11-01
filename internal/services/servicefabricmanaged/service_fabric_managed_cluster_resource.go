package servicefabricmanaged

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"

	"github.com/hashicorp/go-azure-helpers/response"

	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/sdk/2021-05-01/nodetype"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/sdk/2021-05-01/managedcluster"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
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
	ADAuth             ADAuthentication `tfschema:"active_directory"`
	CertAuthentication []ThumbprintAuth `tfschema:"certificate"`
}

type PortRange struct {
	From int64 `tfschema:"from"`
	To   int64 `tfschema:"to"`
}

type VaultCertificates struct {
	Store string `tfschema:"store"`
	Url   string `tfschema:"url"`
}

type VmSecrets struct {
	SourceVault  string              `tfschema:"vault_id"`
	Certificates []VaultCertificates `tfschema:"certificate"`
}

type NodeType struct {
	DataDiskSize            int64  `tfschema:"data_disk_size"`
	MultiplePlacementGroups bool   `tfschema:"multiple_placement_groups"`
	Name                    string `tfschema:"name"`
	Primary                 bool   `tfschema:"primary"`
	Stateless               bool   `tfschema:"stateless"`
	VmImageOffer            string `tfschema:"vm_image_offer"`
	VmImagePublisher        string `tfschema:"vm_image_publisher"`
	VmImageSku              string `tfschema:"vm_image_sku"`
	VmImageVersion          string `tfschema:"vm_image_version"`
	VmInstanceCount         int64  `tfschema:"vm_instance_count"`
	VmSize                  string `tfschema:"vm_size"`

	ApplicationPorts    string            `tfschema:"application_port_range"`
	Capacities          map[string]string `tfschema:"capacities"`
	DataDiskType        nodetype.DiskType `tfschema:"data_disk_type"`
	EphemeralPorts      string            `tfschema:"ephemeral_port_range"`
	PlacementProperties map[string]string `tfschema:"placement_properties"`
	VmSecrets           []VmSecrets       `tfschema:"vm_secrets"`
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
	ClientConnectionPort int64  `tfschema:"client_connection_port"`
	DNSService           bool   `tfschema:"dns_service"`
	HTTPGatewayPort      int64  `tfschema:"http_gateway_port"`
	Location             string `tfschema:"location"`
	Name                 string `tfschema:"name"`
	Username             string `tfschema:"username"`
	Password             string `tfschema:"password"`
	ResourceGroup        string `tfschema:"resource_group_name"`

	Authentication       []Authentication                     `tfschema:"authentication"`
	CustomFabricSettings []CustomFabricSetting                `tfschema:"custom_fabric_setting"`
	LBRules              []LBRule                             `tfschema:"lb_rules"`
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
					"data_disk_size": {
						Type:     pluginsdk.TypeInt,
						Required: true,
					},
					"multiple_placement_groups": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"name": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"primary": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"stateless": {
						Type:     pluginsdk.TypeBool,
						Optional: true,
					},
					"vm_image_offer": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"vm_image_publisher": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"vm_image_sku": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"vm_image_version": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"vm_instance_count": {
						Type:         pluginsdk.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(3, 100),
					},
					"vm_size": {
						Type:     pluginsdk.TypeString,
						Required: true,
					},
					"application_port_range": {
						Type:     pluginsdk.TypeString,
						Required: true,
						//TODO: Add validation
					},
					"capacities": {
						Type:     pluginsdk.TypeMap,
						Optional: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},
					"data_disk_type": {
						Type:     pluginsdk.TypeString,
						Optional: true,
						Default:  string(nodetype.DiskTypeStandardLRS),
						ValidateFunc: validation.StringInSlice([]string{
							string(nodetype.DiskTypeStandardLRS),
							string(nodetype.DiskTypeStandardSSDLRS),
							string(nodetype.DiskTypePremiumLRS),
						}, false),
					},
					"ephemeral_port_range": {
						Type:     pluginsdk.TypeString,
						Required: true,
						//TODO: Add validation
					},
					"placement_properties": {
						Type:     pluginsdk.TypeMap,
						Optional: true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},
					"vm_secrets": {
						Type:     pluginsdk.TypeSet,
						Optional: true,
						Elem: &pluginsdk.Resource{
							Schema: map[string]*pluginsdk.Schema{
								"vault_id": {
									Type:     pluginsdk.TypeString,
									Required: true,
								},
								"certificates": {
									Type:     pluginsdk.TypeSet,
									Required: true,
									Elem: &pluginsdk.Resource{
										Schema: map[string]*pluginsdk.Schema{
											"store": {
												Type:     pluginsdk.TypeString,
												Required: true,
											},
											"url": {
												Type:     pluginsdk.TypeString,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"authentication": {
			Type:     pluginsdk.TypeList,
			Optional: true,
			MaxItems: 1,
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
			return createOrUpdate(ctx, metadata)
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
			nodeTypeClient := metadata.Client.ServiceFabricManaged.NodeTypeClient

			cluster, err := clusterClient.Get(ctx, *resourceId)
			if err != nil {
				if response.WasNotFound(cluster.HttpResponse) {
					return metadata.MarkAsGone(resourceId)
				}
				return fmt.Errorf("while reading data for cluster %q: %+v", resourceId.Name, err)
			}

			nts, err := nodeTypeClient.ListByManagedClustersComplete(ctx, nodetype.ManagedClusterId{
				SubscriptionId: resourceId.SubscriptionId,
				ResourceGroup:  resourceId.ResourceGroup,
				Name:           resourceId.Name,
			})
			if err != nil {
				return fmt.Errorf("while listing NodeTypes for cluster %q: +%v", resourceId.Name, err)
			}

			model := flattenClusterProperties(cluster.Model)
			// Password is read-only
			model.Password = metadata.ResourceData.Get("password").(string)
			model.ResourceGroup = resourceId.ResourceGroup
			model.NodeTypes = make([]NodeType, len(nts.Items))
			for idx, nt := range nts.Items {
				model.NodeTypes[idx] = flattenNodetypeProperties(nt)
			}
			return metadata.Encode(model)
		},
		Timeout: 5 * time.Minute,
	}
}

func (k ClusterResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return createOrUpdate(ctx, metadata)
		},
		Timeout: 30 * time.Minute,
	}
}

func (k ClusterResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {

			resourceId, err := managedcluster.ParseManagedClusterID(metadata.ResourceData.Id())
			if err != nil {
				return fmt.Errorf("while parsing resourceID: %+v", err)
			}
			clusterClient := metadata.Client.ServiceFabricManaged.ManagedClusterClient

			err = clusterClient.DeleteThenPoll(ctx, *resourceId)
			if err != nil {
				return fmt.Errorf("while deleting cluster %q: %+v", resourceId.String(), err)
			}
			return nil
		},
		Timeout: 30 * time.Minute,
	}
}

func (k ClusterResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validate.ServiceFabricManagedClusterID
}

func createOrUpdate(ctx context.Context, metadata sdk.ResourceMetaData) error {
	var model ClusterResourceModel
	if err := metadata.Decode(&model); err != nil {
		return fmt.Errorf("decoding %+v", err)
	}
	ctx, cancel := timeouts.ForCreate(metadata.Client.StopContext, metadata.ResourceData)
	defer cancel()

	clusterClient := metadata.Client.ServiceFabricManaged.ManagedClusterClient
	nodeTypeClient := metadata.Client.ServiceFabricManaged.NodeTypeClient

	subscriptionId := metadata.Client.Account.SubscriptionId

	managedClusterId := managedcluster.NewManagedClusterID(subscriptionId, model.ResourceGroup, model.Name)
	cluster := managedcluster.ManagedCluster{
		Location:   model.Location,
		Name:       utils.String(model.Name),
		Properties: expandClusterProperties(&model),
		Sku:        &managedcluster.Sku{Name: model.Sku},
		//Tags:       tags.Expand(model.Tags),
		// TODO: Fix tags???
	}
	resp, err := clusterClient.CreateOrUpdate(ctx, managedClusterId, cluster)
	if err != nil {
		return fmt.Errorf("while creating cluster %q: %+v", model.Name, err)
	}
	// Wait for the cluster creation operation to be completed
	err = resp.Poller.PollUntilDone()
	if err != nil {
		return fmt.Errorf("while waiting for cluster %q to get created: : %+v", model.Name, err)
	}

	// Send all Create NodeType requests, and store all responses to a list.
	for _, nt := range model.NodeTypes {
		nodeTypeProperties, err := expandNodeTypeProperties(&nt)
		if err != nil {
			return fmt.Errorf("while expanding node type %q: %+v", nt.Name, err)
		}
		nodeTypeId := nodetype.NewNodeTypeID(subscriptionId, model.ResourceGroup, model.Name, nt.Name)
		err = nodeTypeClient.CreateOrUpdateThenPoll(ctx, nodeTypeId, nodetype.NodeType{
			Name:       nil,
			Properties: nodeTypeProperties,
		})
		if err != nil {
			return fmt.Errorf("while creating NodeType %q: %+v", nt.Name, err)
		}
	}
	metadata.SetID(managedClusterId)
	return nil
}

func flattenClusterProperties(cluster *managedcluster.ManagedCluster) *ClusterResourceModel {
	model := &ClusterResourceModel{}
	if cluster == nil {
		return model
	}

	model.Name = utils.NormalizeNilableString(cluster.Name)
	model.Location = cluster.Location
	if sku := cluster.Sku; sku != nil {
		model.Sku = sku.Name
	}

	properties := cluster.Properties
	if properties == nil {
		return model
	}

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

	if aad := properties.AzureActiveDirectory; aad != nil {

		adModel := ADAuthentication{}
		adModel.ClientApp = utils.NormalizeNilableString(aad.ClientApplication)
		adModel.ClusterApp = utils.NormalizeNilableString(aad.ClusterApplication)
		adModel.TenantId = utils.NormalizeNilableString(aad.TenantId)
		model.Authentication[0].ADAuth = adModel
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
		model.Authentication[0].CertAuthentication = certs
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

	model.ClientConnectionPort = utils.NormaliseNilableInt64(properties.ClientConnectionPort)
	model.HTTPGatewayPort = utils.NormaliseNilableInt64(properties.HttpGatewayConnectionPort)

	if lbrules := properties.LoadBalancingRules; lbrules != nil {
		model.LBRules = make([]LBRule, len(*lbrules))
		for idx, rule := range *lbrules {
			model.LBRules[idx] = LBRule{
				BackendPort:      rule.BackendPort,
				FrontendPort:     rule.FrontendPort,
				ProbeProtocol:    rule.ProbeProtocol,
				ProbeRequestPath: utils.NormalizeNilableString(rule.ProbeRequestPath),
				Protocol:         rule.Protocol,
			}
		}
	}

	if upgradeWave := properties.ClusterUpgradeCadence; upgradeWave != nil {
		model.UpgradeWave = *upgradeWave
	}

	return model
}

func flattenNodetypeProperties(nt nodetype.NodeType) NodeType {
	props := nt.Properties
	if props == nil {
		return NodeType{Name: utils.NormalizeNilableString(nt.Name)}
	}

	//from, to, err := parsePortRange(n)
	out := NodeType{
		DataDiskSize:     nt.Properties.DataDiskSizeGB,
		Name:             utils.NormalizeNilableString(nt.Name),
		Primary:          props.IsPrimary,
		VmImageOffer:     utils.NormalizeNilableString(props.VmImageOffer),
		VmImagePublisher: utils.NormalizeNilableString(props.VmImagePublisher),
		VmImageSku:       utils.NormalizeNilableString(props.VmImageSku),
		VmImageVersion:   utils.NormalizeNilableString(props.VmImageVersion),
		VmInstanceCount:  props.VmInstanceCount,
		VmSize:           utils.NormalizeNilableString(props.VmSize),
		ApplicationPorts: fmt.Sprintf("%d-%d", props.ApplicationPorts.StartPort, props.ApplicationPorts.EndPort),
		EphemeralPorts:   fmt.Sprintf("%d-%d", props.EphemeralPorts.StartPort, props.EphemeralPorts.EndPort),
	}

	if mpg := props.MultiplePlacementGroups; mpg != nil {
		out.MultiplePlacementGroups = *mpg
	}

	if stateless := props.IsStateless; stateless != nil {
		out.Stateless = *stateless
	}

	if capacities := props.Capacities; capacities != nil {
		caps := make(map[string]string)
		for k, v := range *capacities {
			caps[k] = v
		}
		out.Capacities = caps
	}

	if diskType := props.DataDiskType; diskType != nil {
		out.DataDiskType = *diskType
	}

	if placementProps := props.PlacementProperties; placementProps != nil {
		placements := make(map[string]string)
		for k, v := range *placementProps {
			placements[k] = v
		}
		out.PlacementProperties = placements
	}

	if secrets := props.VmSecrets; secrets != nil {
		secs := make([]VmSecrets, len(*secrets))
		for idx, sec := range *secrets {
			certs := make([]VaultCertificates, len(sec.VaultCertificates))
			for idx, cert := range sec.VaultCertificates {
				certs[idx] = VaultCertificates{
					Store: cert.CertificateStore,
					Url:   cert.CertificateUrl,
				}
			}
			secs[idx] = VmSecrets{
				SourceVault:  utils.NormalizeNilableString(sec.SourceVault.Id),
				Certificates: certs,
			}
		}
		out.VmSecrets = secs
	}
	return out
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

	if auth := model.Authentication; len(auth) > 0 {
		adAuth := auth[0].ADAuth
		if adAuth.ClientApp != "" && adAuth.ClusterApp != "" && adAuth.TenantId != "" {
			out.AzureActiveDirectory = &managedcluster.AzureActiveDirectory{
				ClientApplication:  utils.String(adAuth.ClientApp),
				ClusterApplication: utils.String(adAuth.ClusterApp),
				TenantId:           utils.String(adAuth.TenantId),
			}
		}
		if certs := auth[0].CertAuthentication; len(certs) > 0 {
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
	}

	out.ClientConnectionPort = &model.ClientConnectionPort
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

	out.HttpGatewayConnectionPort = &model.HTTPGatewayPort

	if rules := model.LBRules; len(rules) > 0 {
		lbRules := make([]managedcluster.LoadBalancingRule, len(rules))
		nsRules := make([]managedcluster.NetworkSecurityRule, len(rules))

		for idx, rule := range rules {
			lbRules[idx] = managedcluster.LoadBalancingRule{
				BackendPort:      rule.BackendPort,
				FrontendPort:     rule.FrontendPort,
				ProbeProtocol:    rule.ProbeProtocol,
				ProbeRequestPath: utils.String(rule.ProbeRequestPath),
				Protocol:         rule.Protocol,
			}

			fePortStr := strconv.FormatInt(rule.FrontendPort, 10)
			var sgProto managedcluster.NsgProtocol
			switch rule.Protocol {
			case managedcluster.ProtocolTcp:
				sgProto = managedcluster.NsgProtocolTcp
			case managedcluster.ProtocolUdp:
				sgProto = managedcluster.NsgProtocolUdp
			}
			nsRules[idx] = managedcluster.NetworkSecurityRule{
				Access:                     managedcluster.AccessAllow,
				SourceAddressPrefixes:      &[]string{"0.0.0.0/0"},
				SourcePortRanges:           &[]string{"1-65535"},
				DestinationPortRanges:      &[]string{fePortStr},
				DestinationAddressPrefixes: &[]string{"0.0.0.0/0"},
				Direction:                  managedcluster.DirectionInbound,
				Name:                       fmt.Sprintf("rule%d-allow-fe", rule.FrontendPort),
				Priority:                   1000,
				Protocol:                   sgProto,
			}
		}
		out.LoadBalancingRules = &lbRules
		out.NetworkSecurityRules = &nsRules

	}

	out.DnsName = model.Name

	return out
}

func expandNodeTypeProperties(nt *NodeType) (*nodetype.NodeTypeProperties, error) {
	vmSecrets := make([]nodetype.VaultSecretGroup, len(nt.VmSecrets))
	for idx, secret := range nt.VmSecrets {
		vcs := make([]nodetype.VaultCertificate, len(secret.Certificates))
		for cidx, cert := range secret.Certificates {
			vcs[cidx] = nodetype.VaultCertificate{
				CertificateStore: cert.Store,
				CertificateUrl:   cert.Url,
			}
		}
		vmSecrets[idx] = nodetype.VaultSecretGroup{
			SourceVault:       nodetype.SubResource{Id: &secret.SourceVault},
			VaultCertificates: vcs,
		}
	}

	appFrom, appTo, err := parsePortRange(nt.ApplicationPorts)
	if err != nil {
		return nil, fmt.Errorf("while parsing application port range (%q): %+v", nt.ApplicationPorts, err)
	}

	ephemeralFrom, ephemeralTo, err := parsePortRange(nt.EphemeralPorts)
	if err != nil {
		return nil, fmt.Errorf("while parsing ephemeral port range (%q): %+v", nt.EphemeralPorts, err)
	}
	nodeTypeProperties := &nodetype.NodeTypeProperties{
		ApplicationPorts: &nodetype.EndpointRangeDescription{
			EndPort:   appTo,
			StartPort: appFrom,
		},
		Capacities:     &nt.Capacities,
		DataDiskSizeGB: nt.DataDiskSize,
		DataDiskType:   &nt.DataDiskType,
		EphemeralPorts: &nodetype.EndpointRangeDescription{
			EndPort:   ephemeralTo,
			StartPort: ephemeralFrom,
		},
		IsPrimary:               nt.Primary,
		IsStateless:             &nt.Stateless,
		MultiplePlacementGroups: &nt.MultiplePlacementGroups,
		PlacementProperties:     &nt.PlacementProperties,
		VmImageOffer:            &nt.VmImageOffer,
		VmImagePublisher:        &nt.VmImagePublisher,
		VmImageSku:              &nt.VmImageSku,
		VmImageVersion:          &nt.VmImageVersion,
		VmInstanceCount:         nt.VmInstanceCount,
		VmSecrets:               &vmSecrets,
		VmSize:                  &nt.VmSize,
	}

	return nodeTypeProperties, nil
}

func parsePortRange(input string) (int64, int64, error) {
	toks := strings.Split(input, "-")
	if len(toks) != 2 {
		return 0, 0, fmt.Errorf("invalid port range format in string %q", input)
	}
	from, err := strconv.ParseInt(toks[0], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("while parsing %q as integer: %s", toks[0], err)
	}

	to, err := strconv.ParseInt(toks[1], 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("while parsing %q as integer: %s", toks[1], err)
	}
	return from, to, nil
}
