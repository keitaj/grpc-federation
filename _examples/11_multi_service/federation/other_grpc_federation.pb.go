// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
// versions:
//
//	protoc-gen-grpc-federation: dev
//
// source: federation/other.proto
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

	comment "example/comment"
	favorite "example/favorite"
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Federation_GetResponseArgument is argument for "federation.GetResponse" message.
type Federation_GetResponseArgument struct {
	Id string
	P  *Post
}

// Federation_GetResponse_PostArgument is custom resolver's argument for "post" field of "federation.GetResponse" message.
type Federation_GetResponse_PostArgument struct {
	*Federation_GetResponseArgument
}

// OtherServiceConfig configuration required to initialize the service that use GRPC Federation.
type OtherServiceConfig struct {
	// Resolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
	// If this interface is not provided, an error is returned during initialization.
	Resolver OtherServiceResolver // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// OtherServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type OtherServiceClientFactory interface {
}

// OtherServiceClientConfig helper to create gRPC client.
// Hints for creating a gRPC Client.
type OtherServiceClientConfig struct {
	// Service FQDN ( `<package-name>.<service-name>` ) of the service on Protocol Buffers.
	Service string
}

// OtherServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type OtherServiceDependentClientSet struct {
}

// OtherServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type OtherServiceResolver interface {
	// Resolve_Federation_GetResponse_Post implements resolver for "federation.GetResponse.post".
	Resolve_Federation_GetResponse_Post(context.Context, *Federation_GetResponse_PostArgument) (*Post, error)
}

// OtherServiceCELPluginWasmConfig type alias for grpcfedcel.WasmConfig.
type OtherServiceCELPluginWasmConfig = grpcfedcel.WasmConfig

// OtherServiceCELPluginConfig hints for loading a WebAssembly based plugin.
type OtherServiceCELPluginConfig struct {
}

// OtherServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type OtherServiceUnimplementedResolver struct{}

// Resolve_Federation_GetResponse_Post resolve "federation.GetResponse.post".
// This method always returns Unimplemented error.
func (OtherServiceUnimplementedResolver) Resolve_Federation_GetResponse_Post(context.Context, *Federation_GetResponse_PostArgument) (ret *Post, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_GetResponse_Post not implemented")
	return
}

// OtherService represents Federation Service.
type OtherService struct {
	*UnimplementedOtherServiceServer
	cfg           OtherServiceConfig
	logger        *slog.Logger
	errorHandler  grpcfed.ErrorHandler
	celCacheMap   *grpcfed.CELCacheMap
	tracer        trace.Tracer
	resolver      OtherServiceResolver
	celTypeHelper *grpcfed.CELTypeHelper
	envOpts       []grpcfed.CELEnvOption
	celPlugins    []*grpcfedcel.CELPlugin
	client        *OtherServiceDependentClientSet
}

