package kieclient

//KVDoc is database struct to store kv
type KVDoc struct {
	ID             string `json:"id,omitempty" bson:"id,omitempty" yaml:"id,omitempty" swag:"string"`
	LabelFormat    string `json:"label_format,omitempty" bson:"label_format,omitempty" yaml:"label_format,omitempty"`
	Key            string `json:"key" yaml:"key" validate:"key"`
	Value          string `json:"value" yaml:"value" validate:"value"`
	ValueType      string `json:"value_type,omitempty" bson:"value_type,omitempty" yaml:"value_type,omitempty" validate:"valueType"` //ini,json,text,yaml,properties
	Checker        string `json:"check,omitempty" yaml:"check,omitempty" validate:"check"`                                           //python script
	CreateRevision int64  `json:"create_revision,omitempty" bson:"create_revision," yaml:"create_revision,omitempty"`
	UpdateRevision int64  `json:"update_revision,omitempty" bson:"update_revision," yaml:"update_revision,omitempty"`
	Project        string `json:"project,omitempty" yaml:"project,omitempty" validate:"commonName"`
	Status         string `json:"status,omitempty" yaml:"status,omitempty" validate:"kvStatus"`
	CreateTime     int64  `json:"create_time,omitempty" bson:"create_time," yaml:"create_time,omitempty"`
	UpdateTime     int64  `json:"update_time,omitempty" bson:"update_time," yaml:"update_time,omitempty"`

	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty" validate:"max=6,dive,keys,labelKV,endkeys,labelKV"` //redundant
	Domain string            `json:"domain,omitempty" yaml:"domain,omitempty" validate:"commonName"`                              //redundant
}

//KVRequest is http request body
type KVRequest struct {
	Key       string            `json:"key" yaml:"key"`
	Value     string            `json:"value,omitempty" yaml:"value,omitempty"`
	ValueType string            `json:"value_type,omitempty" bson:"value_type,omitempty" yaml:"value_type,omitempty"` //ini,json,text,yaml,properties
	Checker   string            `json:"check,omitempty" yaml:"check,omitempty"`                                       //python script
	Labels    map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`                                     //redundant
}

//KVResponse represents the key value list
type KVResponse struct {
	LabelDoc *LabelDocResponse `json:"label,omitempty"`
	Total    int               `json:"total"`
	Data     []*KVDoc          `json:"data"`
}

//LabelDocResponse is label struct
type LabelDocResponse struct {
	Labels map[string]string `json:"labels,omitempty"`
}
