package main

import (
	"encoding/json"
	"fmt"
	schema2 "go-conv/schema"
	"io/ioutil"
	"os"
	"strings"
)

var serviceMetadataIdentifiers = map[string]map[string]string{
	"WMS": {
		"2016": "1b2e3afd-e5dd-4d05-aeee-a75091a76beb",
		"2017": "f874407b-25f9-4c1b-abe6-363467683bb1",
		"2018": "97590436-ecaf-422b-a95e-6aa9a00b47b0",
		"2019": "b965603f-7354-4d5c-9357-68c1c3777117",
		"2020": "991033b9-73de-4abb-9e5e-269a39152852",
		"2021": "43bb4581-0720-426f-a771-b7f9ded50fd5",
		"2022": "9c50592e-57be-4d74-91fd-bcaee20bb14e",
	},
	"WFS": {
		"2016": "150e3a18-44ed-43bd-b0fd-70ff7a3e6906",
		"2017": "8c280cce-ef1b-49b6-93dc-6cee38956101",
		"2018": "b39970a9-e1d7-4a20-9fbd-57661f6d6849",
		"2019": "8c61b066-de2a-4b80-bd20-c878e4edae86",
		"2020": "db353b96-5d0b-453c-87c6-4466dfb65a9b",
		"2021": "9c4c9bbc-4746-414b-9599-b67db6de1d6d",
		"2022": "5dacef44-361e-4a3d-b86f-1d6d26297337",
	},
}

func main() {
	for _, year := range []string{"2016", "2017", "2018", "2019", "2020", "2021", "2022"} {
		schemaFile, err := os.Open(fmt.Sprintf("resources/schema_%v.json", year))
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(schemaFile)
		var schema schema2.ValidationResult
		json.Unmarshal(byteValue, &schema)

		os.WriteFile(fmt.Sprintf("deploy/wms/%s.yaml", year), buildWMS(schema, year), 0644)
		os.WriteFile(fmt.Sprintf("deploy/wfs/%s.yaml", year), buildWFS(schema, year), 0644)
		fmt.Println("ready")
	}
}

func filterTables(slice []schema2.Tables, filter string) []schema2.Tables {
	var result []schema2.Tables
	for _, table := range slice {
		if strings.Contains(table.Name, filter) {
			result = append(result, table)
		}
	}
	return result
}

func trim(name string, year string) string {
	result := name
	result = strings.TrimPrefix(result, "cbsgebiedsindelingen:")
	result = strings.TrimPrefix(result, "cbs_")
	result = strings.TrimPrefix(result, "CBS_")
	result = strings.ReplaceAll(result, "_"+year, "")
	return result
}

func geometryTypesToCC(geometryType string) string {
	geometryType = strings.ReplaceAll(geometryType, "MULTI", "Multi")
	geometryType = strings.ReplaceAll(geometryType, "LINE", "Line")
	geometryType = strings.ReplaceAll(geometryType, "STRING", "String")
	geometryType = strings.ReplaceAll(geometryType, "POLYGON", "Polygon")
	geometryType = strings.ReplaceAll(geometryType, "POINT", "Point")
	return geometryType
}

func columnNamesAndGeomType(table schema2.Tables) ([]string, string) {
	var names []string
	var geomType string
	for _, column := range table.Columns {
		if column.Name == "fid" {
			//ignore
		} else if column.Name == table.GeometryColumn {
			geomType = geometryTypesToCC(column.Type)
		} else {
			names = append(names, column.Name)
		}
	}
	return names, geomType
}
