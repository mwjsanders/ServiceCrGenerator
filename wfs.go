package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"go-conv/cr"
	schema2 "go-conv/schema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func buildWFS(schema schema2.ValidationResult, year string) []byte {
	wfsCrBaseFile, err := os.Open("resources/wfs_base.yaml")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(wfsCrBaseFile)
	var wfsBase cr.WFS_CR
	yaml.Unmarshal(byteValue, &wfsBase)

	var features []cr.FeatureType
	for _, table := range filterTables(schema.Tables, year) {
		var columnNames, geomType = columnNamesAndGeomType(table)
		var featureName = trim(table.Name, year)
		var feature = cr.FeatureType{
			Name:  featureName,
			Title: featureName,
			Data: cr.Data{
				Gpkg: cr.GPKGData{
					GeometryType: geomType,
					S3Key:        "${S3_GEOPACKAGES_BUCKET}/cbs/gebiedsindelingen/2/cbsgebiedsindelingen" + year + ".gpkg",
					Table:        table.Name,
					Columns:      columnNames,
				}},
			Abstract: featureName,
			Keywords: []string{featureName},
		}
		mergo.Merge(&feature, wfsBase.Spec.Service.FeatureTypes[0])
		features = append(features, feature)
	}

	var featureNames []string
	for _, feature := range features[:1] {
		featureNames = append(featureNames, feature.Name)
	}
	wfs_cr := cr.WFS_CR{
		Metadata: cr.Metadata{
			Name: year,
			Labels: cr.Labels{
				Theme: year,
			},
		},
		Spec: cr.WFSSpec{
			General: cr.General{
				Theme: year,
			},
			Service: cr.WFSService{
				Title:              "CBS Gebiedsindelingen " + year + " WFS",
				Abstract:           "Deze service bevat de CBS Gebiedsindelingen van " + year + ". De volgende gebiedsindelingen kunnen voorkomen: Gemeente Landsdeel Provincie Buurt COROP-gebied GGD-regio NUTS1 NUTS2 NUTS3 Wijk COROP-plusgebied COROP-subgebied Landbouwgebied Landbouwgroep RPA-gebied Toeristengroep Toeristengebied Arrondissementsgebied KamervanKoophandelregio Regionale-Eenheid Veiligheidsregio Veiligthuisregio Zorgkantoorregio Arbeidsmarktregio Regioplus-arbeidsmarktregio Jeugdregio Ressort RES-regio Regionaalmeldcoördinatiepunt Regionale-energiestrategie SubRES-regio",
				MetadataIdentifier: serviceMetadataIdentifiers["WFS"][year],
				FeatureTypes:       features,
			},
		},
	}

	err = mergo.Merge(&wfs_cr, wfsBase)
	if err != nil {
		fmt.Printf("Error merging wms cr's")
	}

	yamlData, err := yaml.Marshal(&wfs_cr)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}
	return yamlData
}
