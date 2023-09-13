// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	post "example/post"
	user "example/user"
)

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler FederationServiceErrorHandler
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

// FederationServiceClientConfig information set in `dependencies` of the `grpc.federation.service` option.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service returns the name of the service on Protocol Buffers.
	Service string
	// Name is the value set for `name` in `dependencies` of the `grpc.federation.service` option.
	// It must be unique among the services on which the Federation Service depends.
	Name string
}

// FederationServiceDependencyServiceClient has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependencyServiceClient struct {
	Post_PostServiceClient post.PostServiceClient
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// FederationServiceErrorHandler Federation Service often needs to convert errors received from downstream services.
// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
type FederationServiceErrorHandler func(ctx context.Context, methodName string, err error) error

const (
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
	FederationService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
)

// FederationServiceRecoveredError represents recovered error.
type FederationServiceRecoveredError struct {
	Message string
	Stack   []string
}

func (e *FederationServiceRecoveredError) Error() string {
	return fmt.Sprintf("recovered error: %s", e.Message)
}

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler FederationServiceErrorHandler
	client       *FederationServiceDependencyServiceClient
}

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument struct {
	Client *FederationServiceDependencyServiceClient
	Id     string
	Post   *Post
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument struct {
	Client *FederationServiceDependencyServiceClient
	Id     string
	Post   *post.Post
	User   *User
}

// Federation_UserArgument is argument for "federation.User" message.
type Federation_UserArgument struct {
	Client  *FederationServiceDependencyServiceClient
	Content string
	Id      string
	Title   string
	User    *user.User
	UserId  string
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if err := validateFederationServiceConfig(cfg); err != nil {
		return nil, err
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationServiceClientConfig{
		Service: "post.PostService",
		Name:    "post_service",
	})
	if err != nil {
		return nil, err
	}
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationServiceClientConfig{
		Service: "user.UserService",
		Name:    "user_service",
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
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		client: &FederationServiceDependencyServiceClient{
			Post_PostServiceClient: Post_PostServiceClient,
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

func validateFederationServiceConfig(cfg FederationServiceConfig) error {
	if cfg.Client == nil {
		return fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	return nil
}

func withTimeoutFederationService[T any](ctx context.Context, method string, timeout time.Duration, fn func(context.Context) (*T, error)) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var (
		ret   *T
		errch = make(chan error)
	)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errch <- recoverErrorFederationService(r, debug.Stack())
			}
		}()

		res, err := fn(ctx)
		ret = res
		errch <- err
	}()
	select {
	case <-ctx.Done():
		status := grpcstatus.New(grpccodes.DeadlineExceeded, ctx.Err().Error())
		withDetails, err := status.WithDetails(&errdetails.ErrorInfo{
			Metadata: map[string]string{
				"method":  method,
				"timeout": timeout.String(),
			},
		})
		if err != nil {
			return nil, status.Err()
		}
		return nil, withDetails.Err()
	case err := <-errch:
		return ret, err
	}
}

func withRetryFederationService[T any](b backoff.BackOff, fn func() (*T, error)) (*T, error) {
	var res *T
	if err := backoff.Retry(func() (err error) {
		res, err = fn()
		return
	}, b); err != nil {
		return nil, err
	}
	return res, nil
}

func recoverErrorFederationService(v interface{}, rawStack []byte) *FederationServiceRecoveredError {
	msg := fmt.Sprint(v)
	lines := strings.Split(msg, "\n")
	if len(lines) <= 1 {
		lines := strings.Split(string(rawStack), "\n")
		stack := make([]string, 0, len(lines))
		for _, line := range lines {
			if line == "" {
				continue
			}
			stack = append(stack, strings.TrimPrefix(line, "\t"))
		}
		return &FederationServiceRecoveredError{
			Message: msg,
			Stack:   stack,
		}
	}
	// If panic occurs under singleflight, singleflight's recover catches the error and gives a stack trace.
	// Therefore, once the stack trace is removed.
	stack := make([]string, 0, len(lines))
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		stack = append(stack, strings.TrimPrefix(line, "\t"))
	}
	return &FederationServiceRecoveredError{
		Message: lines[0],
		Stack:   stack,
	}
}

func (s *FederationService) goWithRecover(eg *errgroup.Group, fn func() (interface{}, error)) {
	eg.Go(func() (e error) {
		defer func() {
			if r := recover(); r != nil {
				e = recoverErrorFederationService(r, debug.Stack())
			}
		}()
		_, err := fn()
		return err
	})
}

