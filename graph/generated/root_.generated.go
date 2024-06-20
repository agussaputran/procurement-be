// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"procurement-be/graph/model"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	AuthMutation() AuthMutationResolver
	Mutation() MutationResolver
	Query() QueryResolver
	UserMutation() UserMutationResolver
	UserQuery() UserQueryResolver
}

type DirectiveRoot struct {
	Auth     func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	User     func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	LoggedIn func(ctx context.Context, obj interface{}, next graphql.Resolver, loggedIn model.Login) (res interface{}, err error)
}

type ComplexityRoot struct {
	AuthMutation struct {
		Login func(childComplexity int, in *model.LoginRequestInput) int
	}

	Data struct {
		Token    func(childComplexity int) int
		UserData func(childComplexity int) int
	}

	FetchResponse struct {
		Data      func(childComplexity int) int
		IsSuccess func(childComplexity int) int
		Message   func(childComplexity int) int
		Status    func(childComplexity int) int
	}

	LoginResponse struct {
		Data      func(childComplexity int) int
		IsSuccess func(childComplexity int) int
		Message   func(childComplexity int) int
		Status    func(childComplexity int) int
	}

	Mutation struct {
		Auth func(childComplexity int) int
		User func(childComplexity int) int
	}

	Query struct {
		User               func(childComplexity int) int
		__resolve__service func(childComplexity int) int
	}

	UserData struct {
		CreatedAt func(childComplexity int) int
		Email     func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Role      func(childComplexity int) int
	}

	UserItems struct {
		Items func(childComplexity int) int
	}

	UserMutation struct {
		Store  func(childComplexity int, in *model.UserDataInput) int
		Update func(childComplexity int, in *model.UserDataInput) int
	}

	UserQuery struct {
		Fetch func(childComplexity int, in *model.FetchRequestInput) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "AuthMutation.Login":
		if e.complexity.AuthMutation.Login == nil {
			break
		}

		args, err := ec.field_AuthMutation_Login_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.AuthMutation.Login(childComplexity, args["in"].(*model.LoginRequestInput)), true

	case "Data.token":
		if e.complexity.Data.Token == nil {
			break
		}

		return e.complexity.Data.Token(childComplexity), true

	case "Data.userData":
		if e.complexity.Data.UserData == nil {
			break
		}

		return e.complexity.Data.UserData(childComplexity), true

	case "FetchResponse.data":
		if e.complexity.FetchResponse.Data == nil {
			break
		}

		return e.complexity.FetchResponse.Data(childComplexity), true

	case "FetchResponse.isSuccess":
		if e.complexity.FetchResponse.IsSuccess == nil {
			break
		}

		return e.complexity.FetchResponse.IsSuccess(childComplexity), true

	case "FetchResponse.message":
		if e.complexity.FetchResponse.Message == nil {
			break
		}

		return e.complexity.FetchResponse.Message(childComplexity), true

	case "FetchResponse.status":
		if e.complexity.FetchResponse.Status == nil {
			break
		}

		return e.complexity.FetchResponse.Status(childComplexity), true

	case "LoginResponse.data":
		if e.complexity.LoginResponse.Data == nil {
			break
		}

		return e.complexity.LoginResponse.Data(childComplexity), true

	case "LoginResponse.isSuccess":
		if e.complexity.LoginResponse.IsSuccess == nil {
			break
		}

		return e.complexity.LoginResponse.IsSuccess(childComplexity), true

	case "LoginResponse.message":
		if e.complexity.LoginResponse.Message == nil {
			break
		}

		return e.complexity.LoginResponse.Message(childComplexity), true

	case "LoginResponse.status":
		if e.complexity.LoginResponse.Status == nil {
			break
		}

		return e.complexity.LoginResponse.Status(childComplexity), true

	case "Mutation.Auth":
		if e.complexity.Mutation.Auth == nil {
			break
		}

		return e.complexity.Mutation.Auth(childComplexity), true

	case "Mutation.User":
		if e.complexity.Mutation.User == nil {
			break
		}

		return e.complexity.Mutation.User(childComplexity), true

	case "Query.User":
		if e.complexity.Query.User == nil {
			break
		}

		return e.complexity.Query.User(childComplexity), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "UserData.createdAt":
		if e.complexity.UserData.CreatedAt == nil {
			break
		}

		return e.complexity.UserData.CreatedAt(childComplexity), true

	case "UserData.email":
		if e.complexity.UserData.Email == nil {
			break
		}

		return e.complexity.UserData.Email(childComplexity), true

	case "UserData.id":
		if e.complexity.UserData.ID == nil {
			break
		}

		return e.complexity.UserData.ID(childComplexity), true

	case "UserData.name":
		if e.complexity.UserData.Name == nil {
			break
		}

		return e.complexity.UserData.Name(childComplexity), true

	case "UserData.role":
		if e.complexity.UserData.Role == nil {
			break
		}

		return e.complexity.UserData.Role(childComplexity), true

	case "UserItems.items":
		if e.complexity.UserItems.Items == nil {
			break
		}

		return e.complexity.UserItems.Items(childComplexity), true

	case "UserMutation.Store":
		if e.complexity.UserMutation.Store == nil {
			break
		}

		args, err := ec.field_UserMutation_Store_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.UserMutation.Store(childComplexity, args["in"].(*model.UserDataInput)), true

	case "UserMutation.Update":
		if e.complexity.UserMutation.Update == nil {
			break
		}

		args, err := ec.field_UserMutation_Update_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.UserMutation.Update(childComplexity, args["in"].(*model.UserDataInput)), true

	case "UserQuery.Fetch":
		if e.complexity.UserQuery.Fetch == nil {
			break
		}

		args, err := ec.field_UserQuery_Fetch_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.UserQuery.Fetch(childComplexity, args["in"].(*model.FetchRequestInput)), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputFetchRequestInput,
		ec.unmarshalInputLoginRequestInput,
		ec.unmarshalInputUserDataInput,
		ec.unmarshalInputlogin,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../schema/auth.graphql", Input: `directive @Auth on FIELD_DEFINITION
type AuthMutation {
	Login(in: LoginRequestInput): LoginResponse
}
type Data {
	token: String
	userData: UserData
}
type LoginResponse {
	isSuccess: Boolean
	message: String
	data: Data
	status: Int
}
input LoginRequestInput {
	email: String
	password: String
}
`, BuiltIn: false},
	{Name: "../../schema/core.graphql", Input: `directive @loggedIn(loggedIn: login!) on FIELD_DEFINITION

input login {
    access: String!
}

type Mutation {
	Auth: AuthMutation!
	User: UserMutation!
}
type Query {
	User: UserQuery!
}`, BuiltIn: false},
	{Name: "../../schema/user.graphql", Input: `directive @User on FIELD_DEFINITION
type UserQuery {
	Fetch(in: FetchRequestInput): FetchResponse @loggedIn(loggedIn: { access: "Authenticated"})
}
type UserMutation {
	Store(in: UserDataInput): LoginResponse @loggedIn(loggedIn: { access: "Authenticated"})
	Update(in: UserDataInput): LoginResponse @loggedIn(loggedIn: { access: "Authenticated"})
}
input FetchRequestInput {
	role: String
	name: String
}
type FetchResponse {
	isSuccess: Boolean
	message: String
	data: UserItems
	status: Int
}
type UserItems {
	items: [UserData!]
}
type UserData{
	id: String
	role: String
    email: String
	name: String
	createdAt: String
}
input UserDataInput{
	id: String
	role: String
    email: String
	name: String
	password: String
}`, BuiltIn: false},
	{Name: "../../federation/directives.graphql", Input: `
	directive @key(fields: _FieldSet!) repeatable on OBJECT | INTERFACE
	directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
	directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
	directive @extends on OBJECT | INTERFACE
	directive @external on FIELD_DEFINITION
	scalar _Any
	scalar _FieldSet
`, BuiltIn: true},
	{Name: "../../federation/entity.graphql", Input: `
type _Service {
  sdl: String
}

extend type Query {
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
