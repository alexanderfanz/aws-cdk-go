package awswafv2


// The part of the web request that you want AWS WAF to inspect.
//
// Include the single `FieldToMatch` type that you want to inspect, with additional specifications as needed, according to the type. You specify a single request component in `FieldToMatch` for each rule statement that requires it. To inspect more than one component of the web request, create a separate rule statement for each component.
//
// Example JSON for a `QueryString` field to match:
//
// `"FieldToMatch": { "QueryString": {} }`
//
// Example JSON for a `Method` field to match specification:
//
// `"FieldToMatch": { "Method": { "Name": "DELETE" } }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   fieldToMatchProperty := &FieldToMatchProperty{
//   	AllQueryArguments: allQueryArguments,
//   	Body: &BodyProperty{
//   		OversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	Cookies: &CookiesProperty{
//   		MatchPattern: &CookieMatchPatternProperty{
//   			All: all,
//   			ExcludedCookies: []*string{
//   				jsii.String("excludedCookies"),
//   			},
//   			IncludedCookies: []*string{
//   				jsii.String("includedCookies"),
//   			},
//   		},
//   		MatchScope: jsii.String("matchScope"),
//   		OversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	Headers: &HeadersProperty{
//   		MatchPattern: &HeaderMatchPatternProperty{
//   			All: all,
//   			ExcludedHeaders: []*string{
//   				jsii.String("excludedHeaders"),
//   			},
//   			IncludedHeaders: []*string{
//   				jsii.String("includedHeaders"),
//   			},
//   		},
//   		MatchScope: jsii.String("matchScope"),
//   		OversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	JsonBody: &JsonBodyProperty{
//   		MatchPattern: &JsonMatchPatternProperty{
//   			All: all,
//   			IncludedPaths: []*string{
//   				jsii.String("includedPaths"),
//   			},
//   		},
//   		MatchScope: jsii.String("matchScope"),
//
//   		// the properties below are optional
//   		InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   		OversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	Method: method,
//   	QueryString: queryString,
//   	SingleHeader: singleHeader,
//   	SingleQueryArgument: singleQueryArgument,
//   	UriPath: uriPath,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html
//
type CfnWebACL_FieldToMatchProperty struct {
	// Inspect all query arguments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-allqueryarguments
	//
	AllQueryArguments interface{} `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// Inspect the request body as plain text.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// A limited amount of the request body is forwarded to AWS WAF for inspection by the underlying host service. For regional resources, the limit is 8 KB (8,192 bytes) and for CloudFront distributions, the limit is 16 KB (16,384 bytes). For CloudFront distributions, you can increase the limit in the web ACL's `AssociationConfig` , for additional processing fees.
	//
	// For information about how to handle oversized request bodies, see the `Body` object configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-body
	//
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// Inspect the request cookies.
	//
	// You must configure scope and pattern matching filters in the `Cookies` object, to define the set of cookies and the parts of the cookies that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's cookies and only the first 200 cookies are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize cookie content in the `Cookies` object. AWS WAF applies the pattern matching filters to the cookies that it receives from the underlying host service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-cookies
	//
	Cookies interface{} `field:"optional" json:"cookies" yaml:"cookies"`
	// Inspect the request headers.
	//
	// You must configure scope and pattern matching filters in the `Headers` object, to define the set of headers to and the parts of the headers that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's headers and only the first 200 headers are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize header content in the `Headers` object. AWS WAF applies the pattern matching filters to the headers that it receives from the underlying host service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-headers
	//
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Inspect the request body as JSON.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// A limited amount of the request body is forwarded to AWS WAF for inspection by the underlying host service. For regional resources, the limit is 8 KB (8,192 bytes) and for CloudFront distributions, the limit is 16 KB (16,384 bytes). For CloudFront distributions, you can increase the limit in the web ACL's `AssociationConfig` , for additional processing fees.
	//
	// For information about how to handle oversized request bodies, see the `JsonBody` object configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-jsonbody
	//
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// Inspect the HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-method
	//
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Inspect the query string.
	//
	// This is the part of a URL that appears after a `?` character, if any.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-querystring
	//
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Inspect a single header.
	//
	// Provide the name of the header to inspect, for example, `User-Agent` or `Referer` . This setting isn't case sensitive.
	//
	// Example JSON: `"SingleHeader": { "Name": "haystack" }`
	//
	// Alternately, you can filter and inspect all headers with the `Headers` `FieldToMatch` setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-singleheader
	//
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Inspect a single query argument.
	//
	// Provide the name of the query argument to inspect, such as *UserName* or *SalesRegion* . The name can be up to 30 characters long and isn't case sensitive.
	//
	// Example JSON: `"SingleQueryArgument": { "Name": "myArgument" }`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-singlequeryargument
	//
	SingleQueryArgument interface{} `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// Inspect the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldtomatch.html#cfn-wafv2-webacl-fieldtomatch-uripath
	//
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

