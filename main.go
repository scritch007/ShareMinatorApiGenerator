package main

import (
	"encoding/json"
	"fmt"
	"github.com/scritch007/ShareMinatorApiGenerator/golang"
	"github.com/scritch007/ShareMinatorApiGenerator/js"
	"github.com/scritch007/ShareMinatorApiGenerator/types"
	"github.com/scritch007/go-tools"
	"io/ioutil"
	"os"
	"flag"
)

func main() {
	var help = false
	var configFile = ""
	var bGolang = false
	var bJS = false
	flag.StringVar(&configFile, "config", "", "Configuration file to use")
	flag.StringVar(&configFile, "c", "", "Configuration file to use")
	flag.BoolVar(&help, "help", false, "Display Help")
	flag.BoolVar(&help, "h", false, "Display Help")
	flag.BoolVar(&bGolang, "golang", false, "Build Api fo GO")
	flag.BoolVar(&bJS, "js", false, "Build Api for JS")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var api types.APIDefinitions
	var objects []types.ObjectDefinition
	tools.LogInit(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	file, err := ioutil.ReadFile("example.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &objects)
	if nil != err {
		fmt.Println("Couldn't deserialize the JSON with error" + err.Error())
		return
	}

	api.Commands = make(map[string]*[]types.ObjectDefinition)
	api.Objects = make(map[string]types.ObjectDefinition)
	api.Enums = make(map[string]types.ObjectDefinition)

	var generators []types.GeneratorInterface = make([]types.GeneratorInterface, 0, 2)

	if bGolang{
		generator, err := golang.NewGolangGenerator(new(types.Config))
		if nil != err {
			fmt.Println("Couldn't create Generator with error " + err.Error())
			return
		}
		generators = generators[:len(generators) + 1]
		generators[len(generators) - 1] = generator
	}

	if bJS{
		generator, err := js.NewJSGenerator(new(types.Config))
		if nil != err {
			fmt.Println("Couldn't create Generator with error " + err.Error())
			return
		}
		generators = generators[:len(generators) + 1]
		generators[len(generators) - 1] = generator
	}

	for _, object := range objects {
		switch object.Type {
		case types.ObjectTypeObject:
			if nil == object.Fields {
				fmt.Println("Object " + object.Name + " has no fields definition")
				return
			}
			api.Objects[object.Name] = object
		case types.ObjectTypeCommand:
			if nil == object.Input {
				fmt.Println("Command " + object.Name + " has no input definition")
				return
			}
			if nil == object.Output {
				fmt.Println("Command " + object.Name + " has no output definition")
				return
			}
			category, _, err := object.CommandSplit()
			if nil != err {
				fmt.Println("Command " + object.Name + " couldn't be splitted")
				return
			}
			actions, found := api.Commands[category]
			if !found {
				actions = new([]types.ObjectDefinition)
				*actions = make([]types.ObjectDefinition, 0, len(objects))
				api.Commands[category] = actions
			}
			*actions = (*actions)[0 : len(*actions)+1]
			(*actions)[len(*actions)-1] = object
		case types.ObjectTypeEnum:
			if nil == object.Values {
				fmt.Println("Enum " + object.Name + " has no values definition")
				return
			}
			api.Enums[object.Name] = object
		}
	}
	for _, generator := range generators{

		err = generator.GenerateObjects(&api)
		if nil != err {
			fmt.Println("Failed to generate Objects\n")
			fmt.Println(err)
			return
		}
		err = generator.GenerateCommands(&api)
		if nil != err {
			fmt.Println("Failed to generate Commands\n")
			fmt.Println(err)
			return
		}
		err = generator.GenerateEnums(&api)
		if nil != err {
			fmt.Println("Failed to generate Enums\n")
			fmt.Println(err)
			return
		}
	}
}
