package schema

import (
	"github.com/graphql-go/graphql"
)

type Mutation struct {
	Cash *CashMutation `description:"现金账户服务"`
}

type CashMutation struct {
	createAccount string `description:"创建现金账户"`
}

func (*CashMutation) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
