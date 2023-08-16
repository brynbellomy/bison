package distru

import (
	"fmt"
	"net/http"
	"strconv"

	"bison/distru/company"
	"bison/utils"
)

type Client struct {
	AuthToken string
}

func (c Client) request(method string, url string, body any, headers http.Header, response any) (http.Header, error) {
	if headers == nil {
		headers = http.Header{}
	}
	headers["Authorization"] = []string{"Bearer " + c.AuthToken}
	return utils.JSONReq(method, url, body, headers, response)
}

func (c Client) CompanyCreate(req company.CreateRequest) (int, error) {
	type Response struct {
		ID int `json:"id"`
	}

	var resp Response
	_, err := c.request("POST", "https://app.distru.com/api/v1/company_relationships", req, nil, &resp)
	if err != nil {
		return 0, err
	}
	return resp.ID, nil
}

func (c Client) CompanyGet(id int) (company.Company, error) {
	var resp company.Company
	_, err := c.request("GET", fmt.Sprintf("https://app.distru.com/api/v1/company_relationships/%d", id), nil, nil, &resp)
	if err != nil {
		return company.Company{}, err
	}
	return resp, err
}

func (c Client) CompanyGetAllIDs() ([]uint64, error) {
	data, err := c.CompanyGetAllTableData()
	if err != nil {
		return nil, err
	}
	return utils.Map(data, func(data company.TableData) uint64 { return uint64(data.ID) }), nil
}

func (c Client) CompanyUpdate(id int, req company.UpdateRequest) error {
	_, err := c.request("PUT", fmt.Sprintf("https://app.distru.com/api/v1/company_relationships/%d", id), req, nil, nil)
	return err
}

func (c Client) CompanyDelete(ids []uint64) error {
	for _, id := range ids {
		_, err := c.request("DELETE", fmt.Sprintf("https://app.distru.com/api/v1/company_relationships/%d", id), nil, nil, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c Client) CompanyGetAllTableData() ([]company.TableData, error) {
	respHeaders, err := c.request("GET", "https://app.distru.com/api/v1/company_relationships/table_data?page_size=120", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	totalPages, err := strconv.ParseUint(respHeaders["Total-Pages"][0], 10, 64)
	if err != nil {
		return nil, err
	}

	allPageNumbers := utils.Range(1, totalPages+1)

	return utils.ReduceWithError(allPageNumbers, func(into []company.TableData, page uint64) ([]company.TableData, error) {
		var resp company.TableDataResponse
		_, err := c.request("GET", fmt.Sprintf("https://app.distru.com/api/v1/company_relationships/table_data?page_size=120&page=%d", page), nil, nil, &resp)
		if err != nil {
			return nil, err
		}
		return append(into, resp.TableData...), nil
	}, nil)
}
