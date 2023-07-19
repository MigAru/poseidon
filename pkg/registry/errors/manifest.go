package errors

var (
	ManifestBlobUnknown = DockerError{
		Code:    "MANIFEST_BLOB_UNKNOWN",
		Message: "blob unknown to registry",
		Detail:  "This error may be returned when a manifest blob is unknown to the registry.",
	}
	ManifestInvalid = DockerError{
		Code:    "MANIFEST_INVALID",
		Message: "manifest invalid",
		Detail: "During upload, manifests undergo several checks ensuring validity. " +
			"If those checks fail, this error may be returned, " +
			"unless a more specific error is included. " +
			"The detail will contain information the failed validation.",
	}

	ManifestUnverified = DockerError{
		Code:    "MANIFEST_UNVERIFIED",
		Message: "manifest failed signature verification",
		Detail:  "During manifest upload, if the manifest fails signature verification, this error will be returned.",
	}
)
