package proto2gql

import (
	"github.com/pkg/errors"

	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/proto2gql/parser"
)

func (g *Proto2GraphQL) outputMapGraphQLName(mapFile *parsedFile, res *parser.Map) string {
	return g.outputMessageVariable(mapFile, res.Message) + "__" + res.Field.Name
}

func (g *Proto2GraphQL) outputMapVariable(mapFile *parsedFile, res *parser.Map) string {
	return g.outputMessageVariable(mapFile, res.Message) + "__" + res.Field.Name
}

func (g *Proto2GraphQL) fileMapOutputObjects(file *parsedFile) ([]graphql.MapOutputObject, error) {
	var res []graphql.MapOutputObject
	for _, msg := range file.File.Messages {
		for _, mapFld := range msg.MapFields {
			keyTypResolver, err := g.TypeOutputTypeResolver(file, mapFld.Map.KeyType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve key input type resolver")
			}
			valueFile, err := g.parsedFile(mapFld.Map.ValueType.File())
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value type file")
			}
			valueTypResolver, err := g.TypeOutputTypeResolver(valueFile, mapFld.Map.ValueType)
			if err != nil {
				return nil, errors.Wrap(err, "failed to resolve value input type resolver")
			}
			res = append(res, graphql.MapOutputObject{
				VariableName:    g.outputMapVariable(file, mapFld.Map),
				GraphQLName:     g.outputMapGraphQLName(file, mapFld.Map),
				KeyObjectType:   keyTypResolver,
				ValueObjectType: valueTypResolver,
			})
		}
	}
	return res, nil
}
