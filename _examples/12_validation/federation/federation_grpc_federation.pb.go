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
	"runtime/debug"

	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	grpcfedcel "github.com/mercari/grpc-federation/grpc/federation/cel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Org_Federation_CustomHandlerMessageArgument is argument for "org.federation.CustomHandlerMessage" message.
type Org_Federation_CustomHandlerMessageArgument struct {
	Arg string
}

// Org_Federation_CustomMessageArgument is argument for "org.federation.CustomMessage" message.
type Org_Federation_CustomMessageArgument struct {
	Message string
}

// Org_Federation_GetPostResponseArgument is argument for "org.federation.GetPostResponse" message.
type Org_Federation_GetPostResponseArgument struct {
	Id                  string
	Post                *Post
	XDef4ErrDetail0Msg0 *CustomMessage
	XDef4ErrDetail0Msg1 *CustomMessage
}

// Org_Federation_PostArgument is argument for "org.federation.Post" message.
type Org_Federation_PostArgument struct {
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Resolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
	// If this interface is not provided, an error is returned during initialization.
	Resolver FederationServiceResolver // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
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
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
	// Resolve_Org_Federation_CustomHandlerMessage implements resolver for "org.federation.CustomHandlerMessage".
	Resolve_Org_Federation_CustomHandlerMessage(context.Context, *Org_Federation_CustomHandlerMessageArgument) (*CustomHandlerMessage, error)
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

// Resolve_Org_Federation_CustomHandlerMessage resolve "org.federation.CustomHandlerMessage".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Org_Federation_CustomHandlerMessage(context.Context, *Org_Federation_CustomHandlerMessageArgument) (ret *CustomHandlerMessage, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Org_Federation_CustomHandlerMessage not implemented")
	return
}

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg           FederationServiceConfig
	logger        *slog.Logger
	errorHandler  grpcfed.ErrorHandler
	celCacheMap   *grpcfed.CELCacheMap
	tracer        trace.Tracer
	resolver      FederationServiceResolver
	celTypeHelper *grpcfed.CELTypeHelper
	envOpts       []grpcfed.CELEnvOption
	celPlugins    []*grpcfedcel.CELPlugin
	client        *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Resolver == nil {
		return nil, grpcfed.ErrResolverConfig
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
		"grpc.federation.private.CustomHandlerMessageArgument": {
			"arg": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Arg"),
		},
		"grpc.federation.private.CustomMessageArgument": {
			"message": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Message"),
		},
		"grpc.federation.private.GetPostResponseArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {},
	}
	celTypeHelper := grpcfed.NewCELTypeHelper(celTypeHelperFieldMap)
	var envOpts []grpcfed.CELEnvOption
	envOpts = append(envOpts, grpcfed.NewDefaultEnvOptions(celTypeHelper)...)
	return &FederationService{
		cfg:           cfg,
		logger:        logger,
		errorHandler:  errorHandler,
		envOpts:       envOpts,
		celTypeHelper: celTypeHelper,
		celCacheMap:   grpcfed.NewCELCacheMap(),
		tracer:        otel.Tracer("org.federation.FederationService"),
		resolver:      cfg.Resolver,
		client:        &FederationServiceDependentClientSet{},
	}, nil
}

// GetPost implements "org.federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/GetPost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	ctx = grpcfed.WithCELCacheMap(ctx, s.celCacheMap)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetPostResponse(ctx, &Org_Federation_GetPostResponseArgument{
		Id: req.Id,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_CustomHandlerMessage resolve "org.federation.CustomHandlerMessage" message.
