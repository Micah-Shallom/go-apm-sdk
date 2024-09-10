package telex

import "log"

func (c *Client) ReportError(err interface{}, username string) APMMetrics {
	errMsg, ok := err.(string)
	if !ok {
		log.Println("Error is not a string")
		return APMMetrics{}
	}

	return APMMetrics{
		EventName: "application error",
		Message:   errMsg,
		Status:    "error",
		Username:  username,
	}
}
