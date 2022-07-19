package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"go-conv/cr"
	schema2 "go-conv/schema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func buildWMS(schema schema2.ValidationResult, year string) []byte {
	wmsCrBaseFile, err := os.Open("resources/wms_base.yaml")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(wmsCrBaseFile)
	var wmsBase cr.WMS_CR
	err = yaml.Unmarshal(byteValue, &wmsBase)
	if err != nil {
		fmt.Println(err)
	}

	var crLayers []cr.Layers
	for _, table := range filterTables(schema.Tables, year) {
		var columnNames, geomType = columnNamesAndGeomType(table)
		var layerName = trim(table.Name, year)
		var styleName = getStyleFile(layerName)
		var crLayer = cr.Layers{
			Name:  layerName,
			Title: layerName,
			Data: cr.Data{
				Gpkg: cr.GPKGData{
					GeometryType: geomType,
					S3Key:        "geopackages/cbs/gebiedsindelingen/2/cbsgebiedsindelingen" + year + ".gpkg",
					Table:        table.Name,
					Columns:      columnNames,
				},
			},
			Abstract: `Deze service bevat de CBS Gebiedsindelingen van ` + year + `. De volgende gebiedsindelingen kunnen voorkomen: Gemeente Landsdeel Provincie Buurt COROP-gebied GGD-regio NUTS1 NUTS2 NUTS3 Wijk COROP-plusgebied COROP-subgebied Landbouwgebied Landbouwgroep RPA-gebied Toeristengroep Toeristengebied Arrondissementsgebied KamervanKoophandelregio Regionale-Eenheid Veiligheidsregio Veiligthuisregio Zorgkantoorregio Arbeidsmarktregio Regioplus-arbeidsmarktregio Jeugdregio Ressort RES-regio Regionaalmeldco\xF6rdinatiepunt Regionale-energiestrategie SubRES-regio`,
			Visible:  true,
			Keywords: []string{layerName},
			Styles: []cr.Styles{cr.Styles{
				Title: strings.Title(styleName),
				Name:  styleName,
				File:  styleName + ".style",
			}},
		}
		mergo.Merge(&crLayer, wmsBase.Spec.Service.Layers[0])
		crLayers = append(crLayers, crLayer)
	}

	var layerNames []string
	for _, crLayer := range crLayers[:1] {
		layerNames = append(layerNames, crLayer.Name)
	}

	wms_cr := cr.WMS_CR{
		Metadata: cr.Metadata{
			Name: year,
			Labels: cr.Labels{
				Theme: year,
			},
		},
		Spec: cr.WMSSpec{
			General: cr.General{
				Theme: year,
			},
			Service: cr.WMSService{
				Title:              "CBS Gebiedsindelingen WMS voor " + year,
				Abstract:           `Deze service bevat de CBS Gebiedsindelingen van ` + year + `. De volgende gebiedsindelingen kunnen voorkomen: Gemeente Landsdeel Provincie Buurt COROP-gebied GGD-regio NUTS1 NUTS2 NUTS3 Wijk COROP-plusgebied COROP-subgebied Landbouwgebied Landbouwgroep RPA-gebied Toeristengroep Toeristengebied Arrondissementsgebied KamervanKoophandelregio Regionale-Eenheid Veiligheidsregio Veiligthuisregio Zorgkantoorregio Arbeidsmarktregio Regioplus-arbeidsmarktregio Jeugdregio Ressort RES-regio Regionaalmeldcoördinatiepunt Regionale-energiestrategie SubRES-regio`,
				MetadataIdentifier: serviceMetadataIdentifiers["WMS"][year],
				Layers:             crLayers,
			},
		},
	}

	err = mergo.Merge(&wms_cr, wmsBase)
	if err != nil {
		fmt.Printf("Error merging wms cr's")
	}

	yamlData, err := yaml.Marshal(&wms_cr)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}
	return yamlData
}

func getStyleFile(layerName string) string {
	layerName = trim(layerName, "!!!!")
	// all label layers get the label style
	if strings.HasSuffix(layerName, "labelpoint") {
		return "cbs_label"
	}
	// all gegeneraliseerd en niet_gegeneraliseerd layer:
	if strings.HasPrefix(layerName, "buurt_") {
		return "cbs_buurten"
	}
	if strings.HasPrefix(layerName, "wijk_") {
		return "cbs_wijken"
	}
	return "cbs_gebiedsindeling"
}
