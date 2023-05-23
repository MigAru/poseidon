package errors

var (
	BlobUnknown = Error{
		Code:    "BLOB_UNKNOWN",
		Message: "blob unknown to registry",
		Detail: "This error may be returned when a blob is unknown " +
			"to the registry in a specified repository. " +
			"This can be returned with a standard get or " +
			"if a manifest references an unknown layer during upload",
	}
	BlobUploadInvalid = Error{
		Code:    "BLOB_UPLOAD_INVALID",
		Message: "blob upload invalid",
		Detail:  "The blob upload encountered an error and can no longer proceed.",
	}
	BlobUploadUnknown = Error{
		Code:    "BLOB_UPLOAD_UNKNOWN",
		Message: "blob upload unknown to registry",
		Detail:  "If a blob upload has been cancelled or was never started, this error code may be returned.",
	}
	DigestInvalid = Error{
		Code:    "DIGEST_INVALID",
		Message: "provided digest did not match uploaded content",
		Detail: "When a blob is uploaded, " +
			"the registry will check that the content matches the digest provided by the client. " +
			"The error may include a detail structure with the key “digest”, " +
			"including the invalid digest string. This error may also be returned " +
			"when a manifest includes an invalid layer digest.",
	}
)
