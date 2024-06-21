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

type RoleMutationResolver interface {
	Store(ctx context.Context, obj *model.EmptyObject, in []*model.RoleDataInput) (*model.RoleFetchResponse, error)
	Update(ctx context.Context, obj *model.EmptyObject, in []*model.RoleDataInput) (*model.RoleFetchResponse, error)
}
type RoleQueryResolver interface {
	Fetch(ctx context.Context, obj *model.EmptyObject, in *model.FetchRequestInput) (*model.RoleFetchResponse, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_RoleMutation_Store_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []*model.RoleDataInput
	if tmp, ok := rawArgs["in"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("in"))
		arg0, err = ec.unmarshalORoleDataInput2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataInputᚄ(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["in"] = arg0
	return args, nil
}

func (ec *executionContext) field_RoleMutation_Update_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []*model.RoleDataInput
	if tmp, ok := rawArgs["in"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("in"))
		arg0, err = ec.unmarshalORoleDataInput2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataInputᚄ(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["in"] = arg0
	return args, nil
}

func (ec *executionContext) field_RoleQuery_Fetch_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *model.FetchRequestInput
	if tmp, ok := rawArgs["in"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("in"))
		arg0, err = ec.unmarshalOFetchRequestInput2ᚖprocurementᚑbeᚋgraphᚋmodelᚐFetchRequestInput(ctx, tmp)
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

func (ec *executionContext) _RoleData_id(ctx context.Context, field graphql.CollectedField, obj *model.RoleData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleData_id(ctx, field)
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

func (ec *executionContext) fieldContext_RoleData_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleData_name(ctx context.Context, field graphql.CollectedField, obj *model.RoleData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleData_name(ctx, field)
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

func (ec *executionContext) fieldContext_RoleData_name(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleData_createdAt(ctx context.Context, field graphql.CollectedField, obj *model.RoleData) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleData_createdAt(ctx, field)
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
		return obj.CreatedAt, nil
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

func (ec *executionContext) fieldContext_RoleData_createdAt(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleData",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleFetchResponse_isSuccess(ctx context.Context, field graphql.CollectedField, obj *model.RoleFetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleFetchResponse_isSuccess(ctx, field)
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

func (ec *executionContext) fieldContext_RoleFetchResponse_isSuccess(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleFetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleFetchResponse_message(ctx context.Context, field graphql.CollectedField, obj *model.RoleFetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleFetchResponse_message(ctx, field)
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

func (ec *executionContext) fieldContext_RoleFetchResponse_message(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleFetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleFetchResponse_data(ctx context.Context, field graphql.CollectedField, obj *model.RoleFetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleFetchResponse_data(ctx, field)
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
	res := resTmp.(*model.RoleItems)
	fc.Result = res
	return ec.marshalORoleItems2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleItems(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_RoleFetchResponse_data(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleFetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "items":
				return ec.fieldContext_RoleItems_items(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type RoleItems", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleFetchResponse_status(ctx context.Context, field graphql.CollectedField, obj *model.RoleFetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleFetchResponse_status(ctx, field)
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

func (ec *executionContext) fieldContext_RoleFetchResponse_status(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleFetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleItems_items(ctx context.Context, field graphql.CollectedField, obj *model.RoleItems) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleItems_items(ctx, field)
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
		return obj.Items, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.RoleData)
	fc.Result = res
	return ec.marshalORoleData2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_RoleItems_items(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleItems",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_RoleData_id(ctx, field)
			case "name":
				return ec.fieldContext_RoleData_name(ctx, field)
			case "createdAt":
				return ec.fieldContext_RoleData_createdAt(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type RoleData", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _RoleMutation_Store(ctx context.Context, field graphql.CollectedField, obj *model.EmptyObject) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleMutation_Store(ctx, field)
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
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.RoleMutation().Store(rctx, obj, fc.Args["in"].([]*model.RoleDataInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			loggedIn, err := ec.unmarshalNlogin2procurementᚑbeᚋgraphᚋmodelᚐLogin(ctx, map[string]interface{}{"access": "Authenticated"})
			if err != nil {
				return nil, err
			}
			if ec.directives.LoggedIn == nil {
				return nil, errors.New("directive loggedIn is not implemented")
			}
			return ec.directives.LoggedIn(ctx, obj, directive0, loggedIn)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*model.RoleFetchResponse); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *procurement-be/graph/model.RoleFetchResponse`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.RoleFetchResponse)
	fc.Result = res
	return ec.marshalORoleFetchResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleFetchResponse(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_RoleMutation_Store(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleMutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "isSuccess":
				return ec.fieldContext_RoleFetchResponse_isSuccess(ctx, field)
			case "message":
				return ec.fieldContext_RoleFetchResponse_message(ctx, field)
			case "data":
				return ec.fieldContext_RoleFetchResponse_data(ctx, field)
			case "status":
				return ec.fieldContext_RoleFetchResponse_status(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type RoleFetchResponse", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_RoleMutation_Store_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _RoleMutation_Update(ctx context.Context, field graphql.CollectedField, obj *model.EmptyObject) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleMutation_Update(ctx, field)
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
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.RoleMutation().Update(rctx, obj, fc.Args["in"].([]*model.RoleDataInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			loggedIn, err := ec.unmarshalNlogin2procurementᚑbeᚋgraphᚋmodelᚐLogin(ctx, map[string]interface{}{"access": "Authenticated"})
			if err != nil {
				return nil, err
			}
			if ec.directives.LoggedIn == nil {
				return nil, errors.New("directive loggedIn is not implemented")
			}
			return ec.directives.LoggedIn(ctx, obj, directive0, loggedIn)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*model.RoleFetchResponse); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *procurement-be/graph/model.RoleFetchResponse`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.RoleFetchResponse)
	fc.Result = res
	return ec.marshalORoleFetchResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleFetchResponse(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_RoleMutation_Update(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleMutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "isSuccess":
				return ec.fieldContext_RoleFetchResponse_isSuccess(ctx, field)
			case "message":
				return ec.fieldContext_RoleFetchResponse_message(ctx, field)
			case "data":
				return ec.fieldContext_RoleFetchResponse_data(ctx, field)
			case "status":
				return ec.fieldContext_RoleFetchResponse_status(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type RoleFetchResponse", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_RoleMutation_Update_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _RoleQuery_Fetch(ctx context.Context, field graphql.CollectedField, obj *model.EmptyObject) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_RoleQuery_Fetch(ctx, field)
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
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return ec.resolvers.RoleQuery().Fetch(rctx, obj, fc.Args["in"].(*model.FetchRequestInput))
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			loggedIn, err := ec.unmarshalNlogin2procurementᚑbeᚋgraphᚋmodelᚐLogin(ctx, map[string]interface{}{"access": "Authenticated"})
			if err != nil {
				return nil, err
			}
			if ec.directives.LoggedIn == nil {
				return nil, errors.New("directive loggedIn is not implemented")
			}
			return ec.directives.LoggedIn(ctx, obj, directive0, loggedIn)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(*model.RoleFetchResponse); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be *procurement-be/graph/model.RoleFetchResponse`, tmp)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.RoleFetchResponse)
	fc.Result = res
	return ec.marshalORoleFetchResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleFetchResponse(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_RoleQuery_Fetch(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "RoleQuery",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "isSuccess":
				return ec.fieldContext_RoleFetchResponse_isSuccess(ctx, field)
			case "message":
				return ec.fieldContext_RoleFetchResponse_message(ctx, field)
			case "data":
				return ec.fieldContext_RoleFetchResponse_data(ctx, field)
			case "status":
				return ec.fieldContext_RoleFetchResponse_status(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type RoleFetchResponse", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_RoleQuery_Fetch_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputRoleDataInput(ctx context.Context, obj interface{}) (model.RoleDataInput, error) {
	var it model.RoleDataInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "name"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var roleDataImplementors = []string{"RoleData"}

func (ec *executionContext) _RoleData(ctx context.Context, sel ast.SelectionSet, obj *model.RoleData) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, roleDataImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("RoleData")
		case "id":
			out.Values[i] = ec._RoleData_id(ctx, field, obj)
		case "name":
			out.Values[i] = ec._RoleData_name(ctx, field, obj)
		case "createdAt":
			out.Values[i] = ec._RoleData_createdAt(ctx, field, obj)
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

var roleFetchResponseImplementors = []string{"RoleFetchResponse"}

func (ec *executionContext) _RoleFetchResponse(ctx context.Context, sel ast.SelectionSet, obj *model.RoleFetchResponse) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, roleFetchResponseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("RoleFetchResponse")
		case "isSuccess":
			out.Values[i] = ec._RoleFetchResponse_isSuccess(ctx, field, obj)
		case "message":
			out.Values[i] = ec._RoleFetchResponse_message(ctx, field, obj)
		case "data":
			out.Values[i] = ec._RoleFetchResponse_data(ctx, field, obj)
		case "status":
			out.Values[i] = ec._RoleFetchResponse_status(ctx, field, obj)
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

var roleItemsImplementors = []string{"RoleItems"}

func (ec *executionContext) _RoleItems(ctx context.Context, sel ast.SelectionSet, obj *model.RoleItems) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, roleItemsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("RoleItems")
		case "items":
			out.Values[i] = ec._RoleItems_items(ctx, field, obj)
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

var roleMutationImplementors = []string{"RoleMutation"}

func (ec *executionContext) _RoleMutation(ctx context.Context, sel ast.SelectionSet, obj *model.EmptyObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, roleMutationImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("RoleMutation")
		case "Store":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._RoleMutation_Store(ctx, field, obj)
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
		case "Update":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._RoleMutation_Update(ctx, field, obj)
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

var roleQueryImplementors = []string{"RoleQuery"}

func (ec *executionContext) _RoleQuery(ctx context.Context, sel ast.SelectionSet, obj *model.EmptyObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, roleQueryImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("RoleQuery")
		case "Fetch":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._RoleQuery_Fetch(ctx, field, obj)
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

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNRoleData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleData(ctx context.Context, sel ast.SelectionSet, v *model.RoleData) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._RoleData(ctx, sel, v)
}

func (ec *executionContext) unmarshalNRoleDataInput2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataInput(ctx context.Context, v interface{}) (*model.RoleDataInput, error) {
	res, err := ec.unmarshalInputRoleDataInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNRoleMutation2procurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v model.EmptyObject) graphql.Marshaler {
	return ec._RoleMutation(ctx, sel, &v)
}

func (ec *executionContext) marshalNRoleMutation2ᚖprocurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v *model.EmptyObject) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._RoleMutation(ctx, sel, v)
}

func (ec *executionContext) marshalNRoleQuery2procurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v model.EmptyObject) graphql.Marshaler {
	return ec._RoleQuery(ctx, sel, &v)
}

func (ec *executionContext) marshalNRoleQuery2ᚖprocurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v *model.EmptyObject) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._RoleQuery(ctx, sel, v)
}

func (ec *executionContext) marshalORoleData2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.RoleData) graphql.Marshaler {
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
			ret[i] = ec.marshalNRoleData2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleData(ctx, sel, v[i])
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

func (ec *executionContext) unmarshalORoleDataInput2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataInputᚄ(ctx context.Context, v interface{}) ([]*model.RoleDataInput, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.RoleDataInput, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNRoleDataInput2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleDataInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalORoleFetchResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleFetchResponse(ctx context.Context, sel ast.SelectionSet, v *model.RoleFetchResponse) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._RoleFetchResponse(ctx, sel, v)
}

func (ec *executionContext) marshalORoleItems2ᚖprocurementᚑbeᚋgraphᚋmodelᚐRoleItems(ctx context.Context, sel ast.SelectionSet, v *model.RoleItems) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._RoleItems(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
