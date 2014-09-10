package types

import (
	"encoding/json"
	"errors"
	"strings"
)

type EnumObjectType string

const (
	ObjectTypeEnum    EnumObjectType = "Enum"
	ObjectTypeObject  EnumObjectType = "Object"
	ObjectTypeCommand EnumObjectType = "Command"
	ObjectTypeRequest EnumObjectType = "Request"
)

type EnumRequestType string

const (
	RequestTypeConfig EnumRequestType = "config"
	RequestTypeAuth   EnumRequestType = "auth"
)

type ObjectField struct {
	Optional *bool   `json:"optional, omitempty"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	SubType  *string `json:"sub_type, omitempty"`
	Comment  *string `json:"comment, omitempty"`
}

func (o *ObjectField) String() string {
	res, err := json.Marshal(o)
	if nil != err {
		return "Failed to serialize object with error " + err.Error()
	}
	return string(res)
}

type EnumValue struct {
	Name    string  `json:"name"`
	Value   *int    `json:"value,omitempty"` //The value can be omitted in that case we should be in a iota mode
	Comment *string `json:"comment, omitempty"`
}

func (o *EnumValue) String() string {
	res, err := json.Marshal(o)
	if nil != err {
		return "Failed to serialize object with error " + err.Error()
	}
	return string(res)
}

type ObjectDefinition struct {
	Name        string           `json:"name"`
	Type        EnumObjectType   `json:"type"`
	Methods     *[]string        `json:\"methods\"`
	QueryParams *[]string        `json:\"query_params\"`
	Fields      *[]ObjectField   `json:"fields,omitempty"`
	Input       *[]ObjectField   `json:"input,omitempty"`
	Output      *[]ObjectField   `json:"output,omitempty"`
	Values      *[]EnumValue     `json:"values, omitempty"`
	Comment     *string          `json:"comment, omitempty"`
	RequestType *EnumRequestType `json:"request_type,omitempty"` //Should be set when type is Request
}

func (o *ObjectDefinition) String() string {
	res, err := json.Marshal(o)
	if nil != err {
		return "Failed to serialize object with error " + err.Error()
	}
	return string(res)
}

func (o *ObjectDefinition) CommandSplit() (category string, action string, err error) {
	if ObjectTypeCommand != o.Type {
		return "", "", errors.New("Category only available for Commands")
	}
	split := strings.Split(o.Name, ".")
	return split[0], split[1], nil
}

func (o *ObjectDefinition) RequestSplit() (string, error) {
	split := strings.Split(o.Name, ".")
	if 2 == len(split) {
		return split[1], nil
	} else if 1 == len(split) {
		return split[0], nil
	} else {
		return "", errors.New("This is not correct split " + o.Name)
	}

}

func (o *ObjectDefinition) Dependencies() (deps []string, err error) {
	switch o.Type {
	case ObjectTypeObject:
		if nil == o.Fields {
			return nil, errors.New("No fields attribute for object " + o.Name)
		}
		deps = make([]string, 0, len(*o.Fields))
		for _, f := range *o.Fields {
			switch f.Type {
			case "string":
			case "int64":
			case "bool":
			default:
				deps = deps[0 : len(deps)+1]
				deps[len(deps)-1] = f.Type
			}
		}
		return deps, nil
	case ObjectTypeCommand:
		if nil == o.Input {
			return nil, errors.New("Command " + o.Name + " has no input definition")
		}
		if nil == o.Output {
			return nil, errors.New("Command " + o.Name + " has no output definition")
		}
		deps = make([]string, 0, len(*o.Input)+len(*o.Output))
		for _, f := range *o.Input {
			switch f.Type {
			case "string":
			case "int64":
			case "bool":
			default:
				deps = deps[0 : len(deps)+1]
				deps[len(deps)-1] = f.Type
			}
		}
		for _, f := range *o.Output {
			switch f.Type {
			case "string":
			case "int64":
			case "bool":
			default:
				deps = deps[0 : len(deps)+1]
				deps[len(deps)-1] = f.Type
			}
		}
		return deps, nil
	case ObjectTypeEnum:
		if nil == o.Values {
			return nil, errors.New("Enum " + o.Name + " has no values definition")
		}
		return nil, nil
	}
	return nil, errors.New("Not a valid Object type " + string(o.Type) + " for " + o.Name)
}

type GeneratorInterface interface {
	GenerateObjects(a *APIDefinitions) error
	GenerateCommands(a *APIDefinitions) error
	GenerateEnums(a *APIDefinitions) error
	GenerateRequests(a *APIDefinitions) error
}

type APIDefinitions struct {
	Objects  map[string]ObjectDefinition
	Commands map[string]*[]ObjectDefinition
	Enums    map[string]ObjectDefinition
	Requests map[string]ObjectDefinition
}

type Config struct{}
