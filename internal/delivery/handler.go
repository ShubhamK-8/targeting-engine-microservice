package delivery

import (
    "encoding/json"
    "net/http"
    "strings"
    "targeting-engine/internal/models"
    "targeting-engine/internal/rules"
)

func contains(slice []string, val string) bool {
    for _, v := range slice {
        if v == val {
            return true
        }
    }
    return false
}

func matchRule(rule models.TargetingRule, appID, os, country string) bool {
    if len(rule.IncludeAppIDs) > 0 && !contains(rule.IncludeAppIDs, appID) {
        return false
    }
    if len(rule.ExcludeAppIDs) > 0 && contains(rule.ExcludeAppIDs, appID) {
        return false
    }
    if len(rule.IncludeOS) > 0 && !contains(rule.IncludeOS, os) {
        return false
    }
    if len(rule.ExcludeOS) > 0 && contains(rule.ExcludeOS, os) {
        return false
    }
    if len(rule.IncludeCountry) > 0 && !contains(rule.IncludeCountry, country) {
        return false
    }
    if len(rule.ExcludeCountry) > 0 && contains(rule.ExcludeCountry, country) {
        return false
    }
    return true
}

func HandleDelivery(w http.ResponseWriter, r *http.Request) {
    app := r.URL.Query().Get("app")
    os := r.URL.Query().Get("os")
    country := r.URL.Query().Get("country")

    if app == "" || os == "" || country == "" {
        http.Error(w, `{"error": "missing one or more params"}`, http.StatusBadRequest)
        return
    }

    var matched []models.Campaign
    for _, rule := range rules.Rules {
        if matchRule(rule, app, os, country) {
            for _, camp := range rules.Campaigns {
                if camp.ID == rule.CampaignID && strings.ToUpper(camp.Status) == "ACTIVE" {
                    matched = append(matched, camp)
                }
            }
        }
    }

    if len(matched) == 0 {
        w.WriteHeader(http.StatusNoContent)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(matched)
}
