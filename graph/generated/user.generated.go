// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"procurement-be/graph/model"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type UserQueryResolver interface {
	Fetch(ctx context.Context, obj *model.EmptyObject, in *model.FetchRequestInput) (*model.FetchResponse, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_UserQuery_Fetch_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
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

func (ec *executionContext) _FetchResponse_isSuccess(ctx context.Context, field graphql.CollectedField, obj *model.FetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_FetchResponse_isSuccess(ctx, field)
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

func (ec *executionContext) fieldContext_FetchResponse_isSuccess(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "FetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _FetchResponse_message(ctx context.Context, field graphql.CollectedField, obj *model.FetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_FetchResponse_message(ctx, field)
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

func (ec *executionContext) fieldContext_FetchResponse_message(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "FetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _FetchResponse_data(ctx context.Context, field graphql.CollectedField, obj *model.FetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_FetchResponse_data(ctx, field)
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
	res := resTmp.(*model.UserItems)
	fc.Result = res
	return ec.marshalOUserItems2ᚖprocurementᚑbeᚋgraphᚋmodelᚐUserItems(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_FetchResponse_data(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "FetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "items":
				return ec.fieldContext_UserItems_items(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type UserItems", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _FetchResponse_status(ctx context.Context, field graphql.CollectedField, obj *model.FetchResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_FetchResponse_status(ctx, field)
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

func (ec *executionContext) fieldContext_FetchResponse_status(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "FetchResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _UserItems_items(ctx context.Context, field graphql.CollectedField, obj *model.UserItems) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserItems_items(ctx, field)
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
	res := resTmp.([]*model.UserLoginData)
	fc.Result = res
	return ec.marshalOUserLoginData2ᚕᚖprocurementᚑbeᚋgraphᚋmodelᚐUserLoginDataᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserItems_items(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserItems",
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

func (ec *executionContext) _UserQuery_Fetch(ctx context.Context, field graphql.CollectedField, obj *model.EmptyObject) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UserQuery_Fetch(ctx, field)
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
		return ec.resolvers.UserQuery().Fetch(rctx, obj, fc.Args["in"].(*model.FetchRequestInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*model.FetchResponse)
	fc.Result = res
	return ec.marshalOFetchResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐFetchResponse(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UserQuery_Fetch(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UserQuery",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "isSuccess":
				return ec.fieldContext_FetchResponse_isSuccess(ctx, field)
			case "message":
				return ec.fieldContext_FetchResponse_message(ctx, field)
			case "data":
				return ec.fieldContext_FetchResponse_data(ctx, field)
			case "status":
				return ec.fieldContext_FetchResponse_status(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type FetchResponse", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_UserQuery_Fetch_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputFetchRequestInput(ctx context.Context, obj interface{}) (model.FetchRequestInput, error) {
	var it model.FetchRequestInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"limit", "offset", "sortBy", "sort"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "limit":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("limit"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.Limit = data
		case "offset":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("offset"))
			data, err := ec.unmarshalOInt2ᚖint(ctx, v)
			if err != nil {
				return it, err
			}
			it.Offset = data
		case "sortBy":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sortBy"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.SortBy = data
		case "sort":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sort"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Sort = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var fetchResponseImplementors = []string{"FetchResponse"}

func (ec *executionContext) _FetchResponse(ctx context.Context, sel ast.SelectionSet, obj *model.FetchResponse) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, fetchResponseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("FetchResponse")
		case "isSuccess":
			out.Values[i] = ec._FetchResponse_isSuccess(ctx, field, obj)
		case "message":
			out.Values[i] = ec._FetchResponse_message(ctx, field, obj)
		case "data":
			out.Values[i] = ec._FetchResponse_data(ctx, field, obj)
		case "status":
			out.Values[i] = ec._FetchResponse_status(ctx, field, obj)
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

var userItemsImplementors = []string{"UserItems"}

func (ec *executionContext) _UserItems(ctx context.Context, sel ast.SelectionSet, obj *model.UserItems) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userItemsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UserItems")
		case "items":
			out.Values[i] = ec._UserItems_items(ctx, field, obj)
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

var userQueryImplementors = []string{"UserQuery"}

func (ec *executionContext) _UserQuery(ctx context.Context, sel ast.SelectionSet, obj *model.EmptyObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userQueryImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UserQuery")
		case "Fetch":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._UserQuery_Fetch(ctx, field, obj)
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

func (ec *executionContext) marshalNUserQuery2procurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v model.EmptyObject) graphql.Marshaler {
	return ec._UserQuery(ctx, sel, &v)
}

func (ec *executionContext) marshalNUserQuery2ᚖprocurementᚑbeᚋgraphᚋmodelᚐEmptyObject(ctx context.Context, sel ast.SelectionSet, v *model.EmptyObject) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._UserQuery(ctx, sel, v)
}

func (ec *executionContext) unmarshalOFetchRequestInput2ᚖprocurementᚑbeᚋgraphᚋmodelᚐFetchRequestInput(ctx context.Context, v interface{}) (*model.FetchRequestInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputFetchRequestInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOFetchResponse2ᚖprocurementᚑbeᚋgraphᚋmodelᚐFetchResponse(ctx context.Context, sel ast.SelectionSet, v *model.FetchResponse) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._FetchResponse(ctx, sel, v)
}

func (ec *executionContext) marshalOUserItems2ᚖprocurementᚑbeᚋgraphᚋmodelᚐUserItems(ctx context.Context, sel ast.SelectionSet, v *model.UserItems) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._UserItems(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************