// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package optimization

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	optimizationpb "cloud.google.com/go/optimization/apiv1/optimizationpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newFleetRoutingClientHook clientHook

// FleetRoutingCallOptions contains the retry settings for each method of FleetRoutingClient.
type FleetRoutingCallOptions struct {
	OptimizeTours      []gax.CallOption
	BatchOptimizeTours []gax.CallOption
	GetOperation       []gax.CallOption
}

func defaultFleetRoutingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudoptimization.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudoptimization.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudoptimization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFleetRoutingCallOptions() *FleetRoutingCallOptions {
	return &FleetRoutingCallOptions{
		OptimizeTours: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchOptimizeTours: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetOperation: []gax.CallOption{},
	}
}

func defaultFleetRoutingRESTCallOptions() *FleetRoutingCallOptions {
	return &FleetRoutingCallOptions{
		OptimizeTours: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		BatchOptimizeTours: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetOperation: []gax.CallOption{},
	}
}

// internalFleetRoutingClient is an interface that defines the methods available from Cloud Optimization API.
type internalFleetRoutingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	OptimizeTours(context.Context, *optimizationpb.OptimizeToursRequest, ...gax.CallOption) (*optimizationpb.OptimizeToursResponse, error)
	BatchOptimizeTours(context.Context, *optimizationpb.BatchOptimizeToursRequest, ...gax.CallOption) (*BatchOptimizeToursOperation, error)
	BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// FleetRoutingClient is a client for interacting with Cloud Optimization API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for optimizing vehicle tours.
