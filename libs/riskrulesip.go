package libs

import (
	"encoding/json"
	"sort"

	"go.uber.org/zap"
)

type Riskrules_ip struct {
	Data struct {
		Results []struct {
			CriticalityLabel string `json:"criticalityLabel"`
			Description      string `json:"description"`
			Categories       []struct {
				Name      string `json:"name"`
				Framework string `json:"framework"`
			} `json:"categories,omitempty"`
			Criticality     int           `json:"criticality"`
			RelatedEntities []interface{} `json:"relatedEntities"`
			Name            string        `json:"name"`
			Count           int           `json:"count"`
		} `json:"results"`
	} `json:"data"`
}

func RiskRulesUnwrapIP(body []byte) (Riskrules_ip, []string) {
	var result Riskrules_ip
	json.Unmarshal(body, &result)
	if result.Data.Results == nil {
		logmaker().Error("Error in RiskRulesUnwrapIP: ", zap.String("error", "No data returned"))
	}

	// Make a list of all the riskrules
	var riskrules []string
	for _, v := range result.Data.Results {
		riskrules = append(riskrules, v.Name)
	}
	// Sort the list of riskrules alphabetically
	sort.Strings(riskrules)
	return result, riskrules
}