// NewOtherService creates OtherService instance by OtherServiceConfig.
func NewOtherService(cfg OtherServiceConfig) (*OtherService, error) {
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
		"grpc.federation.private.GetResponseArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {},
		"grpc.federation.private.ReactionArgument": {
			"v": grpcfed.NewCELFieldType(grpcfed.CELIntType, "V"),
		},
		"grpc.federation.private.UserArgument": {
			"id":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
			"name": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Name"),
		},
	}
	celTypeHelper := grpcfed.NewCELTypeHelper(celTypeHelperFieldMap)
	var envOpts []grpcfed.CELEnvOption
	envOpts = append(envOpts, grpcfed.NewDefaultEnvOptions(celTypeHelper)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("comment.CommentType", comment.CommentType_value, comment.CommentType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("favorite.FavoriteType", favorite.FavoriteType_value, favorite.FavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("favorite.FavoriteType", favorite.FavoriteType_value, favorite.FavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	return &OtherService{
		cfg:           cfg,
		logger:        logger,
		errorHandler:  errorHandler,
		envOpts:       envOpts,
		celTypeHelper: celTypeHelper,
		celCacheMap:   grpcfed.NewCELCacheMap(),
		tracer:        otel.Tracer("federation.OtherService"),
		resolver:      cfg.Resolver,
		client:        &OtherServiceDependentClientSet{},
	}, nil
}

// Get implements "federation.OtherService/Get" method.
func (s *OtherService) Get(ctx context.Context, req *GetRequest) (res *GetResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.OtherService/Get")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	ctx = grpcfed.WithCELCacheMap(ctx, s.celCacheMap)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Federation_GetResponse(ctx, &Federation_GetResponseArgument{
		Id: req.Id,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetResponse resolve "federation.GetResponse" message.
func (s *OtherService) resolve_Federation_GetResponse(ctx context.Context, req *Federation_GetResponseArgument) (*GetResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetResponse")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.GetResponse", slog.Any("message_args", s.logvalue_Federation_GetResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			p *Post
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.GetResponseArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "p"
	     message {
	       name: "Post"
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*Post, *localValueType]{
		Name:   "p",
		Type:   grpcfed.CELObjectType("federation.Post"),
		Setter: func(value *localValueType, v *Post) { value.vars.p = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Federation_PostArgument{}
			return s.resolve_Federation_Post(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.P = value.vars.p

	// create a message value to be returned.
	ret := &GetResponse{}

	// field binding section.
	{
		// (grpc.federation.field).custom_resolver = true
		ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx)) // create a new reference to logger.
		var err error
		ret.Post, err = s.resolver.Resolve_Federation_GetResponse_Post(ctx, &Federation_GetResponse_PostArgument{
			Federation_GetResponseArgument: req,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.GetResponse", slog.Any("federation.GetResponse", s.logvalue_Federation_GetResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *OtherService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Post")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			cmp            bool
			favorite_value favorite.FavoriteType
			reaction       *Reaction
			u              *User
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
	   favorite_value ─┐
	                        cmp ─┐
	   favorite_value ─┐         │
	                   reaction ─┤
	                          u ─┤
	*/
	eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "favorite_value"
		     by: "favorite.FavoriteType.value('TYPE1')"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[favorite.FavoriteType, *localValueType]{
			Name:                "favorite_value",
			Type:                grpcfed.CELIntType,
			Setter:              func(value *localValueType, v favorite.FavoriteType) { value.vars.favorite_value = v },
			By:                  "favorite.FavoriteType.value('TYPE1')",
			ByUseContextLibrary: false,
			ByCacheIndex:        1,
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "cmp"
		     by: "favorite_value == favorite.FavoriteType.TYPE1"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[bool, *localValueType]{
			Name:                "cmp",
			Type:                grpcfed.CELBoolType,
			Setter:              func(value *localValueType, v bool) { value.vars.cmp = v },
			By:                  "favorite_value == favorite.FavoriteType.TYPE1",
			ByUseContextLibrary: false,
			ByCacheIndex:        2,
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
		     name: "favorite_value"
		     by: "favorite.FavoriteType.value('TYPE1')"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[favorite.FavoriteType, *localValueType]{
			Name:                "favorite_value",
			Type:                grpcfed.CELIntType,
			Setter:              func(value *localValueType, v favorite.FavoriteType) { value.vars.favorite_value = v },
			By:                  "favorite.FavoriteType.value('TYPE1')",
			ByUseContextLibrary: false,
			ByCacheIndex:        3,
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "reaction"
		     message {
		       name: "Reaction"
		       args { name: "v", by: "favorite_value" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Reaction, *localValueType]{
			Name:   "reaction",
			Type:   grpcfed.CELObjectType("federation.Reaction"),
			Setter: func(value *localValueType, v *Reaction) { value.vars.reaction = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_ReactionArgument{}
				// { name: "v", by: "favorite_value" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[favorite.FavoriteType]{
					Value:             value,
					Expr:              "favorite_value",
					UseContextLibrary: false,
					CacheIndex:        4,
					Setter: func(v favorite.FavoriteType) {
						args.V = v
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Federation_Reaction(ctx, args)
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
		     name: "u"
		     message {
		       name: "User"
		       args: [
		         { name: "id", string: "foo" },
		         { name: "name", string: "bar" }
		       ]
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*User, *localValueType]{
			Name:   "u",
			Type:   grpcfed.CELObjectType("federation.User"),
			Setter: func(value *localValueType, v *User) { value.vars.u = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_UserArgument{
					Id:   "foo", // { name: "id", string: "foo" }
					Name: "bar", // { name: "name", string: "bar" }
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
	req.Cmp = value.vars.cmp
	req.FavoriteValue = value.vars.favorite_value
	req.Reaction = value.vars.reaction
	req.U = value.vars.u

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = "post-id"      // (grpc.federation.field).string = "post-id"
	ret.Title = "title"     // (grpc.federation.field).string = "title"
	ret.Content = "content" // (grpc.federation.field).string = "content"
	// (grpc.federation.field).by = "u"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*User]{
		Value:             value,
		Expr:              "u",
		UseContextLibrary: false,
		CacheIndex:        5,
		Setter:            func(v *User) { ret.User = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "reaction"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*Reaction]{
		Value:             value,
		Expr:              "reaction",
		UseContextLibrary: false,
		CacheIndex:        6,
		Setter:            func(v *Reaction) { ret.Reaction = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "favorite_value"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[favorite.FavoriteType]{
		Value:             value,
		Expr:              "favorite_value",
		UseContextLibrary: false,
		CacheIndex:        7,
		Setter: func(v favorite.FavoriteType) {
			ret.FavoriteValue = s.cast_Favorite_FavoriteType__to__Federation_MyFavoriteType(v)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "cmp"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[bool]{
		Value:             value,
		Expr:              "cmp",
		UseContextLibrary: false,
		CacheIndex:        8,
		Setter:            func(v bool) { ret.Cmp = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

// resolve_Federation_Reaction resolve "federation.Reaction" message.
func (s *OtherService) resolve_Federation_Reaction(ctx context.Context, req *Federation_ReactionArgument) (*Reaction, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Reaction")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.Reaction", slog.Any("message_args", s.logvalue_Federation_ReactionArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			cmp bool
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.ReactionArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "cmp"
	     by: "$.v == favorite.FavoriteType.TYPE1"
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[bool, *localValueType]{
		Name:                "cmp",
		Type:                grpcfed.CELBoolType,
		Setter:              func(value *localValueType, v bool) { value.vars.cmp = v },
		By:                  "$.v == favorite.FavoriteType.TYPE1",
		ByUseContextLibrary: false,
		ByCacheIndex:        9,
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Cmp = value.vars.cmp

	// create a message value to be returned.
	ret := &Reaction{}

	// field binding section.
	// (grpc.federation.field).by = "favorite.FavoriteType.TYPE1"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[favorite.FavoriteType]{
		Value:             value,
		Expr:              "favorite.FavoriteType.TYPE1",
		UseContextLibrary: false,
		CacheIndex:        10,
		Setter:            func(v favorite.FavoriteType) { ret.FavoriteType = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "favorite.FavoriteType.name(favorite.FavoriteType.value('TYPE1'))"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
		Value:             value,
		Expr:              "favorite.FavoriteType.name(favorite.FavoriteType.value('TYPE1'))",
		UseContextLibrary: false,
		CacheIndex:        11,
		Setter:            func(v string) { ret.FavoriteTypeStr = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "cmp"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[bool]{
		Value:             value,
		Expr:              "cmp",
		UseContextLibrary: false,
		CacheIndex:        12,
		Setter:            func(v bool) { ret.Cmp = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.Reaction", slog.Any("federation.Reaction", s.logvalue_Federation_Reaction(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *OtherService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.User")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.UserArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.id"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
		Value:             value,
		Expr:              "$.id",
		UseContextLibrary: false,
		CacheIndex:        13,
		Setter:            func(v string) { ret.Id = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "$.name"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
		Value:             value,
		Expr:              "$.name",
		UseContextLibrary: false,
		CacheIndex:        14,
		Setter:            func(v string) { ret.Name = v },
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

// cast_Favorite_FavoriteType__to__Federation_MyFavoriteType cast from "favorite.FavoriteType" to "federation.MyFavoriteType".
func (s *OtherService) cast_Favorite_FavoriteType__to__Federation_MyFavoriteType(from favorite.FavoriteType) MyFavoriteType {
	switch from {
	case favorite.FavoriteType_UNKNOWN:
		return MyFavoriteType_UNKNOWN
	case favorite.FavoriteType_TYPE1:
		return MyFavoriteType_TYPE1
	default:
		return 0
	}
}

// cast_int64__to__Favorite_FavoriteType cast from "int64" to "favorite.FavoriteType".
func (s *OtherService) cast_int64__to__Favorite_FavoriteType(from int64) favorite.FavoriteType {
	return favorite.FavoriteType(from)
}

func (s *OtherService) logvalue_Favorite_FavoriteType(v favorite.FavoriteType) slog.Value {
	switch v {
	case favorite.FavoriteType_UNKNOWN:
		return slog.StringValue("UNKNOWN")
	case favorite.FavoriteType_TYPE1:
		return slog.StringValue("TYPE1")
	case favorite.FavoriteType_TYPE2:
		return slog.StringValue("TYPE2")
	}
	return slog.StringValue("")
}

func (s *OtherService) logvalue_Federation_GetResponse(v *GetResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
	)
}

func (s *OtherService) logvalue_Federation_GetResponseArgument(v *Federation_GetResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *OtherService) logvalue_Federation_MyFavoriteType(v MyFavoriteType) slog.Value {
	switch v {
	case MyFavoriteType_UNKNOWN:
		return slog.StringValue("UNKNOWN")
	case MyFavoriteType_TYPE1:
		return slog.StringValue("TYPE1")
	}
	return slog.StringValue("")
}

func (s *OtherService) logvalue_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
		slog.Any("reaction", s.logvalue_Federation_Reaction(v.GetReaction())),
		slog.String("favorite_value", s.logvalue_Federation_MyFavoriteType(v.GetFavoriteValue()).String()),
		slog.Bool("cmp", v.GetCmp()),
	)
}

func (s *OtherService) logvalue_Federation_PostArgument(v *Federation_PostArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *OtherService) logvalue_Federation_Reaction(v *Reaction) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("favorite_type", s.logvalue_Favorite_FavoriteType(v.GetFavoriteType()).String()),
		slog.String("favorite_type_str", v.GetFavoriteTypeStr()),
		slog.Bool("cmp", v.GetCmp()),
	)
}

func (s *OtherService) logvalue_Federation_ReactionArgument(v *Federation_ReactionArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("v", s.logvalue_Favorite_FavoriteType(v.V).String()),
	)
}

func (s *OtherService) logvalue_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *OtherService) logvalue_Federation_UserArgument(v *Federation_UserArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("name", v.Name),
	)
}
