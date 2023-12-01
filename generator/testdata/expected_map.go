// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"
	"sync"

	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"

	post "example/post"
	user "example/user"
)

// Org_Federation_GetPostsResponseArgument is argument for "org.federation.GetPostsResponse" message.
type Org_Federation_GetPostsResponseArgument[T any] struct {
	Ids    []string
	Posts  *Posts
	Client T
}

// Org_Federation_PostsArgument is argument for "org.federation.Posts" message.
type Org_Federation_PostsArgument[T any] struct {
	Ids     []string
	PostIds []string
	Posts   []*post.Post
	Res     *post.GetPostsResponse
	Users   []*User
	Client  T
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument[T any] struct {
	Res    *user.GetUserResponse
	User   *user.User
	UserId string
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
	// Org_User_UserServiceClient create a gRPC Client to be used to call methods in org.user.UserService.
	Org_User_UserServiceClient(FederationServiceClientConfig) (user.UserServiceClient, error)
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
	Org_User_UserServiceClient user.UserServiceClient
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
	FederationService_DependentMethod_Org_Post_PostService_GetPosts = "/org.post.PostService/GetPosts"
	FederationService_DependentMethod_Org_User_UserService_GetUser  = "/org.user.UserService/GetUser"
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
	Org_Post_PostServiceClient, err := cfg.Client.Org_Post_PostServiceClient(FederationServiceClientConfig{
		Service: "org.post.PostService",
		Name:    "",
	})
	if err != nil {
		return nil, err
	}
	Org_User_UserServiceClient, err := cfg.Client.Org_User_UserServiceClient(FederationServiceClientConfig{
		Service: "org.user.UserService",
		Name:    "",
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
		"grpc.federation.private.GetPostsResponseArgument": {
			"ids": grpcfed.NewCELFieldType(celtypes.NewListType(celtypes.StringType), "Ids"),
		},
		"grpc.federation.private.PostsArgument": {
			"post_ids": grpcfed.NewCELFieldType(celtypes.NewListType(celtypes.StringType), "PostIds"),
		},
		"grpc.federation.private.UserArgument": {
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
	})
	env, err := cel.NewCustomEnv(
		cel.StdLib(),
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
		tracer:       otel.Tracer("org.federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Org_Post_PostServiceClient: Org_Post_PostServiceClient,
			Org_User_UserServiceClient: Org_User_UserServiceClient,
		},
	}, nil
}

// GetPosts implements "org.federation.FederationService/GetPosts" method.
func (s *FederationService) GetPosts(ctx context.Context, req *GetPostsRequest) (res *GetPostsResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/GetPosts")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetPostsResponse(ctx, &Org_Federation_GetPostsResponseArgument[*FederationServiceDependentClientSet]{
		Client: s.client,
		Ids:    req.Ids,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetPostsResponse resolve "org.federation.GetPostsResponse" message.
func (s *FederationService) resolve_Org_Federation_GetPostsResponse(ctx context.Context, req *Org_Federation_GetPostsResponseArgument[*FederationServiceDependentClientSet]) (*GetPostsResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.GetPostsResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.GetPostsResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetPostsResponseArgument(req)))
	var (
		sg         singleflight.Group
		valueMu    sync.RWMutex
		valuePosts *Posts
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.GetPostsResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "posts"
	     message {
	       name: "Posts"
	       args { name: "post_ids", by: "$.ids" }
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("posts", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_PostsArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "post_ids", by: "$.ids" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.ids", envOpts, evalValues, reflect.TypeOf([]string(nil)))
				if err != nil {
					valueMu.RUnlock()
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.PostIds = value.([]string)
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_Posts(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*Posts)
		valueMu.Lock()
		valuePosts = value
		envOpts = append(envOpts, cel.Variable("posts", cel.ObjectType("org.federation.Posts")))
		evalValues["posts"] = valuePosts
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Posts = valuePosts

	// create a message value to be returned.
	ret := &GetPostsResponse{}

	// field binding section.
	// (grpc.federation.field).by = "posts"
	{
		value, err := grpcfed.EvalCEL(s.env, "posts", envOpts, evalValues, reflect.TypeOf((*Posts)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Posts = value.(*Posts)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.GetPostsResponse", slog.Any("org.federation.GetPostsResponse", s.logvalue_Org_Federation_GetPostsResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_Posts resolve "org.federation.Posts" message.
func (s *FederationService) resolve_Org_Federation_Posts(ctx context.Context, req *Org_Federation_PostsArgument[*FederationServiceDependentClientSet]) (*Posts, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.Posts")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.Posts", slog.Any("message_args", s.logvalue_Org_Federation_PostsArgument(req)))
	var (
		sg         singleflight.Group
		valueIds   []string
		valueMu    sync.RWMutex
		valuePosts []*post.Post
		valueRes   *post.GetPostsResponse
		valueUsers []*User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.PostsArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}
	// A tree view of message dependencies is shown below.
	/*
	   res ─┐
	        posts ─┐
	                 ids ─┐
	   res ─┐             │
	        posts ─┐      │
	               users ─┤
	*/
	eg, ctx1 := errgroup.WithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     call {
		       method: "org.post.PostService/GetPosts"
		       request { field: "ids", by: "$.post_ids" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("res", func() (any, error) {
				valueMu.RLock()
				args := &post.GetPostsRequest{}
				// { field: "ids", by: "$.post_ids" }
				{
					value, err := grpcfed.EvalCEL(s.env, "$.post_ids", envOpts, evalValues, reflect.TypeOf([]string(nil)))
					if err != nil {
						valueMu.RUnlock()
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.Ids = value.([]string)
				}
				valueMu.RUnlock()
				return s.client.Org_Post_PostServiceClient.GetPosts(ctx1, args)
			})
			if err != nil {
				if err := s.errorHandler(ctx1, FederationService_DependentMethod_Org_Post_PostService_GetPosts, err); err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
			}
			value := valueIface.(*post.GetPostsResponse)
			valueMu.Lock()
			valueRes = value
			envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("org.post.GetPostsResponse")))
			evalValues["res"] = valueRes
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "posts"
		     by: "res.posts"
		   }
		*/
		{
			valueIface, err, _ := sg.Do("posts", func() (any, error) {
				valueMu.RLock()
				valueMu.RUnlock()
				valueMu.RLock()
				value, err := grpcfed.EvalCEL(s.env, "res.posts", envOpts, evalValues, reflect.TypeOf([]*post.Post(nil)))
				valueMu.RUnlock()
				return value, err
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.([]*post.Post)
			valueMu.Lock()
			valuePosts = value
			envOpts = append(envOpts, cel.Variable("posts", cel.ListType(cel.ObjectType("org.post.Post"))))
			evalValues["posts"] = valuePosts
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "ids"
		     map {
		       iterator {
		         name: "post"
		         src: "posts"
		       }
		       by: "post.id"
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("ids", func() (any, error) {
				valueMu.RLock()
				valueMu.RUnlock()
				env, err := s.env.Extend(cel.Variable("post", cel.ObjectType("org.post.Post")))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				valueMu.RLock()
				defer valueMu.RUnlock()
				var value []string
				for _, iter := range valuePosts {
					iterValues := make(map[string]any)
					for k, v := range evalValues {
						iterValues[k] = v
					}
					iterValues["post"] = iter
					resultValue, err := grpcfed.EvalCEL(env, "post.id", envOpts, iterValues, reflect.TypeOf(""))
					if err != nil {
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					iterValue := resultValue.(string)
					value = append(value, iterValue)
				}
				return value, nil
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.([]string)
			valueMu.Lock()
			valueIds = value
			envOpts = append(envOpts, cel.Variable("ids", cel.ListType(celtypes.StringType)))
			evalValues["ids"] = valueIds
			valueMu.Unlock()
		}
		return nil, nil
	})

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     call {
		       method: "org.post.PostService/GetPosts"
		       request { field: "ids", by: "$.post_ids" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("res", func() (any, error) {
				valueMu.RLock()
				args := &post.GetPostsRequest{}
				// { field: "ids", by: "$.post_ids" }
				{
					value, err := grpcfed.EvalCEL(s.env, "$.post_ids", envOpts, evalValues, reflect.TypeOf([]string(nil)))
					if err != nil {
						valueMu.RUnlock()
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.Ids = value.([]string)
				}
				valueMu.RUnlock()
				return s.client.Org_Post_PostServiceClient.GetPosts(ctx1, args)
			})
			if err != nil {
				if err := s.errorHandler(ctx1, FederationService_DependentMethod_Org_Post_PostService_GetPosts, err); err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
			}
			value := valueIface.(*post.GetPostsResponse)
			valueMu.Lock()
			valueRes = value
			envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("org.post.GetPostsResponse")))
			evalValues["res"] = valueRes
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "posts"
		     by: "res.posts"
		   }
		*/
		{
			valueIface, err, _ := sg.Do("posts", func() (any, error) {
				valueMu.RLock()
				valueMu.RUnlock()
				valueMu.RLock()
				value, err := grpcfed.EvalCEL(s.env, "res.posts", envOpts, evalValues, reflect.TypeOf([]*post.Post(nil)))
				valueMu.RUnlock()
				return value, err
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.([]*post.Post)
			valueMu.Lock()
			valuePosts = value
			envOpts = append(envOpts, cel.Variable("posts", cel.ListType(cel.ObjectType("org.post.Post"))))
			evalValues["posts"] = valuePosts
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "users"
		     map {
		       iterator {
		         name: "iter"
		         src: "posts"
		       }
		       message {
		         name: "User"
		         args { name: "user_id", by: "iter.user_id" }
		       }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("users", func() (any, error) {
				valueMu.RLock()
				valueMu.RUnlock()
				env, err := s.env.Extend(cel.Variable("iter", cel.ObjectType("org.post.Post")))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				valueMu.RLock()
				defer valueMu.RUnlock()
				var value []*User
				for _, iter := range valuePosts {
					iterValues := make(map[string]any)
					for k, v := range evalValues {
						iterValues[k] = v
					}
					iterValues["iter"] = iter
					args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
						Client: s.client,
					}
					// { name: "user_id", by: "iter.user_id" }
					{
						value, err := grpcfed.EvalCEL(env, "iter.user_id", envOpts, iterValues, reflect.TypeOf(""))
						if err != nil {
							grpcfed.RecordErrorToSpan(ctx, err)
							return nil, err
						}
						args.UserId = value.(string)
					}
					iterValue, err := s.resolve_Org_Federation_User(ctx1, args)
					if err != nil {
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					value = append(value, iterValue)
				}
				return value, nil
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.([]*User)
			valueMu.Lock()
			valueUsers = value
			envOpts = append(envOpts, cel.Variable("users", cel.ListType(cel.ObjectType("org.federation.User"))))
			evalValues["users"] = valueUsers
			valueMu.Unlock()
		}
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Ids = valueIds
	req.Posts = valuePosts
	req.Res = valueRes
	req.Users = valueUsers

	// create a message value to be returned.
	ret := &Posts{}

	// field binding section.
	// (grpc.federation.field).by = "ids"
	{
		value, err := grpcfed.EvalCEL(s.env, "ids", envOpts, evalValues, reflect.TypeOf([]string(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Ids = value.([]string)
	}
	// (grpc.federation.field).by = "posts.map(post, post.title)"
	{
		value, err := grpcfed.EvalCEL(s.env, "posts.map(post, post.title)", envOpts, evalValues, reflect.TypeOf([]string(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Titles = value.([]string)
	}
	// (grpc.federation.field).by = "posts.map(post, post.content)"
	{
		value, err := grpcfed.EvalCEL(s.env, "posts.map(post, post.content)", envOpts, evalValues, reflect.TypeOf([]string(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Contents = value.([]string)
	}
	// (grpc.federation.field).by = "users"
	{
		value, err := grpcfed.EvalCEL(s.env, "users", envOpts, evalValues, reflect.TypeOf([]*User(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Users = value.([]*User)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.Posts", slog.Any("org.federation.Posts", s.logvalue_Org_Federation_Posts(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valueRes  *user.GetUserResponse
		valueUser *user.User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "res"
	     call {
	       method: "org.user.UserService/GetUser"
	       request { field: "id", by: "$.user_id" }
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("res", func() (any, error) {
			valueMu.RLock()
			args := &user.GetUserRequest{}
			// { field: "id", by: "$.user_id" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					valueMu.RUnlock()
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Id = value.(string)
			}
			valueMu.RUnlock()
			return s.client.Org_User_UserServiceClient.GetUser(ctx, args)
		})
		if err != nil {
			if err := s.errorHandler(ctx, FederationService_DependentMethod_Org_User_UserService_GetUser, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
		}
		value := valueIface.(*user.GetUserResponse)
		valueMu.Lock()
		valueRes = value
		envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("org.user.GetUserResponse")))
		evalValues["res"] = valueRes
		valueMu.Unlock()
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "user"
	     autobind: true
	     by: "res.user"
	   }
	*/
	{
		valueIface, err, _ := sg.Do("user", func() (any, error) {
			valueMu.RLock()
			valueMu.RUnlock()
			valueMu.RLock()
			value, err := grpcfed.EvalCEL(s.env, "res.user", envOpts, evalValues, reflect.TypeOf((*user.User)(nil)))
			valueMu.RUnlock()
			return value, err
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*user.User)
		valueMu.Lock()
		valueUser = value
		envOpts = append(envOpts, cel.Variable("user", cel.ObjectType("org.user.User")))
		evalValues["user"] = valueUser
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Res = valueRes
	req.User = valueUser

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	ret.Id = valueUser.GetId()     // { name: "user", autobind: true }
	ret.Name = valueUser.GetName() // { name: "user", autobind: true }

	s.logger.DebugContext(ctx, "resolved org.federation.User", slog.Any("org.federation.User", s.logvalue_Org_Federation_User(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_GetPostsResponse(v *GetPostsResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("posts", s.logvalue_Org_Federation_Posts(v.GetPosts())),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetPostsResponseArgument(v *Org_Federation_GetPostsResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.Ids),
	)
}

func (s *FederationService) logvalue_Org_Federation_Posts(v *Posts) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.GetIds()),
		slog.Any("titles", v.GetTitles()),
		slog.Any("contents", v.GetContents()),
		slog.Any("users", s.logvalue_repeated_Org_Federation_User(v.GetUsers())),
	)
}

func (s *FederationService) logvalue_Org_Federation_PostsArgument(v *Org_Federation_PostsArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post_ids", v.PostIds),
	)
}

func (s *FederationService) logvalue_Org_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserArgument(v *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_repeated_Org_Federation_User(v []*User) slog.Value {
	attrs := make([]slog.Attr, 0, len(v))
	for idx, vv := range v {
		attrs = append(attrs, slog.Attr{
			Key:   fmt.Sprint(idx),
			Value: s.logvalue_Org_Federation_User(vv),
		})
	}
	return slog.GroupValue(attrs...)
}
