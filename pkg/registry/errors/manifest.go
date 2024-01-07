package errors

var (
	GetManifest = DockerError{
		Code:    "MANIFEST_GET",
		Message: "server error to get manifest",
		Detail:  "This error may be returned when a get manifest server error.",
	}
	CreateManifest = DockerError{
		Code:    "MANIFEST_CREATE",
		Message: "server error to update/create manifest",
		Detail:  "This error may be returned when a filesystem on server return error",
	}
	ManifestBlobUnknown = DockerError{
		Code:    "MANIFEST_BLOB_UNKNOWN",
		Message: "blob unknown to registry",
		Detail:  "This error may be returned when a manifest blob is unknown to the registry.",
	}
	ManifestInvalid = DockerError{
		Code:    "MANIFEST_INVALID",
		Message: "manifest invalid",
		Detail: `During uploads, manifests undergo several checks ensuring validity.
			If those checks fail, this error may be returned,
			unless a more specific error is included.
			The detail will contain information the failed validation.`,
	}
	ManifestUnverified = DockerError{
		Code:    "MANIFEST_UNVERIFIED",
		Message: "manifest failed signature verification",
		Detail:  "During manifest uploads, if the manifest fails signature verification, this error will be returned.",
	}
)
