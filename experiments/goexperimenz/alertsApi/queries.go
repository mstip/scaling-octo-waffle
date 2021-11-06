package main

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type SystemInfo struct {
	CheckVal    string
	SystemId     string
	SystemName   string
	CustomerName string
}

func querySystemsWhereLastContactWasMinutesAgo(db *sql.DB, minutes int) ([]SystemInfo, error) {
	query := `
		SELECT 
			system_status.updated_at AS checkVal, systems.id AS systemId, systems.name as systemName, customers.name as customerName
		FROM 
			system_status 
		JOIN systems on system_status.system_id = systems.id 
		JOIN customers on systems.customer_id = customers.id
		WHERE 
			system_status.updated_at < DATE_SUB(now(), INTERVAL ? MINUTE);
	`
	return querySystems(db, query, minutes)
}

func querySystemsWithLowDisk(db *sql.DB, thresholdGB int) ([]SystemInfo, error) {
	query := `
	SELECT 
		system_status.disk_free_kb AS checkVal, systems.id AS systemId, systems.name as systemName, customers.name as customerName
	FROM 
		system_status 
	JOIN systems on system_status.system_id = systems.id 
	JOIN customers on systems.customer_id = customers.id
	WHERE 
		system_status.disk_free_kb < ?;
	`
	return querySystems(db, query, thresholdGB*1000000)
}

func querySystemsWithLowMemory(db *sql.DB, thresholdGB int) ([]SystemInfo, error) {
	query := `
	SELECT 
		system_status.disk_free_kb AS checkVal, systems.id AS systemId, systems.name as systemName, customers.name as customerName
	FROM 
		system_status 
	JOIN systems on system_status.system_id = systems.id 
	JOIN customers on systems.customer_id = customers.id
	WHERE 
		system_status.memory_free_kb < ?;
	`
	return querySystems(db, query, thresholdGB*1000000)
}

func querySystemsWithLowSwap(db *sql.DB, thresholdGB int) ([]SystemInfo, error) {
	query := `
	SELECT 
		system_status.disk_free_kb AS checkVal, systems.id AS systemId, systems.name as systemName, customers.name as customerName
	FROM 
		system_status 
	JOIN systems on system_status.system_id = systems.id 
	JOIN customers on systems.customer_id = customers.id
	WHERE 
		system_status.swap_free_kb < ? and system_status.swap_free_kb > 0;
	`
	return querySystems(db, query, thresholdGB*1000000)
}

func querySystems(db *sql.DB, query string, args ...interface{}) ([]SystemInfo, error) {
	rows, err := dbQuery(db, query, args...)
	if err != nil {
		return nil, errors.New("query systems error: " + err.Error())
	}

	var systems []SystemInfo

	for rows.Next() {
		var checkVal string
		var systemId string
		var systemName string
		var customerName string

		if err := rows.Scan(&checkVal, &systemId, &systemName, &customerName); err != nil {
			return nil, errors.New("Failed to scan: " + err.Error())
		}

		systems = append(systems, SystemInfo{
			CheckVal:     checkVal,
			SystemId:     systemId,
			SystemName:   systemName,
			CustomerName: customerName,
		})
	}
	return systems, nil
}

func dbQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	if err := db.PingContext(ctx); err != nil {
		cancel()
		return nil, errors.New("Failed to PingContext " + err.Error())
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		cancel()
		return nil, errors.New("Failed to QueryContext " + err.Error())
	}

	return rows, nil
}
