// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"procurement-be/graph/model"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type AuthMutationResolver interface {
	Login(ctx context.Context, obj *model.EmptyObject, in *model.LoginRequestInput) (*model.LoginResponse, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_AuthMutation_Login_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *model.LoginRequestInput
	if tmp, ok := rawArgs["in"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("in"))
		arg0, err = ec.unmarshalOLoginRequestInput2ᚖprocurementᚑbeᚋgraphᚋmodelᚐLoginRequestInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["in"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _AuthMutation_Login(ctx context.Context, field graphql.CollectedField, obj *model.EmptyObject) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_AuthMutation_Login(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.AuthMutation().Login(rctx, obj, fc.Args["in"].(*model.LoginRequestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.LoginResponse)
	fc.Result = res
	return ec.marshalOLoginResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐLoginResponse(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_AuthMutation_Login(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "AuthMutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "isSuccess":
				return ec.fieldContext_LoginResponse_isSuccess(ctx, field)
			case "message":
				return ec.fieldContext_LoginResponse_message(ctx, field)
			case "data":
				return ec.fieldContext_LoginResponse_data(ctx, field)
			case "status":
				return ec.fieldContext_LoginResponse_status(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type LoginResponse", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_AuthMutation_Login_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Data_token(ctx context.Context, field graphql.CollectedField, obj *model.Data) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Data_token(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Token, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Data_token(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Data",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Data_refreshToken(ctx context.Context, field graphql.CollectedField, obj *model.Data) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Data_refreshToken(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.RefreshToken, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Data_refreshToken(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Data",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Data_userData(ctx context.Context, field graphql.CollectedField, obj *model.Data) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Data_userData(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UserData, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.UserLoginData)
	fc.Result = res
	return ec.marshalOUserLoginData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐUserLoginData(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Data_userData(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Data",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_UserLoginData_id(ctx, field)
			case "username":
				return ec.fieldContext_UserLoginData_username(ctx, field)
			case "role":
				return ec.fieldContext_UserLoginData_role(ctx, field)
			case "email":
				return ec.fieldContext_UserLoginData_email(ctx, field)
			case "name":
				return ec.fieldContext_UserLoginData_name(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type UserLoginData", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _LoginResponse_isSuccess(ctx context.Context, field graphql.CollectedField, obj *model.LoginResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LoginResponse_isSuccess(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsSuccess, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LoginResponse_isSuccess(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LoginResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _LoginResponse_message(ctx context.Context, field graphql.CollectedField, obj *model.LoginResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LoginResponse_message(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Message, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LoginResponse_message(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LoginResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _LoginResponse_data(ctx context.Context, field graphql.CollectedField, obj *model.LoginResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LoginResponse_data(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Data, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.Data)
	fc.Result = res
	return ec.marshalOData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐData(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LoginResponse_data(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LoginResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "token":
				return ec.fieldContext_Data_token(ctx, field)
			case "refreshToken":
				return ec.fieldContext_Data_refreshToken(ctx, field)
			case "userData":
				return ec.fieldContext_Data_userData(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Data", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _LoginResponse_status(ctx context.Context, field graphql.CollectedField, obj *model.LoginResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LoginResponse_status(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Status, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	fc.Result = res
	return ec.marshalOInt2ᚖint(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LoginResponse_status(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LoginResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserLoginData_id(ctx context.Context, field graphql.CollectedField, obj *model.UserLoginData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserLoginData_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserLoginData_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserLoginData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserLoginData_username(ctx context.Context, field graphql.CollectedField, obj *model.UserLoginData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserLoginData_username(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Username, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserLoginData_username(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserLoginData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserLoginData_role(ctx context.Context, field graphql.CollectedField, obj *model.UserLoginData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserLoginData_role(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Role, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserLoginData_role(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserLoginData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserLoginData_email(ctx context.Context, field graphql.CollectedField, obj *model.UserLoginData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserLoginData_email(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Email, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserLoginData_email(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserLoginData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserLoginData_name(ctx context.Context, field graphql.CollectedField, obj *model.UserLoginData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserLoginData_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserLoginData_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserLoginData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputLoginRequestInput(ctx context.Context, obj interface{}) (model.LoginRequestInput, error) {
	var it model.LoginRequestInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"username", "password"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "username":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("username"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Username = data
		case "password":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Password = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var authMutationImplementors = []string{"AuthMutation"}

func (ec *executionContext) _AuthMutation(ctx context.Context, sel ast.SelectionSet, obj *model.EmptyObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, authMutationImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("AuthMutation")
		case "Login":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._AuthMutation_Login(ctx, field, obj)
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var dataImplementors = []string{"Data"}

func (ec *executionContext) _Data(ctx context.Context, sel ast.SelectionSet, obj *model.Data) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, dataImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Data")
		case "token":
			out.Values[i] = ec._Data_token(ctx, field, obj)
		case "refreshToken":
			out.Values[i] = ec._Data_refreshToken(ctx, field, obj)
		case "userData":
			out.Values[i] = ec._Data_userData(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var loginResponseImplementors = []string{"LoginResponse"}

func (ec *executionContext) _LoginResponse(ctx context.Context, sel ast.SelectionSet, obj *model.LoginResponse) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, loginResponseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("LoginResponse")
		case "isSuccess":
			out.Values[i] = ec._LoginResponse_isSuccess(ctx, field, obj)
		case "message":
			out.Values[i] = ec._LoginResponse_message(ctx, field, obj)
		case "data":
			out.Values[i] = ec._LoginResponse_data(ctx, field, obj)
		case "status":
			out.Values[i] = ec._LoginResponse_status(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var userLoginDataImplementors = []string{"UserLoginData"}

func (ec *executionContext) _UserLoginData(ctx context.Context, sel ast.SelectionSet, obj *model.UserLoginData) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userLoginDataImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UserLoginData")
		case "id":
			out.Values[i] = ec._UserLoginData_id(ctx, field, obj)
		case "username":
			out.Values[i] = ec._UserLoginData_username(ctx, field, obj)
		case "role":
			out.Values[i] = ec._UserLoginData_role(ctx, field, obj)
		case "email":
			out.Values[i] = ec._UserLoginData_email(ctx, field, obj)
		case "name":
			out.Values[i] = ec._UserLoginData_name(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNAuthMutation2procurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v model.EmptyObject) graphql.Marshaler {
	return ec._AuthMutation(ctx, sel, &v)
}

func (ec *executionContext) marshalNAuthMutation2ᚖprocurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v *model.EmptyObject) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._AuthMutation(ctx, sel, v)
}

func (ec *executionContext) marshalNUserLoginData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐUserLoginData(ctx context.Context, sel ast.SelectionSet, v *model.UserLoginData) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._UserLoginData(ctx, sel, v)
}

func (ec *executionContext) marshalOData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐData(ctx context.Context, sel ast.SelectionSet, v *model.Data) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Data(ctx, sel, v)
}

func (ec *executionContext) unmarshalOLoginRequestInput2ᚖprocurementᚑbeᚋgraphᚋmodelᚐLoginRequestInput(ctx context.Context, v interface{}) (*model.LoginRequestInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputLoginRequestInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOLoginResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐLoginResponse(ctx context.Context, sel ast.SelectionSet, v *model.LoginResponse) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._LoginResponse(ctx, sel, v)
}

func (ec *executionContext) marshalOUserLoginData2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐUserLoginDataᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.UserLoginData) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNUserLoginData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐUserLoginData(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalOUserLoginData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐUserLoginData(ctx context.Context, sel ast.SelectionSet, v *model.UserLoginData) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._UserLoginData(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
