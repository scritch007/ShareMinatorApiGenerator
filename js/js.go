package js

import (
	"errors"
	"fmt"
	"github.com/scritch007/ShareMinatorApiGenerator/types"
	"github.com/scritch007/go-tools"
	"os"
	"strings"
)

type JSGenerator struct {
	apiFile *os.File
}

func NewJSGenerator(config *types.Config) (*JSGenerator, error) {
	g := new(JSGenerator)
	cF, err := os.OpenFile("api/api.js", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)

	if nil != err {
		return nil, err
	}

	g.apiFile = cF
	return g, nil
}

func (g *JSGenerator) generateObject(o *types.ObjectDefinition) error {
	return generateObject(g.apiFile, o.Name, o.Fields)
}

func generateObject(w *os.File, name string, fields *[]types.ObjectField) error {

	return nil
}

func getCommandName(commandName *string) (out string) {
	res := strings.Split(*commandName, ".")
	out = "Command"
	for _, s := range res {
		out += tools.Capitalize(s)
	}
	return tools.JsonToGolang(&out)
}

func (g *JSGenerator) generateCommand(o *types.ObjectDefinition) (err error) {
	cname := getCommandName(&o.Name)
	g.apiFile.WriteString(fmt.Sprintf("function send%s(", cname))
	//mandatoryInputs := make([]string, 0, len(*o.Input))
	//for _, input := range *o.Input{
	//	if nil == input.Optional || !*input.Optional{
	//		mandatoryInputs = mandatoryInputs[:len(mandatoryInputs) + 1]
	//		mandatoryInputs[len(mandatoryInputs) - 1] = input.Name
	//	}
	//}
	//g.apiFile.WriteString(strings.Join(mandatoryInputs, ", "))
	if nil != o.Input  && 0 != len(*o.Input){
		g.apiFile.WriteString("input, ")
	}
	g.apiFile.WriteString("config, onSuccess, onError, onPending){\n")
	if nil != o.Input  && 0 != len(*o.Input){
		g.apiFile.WriteString("\t//TODO check that all the mandatory inputs are there\n")
	}else{
		g.apiFile.WriteString("\tvar input = null;\n")
	}
	g.apiFile.WriteString(fmt.Sprintf("\tsendCommand(\"%s\", input, config, onSuccess, onError, onPending);\n", o.Name))
	g.apiFile.WriteString("}\n")
	return nil
}
func (g *JSGenerator) generateEnum(o *types.ObjectDefinition) error {

	g.apiFile.WriteString(fmt.Sprintf("/******************%s******************/\n", o.Name))
	g.apiFile.WriteString(fmt.Sprintf("var %s = function(){};\n", o.Name))
	isIota := false
	for i, enum := range *o.Values {
		var value = i
		fmt.Println("i = ", i, " ", enum.Name)
		if nil == enum.Value {

			if 0 == i {
				isIota = true
			} else {
				if !isIota {
					return errors.New("Value " + enum.Name + " has not value but should have")
				}
			}
		} else {
			if isIota {
				return errors.New("Value " + enum.Name + " has a value but shouldn't")
			}
		}
		g.apiFile.WriteString(fmt.Sprintf("%s.%s = %d;\n", o.Name, enum.Name, value))
	}
	g.apiFile.WriteString(fmt.Sprintf("/******************%s******************/\n", o.Name))
	return nil
}

func (g *JSGenerator) GenerateObjects(a *types.APIDefinitions) error {
	for _, o := range a.Objects {
		err := g.generateObject(&o)
		if nil != err {
			return err
		}

	}
	return nil
}
func (g *JSGenerator) GenerateCommands(a *types.APIDefinitions) error {
	for category, cmds := range a.Commands {
		fmt.Println(category)
		for _, o := range *cmds {
			err := g.generateCommand(&o)
			if nil != err {
				return err
			}
		}
	}
	return nil
}
func (g *JSGenerator) GenerateEnums(a *types.APIDefinitions) error {
	for _, o := range a.Enums {
		err := g.generateEnum(&o)
		if nil != err {
			return err
		}
	}
	return nil
}
