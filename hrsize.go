package main

import (
	"fmt"
)

// Human Readable Size
type HRSize uint64

const (
	KB = 1 << 10
	MB = KB << 10
	GB = MB << 10
	TB = GB << 10
)

func (s HRSize) String() string {
	switch {
	case s >= TB:
		return fmt.Sprintf("%.1fT", float64(s)/float64(TB))
	case s >= GB:
		return fmt.Sprintf("%.1fG", float64(s)/float64(GB))
	case s >= MB:
		return fmt.Sprintf("%.1fM", float64(s)/float64(MB))
	case s >= KB:
		return fmt.Sprintf("%.1fK", float64(s)/float64(KB))
	default:
		return fmt.Sprintf("%d", s)
	}
}
