package hubplanner

import (
	"fmt"
	"io"
	"net/http"
	"recordari/envs"
)

func DoRequest(method string, endpoint string, body io.Reader) (*http.Response, error) {

	c := &http.Client{}

	req, errReq := http.NewRequest(method, envs.Get("HUBPLANNER_BASE_URL")+endpoint, body)
	req.Header.Add("Authorization", envs.Get("HUBPLANNER_API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	res, errRes := c.Do(req)

	if errReq != nil || res.StatusCode != 200 || errRes != nil {
		if res.StatusCode == 401 {
			return nil, fmt.Errorf("Petición no autorizada ", req.Response.Status)
		}
		if res.StatusCode == 404 {
			return nil, fmt.Errorf("Recurso no encontrado ", req.Response.Status)
		}
		if res.StatusCode == 500 {
			return nil, fmt.Errorf("Error interno del servidor ", req.Response.Status)
		}
		if errReq != nil {
			return nil, fmt.Errorf("Ha ocurrido un error al realizar la request ", errReq)
		}
		if errRes != nil {
			return nil, fmt.Errorf("Ha ocurrido un error al realizar la petición ", errRes)
		}
	}

	return res, nil
}
