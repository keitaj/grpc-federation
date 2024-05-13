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

	post "example/post"
	user "example/user"
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Federation_ForNamelessArgument is argument for "federation.ForNameless" message.
type Federation_ForNamelessArgument struct {
	Bar string
}

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument struct {
	Id   string
	Post *Post
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument struct {
	Id     string
	Post   *post.Post
	Res    *post.GetPostResponse
	Unused *Unused
	User   *User
	XDef4  *ForNameless
}

// Federation_Post_UserArgument is custom resolver's argument for "user" field of "federation.Post" message.
type Federation_Post_UserArgument struct {
	*Federation_PostArgument
}

// Federation_UnusedArgument is argument for "federation.Unused" message.
type Federation_UnusedArgument struct {
	Foo string
}

// Federation_UserArgument is argument for "federation.User" message.
type Federation_UserArgument struct {
	Content string
	Id      string
	Res     *user.GetUserResponse
	Title   string
	U       *user.User
	UserId  string
}

// Federation_User_NameArgument is custom resolver's argument for "name" field of "federation.User" message.
type Federation_User_NameArgument struct {
	*Federation_UserArgument
	Federation_User *User
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
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
	// Post_PostServiceClient create a gRPC Client to be used to call methods in post.PostService.
	Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
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
	Post_PostServiceClient post.PostServiceClient
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
	// Resolve_Federation_ForNameless implements resolver for "federation.ForNameless".
	Resolve_Federation_ForNameless(context.Context, *Federation_ForNamelessArgument) (*ForNameless, error)
	// Resolve_Federation_Post_User implements resolver for "federation.Post.user".
	Resolve_Federation_Post_User(context.Context, *Federation_Post_UserArgument) (*User, error)
	// Resolve_Federation_Unused implements resolver for "federation.Unused".
	Resolve_Federation_Unused(context.Context, *Federation_UnusedArgument) (*Unused, error)
	// Resolve_Federation_User implements resolver for "federation.User".
	Resolve_Federation_User(context.Context, *Federation_UserArgument) (*User, error)
	// Resolve_Federation_User_Name implements resolver for "federation.User.name".
	Resolve_Federation_User_Name(context.Context, *Federation_User_NameArgument) (string, error)
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

// Resolve_Federation_ForNameless resolve "federation.ForNameless".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_ForNameless(context.Context, *Federation_ForNamelessArgument) (ret *ForNameless, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_ForNameless not implemented")
	return
}

// Resolve_Federation_Post_User resolve "federation.Post.user".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_Post_User(context.Context, *Federation_Post_UserArgument) (ret *User, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_Post_User not implemented")
	return
}

// Resolve_Federation_Unused resolve "federation.Unused".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_Unused(context.Context, *Federation_UnusedArgument) (ret *Unused, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_Unused not implemented")
	return
}

// Resolve_Federation_User resolve "federation.User".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_User(context.Context, *Federation_UserArgument) (ret *User, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_User not implemented")
	return
}

// Resolve_Federation_User_Name resolve "federation.User.name".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Federation_User_Name(context.Context, *Federation_User_NameArgument) (ret string, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_User_Name not implemented")
	return
}

