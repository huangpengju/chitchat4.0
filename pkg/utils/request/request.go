package request

import (
	"net/http"
	"strings"

	"chitchat4.0/pkg/utils/set"
)

const (
	NamespaceNone = ""
	NamespaceRoot = "root"
)

const (
	GetOperation    = "get"    // get获取单个
	ListOperation   = "list"   // list获取列表
	CreateOperation = "create" // create创建
	UpdateOperation = "update" // update更新
	PatchOperation  = "patch"  // patch更新一部分
	DeleteOperation = "delete" // delete删除
)

// RequestInfoResolver 请求信息解析接口
type RequestInfoResolver interface {
	// NewRequestInfo 返回http请求中的信息。
	NewRequestInfo(req *http.Request) (*RequestInfo, error)
}

// RequestInfo 请求信息结构体
type RequestInfo struct {
	IsResourceRequest bool   // 是否资源请求
	Path              string // 请求的URL
	Verb              string // 请求的方法

	APIPrefix  string // API 前缀（URL前缀）
	APIGroup   string // API 分组
	APIVersion string // API 版本
	Namespace  string // 命名空间
	// Resource 是被请求的资源的名称.  This is not the kind.  For example: pods
	Resource    string // 资源
	Subresource string
	// Name is empty for some verbs, but if the request directly indicates a name (not in body content) then this field is filled in.
	// 对于某些动词，Name为空，但如果请求直接指示一个名称(不在正文内容中)，则填充此字段。
	Name string
	// Parts are the path parts for the request, always starting with /{resource}/{name}
	// Parts是请求的路径部分，总是以/{resource}/{name}开头。
	Parts []string
}

// RequestInfoFactory 请求信息工厂结构体
type RequestInfoFactory struct {
	// API(接口前缀) 比如：map[api:****/***/***]
	APIPrefixes set.String // APIPrefixes map[string]Empty // 空的字符串映射
}

// TODO write an integration test against the swagger doc to test the RequestInfo and match up behavior to responses
// TODO针对swagger文档编写集成测试，以测试RequestInfo并将行为与响应匹配
// NewRequestInfo returns the information from the http request.  If error is not nil, RequestInfo holds the information as best it is known before the failure
// NewRequestInfo返回http请求中的信息。如果error不是nil，RequestInfo将保留故障前已知的最佳信息
// It handles both resource and non-resource requests and fills in all the pertinent information for each.
// 它处理资源和非资源请求，并为每个请求填充所有相关信息。
// Valid Inputs(有效输入):
// Resource paths (资源路径)
// /apis/{api-group}/{version}/namespaces
// /api/{version}/namespaces
// /api/{version}/namespaces/{namespace}
// /api/{version}/namespaces/{namespace}/{resource}
// /api/{version}/namespaces/{namespace}/{resource}/{resourceName}
// /api/{version}/{resource}
// /api/{version}/{resource}/{resourceName}
//
// Special verbs with subresources:
// /api/{version}/watch/{resource}
// /api/{version}/watch/namespaces/{namespace}/{resource}
//
// NonResource paths
// /apis/{api-group}/{version}
// /apis/{api-group}
// /apis
// /api/{version}
// /api
// /healthz
// /
func (r *RequestInfoFactory) NewRequestInfo(req *http.Request) (*RequestInfo, error) {
	// start with a non-resource request until proven otherwise
	requestInfo := RequestInfo{
		IsResourceRequest: false,                       //  不是资源请求
		Path:              req.URL.Path,                // URL路径
		Verb:              strings.ToLower(req.Method), // 请求方法
	}

	currentParts := splitPath(req.URL.Path) // 把URL处理为[]string切片
	// URL 路径处理为切片后，判断是不是资源请求
	if len(currentParts) < 3 {
		// return a non-resource request
		// return 非资源请求
		return &requestInfo, nil
	}

	// 判断 currentParts[0] URL切片的第1个元素，是不是在 APIPrefixes 接口前缀中（字符串-映射）中
	if !r.APIPrefixes.Has(currentParts[0]) {
		// return a non-resource request
		// return 非资源请求
		return &requestInfo, nil
	}

	requestInfo.APIPrefix = currentParts[0] // 接口前缀=URL第1个元素
	currentParts = currentParts[1:]         // 对当前URL切片进行裁剪[1:]

	requestInfo.IsResourceRequest = true     // 是资源请求
	requestInfo.APIVersion = currentParts[0] // API 版本是当前URL第1个元素
	currentParts = currentParts[1:]          // 对当前URL切片进行裁剪[1:]

	switch req.Method { // 判断请求的方法
	case "POST":
		requestInfo.Verb = CreateOperation // 对请求方法进行描述，create 创建
	case "GET", "HEAD":
		requestInfo.Verb = GetOperation // 对请求方法进行描述，get 获取
	case "PUT":
		requestInfo.Verb = UpdateOperation // 对请求方法进行描述，update 修改
	case "PATCH":
		requestInfo.Verb = PatchOperation // 对请求方法进行描述，patch
	case "DELETE":
		requestInfo.Verb = DeleteOperation // 对请求方法进行描述，delete 删除
	default:
		requestInfo.Verb = ""
	}

	// URL forms: /namespaces/{namespace}/{kind}/*, where parts are adjusted to be relative to kind
	// URL 表单:/namespaces/{namespace}/{kind}/*，其中的part被调整为相对于kind

	if currentParts[0] == "namespaces" {
		if len(currentParts) > 1 {
			requestInfo.Namespace = currentParts[1]

			// if there is another step after the namespace name and it is not a known namespace subresource
			// move currentParts to include it as a resource in its own right
			// 如果在命名空间名称后面有另一个步骤，并且它不是一个已知的命名空间子资源，移动currentParts以将其作为一个资源包含在它自己的权利中
			if len(currentParts) > 2 {
				currentParts = currentParts[2:]
			}
		}
	} else {
		requestInfo.Namespace = NamespaceRoot // 根命名空间
	}

	// parsing successful, so we now know the proper value for .Parts
	// 解析成功了，所以我们现在知道了.Parts的正确值
	requestInfo.Parts = currentParts

	// parts look like（部分看起来像）: resource/resourceName/subresource/other/stuff/we/don't/interpret
	switch {
	case len(requestInfo.Parts) >= 3:
		requestInfo.Subresource = requestInfo.Parts[2]
		fallthrough
	case len(requestInfo.Parts) >= 2:
		requestInfo.Name = requestInfo.Parts[1]
		fallthrough
	case len(requestInfo.Parts) >= 1:
		requestInfo.Resource = requestInfo.Parts[0]
	}

	// if there's no name on the request and we thought it was a get before, then the actual verb is a list or a watch
	// 如果请求中没有名字，而我们之前认为它是一个get，那么实际的动词是一个列表或一个手表
	if len(requestInfo.Name) == 0 && requestInfo.Verb == GetOperation {
		requestInfo.Verb = ListOperation
	}

	return &requestInfo, nil
}

// splitPath 返回URL路径的分段(处理为切片)
func splitPath(path string) []string {
	path = strings.Trim(path, "/") // Trim 将字符串中指定的前缀和后缀去除，并返回去除后的字符串。
	if path == "" {
		return []string{}
	}
	return strings.Split(path, "/") // Split 函数用于将字符串按照指定的分隔符进行分割，并返回分割后的字符串切片。
}
