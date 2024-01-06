package vlc

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"

	"github.com/bazdalaz/archiver/lib/compression/vlc/table"

)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}


func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)

	encoded := encodeBin(str, tbl)

	return buildEncodedFile(tbl, encoded)

}

func buildEncodedFile(tbl  table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var buf bytes.Buffer

	buf.Write(encodeInt(len(encodedTbl)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)
	buf.Write(splitByChunks(data, chunkSize).Bytes())
	return buf.Bytes()

}

func encodeInt(num int) []byte {
	res:= make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err:=gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can't encode table", err)
	}

	return tableBuf.Bytes()
}

func decodeTable(tblBinary []byte) (table.EncodingTable) {
	var tbl table.EncodingTable

	if err := gob.NewDecoder(bytes.NewReader(tblBinary)).Decode(&tbl); err != nil {
		log.Fatal("can't decode table", err)
	}

	return tbl
}

// Decode decodes string from hex implemented by vlc algorithm
func (ed EncoderDecoder) Decode(encodedData []byte) string {

	tbl, data := parseFile(encodedData)
	
	return tbl.Decode(data)


}

func parseFile(data []byte) (table.EncodingTable, string) {
	const(
		tableSizeBytesCount = 4
		dataSizeBytesCount = 4
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]
	tbl := decodeTable(tblBinary)
	body := NeBinChunks(data).Join()

	return tbl, body[:dataSize]
}


func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, table))
	}

	return buf.String()
}


func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown character " + string(ch))
	}

	return res

}

