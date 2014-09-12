package golang

import (
	"errors"
	"fmt"
	"github.com/scritch007/ShareMinatorApiGenerator/types"
	"github.com/scritch007/go-tools"
	"os"
	"strings"
)

type GolangGenerator struct {
	commandFile *os.File
	objectFile  *os.File
	enumFile    *os.File
	requestFile *os.File
}

func NewGolangGenerator(config *types.Config) (*GolangGenerator, error) {
	g := new(GolangGenerator)
	cF, err := os.OpenFile("api/commands.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)

	if nil != err {
		return nil, err
	}
	cF.WriteString("package api\n")

	oF, err := os.OpenFile("api/objects.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		cF.Close()
		return nil, err
	}
	oF.WriteString("package api\n")

	eF, err := os.OpenFile("api/enum.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		cF.Close()
		oF.Close()
		return nil, err
	}
	eF.WriteString("package api\n")
	rF, err := os.OpenFile("api/request.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		cF.Close()
		oF.Close()
		eF.Close()
		return nil, err
	}
	rF.WriteString("package api\n")
	g.commandFile = cF
	g.objectFile = oF
	g.enumFile = eF
	g.requestFile = rF
	return g, nil
}

func (g *GolangGenerator) generateObject(o *types.ObjectDefinition) error {
	return generateObject(g.objectFile, o.Name, o.Fields)
}

func generateObject(w *os.File, name string, fields *[]types.ObjectField) error {
	w.WriteString(fmt.Sprintf("type %s struct{\n", name))
	for _, field := range *fields {
		typeString := ""
		jsonString := "`json:\"" + field.Name
		commentString := ""
		if nil != field.Optional && *field.Optional {
			typeString += "*"
			jsonString += ",omitempty"
		}
		if "list" == field.Type {
			if nil == field.SubType {
				return errors.New(fmt.Sprintf("Couldn't find sub_type for field %s in object %s", field.Name, name))
			}
			typeString += "[]" + *field.SubType
		} else {
			//Special case where the field is a list of something
			typeString += field.Type
		}
		jsonString += "\"`"
		if nil != field.Comment {
			commentString += "// " + *field.Comment
		}
		w.WriteString(fmt.Sprintf("\t%s %s %s %s\n", tools.JsonToGolang(&field.Name), typeString, jsonString, commentString))
	}
	w.WriteString("}\n")
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

func (g *GolangGenerator) generateCommand(o *types.ObjectDefinition) (err error) {
	cname := getCommandName(&o.Name)

	if 0 != len(*o.Input) {
		err = generateObject(g.commandFile, cname+"Input", o.Input)
		if nil != err {
			return err
		}

	}
	if 0 != len(*o.Output) {
		err = generateObject(g.commandFile, cname+"Output", o.Output)
		if nil != err {
			return err
		}

	}

	g.commandFile.WriteString(fmt.Sprintf("type %s struct{\n", cname))
	if 0 != len(*o.Input) {
		g.commandFile.WriteString(fmt.Sprintf("\tInput %sInput `json:\"input\" bson:\"input\"` \n", cname))
	}
	if 0 != len(*o.Output) {
		g.commandFile.WriteString(fmt.Sprintf("\tOutput %sOutput `json:\"output\" bson:\"output\"`\n", cname))
	}
	g.commandFile.WriteString("}\n")
	return nil
}
func (g *GolangGenerator) generateEnum(o *types.ObjectDefinition) error {

	g.enumFile.WriteString(fmt.Sprintf("type %s int\n", o.Name))
	g.enumFile.WriteString("const (\n")
	isIota := false
	for i, enum := range *o.Values {
		var value = i
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
		g.enumFile.WriteString(fmt.Sprintf("\t%s %s = %d\n", enum.Name, o.Name, value))
	}
	g.enumFile.WriteString(")\n")
	return nil
}

func (g *GolangGenerator) GenerateObjects(a *types.APIDefinitions) error {
	for _, o := range a.Objects {
		err := g.generateObject(&o)
		if nil != err {
			return err
		}

	}
	return nil
}
func (g *GolangGenerator) GenerateCommands(a *types.APIDefinitions) error {
	for category, cmds := range a.Commands {
		for _, o := range *cmds {
			err := g.generateCommand(&o)
			if nil != err {
				return err
			}
		}
		g.commandFile.WriteString(fmt.Sprintf("type %sCommand struct{\n", tools.JsonToGolang(&category)))
		for _, o := range *cmds {
			_, action, err := o.CommandSplit()
			if nil != err {
				return err
			}
			g.commandFile.WriteString(fmt.Sprintf("\t%s *%s `json:\"%s,omitempty\" bson:\"%s,omitempty\"` \n", tools.JsonToGolang(&action), getCommandName(&o.Name), action, action))
		}
		g.commandFile.WriteString("}\n")
	}
	g.commandFile.WriteString(fmt.Sprintf("type EnumAction string\n"))
	g.commandFile.WriteString("const (\n")
	for category, cmds := range a.Commands {
		for _, o := range *cmds {
			_, action, _ := o.CommandSplit()
			g.commandFile.WriteString(fmt.Sprintf("\tEnum%s%s EnumAction = \"%s\"\n", tools.JsonToGolang(&category), tools.JsonToGolang(&action), o.Name))
		}
	}
	g.commandFile.WriteString(")\n")
	g.commandFile.WriteString("type Command struct{\n")
	g.commandFile.WriteString("\tName EnumAction `json:\"name\"`\n")
	g.commandFile.WriteString("\tCommandId string `json:\"command_id\"`\n")
	g.commandFile.WriteString("\tTimeout int64 `json:\"timeout,omitempty\"`\n")
	g.commandFile.WriteString("\tAuthKey *string `json:\"auth_key,omitempty\"` //Used when calling commands on behalf of a sharedlink\n")
	g.commandFile.WriteString("\tPassword *string `json:\"password\"` //Used when a share_link requires a password This should be the hash of AuthKey + Password\n")
	g.commandFile.WriteString("\tState CommandStatus `json:\"state\"`\n")
	for category, _ := range a.Commands {
		g.commandFile.WriteString(fmt.Sprintf("\t%s *%sCommand `json:\"%s,omitempty\"`\n", tools.JsonToGolang(&category), tools.JsonToGolang(&category), category))
	}
	g.commandFile.WriteString("}\n")
	return nil
}
func (g *GolangGenerator) GenerateEnums(a *types.APIDefinitions) error {
	for _, o := range a.Enums {
		err := g.generateEnum(&o)
		if nil != err {
			return err
		}
	}
	return nil
}

func getRequestName(commandName *string) (out string) {
	res := strings.Split(*commandName, ".")
	out = "Request"
	for _, s := range res {
		out += tools.Capitalize(s)
	}
	return tools.JsonToGolang(&out)
}

func (g *GolangGenerator) generateRequest(o *types.ObjectDefinition) (err error) {
	cname := getRequestName(&o.Name)

	var prefix string
	if types.RequestTypeAuth == *o.RequestType {
		prefix = "auths"
	} else {
		prefix = "config"
	}

	g.requestFile.WriteString("var " + cname + "Url" + "= \"" + prefix + "/" + o.Name + "\"\n")

	if 0 != len(*o.Input) {
		err = generateObject(g.requestFile, cname+"Input", o.Input)
		if nil != err {
			return err
		}

	}
	if 0 != len(*o.Output) {
		err = generateObject(g.requestFile, cname+"Output", o.Output)
		if nil != err {
			return err
		}

	}
	return nil
}

func (g *GolangGenerator) GenerateRequests(a *types.APIDefinitions) error {
	for _, o := range a.Requests {
		err := g.generateRequest(&o)
		if nil != err {
			return err
		}
	}
	return nil
}
