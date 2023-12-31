package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazdalaz/archiver/lib/compression"
	"github.com/bazdalaz/archiver/lib/compression/vlc"
	"github.com/bazdalaz/archiver/lib/compression/vlc/table/shannon_fano"
	"github.com/spf13/cobra"
)


var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

// TODO: refactor this to get filenam
const unpackedExtension = ".txt"
func unpack(cmd *cobra.Command, args []string) {


	var decoder compression.Decoder

	if len(args) == 0 {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New(shannon_fano.Generator{})
	
	default:
		cmd.PrintErr("Unknown method")
	}

	filePath := args[0] //BUG: check if file exists

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

// TODO: refactor this function
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + unpackedExtension	
}
func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}


}
