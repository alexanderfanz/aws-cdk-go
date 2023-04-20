package awsappmesh


// Represent the HTTP2 Node Listener prorperty.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var healthCheck healthCheck
//   var mutualTlsValidationTrust mutualTlsValidationTrust
//   var subjectAlternativeNames subjectAlternativeNames
//   var tlsCertificate tlsCertificate
//
//   http2VirtualNodeListenerOptions := &Http2VirtualNodeListenerOptions{
//   	ConnectionPool: &Http2ConnectionPool{
//   		MaxRequests: jsii.Number(123),
//   	},
//   	HealthCheck: healthCheck,
//   	OutlierDetection: &OutlierDetection{
//   		BaseEjectionDuration: duration,
//   		Interval: duration,
//   		MaxEjectionPercent: jsii.Number(123),
//   		MaxServerErrors: jsii.Number(123),
//   	},
//   	Port: jsii.Number(123),
//   	Timeout: &HttpTimeout{
//   		Idle: duration,
//   		PerRequest: duration,
//   	},
//   	Tls: &ListenerTlsOptions{
//   		Certificate: tlsCertificate,
//   		Mode: awscdk.Aws_appmesh.TlsMode_STRICT,
//
//   		// the properties below are optional
//   		MutualTlsValidation: &MutualTlsValidation{
//   			Trust: mutualTlsValidationTrust,
//
//   			// the properties below are optional
//   			SubjectAlternativeNames: subjectAlternativeNames,
//   		},
//   	},
//   }
//
// Experimental.
type Http2VirtualNodeListenerOptions struct {
	// Connection pool for http2 listeners.
	// Experimental.
	ConnectionPool *Http2ConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for HTTP protocol.
	// Experimental.
	Timeout *HttpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}

