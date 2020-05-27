package avrogenv7_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/akokhanovskyi/avrogenv7/gen/avro"

	"gotest.tools/assert"
)

func TestAvroV7(t *testing.T) {
	in := &avro.TestRecord{
		Kvs: map[string]*avro.UnionStringBytesNull{
			"one": {
				UnionType: avro.UnionStringBytesNullTypeEnumBytes,
				Bytes:     []byte{'h', 'e', 'l', 'l', 'o'},
			},
		},
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	_ = in.Serialize(writer) // Errors ignored for brevity
	_ = writer.Flush()

	out, _ := avro.DeserializeTestRecord(&buf)

	assert.Equal(t, out, in)
	// UnionType is suddenly UnionStringBytesNullTypeEnumString and String == "hello"
}