func (s *FederationService) resolve_Org_Federation_CustomHandlerMessage(ctx context.Context, req *Org_Federation_CustomHandlerMessageArgument) (*CustomHandlerMessage, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.CustomHandlerMessage")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.CustomHandlerMessage", slog.Any("message_args", s.logvalue_Org_Federation_CustomHandlerMessageArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			_def0 bool
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.CustomHandlerMessageArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "_def0"
	     validation {
	       error {
	         code: FAILED_PRECONDITION
	         if: "$.arg == 'wrong'"
	       }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[bool, *localValueType]{
		Name:   "_def0",
		Type:   grpcfed.CELBoolType,
		Setter: func(value *localValueType, v bool) { value.vars._def0 = v },
		Validation: func(ctx context.Context, value *localValueType) error {
			var stat *grpcfed.Status
			if err := grpcfed.If(ctx, &grpcfed.IfParam[*localValueType]{
				Value:             value,
				Expr:              "$.arg == 'wrong'",
				UseContextLibrary: false,
				CacheIndex:        1,
				Body: func(value *localValueType) error {
					errorMessage := "error"
					stat = grpcfed.NewGRPCStatus(grpcfed.FailedPreconditionCode, errorMessage)
					return nil
				},
			}); err != nil {
				return err
			}
			return stat.Err()
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithCustomResolverValue(ctx)
	ret, err := s.resolver.Resolve_Org_Federation_CustomHandlerMessage(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	ctx = grpcfed.WithLogger(ctx, grpcfed.GetCustomResolverValue(ctx).Logger)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.CustomHandlerMessage", slog.Any("org.federation.CustomHandlerMessage", s.logvalue_Org_Federation_CustomHandlerMessage(ret)))
	return ret, nil
}

// resolve_Org_Federation_CustomMessage resolve "org.federation.CustomMessage" message.
func (s *FederationService) resolve_Org_Federation_CustomMessage(ctx context.Context, req *Org_Federation_CustomMessageArgument) (*CustomMessage, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.CustomMessage")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.CustomMessage", slog.Any("message_args", s.logvalue_Org_Federation_CustomMessageArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.CustomMessageArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// create a message value to be returned.
	ret := &CustomMessage{}

	// field binding section.
	// (grpc.federation.field).by = "$.message"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
		Value:             value,
		Expr:              "$.message",
		UseContextLibrary: false,
		CacheIndex:        2,
		Setter:            func(v string) { ret.Message = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.CustomMessage", slog.Any("org.federation.CustomMessage", s.logvalue_Org_Federation_CustomMessage(ret)))
	return ret, nil
}

// resolve_Org_Federation_GetPostResponse resolve "org.federation.GetPostResponse" message.
func (s *FederationService) resolve_Org_Federation_GetPostResponse(ctx context.Context, req *Org_Federation_GetPostResponseArgument) (*GetPostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.GetPostResponse")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.GetPostResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetPostResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			_def2                  bool
			_def3                  bool
			_def4                  bool
			_def4_err_detail0_msg0 *CustomMessage
			_def4_err_detail0_msg1 *CustomMessage
			customHandler          *CustomHandlerMessage
			post                   *Post
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.GetPostResponseArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()
	// A tree view of message dependencies is shown below.
	/*
	   post ─┐
	                 _def2 ─┐
	   post ─┐              │
	                 _def3 ─┤
	   post ─┐              │
	                 _def4 ─┤
	         customHandler ─┤
	*/
	eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "post"
		     message {
		       name: "Post"
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Post, *localValueType]{
			Name:   "post",
			Type:   grpcfed.CELObjectType("org.federation.Post"),
			Setter: func(value *localValueType, v *Post) { value.vars.post = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_PostArgument{}
				return s.resolve_Org_Federation_Post(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "_def2"
		     validation {
		       error {
		         code: FAILED_PRECONDITION
		         if: "post.id != 'some-id'"
		         message: "'validation1 failed!'"
		       }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[bool, *localValueType]{
			Name:   "_def2",
			Type:   grpcfed.CELBoolType,
			Setter: func(value *localValueType, v bool) { value.vars._def2 = v },
			Validation: func(ctx context.Context, value *localValueType) error {
				var stat *grpcfed.Status
				if err := grpcfed.If(ctx1, &grpcfed.IfParam[*localValueType]{
					Value:             value,
					Expr:              "post.id != 'some-id'",
					UseContextLibrary: false,
					CacheIndex:        3,
					Body: func(value *localValueType) error {
						errmsg, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
							Value:             value,
							Expr:              "'validation1 failed!'",
							UseContextLibrary: false,
							OutType:           reflect.TypeOf(""),
							CacheIndex:        4,
						})
						if err != nil {
							return err
						}
						errorMessage := errmsg.(string)
						stat = grpcfed.NewGRPCStatus(grpcfed.FailedPreconditionCode, errorMessage)
						return nil
					},
				}); err != nil {
					return err
				}
				return stat.Err()
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
		     name: "post"
		     message {
		       name: "Post"
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Post, *localValueType]{
			Name:   "post",
			Type:   grpcfed.CELObjectType("org.federation.Post"),
			Setter: func(value *localValueType, v *Post) { value.vars.post = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_PostArgument{}
				return s.resolve_Org_Federation_Post(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "_def3"
		     validation {
		       error {
		         code: FAILED_PRECONDITION
		         if: "post.id != 'some-id'"
		         message: "'validation2 failed!'"
		       }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[bool, *localValueType]{
			Name:   "_def3",
			Type:   grpcfed.CELBoolType,
			Setter: func(value *localValueType, v bool) { value.vars._def3 = v },
			Validation: func(ctx context.Context, value *localValueType) error {
				var stat *grpcfed.Status
				if err := grpcfed.If(ctx1, &grpcfed.IfParam[*localValueType]{
					Value:             value,
					Expr:              "post.id != 'some-id'",
					UseContextLibrary: false,
					CacheIndex:        5,
					Body: func(value *localValueType) error {
						errmsg, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
							Value:             value,
							Expr:              "'validation2 failed!'",
							UseContextLibrary: false,
							OutType:           reflect.TypeOf(""),
							CacheIndex:        6,
						})
						if err != nil {
							return err
						}
						errorMessage := errmsg.(string)
						stat = grpcfed.NewGRPCStatus(grpcfed.FailedPreconditionCode, errorMessage)
						return nil
					},
				}); err != nil {
					return err
				}
				return stat.Err()
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
		     name: "post"
		     message {
		       name: "Post"
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Post, *localValueType]{
			Name:   "post",
			Type:   grpcfed.CELObjectType("org.federation.Post"),
			Setter: func(value *localValueType, v *Post) { value.vars.post = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_PostArgument{}
				return s.resolve_Org_Federation_Post(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "_def4"
		     validation {
		       error {
		         code: FAILED_PRECONDITION
		         if: "$.id != 'correct-id'"
		         message: "'validation3 failed!'"
		         details {
		           if: "true"
		           message: [
		             {...},
		             {...}
		           ]
		           precondition_failure {...}
		           bad_request {...}
		           localized_message {...}
		         }
		       }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[bool, *localValueType]{
			Name:   "_def4",
			Type:   grpcfed.CELBoolType,
			Setter: func(value *localValueType, v bool) { value.vars._def4 = v },
			Validation: func(ctx context.Context, value *localValueType) error {
				var stat *grpcfed.Status
				if err := grpcfed.If(ctx1, &grpcfed.IfParam[*localValueType]{
					Value:             value,
					Expr:              "$.id != 'correct-id'",
					UseContextLibrary: false,
					CacheIndex:        7,
					Body: func(value *localValueType) error {
						errmsg, err := grpcfed.EvalCEL(ctx, &grpcfed.EvalCELRequest{
							Value:             value,
							Expr:              "'validation3 failed!'",
							UseContextLibrary: false,
							OutType:           reflect.TypeOf(""),
							CacheIndex:        8,
						})
						if err != nil {
							return err
						}
						errorMessage := errmsg.(string)
						var details []grpcfed.ProtoMessage
						if err := grpcfed.If(ctx1, &grpcfed.IfParam[*localValueType]{
							Value:             value,
							Expr:              "true",
							UseContextLibrary: false,
							CacheIndex:        9,
							Body: func(value *localValueType) error {
								if _, err := func() (any, error) {
									// A tree view of message dependencies is shown below.
									/*
									   _def4_err_detail0_msg0 ─┐
									   _def4_err_detail0_msg1 ─┤
									*/
									eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

									grpcfed.GoWithRecover(eg, func() (any, error) {

										// This section's codes are generated by the following proto definition.
										/*
										   def {
										     name: "_def4_err_detail0_msg0"
										     message {
										       name: "CustomMessage"
										       args { name: "message", string: "message1" }
										     }
										   }
										*/
										if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*CustomMessage, *localValueType]{
											Name:   "_def4_err_detail0_msg0",
											Type:   grpcfed.CELObjectType("org.federation.CustomMessage"),
											Setter: func(value *localValueType, v *CustomMessage) { value.vars._def4_err_detail0_msg0 = v },
											Message: func(ctx context.Context, value *localValueType) (any, error) {
												args := &Org_Federation_CustomMessageArgument{
													Message: "message1", // { name: "message", string: "message1" }
												}
												return s.resolve_Org_Federation_CustomMessage(ctx, args)
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
										     name: "_def4_err_detail0_msg1"
										     message {
										       name: "CustomMessage"
										       args { name: "message", string: "message2" }
										     }
										   }
										*/
										if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*CustomMessage, *localValueType]{
											Name:   "_def4_err_detail0_msg1",
											Type:   grpcfed.CELObjectType("org.federation.CustomMessage"),
											Setter: func(value *localValueType, v *CustomMessage) { value.vars._def4_err_detail0_msg1 = v },
											Message: func(ctx context.Context, value *localValueType) (any, error) {
												args := &Org_Federation_CustomMessageArgument{
													Message: "message2", // { name: "message", string: "message2" }
												}
												return s.resolve_Org_Federation_CustomMessage(ctx, args)
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
									return nil, nil
								}(); err != nil {
									return err
								}
								if detail := grpcfed.CustomMessage(ctx1, &grpcfed.CustomMessageParam{
									Value:            value,
									MessageValueName: "_def4_err_detail0_msg0",
									CacheIndex:       10,
									MessageIndex:     0,
								}); detail != nil {
									details = append(details, detail)
								}
								if detail := grpcfed.CustomMessage(ctx1, &grpcfed.CustomMessageParam{
									Value:            value,
									MessageValueName: "_def4_err_detail0_msg1",
									CacheIndex:       11,
									MessageIndex:     1,
								}); detail != nil {
									details = append(details, detail)
								}
								if detail := grpcfed.PreconditionFailure(ctx, value, []*grpcfed.PreconditionFailureViolation{
									{
										Type:                     "'type1'",
										Subject:                  "post.id",
										Desc:                     "'description1'",
										TypeUseContextLibrary:    false,
										SubjectUseContextLibrary: false,
										DescUseContextLibrary:    false,
										TypeCacheIndex:           12,
										SubjectCacheIndex:        13,
										DescCacheIndex:           14,
									},
								}); detail != nil {
									details = append(details, detail)
								}
								if detail := grpcfed.BadRequest(ctx, value, []*grpcfed.BadRequestFieldViolation{
									{
										Field:                  "post.id",
										Desc:                   "'description2'",
										FieldUseContextLibrary: false,
										DescUseContextLibrary:  false,
										FieldCacheIndex:        15,
										DescCacheIndex:         16,
									},
								}); detail != nil {
									details = append(details, detail)
								}
								if detail := grpcfed.LocalizedMessage(ctx, &grpcfed.LocalizedMessageParam{
									Value:             value,
									Locale:            "en-US",
									Message:           "post.content",
									UseContextLibrary: false,
									CacheIndex:        17,
								}); detail != nil {
									details = append(details, detail)
								}
								return nil
							},
						}); err != nil {
							return err
						}
						status := grpcfed.NewGRPCStatus(grpcfed.FailedPreconditionCode, errorMessage)
						statusWithDetails, err := status.WithDetails(details...)
						if err != nil {
							grpcfed.Logger(ctx1).ErrorContext(ctx1, "failed setting error details", slog.String("error", err.Error()))
							stat = status
						} else {
							stat = statusWithDetails
						}
						return nil
					},
				}); err != nil {
					return err
				}
				return stat.Err()
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
		     name: "customHandler"
		     message {
		       name: "CustomHandlerMessage"
		       args { name: "arg", string: "some-arg" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*CustomHandlerMessage, *localValueType]{
			Name:   "customHandler",
			Type:   grpcfed.CELObjectType("org.federation.CustomHandlerMessage"),
			Setter: func(value *localValueType, v *CustomHandlerMessage) { value.vars.customHandler = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Org_Federation_CustomHandlerMessageArgument{
					Arg: "some-arg", // { name: "arg", string: "some-arg" }
				}
				return s.resolve_Org_Federation_CustomHandlerMessage(ctx, args)
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
	req.Post = value.vars.post
	req.XDef4ErrDetail0Msg0 = value.vars._def4_err_detail0_msg0
	req.XDef4ErrDetail0Msg1 = value.vars._def4_err_detail0_msg1

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "post"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*Post]{
		Value:             value,
		Expr:              "post",
		UseContextLibrary: false,
		CacheIndex:        18,
		Setter:            func(v *Post) { ret.Post = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.GetPostResponse", slog.Any("org.federation.GetPostResponse", s.logvalue_Org_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_Post resolve "org.federation.Post" message.
func (s *FederationService) resolve_Org_Federation_Post(ctx context.Context, req *Org_Federation_PostArgument) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.Post")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve org.federation.Post", slog.Any("message_args", s.logvalue_Org_Federation_PostArgument(req)))

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = "some-id"           // (grpc.federation.field).string = "some-id"
	ret.Title = "some-title"     // (grpc.federation.field).string = "some-title"
	ret.Content = "some-content" // (grpc.federation.field).string = "some-content"

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved org.federation.Post", slog.Any("org.federation.Post", s.logvalue_Org_Federation_Post(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_CustomHandlerMessage(v *CustomHandlerMessage) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Org_Federation_CustomHandlerMessageArgument(v *Org_Federation_CustomHandlerMessageArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("arg", v.Arg),
	)
}

func (s *FederationService) logvalue_Org_Federation_CustomMessage(v *CustomMessage) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("message", v.GetMessage()),
	)
}

func (s *FederationService) logvalue_Org_Federation_CustomMessageArgument(v *Org_Federation_CustomMessageArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("message", v.Message),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Org_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetPostResponseArgument(v *Org_Federation_GetPostResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Org_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
	)
}

func (s *FederationService) logvalue_Org_Federation_PostArgument(v *Org_Federation_PostArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}
