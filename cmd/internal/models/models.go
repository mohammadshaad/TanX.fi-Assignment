package models

type Alert struct {
    ID          int     `json:"id"`
    UserID      int     `json:"user_id"`
    Coin        string  `json:"coin"`
    TargetPrice float64 `json:"target_price"`
    Status      string  `json:"status"`
}
