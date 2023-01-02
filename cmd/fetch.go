package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

const coinApiUrl = "https://rest.coinapi.io/v1/exchangerate"

var (
	crypto, in, key string
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
	RunE:  fetch,
}

var fetch = func(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("%s/%s/%s", coinApiUrl, crypto, in)
	rb := new(ResponseBody)
	c := new(http.Client)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	if err = json.Unmarshal(body, rb); err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request to %s/%s/%s failed with status code %d", coinApiUrl, crypto, in, res.StatusCode)
	}
	fmt.Printf("At the time %s the price of %s in %s was %f\n", rb.Time, rb.AssetIdBase, rb.AssetIdQuote, rb.Rate)

	return nil
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringVarP(&crypto, "crypto", "c", "BTC", "The cryptocurrency to fetch")
	fetchCmd.Flags().StringVarP(&in, "in", "i", "USD", "The fiat currency the price will be returned in")
	fetchCmd.Flags().StringVarP(&key, "key", "k", "", "The API key from https://www.coinapi.io/pricing?apikey")
}
