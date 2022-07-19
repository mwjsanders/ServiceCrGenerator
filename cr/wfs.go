package cr

type WFS_CR struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       WFSSpec  `yaml:"spec"`
}

type WFSSpec struct {
	General    General    `yaml:"general"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
	Options    WFSOptions `yaml:"options"`
	Service    WFSService `yaml:"service"`
}

type WFSOptions struct {
	AutomaticCasing bool `yaml:"automaticCasing"`
	IncludeIngress  bool `yaml:"includeIngress"`
}

type WFSService struct {
	Inspire            bool          `yaml:"inspire"`
	Title              string        `yaml:"title"`
	Abstract           string        `yaml:"abstract"`
	Keywords           []string      `yaml:"keywords"`
	Extent             string        `yaml:"extent"`
	MetadataIdentifier string        `yaml:"metadataIdentifier"`
	Authority          Authority     `yaml:"authority"`
	DataEPSG           string        `yaml:"dataEPSG""`
	FeatureTypes       []FeatureType `yaml:"featureTypes"`
}

type FeatureType struct {
	Name                      string   `yaml:"name"`
	Title                     string   `yaml:"title"`
	Abstract                  string   `yaml:"abstract"`
	Keywords                  []string `yaml:"keywords"`
	SourceMetadataIdentifier  string   `yaml:"sourceMetadataIdentifier"`
	DatasetMetadataIdentifier string   `yaml:"datasetMetadataIdentifier"`
	Data                      Data     `yaml:"data"`
}
