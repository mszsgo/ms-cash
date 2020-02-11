package schema

import (
	"github.com/graphql-go/graphql"
)

type Query struct {
	Cash *CashQuery `description:"现金账户服务"`
}

type CashQuery struct {
	Name    string `description:"Service Name"`
	Version string `description:"Service Version"`
	// Org          OrgPointsType          `description:"积分发行机构积分账户信息"`
}

func (*CashQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return &CashQuery{Name: "cash", Version: "v0.0.1"}, err
	}
}
