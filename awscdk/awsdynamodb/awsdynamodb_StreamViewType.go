package awsdynamodb


// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_StreamSpecification.html
//
// Experimental.
type StreamViewType string

const (
	// The entire item, as it appears after it was modified, is written to the stream.
	// Experimental.
	StreamViewType_NEW_IMAGE StreamViewType = "NEW_IMAGE"
	// The entire item, as it appeared before it was modified, is written to the stream.
	// Experimental.
	StreamViewType_OLD_IMAGE StreamViewType = "OLD_IMAGE"
	// Both the new and the old item images of the item are written to the stream.
	// Experimental.
	StreamViewType_NEW_AND_OLD_IMAGES StreamViewType = "NEW_AND_OLD_IMAGES"
	// Only the key attributes of the modified item are written to the stream.
	// Experimental.
	StreamViewType_KEYS_ONLY StreamViewType = "KEYS_ONLY"
)

