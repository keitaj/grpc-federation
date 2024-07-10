// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
// versions:
//
//	protoc-gen-grpc-federation: dev
//
// source: federation/federation.proto
package federation

import (
	"context"
	"io"
	"log/slog"
	"reflect"

	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	grpcfedcel "github.com/mercari/grpc-federation/grpc/federation/cel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	user "example/user"
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Org_Federation_GetResponseArgument is argument for "org.federation.GetResponse" message.
type Org_Federation_GetResponseArgument struct {
	MsgSel *MessageSelection
	Sel    *UserSelection
}

// Org_Federation_MArgument is argument for "org.federation.M" message.
type Org_Federation_MArgument struct {
}

// Org_Federation_MessageSelectionArgument is argument for "org.federation.MessageSelection" message.
type Org_Federation_MessageSelectionArgument struct {
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument struct {
	Bar    string
	Foo    int64
	UserId string
}

// Org_Federation_UserSelectionArgument is argument for "org.federation.UserSelection" message.
type Org_Federation_UserSelectionArgument struct {
	Ua    *User
	Ub    *User
	Uc    *User
	Value string
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// User_UserServiceClient create a gRPC Client to be used to call methods in user.UserService.
	User_UserServiceClient(FederationServiceClientConfig) (user.UserServiceClient, error)
}

// FederationServiceClientConfig helper to create gRPC client.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service FQDN ( `<package-name>.<service-name>` ) of the service on Protocol Buffers.
	Service string
}

// FederationServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependentClientSet struct {
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceCELPluginWasmConfig type alias for grpcfedcel.WasmConfig.
type FederationServiceCELPluginWasmConfig = grpcfedcel.WasmConfig

// FederationServiceCELPluginConfig hints for loading a WebAssembly based plugin.
type FederationServiceCELPluginConfig struct {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

const (
	FederationService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg           FederationServiceConfig
	logger        *slog.Logger
	errorHandler  grpcfed.ErrorHandler
	celCacheMap   *grpcfed.CELCacheMap
	tracer        trace.Tracer
	celTypeHelper *grpcfed.CELTypeHelper
	celEnvOpts    []grpcfed.CELEnvOption
	celPlugins    []*grpcfedcel.CELPlugin
	client        *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, grpcfed.ErrClientConfig
	}
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationServiceClientConfig{
		Service: "user.UserService",
	})
	if err != nil {
		return nil, err
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	celTypeHelperFieldMap := grpcfed.CELTypeHelperFieldMap{
		"grpc.federation.private.GetResponseArgument":      {},
		"grpc.federation.private.MessageSelectionArgument": {},
		"grpc.federation.private.UserArgument": {
			"user_id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "UserId"),
			"foo":     grpcfed.NewCELFieldType(grpcfed.CELIntType, "Foo"),
			"bar":     grpcfed.NewCELFieldType(grpcfed.CELStringType, "Bar"),
		},
		"grpc.federation.private.UserSelectionArgument": {
			"value": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Value"),
		},
		"org.federation.MessageSelection": {
			"message": grpcfed.NewOneofSelectorFieldType(
				grpcfed.CELStringType, "Message",
				[]reflect.Type{reflect.TypeOf((*MessageSelection_MsgA)(nil)), reflect.TypeOf((*MessageSelection_MsgB)(nil)), reflect.TypeOf((*MessageSelection_MsgC)(nil))},
				[]string{"GetMsgA", "GetMsgB", "GetMsgC"},
				reflect.Zero(reflect.TypeOf("")),
			),
		},
		"org.federation.UserSelection": {
			"user": grpcfed.NewOneofSelectorFieldType(
				grpcfed.NewCELObjectType("org.federation.User"), "User",
				[]reflect.Type{reflect.TypeOf((*UserSelection_UserA)(nil)), reflect.TypeOf((*UserSelection_UserB)(nil)), reflect.TypeOf((*UserSelection_UserC)(nil))},
				[]string{"GetUserA", "GetUserB", "GetUserC"},
				reflect.Zero(reflect.TypeOf((*User)(nil))),
			),
		},
	}
	celTypeHelper := grpcfed.NewCELTypeHelper(celTypeHelperFieldMap)
	var celEnvOpts []grpcfed.CELEnvOption
	celEnvOpts = append(celEnvOpts, grpcfed.NewDefaultEnvOptions(celTypeHelper)...)
	return &FederationService{
		cfg:           cfg,
		logger:        logger,
		errorHandler:  errorHandler,
		celEnvOpts:    celEnvOpts,
		celTypeHelper: celTypeHelper,
		celCacheMap:   grpcfed.NewCELCacheMap(),
		tracer:        otel.Tracer("org.federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

// Get implements "org.federation.FederationService/Get" method.
func (s *FederationService) Get(ctx context.Context, req *GetRequest) (res *GetResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/Get")
	defer span.End()
	ctx = grpcfed.WithLogger(ctx, s.logger)
	ctx = grpcfed.WithCELCacheMap(ctx, s.celCacheMap)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, grpcfed.StackTrace())
			grpcfed.OutputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetResponse(ctx, &Org_Federation_GetResponseArgument{})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetResponse resolve "org.federation.GetResponse" message.
func (s *FederationService) resolve_Org_Federation_GetResponse(ctx context.Context, req *Org_Federation_GetResponseArgument) (*GetResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.GetResponse")
	defer span.End()
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx), grpcfed.LogAttrs(ctx)...)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.GetResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			msg_sel *MessageSelection
			sel     *UserSelection
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.celEnvOpts, s.celPlugins, false, "grpc.federation.private.GetResponseArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()
	// A tree view of message dependencies is shown below.
	/*
	   msg_sel ─┐
	       sel ─┤
	*/
	eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "msg_sel"
		     message {
		       name: "MessageSelection"
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*MessageSelection, *localValueType]{
			Name: `msg_sel`,
			Type: grpcfed.CELObjectType("org.federation.MessageSelection"),
			Setter: func(value *localValueType, v *MessageSelection) error {
				value.vars.msg_sel = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_MessageSelectionArgument{}
				return s.resolve_Org_Federation_MessageSelection(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "sel"
		     message {
		       name: "UserSelection"
		       args { name: "value", by: "'foo'" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*UserSelection, *localValueType]{
			Name: `sel`,
			Type: grpcfed.CELObjectType("org.federation.UserSelection"),
			Setter: func(value *localValueType, v *UserSelection) error {
				value.vars.sel = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_UserSelectionArgument{}
				// { name: "value", by: "'foo'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'foo'`,
					UseContextLibrary: false,
					CacheIndex:        1,
					Setter: func(v string) error {
						args.Value = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Org_Federation_UserSelection(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.MsgSel = value.vars.msg_sel
	req.Sel = value.vars.sel

	// create a message value to be returned.
	ret := &GetResponse{}

	// field binding section.
	// (grpc.federation.field).by = "sel.user"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*User]{
		Value:             value,
		Expr:              `sel.user`,
		UseContextLibrary: false,
		CacheIndex:        2,
		Setter: func(v *User) error {
			ret.User = v
			return nil
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "msg_sel.message"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
		Value:             value,
		Expr:              `msg_sel.message`,
		UseContextLibrary: false,
		CacheIndex:        3,
		Setter: func(v string) error {
			ret.Msg = v
			return nil
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.GetResponse", slog.Any("org.federation.GetResponse", s.logvalue_Org_Federation_GetResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_MessageSelection resolve "org.federation.MessageSelection" message.
func (s *FederationService) resolve_Org_Federation_MessageSelection(ctx context.Context, req *Org_Federation_MessageSelectionArgument) (*MessageSelection, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.MessageSelection")
	defer span.End()
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx), grpcfed.LogAttrs(ctx)...)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.MessageSelection", slog.Any("message_args", s.logvalue_Org_Federation_MessageSelectionArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.celEnvOpts, s.celPlugins, false, "grpc.federation.private.MessageSelectionArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// create a message value to be returned.
	ret := &MessageSelection{}

	// field binding section.
	oneof_MsgA, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
		Value:             value,
		Expr:              `false`,
		UseContextLibrary: false,
		OutType:           reflect.TypeOf(true),
		CacheIndex:        4,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	oneof_MsgB, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
		Value:             value,
		Expr:              `true`,
		UseContextLibrary: false,
		OutType:           reflect.TypeOf(true),
		CacheIndex:        5,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	switch {
	case oneof_MsgA.(bool):

		if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
			Value:             value,
			Expr:              `'aaa'`,
			UseContextLibrary: false,
			CacheIndex:        6,
			Setter: func(v string) error {
				ret.Message = &MessageSelection_MsgA{MsgA: v}
				return nil
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	case oneof_MsgB.(bool):

		if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
			Value:             value,
			Expr:              `'bbb'`,
			UseContextLibrary: false,
			CacheIndex:        7,
			Setter: func(v string) error {
				ret.Message = &MessageSelection_MsgB{MsgB: v}
				return nil
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	default:

		if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
			Value:             value,
			Expr:              `'ccc'`,
			UseContextLibrary: false,
			CacheIndex:        8,
			Setter: func(v string) error {
				ret.Message = &MessageSelection_MsgC{MsgC: v}
				return nil
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.MessageSelection", slog.Any("org.federation.MessageSelection", s.logvalue_Org_Federation_MessageSelection(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.User")
	defer span.End()
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx), grpcfed.LogAttrs(ctx)...)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			_def0 *user.GetUserResponse
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.celEnvOpts, s.celPlugins, false, "grpc.federation.private.UserArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "_def0"
	     call {
	       method: "user.UserService/GetUser"
	       request: [
	         { field: "id", by: "$.user_id" },
	         { field: "foo", by: "$.foo", if: "$.foo != 0" },
	         { field: "bar", by: "$.bar", if: "$.bar != ''" }
	       ]
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.GetUserResponse, *localValueType]{
		Name: `_def0`,
		Type: grpcfed.CELObjectType("user.GetUserResponse"),
		Setter: func(value *localValueType, v *user.GetUserResponse) error {
			value.vars._def0 = v
			return nil
		},
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &user.GetUserRequest{}
			// { field: "id", by: "$.user_id" }
			if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
				Value:             value,
				Expr:              `$.user_id`,
				UseContextLibrary: false,
				CacheIndex:        9,
				Setter: func(v string) error {
					args.Id = v
					return nil
				},
			}); err != nil {
				return nil, err
			}
			// { field: "foo", by: "$.foo", if: "$.foo != 0" }
			if err := grpcfed.If(ctx, &grpcfed.IfParam[*localValueType]{
				Value:             value,
				Expr:              `$.foo != 0`,
				UseContextLibrary: false,
				CacheIndex:        10,
				Body: func(value *localValueType) error {
					return grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[int64]{
						Value:             value,
						Expr:              `$.foo`,
						UseContextLibrary: false,
						CacheIndex:        11,
						Setter: func(v int64) error {
							args.Foobar = &user.GetUserRequest_Foo{
								Foo: v,
							}
							return nil
						},
					})
				},
			}); err != nil {
				return nil, err
			}
			// { field: "bar", by: "$.bar", if: "$.bar != ”" }
			if err := grpcfed.If(ctx, &grpcfed.IfParam[*localValueType]{
				Value:             value,
				Expr:              `$.bar != ''`,
				UseContextLibrary: false,
				CacheIndex:        12,
				Body: func(value *localValueType) error {
					return grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
						Value:             value,
						Expr:              `$.bar`,
						UseContextLibrary: false,
						CacheIndex:        13,
						Setter: func(v string) error {
							args.Foobar = &user.GetUserRequest_Bar{
								Bar: v,
							}
							return nil
						},
					})
				},
			}); err != nil {
				return nil, err
			}
			grpcfed.Logger(ctx).DebugContext(ctx, "call user.UserService/GetUser", slog.Any("user.GetUserRequest", s.logvalue_User_GetUserRequest(args)))
			return s.client.User_UserServiceClient.GetUser(ctx, args)
		},
	}); err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, grpcfed.NewErrorWithLogAttrs(err, grpcfed.LogAttrs(ctx))
		}
	}

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.user_id"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
		Value:             value,
		Expr:              `$.user_id`,
		UseContextLibrary: false,
		CacheIndex:        14,
		Setter: func(v string) error {
			ret.Id = v
			return nil
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.User", slog.Any("org.federation.User", s.logvalue_Org_Federation_User(ret)))
	return ret, nil
}

// resolve_Org_Federation_UserSelection resolve "org.federation.UserSelection" message.
func (s *FederationService) resolve_Org_Federation_UserSelection(ctx context.Context, req *Org_Federation_UserSelectionArgument) (*UserSelection, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.UserSelection")
	defer span.End()
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx), grpcfed.LogAttrs(ctx)...)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.UserSelection", slog.Any("message_args", s.logvalue_Org_Federation_UserSelectionArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			ua *User
			ub *User
			uc *User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.celEnvOpts, s.celPlugins, false, "grpc.federation.private.UserSelectionArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Ua = value.vars.ua
	req.Ub = value.vars.ub
	req.Uc = value.vars.uc

	// create a message value to be returned.
	ret := &UserSelection{}

	// field binding section.
	oneof_UserA, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
		Value:             value,
		Expr:              `false`,
		UseContextLibrary: false,
		OutType:           reflect.TypeOf(true),
		CacheIndex:        15,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	oneof_UserB, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
		Value:             value,
		Expr:              `true`,
		UseContextLibrary: false,
		OutType:           reflect.TypeOf(true),
		CacheIndex:        16,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	switch {
	case oneof_UserA.(bool):

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "ua"
		     message {
		       name: "User"
		       args: [
		         { name: "user_id", by: "'a'" },
		         { name: "foo", by: "0" },
		         { name: "bar", by: "'hello'" }
		       ]
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*User, *localValueType]{
			Name: `ua`,
			Type: grpcfed.CELObjectType("org.federation.User"),
			Setter: func(value *localValueType, v *User) error {
				value.vars.ua = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_UserArgument{}
				// { name: "user_id", by: "'a'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'a'`,
					UseContextLibrary: false,
					CacheIndex:        17,
					Setter: func(v string) error {
						args.UserId = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				// { name: "foo", by: "0" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[int64]{
					Value:             value,
					Expr:              `0`,
					UseContextLibrary: false,
					CacheIndex:        18,
					Setter: func(v int64) error {
						args.Foo = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				// { name: "bar", by: "'hello'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'hello'`,
					UseContextLibrary: false,
					CacheIndex:        19,
					Setter: func(v string) error {
						args.Bar = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Org_Federation_User(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*User]{
			Value:             value,
			Expr:              `ua`,
			UseContextLibrary: false,
			CacheIndex:        20,
			Setter: func(v *User) error {
				ret.User = &UserSelection_UserA{UserA: v}
				return nil
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	case oneof_UserB.(bool):

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "ub"
		     message {
		       name: "User"
		       args: [
		         { name: "user_id", by: "'b'" },
		         { name: "foo", by: "0" },
		         { name: "bar", by: "'hello'" }
		       ]
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*User, *localValueType]{
			Name: `ub`,
			Type: grpcfed.CELObjectType("org.federation.User"),
			Setter: func(value *localValueType, v *User) error {
				value.vars.ub = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_UserArgument{}
				// { name: "user_id", by: "'b'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'b'`,
					UseContextLibrary: false,
					CacheIndex:        21,
					Setter: func(v string) error {
						args.UserId = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				// { name: "foo", by: "0" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[int64]{
					Value:             value,
					Expr:              `0`,
					UseContextLibrary: false,
					CacheIndex:        22,
					Setter: func(v int64) error {
						args.Foo = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				// { name: "bar", by: "'hello'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'hello'`,
					UseContextLibrary: false,
					CacheIndex:        23,
					Setter: func(v string) error {
						args.Bar = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Org_Federation_User(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*User]{
			Value:             value,
			Expr:              `ub`,
			UseContextLibrary: false,
			CacheIndex:        24,
			Setter: func(v *User) error {
				ret.User = &UserSelection_UserB{UserB: v}
				return nil
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	default:

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "uc"
		     message {
		       name: "User"
		       args: [
		         { name: "user_id", by: "$.value" },
		         { name: "foo", by: "0" },
		         { name: "bar", by: "'hello'" }
		       ]
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*User, *localValueType]{
			Name: `uc`,
			Type: grpcfed.CELObjectType("org.federation.User"),
			Setter: func(value *localValueType, v *User) error {
				value.vars.uc = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_UserArgument{}
				// { name: "user_id", by: "$.value" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `$.value`,
					UseContextLibrary: false,
					CacheIndex:        25,
					Setter: func(v string) error {
						args.UserId = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				// { name: "foo", by: "0" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[int64]{
					Value:             value,
					Expr:              `0`,
					UseContextLibrary: false,
					CacheIndex:        26,
					Setter: func(v int64) error {
						args.Foo = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				// { name: "bar", by: "'hello'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'hello'`,
					UseContextLibrary: false,
					CacheIndex:        27,
					Setter: func(v string) error {
						args.Bar = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Org_Federation_User(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*User]{
			Value:             value,
			Expr:              `uc`,
			UseContextLibrary: false,
			CacheIndex:        28,
			Setter: func(v *User) error {
				ret.User = &UserSelection_UserC{UserC: v}
				return nil
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.UserSelection", slog.Any("org.federation.UserSelection", s.logvalue_Org_Federation_UserSelection(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_GetResponse(v *GetResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("user", s.logvalue_Org_Federation_User(v.GetUser())),
		slog.String("msg", v.GetMsg()),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetResponseArgument(v *Org_Federation_GetResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Org_Federation_MessageSelection(v *MessageSelection) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("msg_a", v.GetMsgA()),
		slog.String("msg_b", v.GetMsgB()),
		slog.String("msg_c", v.GetMsgC()),
	)
}

func (s *FederationService) logvalue_Org_Federation_MessageSelectionArgument(v *Org_Federation_MessageSelectionArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Org_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserArgument(v *Org_Federation_UserArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("user_id", v.UserId),
		slog.Int64("foo", v.Foo),
		slog.String("bar", v.Bar),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserSelection(v *UserSelection) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("user_a", s.logvalue_Org_Federation_User(v.GetUserA())),
		slog.Any("user_b", s.logvalue_Org_Federation_User(v.GetUserB())),
		slog.Any("user_c", s.logvalue_Org_Federation_User(v.GetUserC())),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserSelectionArgument(v *Org_Federation_UserSelectionArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("value", v.Value),
	)
}

func (s *FederationService) logvalue_User_GetUserRequest(v *user.GetUserRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.Int64("foo", v.GetFoo()),
		slog.String("bar", v.GetBar()),
	)
}
