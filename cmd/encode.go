package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"os"
)

var qrEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a string into a QR code PNG image.",
	Long:  `Will encode a text to QR code and saves the QR code to file in .PNG format`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := cmd.Flags().GetString("text")
		if err != nil {
			fmt.Printf("The flag has invalid contents: %s\n", err.Error())
			os.Exit(1)
			return
		}

		if len(msg) == 0 {
			fmt.Println("Flag text is mandatory!")
			cmd.Usage()
			os.Exit(1)
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("The flag has invalid contents: %s\n", err.Error())
		}

		if output == "" {
			output = "qr.png"
		}

		err = qrcode.WriteFile(msg, qrcode.Medium, 256, output)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(qrEncodeCmd)
	qrEncodeCmd.Flags().StringP("text", "t", "", "The text you want to encode.\n")
	qrEncodeCmd.Flags().StringP("output", "o", "", "The path to save the generated QR code PNG image.\n")
}
