package sap_api_output_formatter

type WhereUsedListReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       bool   `json:"deleted"`
}

type WhereUsedList struct {
	BillOfMaterialItemUUID         string `json:"BillOfMaterialItemUUID"`
	BillOfMaterialComponent        string `json:"BillOfMaterialComponent"`
	BillOfMaterialItemNumber       string `json:"BillOfMaterialItemNumber"`
	HeaderChangeDocument           string `json:"HeaderChangeDocument"`
	BillOfMaterialCategory         string `json:"BillOfMaterialCategory"`
	BillOfMaterial                 string `json:"BillOfMaterial"`
	BillOfMaterialVariant          string `json:"BillOfMaterialVariant"`
	BillOfMaterialVersion          string `json:"BillOfMaterialVersion"`
	BillOfMaterialItemCategory     string `json:"BillOfMaterialItemCategory"`
	BillOfMaterialItemUnit         string `json:"BillOfMaterialItemUnit"`
	BillOfMaterialItemQuantity     string `json:"BillOfMaterialItemQuantity"`
	EngineeringChangeDocument      string `json:"EngineeringChangeDocument"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	BillOfMaterialItemNodeNumber   string `json:"BillOfMaterialItemNodeNumber"`
	BOMItemDescription             string `json:"BOMItemDescription"`
	Material                       string `json:"Material"`
	MaterialName                   string `json:"MaterialName"`
	PlantName                      string `json:"PlantName"`
	BillOfMaterialVariantUsageDesc string `json:"BillOfMaterialVariantUsageDesc"`
	Plant                          string `json:"Plant"`
	BillOfMaterialVariantUsage     string `json:"BillOfMaterialVariantUsage"`
	BOMVersionStatus               string `json:"BOMVersionStatus"`
	BOMVersionStatusDescription    string `json:"BOMVersionStatusDescription"`
}
