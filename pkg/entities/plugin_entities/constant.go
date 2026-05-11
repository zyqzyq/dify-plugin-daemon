package plugin_entities

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

const (
	SECRET_INPUT   = "secret-input"
	TEXT_INPUT     = "text-input"
	SELECT         = "select"
	STRING         = "string"
	NUMBER         = "number"
	FILE           = "file"
	FILES          = "files"
	BOOLEAN        = "boolean"
	APP_SELECTOR   = "app-selector"
	MODEL_SELECTOR = "model-selector"
	// TOOL_SELECTOR  = "tool-selector"
	TOOLS_SELECTOR = "array[tools]"
	ANY            = "any"
	// DynamicSelect
	DYNAMIC_SELECT      = "dynamic-select"
	DYNAMIC_TREE_SELECT = "dynamic-tree-select"
	ARRAY               = "array"
	OBJECT         = "object"
	CHECKBOX       = "checkbox"
	DATE           = "date"
	DATE_PICKER    = "date-picker"
)

type ParameterOption struct {
	Value    string                      `json:"value" yaml:"value" validate:"required"`
	Label    I18nObject                  `json:"label" yaml:"label" validate:"required"`
	Icon     string                      `json:"icon" yaml:"icon" validate:"omitempty"`
	Children []ParameterOption           `json:"children,omitempty" yaml:"children,omitempty" validate:"omitempty,dive"`
	ShowOn   []ToolParameterShowOnObject `json:"show_on" yaml:"show_on" validate:"omitempty,lte=16,dive"`
}

func (p *ParameterOption) UnmarshalJSON(data []byte) error {
	type Alias ParameterOption
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if p.ShowOn == nil {
		p.ShowOn = []ToolParameterShowOnObject{}
	}

	if p.Children == nil {
		p.Children = []ParameterOption{}
	}

	return nil
}

func (p *ParameterOption) UnmarshalYAML(value *yaml.Node) error {
	type Alias ParameterOption
	aux := &struct {
		*Alias `yaml:",inline"`
	}{
		Alias: (*Alias)(p),
	}

	if err := value.Decode(&aux); err != nil {
		return err
	}

	if p.ShowOn == nil {
		p.ShowOn = []ToolParameterShowOnObject{}
	}

	if p.Children == nil {
		p.Children = []ParameterOption{}
	}

	return nil
}
