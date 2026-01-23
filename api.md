# Poi

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#PoiListResponse">PoiListResponse</a>

Methods:

- <code title="get /poi">client.Poi.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#PoiService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#PoiListParams">PoiListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#PoiListResponse">PoiListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Referencedata

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#Country">Country</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ReferencedataGetResponse">ReferencedataGetResponse</a>

Methods:

- <code title="get /referencedata">client.Referencedata.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ReferencedataService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ReferencedataGetParams">ReferencedataGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ReferencedataGetResponse">ReferencedataGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profile

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ProfileAuthenticateResponse">ProfileAuthenticateResponse</a>

Methods:

- <code title="post /profile/authenticate">client.Profile.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ProfileService.Authenticate">Authenticate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ProfileAuthenticateParams">ProfileAuthenticateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#ProfileAuthenticateResponse">ProfileAuthenticateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Comment

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#CommentSubmitResponse">CommentSubmitResponse</a>

Methods:

- <code title="post /comment">client.Comment.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#CommentService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#CommentSubmitParams">CommentSubmitParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#CommentSubmitResponse">CommentSubmitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Mediaitem

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#MediaitemNewResponse">MediaitemNewResponse</a>

Methods:

- <code title="post /mediaitem">client.Mediaitem.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#MediaitemService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#MediaitemNewParams">MediaitemNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#MediaitemNewResponse">MediaitemNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OpenAPI

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#OpenAPIGetResponse">OpenAPIGetResponse</a>

Methods:

- <code title="get /openapi">client.OpenAPI.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#OpenAPIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go">ocm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/ocm-go#OpenAPIGetResponse">OpenAPIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
