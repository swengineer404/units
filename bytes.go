package units

import "strings"

var (
	bytesUnitMap = MakeUnitMap("iB", "B", 1024)
)

var (
	metricBytesUnitMap = MakeUnitMap("B", "B", 1000)
)

// MetricBytes are SI byte units (1000 bytes in a kilobyte).
type MetricBytes SI

// SI base-10 byte units.
const (
	Kilobyte MetricBytes = 1000
	KB                   = Kilobyte
	Megabyte             = Kilobyte * 1000
	MB                   = Megabyte
	Gigabyte             = Megabyte * 1000
	GB                   = Gigabyte
	Terabyte             = Gigabyte * 1000
	TB                   = Terabyte
	Petabyte             = Terabyte * 1000
	PB                   = Petabyte
	Exabyte              = Petabyte * 1000
	EB                   = Exabyte
)

// ParseMetricBytes parses base-10 metric byte units. That is, KB is 1000 bytes.
func ParseMetricBytes(s string) (MetricBytes, error) {
	n, err := ParseUnit(strings.ReplaceAll(s, ":", ""), metricBytesUnitMap)
	return MetricBytes(n), err
}

// TODO: represents 1000B as uppercase "KB", while SI standard requires "kB".
func (m MetricBytes) String() string {
	return ToString(int64(m), 1000, "B", "B")
}

// ParseStrictBytes supports both iB and B suffixes for base 2 and metric,
// respectively. That is, KiB represents 1024 and kB, KB represent 1000.
func ParseStrictBytes(s string) (int64, error) {
	n, err := ParseUnit(s, bytesUnitMap)
	if err != nil {
		n, err = ParseUnit(s, metricBytesUnitMap)
	}
	return int64(n), err
}