const (
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
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
	resolver      FederationServiceResolver
	celTypeHelper *grpcfed.CELTypeHelper
	envOpts       []grpcfed.CELEnvOption
	celPlugins    []*grpcfedcel.CELPlugin
	client        *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, grpcfed.ErrClientConfig
	}
	if cfg.Resolver == nil {
		return nil, grpcfed.ErrResolverConfig
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationServiceClientConfig{
		Service: "post.PostService",
	})
	if err != nil {
		return nil, err
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
		"grpc.federation.private.ForNamelessArgument": {
			"bar": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Bar"),
		},
		"grpc.federation.private.GetPostResponseArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.UnusedArgument": {
			"foo": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Foo"),
		},
		"grpc.federation.private.UserArgument": {
			"id":      grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
			"title":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Title"),
			"content": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "UserId"),
		},
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
		tracer:        otel.Tracer("federation.FederationService"),
		resolver:      cfg.Resolver,
		client: &FederationServiceDependentClientSet{
			Post_PostServiceClient: Post_PostServiceClient,
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

// GetPost implements "federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.FederationService/GetPost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	ctx = grpcfed.WithCELCacheMap(ctx, s.celCacheMap)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument{
		Id: req.Id,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_ForNameless resolve "federation.ForNameless" message.
func (s *FederationService) resolve_Federation_ForNameless(ctx context.Context, req *Federation_ForNamelessArgument) (*ForNameless, error) {
	ctx, span := s.tracer.Start(ctx, "federation.ForNameless")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.ForNameless", slog.Any("message_args", s.logvalue_Federation_ForNamelessArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithCustomResolverValue(ctx)
	ret, err := s.resolver.Resolve_Federation_ForNameless(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	ctx = grpcfed.WithLogger(ctx, grpcfed.GetCustomResolverValue(ctx).Logger)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.ForNameless", slog.Any("federation.ForNameless", s.logvalue_Federation_ForNameless(ret)))
	return ret, nil
}

// resolve_Federation_GetPostResponse resolve "federation.GetPostResponse" message.
func (s *FederationService) resolve_Federation_GetPostResponse(ctx context.Context, req *Federation_GetPostResponseArgument) (*GetPostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetPostResponse")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.GetPostResponse", slog.Any("message_args", s.logvalue_Federation_GetPostResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			post *Post
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.GetPostResponseArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "post"
	     message {
	       name: "Post"
	       args { name: "id", by: "$.id" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*Post, *localValueType]{
		Name:   "post",
		Type:   grpcfed.CELObjectType("federation.Post"),
		Setter: func(value *localValueType, v *Post) { value.vars.post = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Federation_PostArgument{}
			// { name: "id", by: "$.id" }
			if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
				Value:             value,
				Expr:              "$.id",
				UseContextLibrary: false,
				CacheIndex:        1,
				Setter: func(v string) {
					args.Id = v
				},
			}); err != nil {
				return nil, err
			}
			return s.resolve_Federation_Post(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = value.vars.post

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "post"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*Post]{
		Value:             value,
		Expr:              "post",
		UseContextLibrary: false,
		CacheIndex:        2,
		Setter:            func(v *Post) { ret.Post = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.GetPostResponse", slog.Any("federation.GetPostResponse", s.logvalue_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *FederationService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Post")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			_def4  *ForNameless
			post   *post.Post
			res    *post.GetPostResponse
			unused *Unused
			user   *User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.PostArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()
	// A tree view of message dependencies is shown below.
	/*
	               _def4 ─┐
	              unused ─┤
	   res ─┐             │
	        post ─┐       │
	                user ─┤
	*/
	eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "_def4"
		     message {
		       name: "ForNameless"
		       args { name: "bar", string: "bar" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*ForNameless, *localValueType]{
			Name:   "_def4",
			Type:   grpcfed.CELObjectType("federation.ForNameless"),
			Setter: func(value *localValueType, v *ForNameless) { value.vars._def4 = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_ForNamelessArgument{
					Bar: "bar", // { name: "bar", string: "bar" }
				}
				return s.resolve_Federation_ForNameless(ctx, args)
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
		     name: "unused"
		     message {
		       name: "Unused"
		       args { name: "foo", string: "foo" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Unused, *localValueType]{
			Name:   "unused",
			Type:   grpcfed.CELObjectType("federation.Unused"),
			Setter: func(value *localValueType, v *Unused) { value.vars.unused = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_UnusedArgument{
					Foo: "foo", // { name: "foo", string: "foo" }
				}
				return s.resolve_Federation_Unused(ctx, args)
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
		     name: "res"
		     call {
		       method: "post.PostService/GetPost"
		       request { field: "id", by: "$.id" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*post.GetPostResponse, *localValueType]{
			Name:   "res",
			Type:   grpcfed.CELObjectType("post.GetPostResponse"),
			Setter: func(value *localValueType, v *post.GetPostResponse) { value.vars.res = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &post.GetPostRequest{}
				// { field: "id", by: "$.id" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              "$.id",
					UseContextLibrary: false,
					CacheIndex:        3,
					Setter: func(v string) {
						args.Id = v
					},
				}); err != nil {
					return nil, err
				}
				grpcfed.Logger(ctx).DebugContext(ctx, "call post.PostService/GetPost", slog.Any("post.GetPostRequest", s.logvalue_Post_GetPostRequest(args)))
				return s.client.Post_PostServiceClient.GetPost(ctx, args)
			},
		}); err != nil {
			if err := s.errorHandler(ctx1, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx1, err)
				return nil, err
			}
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "post"
		     autobind: true
		     by: "res.post"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*post.Post, *localValueType]{
			Name:                "post",
			Type:                grpcfed.CELObjectType("post.Post"),
			Setter:              func(value *localValueType, v *post.Post) { value.vars.post = v },
			By:                  "res.post",
			ByUseContextLibrary: false,
			ByCacheIndex:        4,
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "user"
		     message {
		       name: "User"
		       args { inline: "post" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*User, *localValueType]{
			Name:   "user",
			Type:   grpcfed.CELObjectType("federation.User"),
			Setter: func(value *localValueType, v *User) { value.vars.user = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_UserArgument{}
				// { inline: "post" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*post.Post]{
					Value:             value,
					Expr:              "post",
					UseContextLibrary: false,
					CacheIndex:        5,
					Setter: func(v *post.Post) {
						args.Id = v.GetId()
						args.Title = v.GetTitle()
						args.Content = v.GetContent()
						args.UserId = v.GetUserId()
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Federation_User(ctx, args)
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
	req.Res = value.vars.res
	req.Unused = value.vars.unused
	req.User = value.vars.user
	req.XDef4 = value.vars._def4

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = value.vars.post.GetId()           // { name: "post", autobind: true }
	ret.Title = value.vars.post.GetTitle()     // { name: "post", autobind: true }
	ret.Content = value.vars.post.GetContent() // { name: "post", autobind: true }
	{
		// (grpc.federation.field).custom_resolver = true
		ctx = grpcfed.WithCustomResolverValue(ctx)
		var err error
		ret.User, err = s.resolver.Resolve_Federation_Post_User(ctx, &Federation_Post_UserArgument{
			Federation_PostArgument: req,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ctx = grpcfed.WithLogger(ctx, grpcfed.GetCustomResolverValue(ctx).Logger)
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

// resolve_Federation_Unused resolve "federation.Unused" message.
func (s *FederationService) resolve_Federation_Unused(ctx context.Context, req *Federation_UnusedArgument) (*Unused, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Unused")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.Unused", slog.Any("message_args", s.logvalue_Federation_UnusedArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithCustomResolverValue(ctx)
	ret, err := s.resolver.Resolve_Federation_Unused(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	ctx = grpcfed.WithLogger(ctx, grpcfed.GetCustomResolverValue(ctx).Logger)

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.Unused", slog.Any("federation.Unused", s.logvalue_Federation_Unused(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *FederationService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.User")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			res *user.GetUserResponse
			u   *user.User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.UserArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "res"
	     call {
	       method: "user.UserService/GetUser"
	       request { field: "id", by: "$.user_id" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.GetUserResponse, *localValueType]{
		Name:   "res",
		Type:   grpcfed.CELObjectType("user.GetUserResponse"),
		Setter: func(value *localValueType, v *user.GetUserResponse) { value.vars.res = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &user.GetUserRequest{}
			// { field: "id", by: "$.user_id" }
			if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
				Value:             value,
				Expr:              "$.user_id",
				UseContextLibrary: false,
				CacheIndex:        6,
				Setter: func(v string) {
					args.Id = v
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
			return nil, err
		}
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "u"
	     by: "res.user"
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.User, *localValueType]{
		Name:                "u",
		Type:                grpcfed.CELObjectType("user.User"),
		Setter:              func(value *localValueType, v *user.User) { value.vars.u = v },
		By:                  "res.user",
		ByUseContextLibrary: false,
		ByCacheIndex:        7,
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Res = value.vars.res
	req.U = value.vars.u

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithCustomResolverValue(ctx)
	ret, err := s.resolver.Resolve_Federation_User(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	ctx = grpcfed.WithLogger(ctx, grpcfed.GetCustomResolverValue(ctx).Logger)

	// field binding section.
	{
		// (grpc.federation.field).custom_resolver = true
		ctx = grpcfed.WithCustomResolverValue(ctx)
		var err error
		ret.Name, err = s.resolver.Resolve_Federation_User_Name(ctx, &Federation_User_NameArgument{
			Federation_UserArgument: req,
			Federation_User:         ret,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ctx = grpcfed.WithLogger(ctx, grpcfed.GetCustomResolverValue(ctx).Logger)
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Federation_ForNameless(v *ForNameless) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("bar", v.GetBar()),
	)
}

func (s *FederationService) logvalue_Federation_ForNamelessArgument(v *Federation_ForNamelessArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("bar", v.Bar),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponseArgument(v *Federation_GetPostResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_Unused(v *Unused) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("foo", v.GetFoo()),
	)
}

func (s *FederationService) logvalue_Federation_UnusedArgument(v *Federation_UnusedArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("foo", v.Foo),
	)
}

func (s *FederationService) logvalue_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *FederationService) logvalue_Federation_UserArgument(v *Federation_UserArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Post_GetPostRequest(v *post.GetPostRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
	)
}

func (s *FederationService) logvalue_Post_GetPostsRequest(v *post.GetPostsRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.GetIds()),
	)
}

func (s *FederationService) logvalue_User_GetUserRequest(v *user.GetUserRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
	)
}

func (s *FederationService) logvalue_User_GetUsersRequest(v *user.GetUsersRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.GetIds()),
	)
}
