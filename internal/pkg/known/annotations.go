package known

const (
	// This exposes compute information based on the miner type.
	CPUAnnotation    = "apps.go-pro.io/vCPU"
	MemoryAnnotation = "apps.go-pro.io/memoryMb"
)

const (
	SkipVerifyAnnotation = "apps.go-pro.io/skip-verify"
)

var AllImmutableAnnotations = []string{
	CPUAnnotation,
	MemoryAnnotation,
}
