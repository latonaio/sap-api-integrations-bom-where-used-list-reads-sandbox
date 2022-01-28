package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bill-of-material-where-used-list-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToWhereUsedList(raw []byte, l *logger.Logger) ([]WhereUsedList, error) {
	pm := &responses.WhereUsedList{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to WhereUsedList. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	whereUsedList := make([]WhereUsedList, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		whereUsedList = append(whereUsedList, WhereUsedList{
	BillOfMaterialItemUUID:         data.BillOfMaterialItemUUID,
	BillOfMaterialComponent:        data.BillOfMaterialComponent,
	BillOfMaterialItemNumber:       data.BillOfMaterialItemNumber,
	HeaderChangeDocument:           data.HeaderChangeDocument,
	BillOfMaterialCategory:         data.BillOfMaterialCategory,
	BillOfMaterial:                 data.BillOfMaterial,
	BillOfMaterialVariant:          data.BillOfMaterialVariant,
	BillOfMaterialVersion:          data.BillOfMaterialVersion,
	BillOfMaterialItemCategory:     data.BillOfMaterialItemCategory,
	BillOfMaterialItemUnit:         data.BillOfMaterialItemUnit,
	BillOfMaterialItemQuantity:     data.BillOfMaterialItemQuantity,
	EngineeringChangeDocument:      data.EngineeringChangeDocument,
	ValidityStartDate:              data.ValidityStartDate,
	ValidityEndDate:                data.ValidityEndDate,
	BillOfMaterialItemNodeNumber:   data.BillOfMaterialItemNodeNumber,
	BOMItemDescription:             data.BOMItemDescription,
	Material:                       data.Material,
	MaterialName:                   data.MaterialName,
	PlantName:                      data.PlantName,
	BillOfMaterialVariantUsageDesc: data.BillOfMaterialVariantUsageDesc,
	Plant:                          data.Plant,
	BillOfMaterialVariantUsage:     data.BillOfMaterialVariantUsage,
	BOMVersionStatus:               data.BOMVersionStatus,
	BOMVersionStatusDescription:    data.BOMVersionStatusDescription,
		})
	}

	return whereUsedList, nil
}
