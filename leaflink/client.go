package leaflink

import (
	"fmt"
	"net/http"
	"time"

	"bison/utils"
)

type Client struct {
	AppToken string
}

const (
	apiBaseURL = "https://www.leaflink.com/api/v2"
)

func (c Client) request(method string, url string, body any, headers http.Header, response any) (http.Header, error) {
	if headers == nil {
		headers = http.Header{}
	}
	headers["Authorization"] = []string{"App " + c.AppToken}
	return utils.JSONReq(method, apiBaseURL+url, body, headers, response)
}

func (c Client) GetOrdersSince(when time.Time) ([]Order, error) {
	var (
		results []Order
		offset  int
	)
	for {
		type Response struct {
			Results []Order `json:"results"`
			Next    *string `json:"next"`
		}

		var resp Response
		_, err := c.request("GET", fmt.Sprintf("/orders-received/?offset=%d&created_on__gt=%s", offset, when), nil, nil, &resp)
		if err != nil {
			return nil, err
		}
		results = append(results, resp.Results...)
		offset += len(resp.Results)

		if resp.Next == nil {
			break
		}
	}
	return results, nil
}

func (c Client) GetAllOrderSalesReps() ([]OrderSalesRep, error) {
	var (
		results []OrderSalesRep
		offset  int
	)
	for {
		type Response struct {
			Results []OrderSalesRep `json:"results"`
			Next    *string         `json:"next"`
		}

		var resp Response
		_, err := c.request("GET", fmt.Sprintf("/order-sales-reps/?offset=%d", offset), nil, nil, &resp)
		if err != nil {
			return nil, err
		}
		results = append(results, resp.Results...)
		offset += len(resp.Results)

		if resp.Next == nil {
			break
		}
	}
	return results, nil
}

func (c Client) DeleteOrderSalesRep(id int) error {
	// Yeah, the following route silently fails without the trailing slash. Lol, Django... sigh
	// (you can tell because you get a 200 response instead of a 204)
	_, err := c.request("DELETE", fmt.Sprintf("/order-sales-reps/%d/", id), nil, nil, nil)
	return err
}

func (c Client) SetOrderSalesRep(orderNum string, rep int) error {
	// This route will fail if you add a rep already assigned to this order

	type SetSalesRep struct {
		Order string `json:"order"`
		Rep   int    `json:"rep"`
	}

	body := SetSalesRep{
		Order: orderNum,
		Rep:   rep,
	}

	_, err := c.request("POST", "/order-sales-reps/", body, nil, nil)
	return err
}
