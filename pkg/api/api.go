package api

type TransformOptions struct {
	Erase         bool // Ignored space
	CaseSensitive bool // If falsey. Case will be ignored.
}

type TransformResult struct {
	Code []byte
}

func Transform(input string, options TransformOptions) TransformResult {
	return transformImpl(input, options)
}
