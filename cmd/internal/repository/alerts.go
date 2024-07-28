package repository

import (

    "price-alert-app/cmd/internal/models"
    "price-alert-app/cmd/internal/config"
)

func CreateAlert(alert models.Alert) error {
    query := `INSERT INTO alerts (user_id, coin, target_price, status) VALUES ($1, $2, $3, $4)`
    _, err := config.DB.Exec(query, alert.UserID, alert.Coin, alert.TargetPrice, alert.Status)
    return err
}

func DeleteAlert(id string) error {
    query := `DELETE FROM alerts WHERE id = $1`
    _, err := config.DB.Exec(query, id)
    return err
}

func FetchAlerts(status string) ([]models.Alert, error) {
    var alerts []models.Alert
    query := `SELECT id, user_id, coin, target_price, status FROM alerts WHERE status = $1`
    rows, err := config.DB.Query(query, status)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var alert models.Alert
        if err := rows.Scan(&alert.ID, &alert.UserID, &alert.Coin, &alert.TargetPrice, &alert.Status); err != nil {
            return nil, err
        }
        alerts = append(alerts, alert)
    }
    return alerts, nil
}
