package errors

var (
	NameInvalid = Error{
		Code:    "NAME_INVALID",
		Message: "invalid repository name",
		Detail:  "Invalid repository name encountered either during manifest validation or any API operation.",
	}
	NameUnknown = Error{
		Code:    "NAME_UNKNOWN",
		Message: "repository name not known to registry",
		Detail:  "This is returned if the name used during an operation is unknown to the registry.",
	}
	PaginationNumberInvalid = Error{
		Code:    "PAGINATION_NUMBER_INVALID",
		Message: "invalid number of results requested",
		Detail:  "Returned when the “n” parameter (number of results to return) is not an integer, or “n” is negative.",
	}
	RangeInvalid = Error{
		Code:    "RANGE_INVALID",
		Message: "invalid content range",
		Detail: "When a layer is uploaded, the provided range is checked against the uploaded chunk. " +
			"This error is returned if the range is out of order.",
	}
	TagInvalid = Error{
		Code:    "TAG_INVALID",
		Message: "manifest tag did not match URI",
		Detail: "During a manifest upload, if the tag in the manifest does not match " +
			"the uri tag, this error will be returned.",
	}
	Unauthorized = Error{
		Code:    "UNAUTHORIZED",
		Message: "authentication required",
		Detail: "The access controller was unable to authenticate the client. " +
			"Often this will be accompanied by a Www-Authenticate HTTP response header indicating how to authenticate.",
	}
	Denied = Error{
		Code:    "DENIED",
		Message: "requested access to the resource is denied",
		Detail:  "The access controller denied access for the operation on a resource.",
	}
	Unsupported = Error{
		Code:    "UNSUPPORTED",
		Message: "The operation is unsupported.",
		Detail:  "The operation was unsupported due to a missing implementation or invalid set of parameters.",
	}
)
