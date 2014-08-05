package main

import (
	"encoding/json"
	"fmt"
	"github.com/scritch007/ShareMinatorApiGenerator/golang"
	"github.com/scritch007/ShareMinatorApiGenerator/types"
	"io/ioutil"
)

func main() {
	var api types.APIDefinitions
	var objects []types.ObjectDefinition

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

	var generator types.GeneratorInterface
	generator, err = golang.NewGolangGenerator(new(types.Config))
	if nil != err {
		fmt.Println("Couldn't create Generator with error " + err.Error())
		return
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
