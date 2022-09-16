# sap-api-integrations-bom-where-used-list-reads  
sap-api-integrations-bom-where-used-list-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 構成品目 部品表 使用先一覧 データを取得するマイクロサービスです。  
sap-api-integrations-bom-where-used-list-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-bom-where-used-list-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_BOM_WHERE_USED_SRV/overview   

## 動作環境
sap-api-integrations-bom-where-used-list-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-bom-where-used-list-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-bom-where-used-list-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BOM_WHERE_USED_SRV/overview   
* APIサービス名(=baseURL): API_BOM_WHERE_USED_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-bom-where-used-list-reads には、次の API をコールするためのリソースが含まれています。  

* A_BOMWhereUsed（構成品目 部品表 使用先一覧データ）

## API への 値入力条件 の 初期値
sap-api-integrations-bom-where-used-list-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.BillOfMaterialWhereUsedList.BillOfMaterialComponent（構成品目）
* inoutSDC.BillOfMaterialWhereUsedList.Material（品目）
* inoutSDC.BillOfMaterialWhereUsedList.Plant（プラント）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ByComponent" が指定されています。    
  
```
	"api_schema": "A_BOMWhereUsed",
	"accepter": ["ByComponent"],
	"material_code": "",
	"plant": "1010",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "A_BOMWhereUsed",
	"accepter": ["All"],
	"material_code": "",
	"plant": "1010",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 構成品目 部品表 使用先一覧 データ が取得された結果の JSON の例です。  
以下の項目のうち、"BillOfMaterialItemUUID" ～ "BOMVersionStatusDescription" は、/SAP_API_Output_Formatter/type.go 内 の Type WhereUsedList {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-bom-where-used-list-reads/SAP_API_Caller/caller.go#L89",
	"function": "sap-api-integrations-bom-where-used-list-reads/SAP_API_Caller.(*SAPAPICaller).ByMaterial",
	"level": "INFO",
	"message": [
		{
			"BillOfMaterialItemUUID": "00163e19-8846-1ed6-8ebf-608a9b43c3e4",
			"BillOfMaterialComponent": "RM13",
			"BillOfMaterialItemNumber": "0010",
			"HeaderChangeDocument": "",
			"BillOfMaterialCategory": "M",
			"BillOfMaterial": "00000001",
			"BillOfMaterialVariant": "1",
			"BillOfMaterialVersion": "",
			"BillOfMaterialItemCategory": "L",
			"BillOfMaterialItemUnit": "PC",
			"BillOfMaterialItemQuantity": "100",
			"EngineeringChangeDocument": "",
			"ValidityStartDate": "2006-01-01T09:00:00+09:00",
			"ValidityEndDate": "9999-12-31T09:00:00+09:00",
			"BillOfMaterialItemNodeNumber": "1",
			"BOMItemDescription": "",
			"Material": "SG23",
			"MaterialName": "SEMI23,PD,Subcontracting",
			"PlantName": "Plant 1 DE",
			"BillOfMaterialVariantUsageDesc": "Production",
			"Plant": "1010",
			"BillOfMaterialVariantUsage": "1",
			"BOMVersionStatus": "",
			"BOMVersionStatusDescription": ""
		},
		{
			"BillOfMaterialItemUUID": "00163e19-8846-1ed6-8ebf-60ab800583e4",
			"BillOfMaterialComponent": "RM14",
			"BillOfMaterialItemNumber": "0020",
			"HeaderChangeDocument": "",
			"BillOfMaterialCategory": "M",
			"BillOfMaterial": "00000001",
			"BillOfMaterialVariant": "1",
			"BillOfMaterialVersion": "",
			"BillOfMaterialItemCategory": "L",
			"BillOfMaterialItemUnit": "PC",
			"BillOfMaterialItemQuantity": "100",
			"EngineeringChangeDocument": "",
			"ValidityStartDate": "2006-01-01T09:00:00+09:00",
			"ValidityEndDate": "9999-12-31T09:00:00+09:00",
			"BillOfMaterialItemNodeNumber": "2",
			"BOMItemDescription": "",
			"Material": "SG23",
			"MaterialName": "SEMI23,PD,Subcontracting",
			"PlantName": "Plant 1 DE",
			"BillOfMaterialVariantUsageDesc": "Production",
			"Plant": "1010",
			"BillOfMaterialVariantUsage": "1",
			"BOMVersionStatus": "",
			"BOMVersionStatusDescription": ""
		},
		{
			"BillOfMaterialItemUUID": "00163e19-8846-1ee6-9fce-2efd69955443",
			"BillOfMaterialComponent": "RM13",
			"BillOfMaterialItemNumber": "0010",
			"HeaderChangeDocument": "",
			"BillOfMaterialCategory": "M",
			"BillOfMaterial": "00000001",
			"BillOfMaterialVariant": "1",
			"BillOfMaterialVersion": "",
			"BillOfMaterialItemCategory": "L",
			"BillOfMaterialItemUnit": "PC",
			"BillOfMaterialItemQuantity": "100",
			"EngineeringChangeDocument": "",
			"ValidityStartDate": "2006-01-01T09:00:00+09:00",
			"ValidityEndDate": "9999-12-31T09:00:00+09:00",
			"BillOfMaterialItemNodeNumber": "3",
			"BOMItemDescription": "",
			"Material": "SG23",
			"MaterialName": "SEMI23,PD,Subcontracting",
			"PlantName": "Plant 1 DE",
			"BillOfMaterialVariantUsageDesc": "Production",
			"Plant": "1010",
			"BillOfMaterialVariantUsage": "1",
			"BOMVersionStatus": "",
			"BOMVersionStatusDescription": ""
		}
	],
	"time": "2022-01-28T12:44:17+09:00"
}
```