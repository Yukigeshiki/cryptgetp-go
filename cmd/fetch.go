package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

const CoinApiUrl = "https://rest.coinapi.io/v1/exchangerate"

var (
	crypto string
	in     string
	key    string
)

type ResponseBody struct {
	Time         string  `json:"time"`
	AssetIdBase  string  `json:"asset_id_base"`
	AssetIdQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch the price of a given cryptocurrency (--crypto) returned in a given fiat currency (--in).",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err  error
			body []byte
			req  *http.Request
			res  *http.Response
		)

		url := fmt.Sprintf("%s/%s/%s", CoinApiUrl, crypto, in)

		rb := ResponseBody{}
		c := &http.Client{}
		if req, err = http.NewRequest("GET", url, nil); err != nil {
			return err
		}
		req.Header.Set("X-CoinAPI-Key", key)
		if res, err = c.Do(req); err != nil {
			return err
		}

		if body, err = io.ReadAll(res.Body); err != nil {
			return err
		}
		if err = json.Unmarshal(body, &rb); err != nil {
			return err
		}

		t, base, quote, r := rb.getValues()
		if t != "" {
			fmt.Println(fmt.Sprintf("At the time %s the price of %s in %s was %s", t, base, quote, r))
		} else {
			fmt.Println("Unable to fetch price. Check your API key, currency values, or try again later.")
		}

		return nil
	},
}

// getValues returns a tuple of response body values
func (rb ResponseBody) getValues() (string, string, string, string) {
	return rb.Time, rb.AssetIdBase, rb.AssetIdQuote, fmt.Sprintf("%f", rb.Rate)
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringVarP(&crypto, "crypto", "c", "BTC", "The cryptocurrency to fetch")
	fetchCmd.Flags().StringVarP(&in, "in", "i", "USD", "The fiat currency the price will be returned in")
	fetchCmd.Flags().StringVarP(&key, "key", "k", "", "The API key from https://www.coinapi.io/pricing?apikey")
}
