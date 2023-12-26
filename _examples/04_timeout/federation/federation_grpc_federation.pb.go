// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime/debug"

	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	grpcfedcel "github.com/mercari/grpc-federation/grpc/federation/cel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	post "example/post"
)

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument[T any] struct {
	Id     string
	Post   *Post
	Client T
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument[T any] struct {
	Id     string
	Post   *post.Post
	Res    *post.GetPostResponse
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
	// Post_PostServiceClient create a gRPC Client to be used to call methods in post.PostService.
	Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
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
	Post_PostServiceClient post.PostServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

const (
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	tracer       trace.Tracer
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationServiceClientConfig{
		Service: "post.PostService",
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
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*celtypes.FieldType{
		"grpc.federation.private.GetPostResponseArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
	})
	env, err := cel.NewCustomEnv(
		cel.StdLib(),
		cel.Lib(grpcfedcel.NewLibrary()),
		cel.CrossTypeNumericComparisons(true),
		cel.CustomTypeAdapter(celHelper.TypeAdapter()),
		cel.CustomTypeProvider(celHelper.TypeProvider()),
	)
	if err != nil {
		return nil, err
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Post_PostServiceClient: Post_PostServiceClient,
		},
	}, nil
}

// GetPost implements "federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.FederationService/GetPost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := grpcfed.WithTimeout[GetPostResponse](ctx, "federation.FederationService/GetPost", 1000000000 /* 1s */, func(ctx context.Context) (*GetPostResponse, error) {
		return s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]{
			Client: s.client,
			Id:     req.Id,
		})
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetPostResponse resolve "federation.GetPostResponse" message.
func (s *FederationService) resolve_Federation_GetPostResponse(ctx context.Context, req *Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) (*GetPostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetPostResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.GetPostResponse", slog.Any("message_args", s.logvalue_Federation_GetPostResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			post *Post
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.GetPostResponseArgument", req)}

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
		Type:   cel.ObjectType("federation.Post"),
		Setter: func(value *localValueType, v *Post) { value.vars.post = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Federation_PostArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "id", by: "$.id" }
			if err := grpcfed.SetCELValue(ctx, value, "$.id", func(v string) {
				args.Id = v
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
	if err := grpcfed.SetCELValue(ctx, value, "post", func(v *Post) { ret.Post = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.GetPostResponse", slog.Any("federation.GetPostResponse", s.logvalue_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *FederationService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument[*FederationServiceDependentClientSet]) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Post")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			post *post.Post
			res  *post.GetPostResponse
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.PostArgument", req)}

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
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*post.GetPostResponse, *localValueType]{
		Name:   "res",
		Type:   cel.ObjectType("post.GetPostResponse"),
		Setter: func(value *localValueType, v *post.GetPostResponse) { value.vars.res = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &post.GetPostRequest{}
			// { field: "id", by: "$.id" }
			if err := grpcfed.SetCELValue(ctx, value, "$.id", func(v string) {
				args.Id = v
			}); err != nil {
				return nil, err
			}
			return s.client.Post_PostServiceClient.GetPost(ctx, args)
		},
	}); err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
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
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*post.Post, *localValueType]{
		Name:   "post",
		Type:   cel.ObjectType("post.Post"),
		Setter: func(value *localValueType, v *post.Post) { value.vars.post = v },
		By:     "res.post",
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = value.vars.post
	req.Res = value.vars.res

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = value.vars.post.GetId()           // { name: "post", autobind: true }
	ret.Title = value.vars.post.GetTitle()     // { name: "post", autobind: true }
	ret.Content = value.vars.post.GetContent() // { name: "post", autobind: true }

	s.logger.DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponseArgument(v *Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
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
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}