func (s *FederationService) outputErrorLog(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if status, ok := grpcstatus.FromError(err); ok {
		s.logger.ErrorContext(ctx, status.Message(),
			slog.Group("grpc_status",
				slog.String("code", status.Code().String()),
				slog.Any("details", status.Details()),
			),
		)
		return
	}
	var recoveredErr *FederationServiceRecoveredError
	if errors.As(err, &recoveredErr) {
		trace := make([]interface{}, 0, len(recoveredErr.Stack))
		for idx, stack := range recoveredErr.Stack {
			trace = append(trace, slog.String(fmt.Sprint(idx+1), stack))
		}
		s.logger.ErrorContext(ctx, recoveredErr.Message, slog.Group("stack_trace", trace...))
		return
	}
	s.logger.ErrorContext(ctx, err.Error())
}

// GetPost implements "federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = recoverErrorFederationService(r, debug.Stack())
			s.outputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument{
		Client: s.client,
		Id:     req.Id,
	})
	if err != nil {
		s.outputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetPostResponse resolve "federation.GetPostResponse" message.
func (s *FederationService) resolve_Federation_GetPostResponse(ctx context.Context, req *Federation_GetPostResponseArgument) (*GetPostResponse, error) {
	s.logger.DebugContext(ctx, "resolve  federation.GetPostResponse", slog.Any("message_args", s.logvalue_Federation_GetPostResponseArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *Post
	)

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "post"
	     message: "Post"
	     args { name: "id", by: "$.id" }
	   }
	*/
	resPostIface, err, _ := sg.Do("post_federation.Post", func() (interface{}, error) {
		valueMu.RLock()
		args := &Federation_PostArgument{
			Client: s.client,
			Id:     req.Id, // { name: "id", by: "$.id" }
		}
		valueMu.RUnlock()
		return s.resolve_Federation_Post(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resPost := resPostIface.(*Post)
	valueMu.Lock()
	valuePost = resPost // { name: "post", message: "Post" ... }
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	ret.Post = valuePost // (grpc.federation.field).by = "post"
	ret.Str = "hello"    // (grpc.federation.field).string = "hello"

	s.logger.DebugContext(ctx, "resolved federation.GetPostResponse", slog.Any("federation.GetPostResponse", s.logvalue_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *FederationService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument) (*Post, error) {
	s.logger.DebugContext(ctx, "resolve  federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *post.Post
		valueUser *User
	)

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "post.PostService/GetPost"
	     request { field: "id", by: "$.id" }
	     response { name: "post", field: "post", autobind: true }
	   }
	*/
	resGetPostResponseIface, err, _ := sg.Do("post.PostService/GetPost", func() (interface{}, error) {
		valueMu.RLock()
		args := &post.GetPostRequest{
			Id: req.Id, // { field: "id", by: "$.id" }
		}
		valueMu.RUnlock()
		return withTimeoutFederationService[post.GetPostResponse](ctx, "post.PostService/GetPost", 10000000000 /* 10s */, func(ctx context.Context) (*post.GetPostResponse, error) {
			var b backoff.BackOff = backoff.NewConstantBackOff(2000000000 /* 2s */)
			b = backoff.WithMaxRetries(b, 3)
			b = backoff.WithContext(b, ctx)
			return withRetryFederationService[post.GetPostResponse](b, func() (*post.GetPostResponse, error) {
				return s.client.Post_PostServiceClient.GetPost(ctx, args)
			})
		})
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
			return nil, err
		}
	}
	resGetPostResponse := resGetPostResponseIface.(*post.GetPostResponse)
	valueMu.Lock()
	valuePost = resGetPostResponse.GetPost() // { name: "post", field: "post", autobind: true }
	valueMu.Unlock()

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "user"
	     message: "User"
	     args { inline: "post" }
	   }
	*/
	resUserIface, err, _ := sg.Do("user_federation.User", func() (interface{}, error) {
		valueMu.RLock()
		args := &Federation_UserArgument{
			Client:  s.client,
			Id:      valuePost.GetId(),      // { inline: "post" }
			Title:   valuePost.GetTitle(),   // { inline: "post" }
			Content: valuePost.GetContent(), // { inline: "post" }
			UserId:  valuePost.GetUserId(),  // { inline: "post" }
		}
		valueMu.RUnlock()
		return s.resolve_Federation_User(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resUser := resUserIface.(*User)
	valueMu.Lock()
	valueUser = resUser // { name: "user", message: "User" ... }
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost
	req.User = valueUser

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = valuePost.GetId()           // { name: "post", autobind: true }
	ret.Title = valuePost.GetTitle()     // { name: "post", autobind: true }
	ret.Content = valuePost.GetContent() // { name: "post", autobind: true }
	ret.User = valueUser                 // (grpc.federation.field).by = "user"

	s.logger.DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *FederationService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument) (*User, error) {
	s.logger.DebugContext(ctx, "resolve  federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valueUser *user.User
	)

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "user.UserService/GetUser"
	     request { field: "id", by: "$.user_id" }
	     response { name: "user", field: "user", autobind: true }
	   }
	*/
	resGetUserResponseIface, err, _ := sg.Do("user.UserService/GetUser", func() (interface{}, error) {
		valueMu.RLock()
		args := &user.GetUserRequest{
			Id: req.UserId, // { field: "id", by: "$.user_id" }
		}
		valueMu.RUnlock()
		return withTimeoutFederationService[user.GetUserResponse](ctx, "user.UserService/GetUser", 20000000000 /* 20s */, func(ctx context.Context) (*user.GetUserResponse, error) {
			eb := backoff.NewExponentialBackOff()
			eb.InitialInterval = 1000000000 /* 1s */
			eb.RandomizationFactor = 0.7
			eb.Multiplier = 1.7
			eb.MaxInterval = 30000000000    /* 30s */
			eb.MaxElapsedTime = 20000000000 /* 20s */

			var b backoff.BackOff = eb
			b = backoff.WithMaxRetries(b, 3)
			b = backoff.WithContext(b, ctx)
			return withRetryFederationService[user.GetUserResponse](b, func() (*user.GetUserResponse, error) {
				return s.client.User_UserServiceClient.GetUser(ctx, args)
			})
		})
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
			return nil, err
		}
	}
	resGetUserResponse := resGetUserResponseIface.(*user.GetUserResponse)
	valueMu.Lock()
	valueUser = resGetUserResponse.GetUser() // { name: "user", field: "user", autobind: true }
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.User = valueUser

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	ret.Id = valueUser.GetId()                                                                // { name: "user", autobind: true }
	ret.Name = valueUser.GetName()                                                            // { name: "user", autobind: true }
	ret.Items = s.cast_repeated_User_Item__to__repeated_Federation_Item(valueUser.GetItems()) // { name: "user", autobind: true }
	ret.Profile = valueUser.GetProfile()                                                      // { name: "user", autobind: true }

	switch {
	case s.cast_User_User_AttrA___to__Federation_User_AttrA_(valueUser.GetAttrA()) != nil:
		ret.Attr = s.cast_User_User_AttrA___to__Federation_User_AttrA_(valueUser.GetAttrA())
	case s.cast_User_User_B__to__Federation_User_B(valueUser.GetB()) != nil:
		ret.Attr = s.cast_User_User_B__to__Federation_User_B(valueUser.GetB())
	}

	s.logger.DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

// cast_repeated_User_Item__to__repeated_Federation_Item cast from "repeated user.Item" to "repeated federation.Item".
func (s *FederationService) cast_repeated_User_Item__to__repeated_Federation_Item(from []*user.Item) []*Item {
	ret := make([]*Item, 0, len(from))
	for _, v := range from {
		ret = append(ret, s.cast_User_Item__to__Federation_Item(v))
	}
	return ret
}

// cast_User_Item_ItemType__to__Federation_Item_ItemType cast from "user.Item.ItemType" to "federation.Item.ItemType".
func (s *FederationService) cast_User_Item_ItemType__to__Federation_Item_ItemType(from user.Item_ItemType) Item_ItemType {
	switch from {
	case user.Item_ITEM_TYPE_UNSPECIFIED:
		return Item_ITEM_TYPE_UNSPECIFIED
	case user.Item_ITEM_TYPE_1:
		return Item_ITEM_TYPE_1
	case user.Item_ITEM_TYPE_2:
		return Item_ITEM_TYPE_2
	case user.Item_ITEM_TYPE_3:
		return Item_ITEM_TYPE_3
	default:
		return 0
	}
}

// cast_User_Item_Location__to__Federation_Item_Location cast from "user.Item.Location" to "federation.Item.Location".
func (s *FederationService) cast_User_Item_Location__to__Federation_Item_Location(from *user.Item_Location) *Item_Location {
	if from == nil {
		return nil
	}
	return &Item_Location{
		Addr1: from.GetAddr1(),
		Addr2: from.GetAddr2(),
	}
}

// cast_User_Item__to__Federation_Item cast from "user.Item" to "federation.Item".
func (s *FederationService) cast_User_Item__to__Federation_Item(from *user.Item) *Item {
	if from == nil {
		return nil
	}
	return &Item{
		Name:     from.GetName(),
		Type:     s.cast_User_Item_ItemType__to__Federation_Item_ItemType(from.GetType()),
		Value:    from.GetValue(),
		Location: s.cast_User_Item_Location__to__Federation_Item_Location(from.GetLocation()),
	}
}

// cast_User_User_AttrA__to__Federation_User_AttrA cast from "user.User.AttrA" to "federation.User.AttrA".
func (s *FederationService) cast_User_User_AttrA__to__Federation_User_AttrA(from *user.User_AttrA) *User_AttrA {
	if from == nil {
		return nil
	}
	return &User_AttrA{
		Foo: from.GetFoo(),
	}
}

// cast_User_User_AttrB__to__Federation_User_AttrB cast from "user.User.AttrB" to "federation.User.AttrB".
func (s *FederationService) cast_User_User_AttrB__to__Federation_User_AttrB(from *user.User_AttrB) *User_AttrB {
	if from == nil {
		return nil
	}
	return &User_AttrB{
		Bar: from.GetBar(),
	}
}

// cast_User_User_AttrA___to__Federation_User_AttrA_ cast from "user.User.attr_a" to "federation.User.attr_a".
func (s *FederationService) cast_User_User_AttrA___to__Federation_User_AttrA_(from *user.User_AttrA) *User_AttrA_ {
	if from == nil {
		return nil
	}
	return &User_AttrA_{
		AttrA: s.cast_User_User_AttrA__to__Federation_User_AttrA(from),
	}
}

// cast_User_User_B__to__Federation_User_B cast from "user.User.b" to "federation.User.b".
func (s *FederationService) cast_User_User_B__to__Federation_User_B(from *user.User_AttrB) *User_B {
	if from == nil {
		return nil
	}
	return &User_B{
		B: s.cast_User_User_AttrB__to__Federation_User_AttrB(from),
	}
}

func (s *FederationService) logvalue_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
		slog.String("str", v.GetStr()),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponseArgument(v *Federation_GetPostResponseArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_Item(v *Item) slog.Value {
	return slog.GroupValue(
		slog.String("name", v.GetName()),
		slog.String("type", s.logvalue_Federation_Item_ItemType(v.GetType()).String()),
		slog.Int64("value", v.GetValue()),
		slog.Any("location", s.logvalue_Federation_Item_Location(v.GetLocation())),
	)
}

func (s *FederationService) logvalue_Federation_Item_ItemType(v Item_ItemType) slog.Value {
	switch v {
	case Item_ITEM_TYPE_UNSPECIFIED:
		return slog.StringValue("ITEM_TYPE_UNSPECIFIED")
	case Item_ITEM_TYPE_1:
		return slog.StringValue("ITEM_TYPE_1")
	case Item_ITEM_TYPE_2:
		return slog.StringValue("ITEM_TYPE_2")
	case Item_ITEM_TYPE_3:
		return slog.StringValue("ITEM_TYPE_3")
	}
	return slog.StringValue("")
}

func (s *FederationService) logvalue_Federation_Item_Location(v *Item_Location) slog.Value {
	return slog.GroupValue(
		slog.String("addr1", v.GetAddr1()),
		slog.String("addr2", v.GetAddr2()),
	)
}

func (s *FederationService) logvalue_Federation_Post(v *Post) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_User(v *User) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
		slog.Any("items", s.logvalue_repeated_Federation_Item(v.GetItems())),
		slog.Any("profile", s.logvalue_Federation_User_ProfileEntry(v.GetProfile())),
		slog.Any("attr_a", s.logvalue_Federation_User_AttrA(v.GetAttrA())),
		slog.Any("b", s.logvalue_Federation_User_AttrB(v.GetB())),
	)
}

func (s *FederationService) logvalue_Federation_UserArgument(v *Federation_UserArgument) slog.Value {
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Federation_User_AttrA(v *User_AttrA) slog.Value {
	return slog.GroupValue(
		slog.String("foo", v.GetFoo()),
	)
}

func (s *FederationService) logvalue_Federation_User_AttrB(v *User_AttrB) slog.Value {
	return slog.GroupValue(
		slog.Bool("bar", v.GetBar()),
	)
}

func (s *FederationService) logvalue_Federation_User_ProfileEntry(v map[string]*anypb.Any) slog.Value {
	attrs := make([]slog.Attr, 0, len(v))
	for key, value := range v {
		attrs = append(attrs, slog.Attr{
			Key:   fmt.Sprint(key),
			Value: s.logvalue_Google_Protobuf_Any(value),
		})
	}
	return slog.GroupValue(attrs...)
}

func (s *FederationService) logvalue_Google_Protobuf_Any(v *anypb.Any) slog.Value {
	return slog.GroupValue(
		slog.String("type_url", v.GetTypeUrl()),
		slog.String("value", string(v.GetValue())),
	)
}

func (s *FederationService) logvalue_repeated_Federation_Item(v []*Item) slog.Value {
	attrs := make([]slog.Attr, 0, len(v))
	for idx, vv := range v {
		attrs = append(attrs, slog.Attr{
			Key:   fmt.Sprint(idx),
			Value: s.logvalue_Federation_Item(vv),
		})
	}
	return slog.GroupValue(attrs...)
}
