package responses

type WhereUsedList struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
