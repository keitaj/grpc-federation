// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
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
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Org_Federation_CreatePostArgument is argument for "org.federation.CreatePost" message.
type Org_Federation_CreatePostArgument[T any] struct {
	Content string
	Title   string
	UserId  string
	Client  T
}

// Org_Federation_CreatePostResponseArgument is argument for "org.federation.CreatePostResponse" message.
type Org_Federation_CreatePostResponseArgument[T any] struct {
	Content string
	Cp      *CreatePost
	P       *post.Post
	Res     *post.CreatePostResponse
	Title   string
	UserId  string
	Client  T
}

// Org_Federation_PostArgument is argument for "org.federation.Post" message.
type Org_Federation_PostArgument[T any] struct {
	Client T
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
	// Org_Post_PostServiceClient create a gRPC Client to be used to call methods in org.post.PostService.
	Org_Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
}

// FederationServiceClientConfig information set in `dependencies` of the `grpc.federation.service` option.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service returns the name of the service on Protocol Buffers.
	Service string
	// Name is the value set for `name` in `dependencies` of the `grpc.federation.service` option.
	// It must be unique among the services on which the Federation Service depends.
	Name string
}

// FederationServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependentClientSet struct {
	Org_Post_PostServiceClient post.PostServiceClient
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
	FederationService_DependentMethod_Org_Post_PostService_CreatePost = "/org.post.PostService/CreatePost"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *grpcfed.CELEnv
	tracer       trace.Tracer
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, grpcfed.ErrClientConfig
	}
	Org_Post_PostServiceClient, err := cfg.Client.Org_Post_PostServiceClient(FederationServiceClientConfig{
		Service: "org.post.PostService",
		Name:    "post_service",
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
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*grpcfed.CELFieldType{
		"grpc.federation.private.CreatePostArgument": {
			"title":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Title"),
			"content": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "UserId"),
		},
		"grpc.federation.private.CreatePostResponseArgument": {
			"title":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Title"),
			"content": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "UserId"),
		},
	})
	envOpts := grpcfed.NewDefaultEnvOptions(celHelper)
	env, err := grpcfed.NewCELEnv(envOpts...)
	if err != nil {
		return nil, err
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("org.federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Org_Post_PostServiceClient: Org_Post_PostServiceClient,
		},
	}, nil
}

