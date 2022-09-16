package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-bom-where-used-list-reads-sandbox/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBillOfMaterialWhereUsedList(billOfMaterialComponent, material, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ByComponent":
			func() {
				c.ByComponent(billOfMaterialComponent, plant)
				wg.Done()
			}()
		case "ByMaterial":
			func() {
				c.ByMaterial(material, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ByComponent(billOfMaterialComponent, plant string) {
	data, err := c.callBillOfMaterialWhereUsedListSrvAPIRequirementByComponent("A_BOMWhereUsed", billOfMaterialComponent, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callBillOfMaterialWhereUsedListSrvAPIRequirementByComponent(api, billOfMaterialComponent, plant string) ([]sap_api_output_formatter.WhereUsedList, error) {
	url := strings.Join([]string{c.baseURL, "API_BOM_WHERE_USED_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithByComponent(req, billOfMaterialComponent, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToWhereUsedList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ByMaterial(material, plant string) {
	data, err := c.callBillOfMaterialWhereUsedListSrvAPIRequirementByMaterial("A_BOMWhereUsed", material, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callBillOfMaterialWhereUsedListSrvAPIRequirementByMaterial(api, material, plant string) ([]sap_api_output_formatter.WhereUsedList, error) {
	url := strings.Join([]string{c.baseURL, "API_BOM_WHERE_USED_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithByMaterial(req, material, plant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToWhereUsedList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithByComponent(req *http.Request, billOfMaterialComponent, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillOfMaterialComponent eq '%s' and Plant eq '%s'", billOfMaterialComponent, plant))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithByMaterial(req *http.Request, material, plant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Material eq '%s' and Plant eq '%s'", material, plant))
	req.URL.RawQuery = params.Encode()
}
