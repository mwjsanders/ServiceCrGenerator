package cr

type Labels struct {
	DatasetOwner   string `yaml:"dataset-owner"`
	Dataset        string `yaml:"dataset"`
	Theme          string `yaml:"theme,omitempty"`
	ServiceVersion string `yaml:"service-version"`
	ServiceType    string `yaml:"service-type"`
}
type Metadata struct {
	Name   string `yaml:"name"`
	Labels Labels `yaml:"labels"`
}

type General struct {
	DatasetOwner   string `yaml:"datasetOwner"`
	Dataset        string `yaml:"dataset"`
	Theme          string `yaml:"theme,omitempty"`
	ServiceVersion string `yaml:"serviceVersion"`
}

type Authority struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}
type Kubernetes struct {
	HealthCheck HealthCheck    `yaml:"healthCheck"`
	Resources   ResourceLimits `yaml:"resources"`
}

type ResourceLimits struct {
	Limits Limits `yaml:"limits"`
}

type Limits struct {
	EphemeralStorage string `yaml:"ephemeralStorage"`
}

type HealthCheck struct {
	BoundingBox string `yaml:"boundingbox,omitempty"`
	MimeType    string `yaml:"mimetype,omitempty"`
	QueryString string `yaml:"querystring,omitempty"`
}

type Data struct {
	Gpkg GPKGData `yaml:"gpkg"`
}

type GPKGData struct {
	S3Key        string   `yaml:"s3Key"`
	Table        string   `yaml:"table"`
	GeometryType string   `yaml:"geometryType"`
	Columns      []string `yaml:"columns"`
}
