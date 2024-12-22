package known

const (
	// This exposes compute information based on the miner type.
	CPUAnnotation    = "apps.acex.io/vCPU"
	MemoryAnnotation = "apps.acex.io/memoryMb"
)

const (
	SkipVerifyAnnotation = "apps.acex.io/skip-verify"
)

var AllImmutableAnnotations = []string{
	CPUAnnotation,
	MemoryAnnotation,
}
