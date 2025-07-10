package vmx

type DiskType int

const (
	DiskTypeSingleFileGrowable = iota
	DiskTypeMultiFileGrowable
	DiskTypeSingleFilePreallocated
	DiskTypeMultiFilePreallocated
	DiskTypeESXPreallocated
	DiskTypeStreaming
	DiskTypeThin
)

var diskTypeName = map[DiskType]string{
	DiskTypeSingleFileGrowable:     "single_file_growable",
	DiskTypeMultiFileGrowable:      "multiple_file_growable",
	DiskTypeSingleFilePreallocated: "sigle_file_preallocated",
	DiskTypeMultiFilePreallocated:  "multiple_file_preallocated",
	DiskTypeESXPreallocated:        "preallocated_ESX",
	DiskTypeStreaming:              "compressed_streaming_optimized",
	DiskTypeThin:                   "thin_provisioned",
}

func (dt DiskType) String() string {
	return diskTypeName[dt]
}
