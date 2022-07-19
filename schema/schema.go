package schema

type ValidationResult struct {
	GeopackageValidatorVersion string   `json:"geopackage_validator_version"`
	Projection                 int      `json:"projection"`
	Tables                     []Tables `json:"tables"`
}
type Columns struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Tables struct {
	Name           string    `json:"name"`
	GeometryColumn string    `json:"geometry_column"`
	Columns        []Columns `json:"columns"`
}
