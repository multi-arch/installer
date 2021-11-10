package elasticsearch

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Result is a nested struct in elasticsearch response
type Result struct {
	Unit                      string                     `json:"unit" xml:"unit"`
	QueueType                 string                     `json:"queueType" xml:"queueType"`
	Endpoints                 string                     `json:"endpoints" xml:"endpoints"`
	IlmPolicy                 string                     `json:"ilmPolicy" xml:"ilmPolicy"`
	Priority                  int                        `json:"priority" xml:"priority"`
	Tags                      map[string]interface{}     `json:"tags" xml:"tags"`
	Domain                    string                     `json:"domain" xml:"domain"`
	GmtUpdateTime             string                     `json:"gmtUpdateTime" xml:"gmtUpdateTime"`
	BatchSize                 int                        `json:"batchSize" xml:"batchSize"`
	ResVersion                string                     `json:"resVersion" xml:"resVersion"`
	PipelineId                string                     `json:"pipelineId" xml:"pipelineId"`
	RegionId                  string                     `json:"regionId" xml:"regionId"`
	Env                       string                     `json:"env" xml:"env"`
	VpcId                     string                     `json:"vpcId" xml:"vpcId"`
	Connectable               bool                       `json:"connectable" xml:"connectable"`
	Dps                       map[string]interface{}     `json:"dps" xml:"dps"`
	Phases                    map[string]interface{}     `json:"phases" xml:"phases"`
	PublicPort                int                        `json:"publicPort" xml:"publicPort"`
	Enable                    bool                       `json:"enable" xml:"enable"`
	State                     string                     `json:"state" xml:"state"`
	Workers                   int                        `json:"workers" xml:"workers"`
	CreateUrl                 string                     `json:"createUrl" xml:"createUrl"`
	PipelineManagementType    string                     `json:"pipelineManagementType" xml:"pipelineManagementType"`
	ReportId                  string                     `json:"reportId" xml:"reportId"`
	Metric                    string                     `json:"metric" xml:"metric"`
	Summary                   float64                    `json:"summary" xml:"summary"`
	DryRun                    bool                       `json:"dryRun" xml:"dryRun"`
	Type                      string                     `json:"type" xml:"type"`
	PipelineStatus            string                     `json:"pipelineStatus" xml:"pipelineStatus"`
	NodeAmount                int                        `json:"nodeAmount" xml:"nodeAmount"`
	DataStream                bool                       `json:"dataStream" xml:"dataStream"`
	InstanceId                string                     `json:"instanceId" xml:"instanceId"`
	Description               string                     `json:"description" xml:"description"`
	Trigger                   string                     `json:"trigger" xml:"trigger"`
	EsInstanceId              string                     `json:"esInstanceId" xml:"esInstanceId"`
	Scene                     string                     `json:"scene" xml:"scene"`
	CreatedAt                 string                     `json:"createdAt" xml:"createdAt"`
	GmtCreatedTime            string                     `json:"gmtCreatedTime" xml:"gmtCreatedTime"`
	ValidateType              string                     `json:"validateType" xml:"validateType"`
	KibanaPort                int                        `json:"kibanaPort" xml:"kibanaPort"`
	Health                    string                     `json:"health" xml:"health"`
	UserName                  string                     `json:"userName" xml:"userName"`
	ClusterType               string                     `json:"clusterType" xml:"clusterType"`
	ResId                     string                     `json:"resId" xml:"resId"`
	PublicDomain              string                     `json:"publicDomain" xml:"publicDomain"`
	KibanaDomain              string                     `json:"kibanaDomain" xml:"kibanaDomain"`
	QueueCheckPointWrites     int                        `json:"queueCheckPointWrites" xml:"queueCheckPointWrites"`
	Name                      string                     `json:"name" xml:"name"`
	MessageWatermark          int64                      `json:"messageWatermark" xml:"messageWatermark"`
	QuartzRegex               string                     `json:"QuartzRegex" xml:"QuartzRegex"`
	ResType                   string                     `json:"resType" xml:"resType"`
	Integrity                 float64                    `json:"integrity" xml:"integrity"`
	EnablePublic              bool                       `json:"enablePublic" xml:"enablePublic"`
	EsVersion                 string                     `json:"esVersion" xml:"esVersion"`
	ClusterId                 string                     `json:"clusterId" xml:"clusterId"`
	FileSize                  int64                      `json:"fileSize" xml:"fileSize"`
	Config                    string                     `json:"config" xml:"config"`
	Value                     int64                      `json:"value" xml:"value"`
	UpdateTime                int64                      `json:"updateTime" xml:"updateTime"`
	IndexTemplate             string                     `json:"indexTemplate" xml:"indexTemplate"`
	PaymentType               string                     `json:"paymentType" xml:"paymentType"`
	UpdatedAt                 string                     `json:"updatedAt" xml:"updatedAt"`
	OwnerId                   string                     `json:"ownerId" xml:"ownerId"`
	QueueMaxBytes             int                        `json:"queueMaxBytes" xml:"queueMaxBytes"`
	Status                    string                     `json:"status" xml:"status"`
	BatchDelay                int                        `json:"batchDelay" xml:"batchDelay"`
	Version                   string                     `json:"version" xml:"version"`
	CreateTime                int64                      `json:"createTime" xml:"createTime"`
	MasterSpec                []string                   `json:"masterSpec" xml:"masterSpec"`
	CollectorPaths            []string                   `json:"collectorPaths" xml:"collectorPaths"`
	PrivateNetworkIpWhiteList []string                   `json:"privateNetworkIpWhiteList" xml:"privateNetworkIpWhiteList"`
	EsVersions                []string                   `json:"esVersions" xml:"esVersions"`
	EsIPWhitelist             []string                   `json:"esIPWhitelist" xml:"esIPWhitelist"`
	PublicIpWhitelist         []string                   `json:"publicIpWhitelist" xml:"publicIpWhitelist"`
	Zones                     []string                   `json:"zones" xml:"zones"`
	EsIPBlacklist             []string                   `json:"esIPBlacklist" xml:"esIPBlacklist"`
	IndexPatterns             []string                   `json:"indexPatterns" xml:"indexPatterns"`
	ClientNodeSpec            []string                   `json:"clientNodeSpec" xml:"clientNodeSpec"`
	InstanceSupportNodes      []string                   `json:"instanceSupportNodes" xml:"instanceSupportNodes"`
	KibanaIPWhitelist         []string                   `json:"kibanaIPWhitelist" xml:"kibanaIPWhitelist"`
	PipelineIds               []string                   `json:"pipelineIds" xml:"pipelineIds"`
	Node                      Node                       `json:"node" xml:"node"`
	ElasticExpansionTask      ElasticExpansionTask       `json:"elasticExpansionTask" xml:"elasticExpansionTask"`
	NodeSpec                  NodeSpec                   `json:"nodeSpec" xml:"nodeSpec"`
	MasterConfiguration       MasterConfiguration        `json:"masterConfiguration" xml:"masterConfiguration"`
	JvmConfine                JvmConfine                 `json:"jvmConfine" xml:"jvmConfine"`
	MetaInfo                  MetaInfo                   `json:"metaInfo" xml:"metaInfo"`
	ElasticShrinkTask         ElasticShrinkTask          `json:"elasticShrinkTask" xml:"elasticShrinkTask"`
	NetworkConfig             NetworkConfig              `json:"networkConfig" xml:"networkConfig"`
	KibanaConfiguration       KibanaConfiguration        `json:"kibanaConfiguration" xml:"kibanaConfiguration"`
	KibanaNodeProperties      KibanaNodeProperties       `json:"kibanaNodeProperties" xml:"kibanaNodeProperties"`
	ClientNodeAmountRange     ClientNodeAmountRange      `json:"clientNodeAmountRange" xml:"clientNodeAmountRange"`
	Template                  Template                   `json:"template" xml:"template"`
	ValidateResult            ValidateResult             `json:"validateResult" xml:"validateResult"`
	OssObject                 OssObject                  `json:"ossObject" xml:"ossObject"`
	ElasticNodeProperties     ElasticNodeProperties      `json:"elasticNodeProperties" xml:"elasticNodeProperties"`
	WarmNodeProperties        WarmNodeProperties         `json:"warmNodeProperties" xml:"warmNodeProperties"`
	Configs                   []ConfigsItem              `json:"configs" xml:"configs"`
	MasterDiskList            []Disk                     `json:"masterDiskList" xml:"masterDiskList"`
	SynonymsDicts             []SynonymsDictsItem        `json:"synonymsDicts" xml:"synonymsDicts"`
	ClientNodeDiskList        []Disk                     `json:"clientNodeDiskList" xml:"clientNodeDiskList"`
	ExtendConfigs             []ExtendConfigsItem        `json:"extendConfigs" xml:"extendConfigs"`
	DictList                  []DictListItem             `json:"dictList" xml:"dictList"`
	SupportVersions           []CategoryEntity           `json:"supportVersions" xml:"supportVersions"`
	EsVersionsLatestList      []EsVersionsLatestListItem `json:"esVersionsLatestList" xml:"esVersionsLatestList"`
	NodeSpecList              []NodeSpecListItem         `json:"nodeSpecList" xml:"nodeSpecList"`
	DataDiskList              []DataDiskListItem         `json:"dataDiskList" xml:"dataDiskList"`
	DiagnoseItems             []DiagnoseItemsItem        `json:"diagnoseItems" xml:"diagnoseItems"`
}