package cr

type WMS_CR struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       WMSSpec  `yaml:"spec"`
}

type Symbols struct {
	FileName string `yaml:"fileName"`
}
type Styles struct {
	Title string `yaml:"title"`
	Name  string `yaml:"name"`
	File  string `yaml:"visualization"`
}
type Layers struct {
	Name                      string   `yaml:"name"`
	Title                     string   `yaml:"title"`
	Abstract                  string   `yaml:"abstract"`
	Visible                   bool     `yaml:"visible"`
	Keywords                  []string `yaml:"keywords"`
	SourceMetadataIdentifier  string   `yaml:"sourceMetadataIdentifier"`
	DatasetMetadataIdentifier string   `yaml:"datasetMetadataIdentifier"`
	Styles                    []Styles `yaml:"styles"`
	Data                      Data     `yaml:"data"`
}

type Includes struct {
	ConfigMap string `yaml:"configMap"`
}

type WMSSpec struct {
	General    General    `yaml:"general"`
	Kubernetes Kubernetes `yaml:"kubernetes""`
	Options    WMSOptions `yaml:"options"`
	Service    WMSService `yaml:"service"`
}

type WMSOptions struct {
	AutomaticCasing        bool `yaml:"automaticCasing"`
	DisableWebserviceProxy bool `yaml:"disableWebserviceProxy"`
	IncludeIngress         bool `yaml:"includeIngress"`
	ValidateRequests       bool `yaml:"validateRequests"`
}

type WMSService struct {
	Inspire            bool          `yaml:"inspire"`
	Title              string        `yaml:"title"`
	Abstract           string        `yaml:"abstract"`
	Keywords           []string      `yaml:"keywords"`
	MetadataIdentifier string        `yaml:"metadataIdentifier"`
	Authority          Authority     `yaml:"authority"`
	DataEPSG           string        `yaml:"dataEPSG"`
	Extent             string        `yaml:"extent"`
	StylingAssets      StylingAssets `yaml:"stylingAssets"`
	Layers             []Layers      `yaml:"layers"`
}

type StylingAssets struct {
	ConfigMapRefs []ConfigMapRefs `yaml:"configMapRefs"`
	S3Keys        []string        `yaml:"s3Keys"`
}

type ConfigMapRefs struct {
	Name string `yaml:"name"`
}
