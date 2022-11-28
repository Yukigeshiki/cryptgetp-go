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
		url := fmt.Sprintf("%s/%s/%s", CoinApiUrl, crypto, in)

		rb := ResponseBody{}
		c := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Set("X-CoinAPI-Key", key)
		res, err := c.Do(req)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(body, &rb); err != nil {
			return err
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("unable to fetch price - check your API key, currency values, or try again later")
		}
		t, base, quote, r := rb.getValues()
		fmt.Println(fmt.Sprintf("At the time %s the price of %s in %s was %s", t, base, quote, r))

		return nil
	},
}

// getValues returns a tuple of response body values
func (rb *ResponseBody) getValues() (string, string, string, string) {
	return rb.Time, rb.AssetIdBase, rb.AssetIdQuote, fmt.Sprintf("%f", rb.Rate)
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringVarP(&crypto, "crypto", "c", "BTC", "The cryptocurrency to fetch")
	fetchCmd.Flags().StringVarP(&in, "in", "i", "USD", "The fiat currency the price will be returned in")
	fetchCmd.Flags().StringVarP(&key, "key", "k", "", "The API key from https://www.coinapi.io/pricing?apikey")
}