//
// Validity of certain types of fields:
//
//	google.protobuf.Timestamp
//
//	  Times are in Unix time: seconds since 1970-01-01T00:00:00+00:00.
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.protobuf.Duration
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.type.LatLng
//
//	  latitude must be in [-90.0, 90.0].
//
//	  longitude must be in [-180.0, 180.0].
//
//	  at least one of latitude and longitude must be non-zero.
type FleetRoutingClient struct {
	// The internal transport-dependent client.
	internalClient internalFleetRoutingClient

	// The call options for this service.
	CallOptions *FleetRoutingCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FleetRoutingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FleetRoutingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *FleetRoutingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// OptimizeTours sends an OptimizeToursRequest containing a ShipmentModel and returns an
// OptimizeToursResponse containing ShipmentRoutes, which are a set of
// routes to be performed by vehicles minimizing the overall cost.
//
// A ShipmentModel model consists mainly of Shipments that need to be
// carried out and Vehicles that can be used to transport the Shipments.
// The ShipmentRoutes assign Shipments to Vehicles. More specifically,
// they assign a series of Visits to each vehicle, where a Visit
// corresponds to a VisitRequest, which is a pickup or delivery for a
// Shipment.
//
// The goal is to provide an assignment of ShipmentRoutes to Vehicles that
// minimizes the total cost where cost has many components defined in the
// ShipmentModel.
func (c *FleetRoutingClient) OptimizeTours(ctx context.Context, req *optimizationpb.OptimizeToursRequest, opts ...gax.CallOption) (*optimizationpb.OptimizeToursResponse, error) {
	return c.internalClient.OptimizeTours(ctx, req, opts...)
}

// BatchOptimizeTours optimizes vehicle tours for one or more OptimizeToursRequest
// messages as a batch.
//
// This method is a Long Running Operation (LRO). The inputs for optimization
// (OptimizeToursRequest messages) and outputs (OptimizeToursResponse
// messages) are read/written from/to Cloud Storage in user-specified
// format. Like the OptimizeTours method, each OptimizeToursRequest
// contains a ShipmentModel and returns an OptimizeToursResponse
// containing ShipmentRoutes, which are a set of routes to be performed by
// vehicles minimizing the overall cost.
func (c *FleetRoutingClient) BatchOptimizeTours(ctx context.Context, req *optimizationpb.BatchOptimizeToursRequest, opts ...gax.CallOption) (*BatchOptimizeToursOperation, error) {
	return c.internalClient.BatchOptimizeTours(ctx, req, opts...)
}

// BatchOptimizeToursOperation returns a new BatchOptimizeToursOperation from a given name.
// The name must be that of a previously created BatchOptimizeToursOperation, possibly from a different process.
func (c *FleetRoutingClient) BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation {
	return c.internalClient.BatchOptimizeToursOperation(name)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *FleetRoutingClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// fleetRoutingGRPCClient is a client for interacting with Cloud Optimization API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type fleetRoutingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FleetRoutingClient
	CallOptions **FleetRoutingCallOptions

	// The gRPC API client.
	fleetRoutingClient optimizationpb.FleetRoutingClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFleetRoutingClient creates a new fleet routing client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for optimizing vehicle tours.
//
// Validity of certain types of fields:
//
//	google.protobuf.Timestamp
//
//	  Times are in Unix time: seconds since 1970-01-01T00:00:00+00:00.
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.protobuf.Duration
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.type.LatLng
//
//	  latitude must be in [-90.0, 90.0].
//
//	  longitude must be in [-180.0, 180.0].
//
//	  at least one of latitude and longitude must be non-zero.
func NewFleetRoutingClient(ctx context.Context, opts ...option.ClientOption) (*FleetRoutingClient, error) {
	clientOpts := defaultFleetRoutingGRPCClientOptions()
	if newFleetRoutingClientHook != nil {
		hookOpts, err := newFleetRoutingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FleetRoutingClient{CallOptions: defaultFleetRoutingCallOptions()}

	c := &fleetRoutingGRPCClient{
		connPool:           connPool,
		disableDeadlines:   disableDeadlines,
		fleetRoutingClient: optimizationpb.NewFleetRoutingClient(connPool),
		CallOptions:        &client.CallOptions,
		operationsClient:   longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *fleetRoutingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *fleetRoutingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *fleetRoutingGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type fleetRoutingRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing FleetRoutingClient
	CallOptions **FleetRoutingCallOptions
}

// NewFleetRoutingRESTClient creates a new fleet routing rest client.
//
// A service for optimizing vehicle tours.
//
// Validity of certain types of fields:
//
//	google.protobuf.Timestamp
//
//	  Times are in Unix time: seconds since 1970-01-01T00:00:00+00:00.
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.protobuf.Duration
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.type.LatLng
//
//	  latitude must be in [-90.0, 90.0].
//
//	  longitude must be in [-180.0, 180.0].
//
//	  at least one of latitude and longitude must be non-zero.
func NewFleetRoutingRESTClient(ctx context.Context, opts ...option.ClientOption) (*FleetRoutingClient, error) {
	clientOpts := append(defaultFleetRoutingRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultFleetRoutingRESTCallOptions()
	c := &fleetRoutingRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &FleetRoutingClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultFleetRoutingRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudoptimization.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudoptimization.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudoptimization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *fleetRoutingRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *fleetRoutingRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *fleetRoutingRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *fleetRoutingGRPCClient) OptimizeTours(ctx context.Context, req *optimizationpb.OptimizeToursRequest, opts ...gax.CallOption) (*optimizationpb.OptimizeToursResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).OptimizeTours[0:len((*c.CallOptions).OptimizeTours):len((*c.CallOptions).OptimizeTours)], opts...)
	var resp *optimizationpb.OptimizeToursResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.fleetRoutingClient.OptimizeTours(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *fleetRoutingGRPCClient) BatchOptimizeTours(ctx context.Context, req *optimizationpb.BatchOptimizeToursRequest, opts ...gax.CallOption) (*BatchOptimizeToursOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchOptimizeTours[0:len((*c.CallOptions).BatchOptimizeTours):len((*c.CallOptions).BatchOptimizeTours)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.fleetRoutingClient.BatchOptimizeTours(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchOptimizeToursOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *fleetRoutingGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OptimizeTours sends an OptimizeToursRequest containing a ShipmentModel and returns an
// OptimizeToursResponse containing ShipmentRoutes, which are a set of
// routes to be performed by vehicles minimizing the overall cost.
//
// A ShipmentModel model consists mainly of Shipments that need to be
// carried out and Vehicles that can be used to transport the Shipments.
// The ShipmentRoutes assign Shipments to Vehicles. More specifically,
// they assign a series of Visits to each vehicle, where a Visit
// corresponds to a VisitRequest, which is a pickup or delivery for a
// Shipment.
//
// The goal is to provide an assignment of ShipmentRoutes to Vehicles that
// minimizes the total cost where cost has many components defined in the
// ShipmentModel.
func (c *fleetRoutingRESTClient) OptimizeTours(ctx context.Context, req *optimizationpb.OptimizeToursRequest, opts ...gax.CallOption) (*optimizationpb.OptimizeToursResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:optimizeTours", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).OptimizeTours[0:len((*c.CallOptions).OptimizeTours):len((*c.CallOptions).OptimizeTours)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &optimizationpb.OptimizeToursResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// BatchOptimizeTours optimizes vehicle tours for one or more OptimizeToursRequest
// messages as a batch.
//
// This method is a Long Running Operation (LRO). The inputs for optimization
// (OptimizeToursRequest messages) and outputs (OptimizeToursResponse
// messages) are read/written from/to Cloud Storage in user-specified
// format. Like the OptimizeTours method, each OptimizeToursRequest
// contains a ShipmentModel and returns an OptimizeToursResponse
// containing ShipmentRoutes, which are a set of routes to be performed by
// vehicles minimizing the overall cost.
func (c *fleetRoutingRESTClient) BatchOptimizeTours(ctx context.Context, req *optimizationpb.BatchOptimizeToursRequest, opts ...gax.CallOption) (*BatchOptimizeToursOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:batchOptimizeTours", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &BatchOptimizeToursOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *fleetRoutingRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// BatchOptimizeToursOperation manages a long-running operation from BatchOptimizeTours.
type BatchOptimizeToursOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// BatchOptimizeToursOperation returns a new BatchOptimizeToursOperation from a given name.
// The name must be that of a previously created BatchOptimizeToursOperation, possibly from a different process.
func (c *fleetRoutingGRPCClient) BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation {
	return &BatchOptimizeToursOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// BatchOptimizeToursOperation returns a new BatchOptimizeToursOperation from a given name.
// The name must be that of a previously created BatchOptimizeToursOperation, possibly from a different process.
func (c *fleetRoutingRESTClient) BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &BatchOptimizeToursOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchOptimizeToursOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*optimizationpb.BatchOptimizeToursResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp optimizationpb.BatchOptimizeToursResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *BatchOptimizeToursOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*optimizationpb.BatchOptimizeToursResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp optimizationpb.BatchOptimizeToursResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *BatchOptimizeToursOperation) Metadata() (*optimizationpb.AsyncModelMetadata, error) {
	var meta optimizationpb.AsyncModelMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchOptimizeToursOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchOptimizeToursOperation) Name() string {
	return op.lro.Name()
}
