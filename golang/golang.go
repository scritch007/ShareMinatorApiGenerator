package golang

import (
	"errors"
	"fmt"
	"github.com/scritch007/ShareMinatorApiGenerator/types"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type GolangGenerator struct {
	commandFile *os.File
	objectFile  *os.File
	enumFile    *os.File
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
		eF.Close()
		return nil, err
	}
	eF.WriteString("package api\n")
	g.commandFile = cF
	g.objectFile = oF
	g.enumFile = eF
	return g, nil
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func JsonToGolang(in *string) (out string) {
	res := strings.Split(*in, "_")
	out = ""
	for _, s := range res {
		out += Capitalize(s)
	}
	return out
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
		w.WriteString(fmt.Sprintf("\t%s %s %s %s\n", JsonToGolang(&field.Name), typeString, jsonString, commentString))
	}
	w.WriteString("}\n")
	return nil
}
func getCommandName(commandName *string) (out string) {
	res := strings.Split(*commandName, ".")
	out = "Command"
	for _, s := range res {
		out += Capitalize(s)
	}
	return JsonToGolang(&out)
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
		err = generateObject(g.commandFile, cname+"Output", o.Input)
		if nil != err {
			return err
		}

	}

	g.commandFile.WriteString(fmt.Sprintf("type %s struct{\n", cname))
	if 0 != len(*o.Input) {
		g.commandFile.WriteString(fmt.Sprintf("\tInput %sInput `json:\"input\"`\n", cname))
	}
	if 0 != len(*o.Output) {
		g.commandFile.WriteString(fmt.Sprintf("\tOutput %sOutput `json:\"output\"`\n", cname))
	}
	g.commandFile.WriteString("}\n")
	return nil
}
func (g *GolangGenerator) generateEnum(o *types.ObjectDefinition) error {
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
		g.commandFile.WriteString(fmt.Sprintf("type %sCommand struct{\n", JsonToGolang(&category)))
		for _, o := range *cmds {
			_, action, err := o.CommandSplit()
			if nil != err {
				return err
			}
			g.commandFile.WriteString(fmt.Sprintf("\t%s *%s `json:\"%s,omitempty\"`\n", JsonToGolang(&action), getCommandName(&o.Name), action))
		}
		g.commandFile.WriteString("}\n")
	}
	g.commandFile.WriteString("type Command struct{\n")
	g.commandFile.WriteString("\tName string `json:\"name\"`\n")
	for category, _ := range a.Commands {
		g.commandFile.WriteString(fmt.Sprintf("\t%s *%sCommand `json:\"%s ,omitempty\"`\n", JsonToGolang(&category), JsonToGolang(&category), category))
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