// CreatePost implements "org.federation.FederationService/CreatePost" method.
func (s *FederationService) CreatePost(ctx context.Context, req *CreatePostRequest) (res *CreatePostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/CreatePost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_CreatePostResponse(ctx, &Org_Federation_CreatePostResponseArgument[*FederationServiceDependentClientSet]{
		Client:  s.client,
		Title:   req.Title,
		Content: req.Content,
		UserId:  req.UserId,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_CreatePost resolve "org.federation.CreatePost" message.
func (s *FederationService) resolve_Org_Federation_CreatePost(ctx context.Context, req *Org_Federation_CreatePostArgument[*FederationServiceDependentClientSet]) (*CreatePost, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.CreatePost")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.CreatePost", slog.Any("message_args", s.logvalue_Org_Federation_CreatePostArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.CreatePostArgument", req)}

	// create a message value to be returned.
	ret := &CreatePost{}

	// field binding section.
	// (grpc.federation.field).by = "$.title"
	if err := grpcfed.SetCELValue(ctx, value, "$.title", func(v string) { ret.Title = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "$.content"
	if err := grpcfed.SetCELValue(ctx, value, "$.content", func(v string) { ret.Content = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "$.user_id"
	if err := grpcfed.SetCELValue(ctx, value, "$.user_id", func(v string) { ret.UserId = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved org.federation.CreatePost", slog.Any("org.federation.CreatePost", s.logvalue_Org_Federation_CreatePost(ret)))
	return ret, nil
}

// resolve_Org_Federation_CreatePostResponse resolve "org.federation.CreatePostResponse" message.
func (s *FederationService) resolve_Org_Federation_CreatePostResponse(ctx context.Context, req *Org_Federation_CreatePostResponseArgument[*FederationServiceDependentClientSet]) (*CreatePostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.CreatePostResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.CreatePostResponse", slog.Any("message_args", s.logvalue_Org_Federation_CreatePostResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			cp  *CreatePost
			p   *post.Post
			res *post.CreatePostResponse
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.CreatePostResponseArgument", req)}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "cp"
	     message {
	       name: "CreatePost"
	       args: [
	         { name: "title", by: "$.title" },
	         { name: "content", by: "$.content" },
	         { name: "user_id", by: "$.user_id" }
	       ]
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*CreatePost, *localValueType]{
		Name:   "cp",
		Type:   grpcfed.CELObjectType("org.federation.CreatePost"),
		Setter: func(value *localValueType, v *CreatePost) { value.vars.cp = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Org_Federation_CreatePostArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "title", by: "$.title" }
			if err := grpcfed.SetCELValue(ctx, value, "$.title", func(v string) {
				args.Title = v
			}); err != nil {
				return nil, err
			}
			// { name: "content", by: "$.content" }
			if err := grpcfed.SetCELValue(ctx, value, "$.content", func(v string) {
				args.Content = v
			}); err != nil {
				return nil, err
			}
			// { name: "user_id", by: "$.user_id" }
			if err := grpcfed.SetCELValue(ctx, value, "$.user_id", func(v string) {
				args.UserId = v
			}); err != nil {
				return nil, err
			}
			return s.resolve_Org_Federation_CreatePost(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "res"
	     call {
	       method: "org.post.PostService/CreatePost"
	       request { field: "post", by: "cp" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*post.CreatePostResponse, *localValueType]{
		Name:   "res",
		Type:   grpcfed.CELObjectType("org.post.CreatePostResponse"),
		Setter: func(value *localValueType, v *post.CreatePostResponse) { value.vars.res = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &post.CreatePostRequest{}
			// { field: "post", by: "cp" }
			if err := grpcfed.SetCELValue(ctx, value, "cp", func(v *CreatePost) {
				args.Post = s.cast_Org_Federation_CreatePost__to__Org_Post_CreatePost(v)
			}); err != nil {
				return nil, err
			}
			return s.client.Org_Post_PostServiceClient.CreatePost(ctx, args)
		},
	}); err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Org_Post_PostService_CreatePost, err); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "p"
	     by: "res.post"
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*post.Post, *localValueType]{
		Name:   "p",
		Type:   grpcfed.CELObjectType("org.post.Post"),
		Setter: func(value *localValueType, v *post.Post) { value.vars.p = v },
		By:     "res.post",
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Cp = value.vars.cp
	req.P = value.vars.p
	req.Res = value.vars.res

	// create a message value to be returned.
	ret := &CreatePostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "p"
	if err := grpcfed.SetCELValue(ctx, value, "p", func(v *post.Post) { ret.Post = s.cast_Org_Post_Post__to__Org_Federation_Post(v) }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved org.federation.CreatePostResponse", slog.Any("org.federation.CreatePostResponse", s.logvalue_Org_Federation_CreatePostResponse(ret)))
	return ret, nil
}

// cast_Org_Federation_CreatePost__to__Org_Post_CreatePost cast from "org.federation.CreatePost" to "org.post.CreatePost".
func (s *FederationService) cast_Org_Federation_CreatePost__to__Org_Post_CreatePost(from *CreatePost) *post.CreatePost {
	if from == nil {
		return nil
	}

	return &post.CreatePost{
		Title:   from.GetTitle(),
		Content: from.GetContent(),
		UserId:  from.GetUserId(),
	}
}

// cast_Org_Post_Post__to__Org_Federation_Post cast from "org.post.Post" to "org.federation.Post".
func (s *FederationService) cast_Org_Post_Post__to__Org_Federation_Post(from *post.Post) *Post {
	if from == nil {
		return nil
	}

	return &Post{
		Id:      from.GetId(),
		Title:   from.GetTitle(),
		Content: from.GetContent(),
		UserId:  from.GetUserId(),
	}
}

func (s *FederationService) logvalue_Org_Federation_CreatePost(v *CreatePost) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.String("user_id", v.GetUserId()),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostArgument(v *Org_Federation_CreatePostArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostResponse(v *CreatePostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Org_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostResponseArgument(v *Org_Federation_CreatePostResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
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
		slog.String("user_id", v.GetUserId()),
	)
}
