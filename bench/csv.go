package main

import (
	"fmt"
	"time"
)

type Record struct {
	Workload         string
	Ops              int
	Goroutines       int
	Depth            int
	Blockprofilerate int
	Run              int
	Duration         time.Duration
}

type Column struct {
	Name         string
	MarshalValue func(*Record) (string, error)
}

var Columns = []Column{
	{"workload", func(r *Record) (string, error) {
		return fmt.Sprintf("%s", r.Workload), nil
	}},
	{"ops", func(r *Record) (string, error) {
		return fmt.Sprintf("%d", r.Ops), nil
	}},
	{"goroutines", func(r *Record) (string, error) {
		return fmt.Sprintf("%d", r.Goroutines), nil
	}},
	{"depth", func(r *Record) (string, error) {
		return fmt.Sprintf("%d", r.Depth), nil
	}},
	{"blockprofilerate", func(r *Record) (string, error) {
		return fmt.Sprintf("%d", r.Blockprofilerate), nil
	}},
	{"run", func(r *Record) (string, error) {
		return fmt.Sprintf("%d", r.Run), nil
	}},
	{"ms", func(r *Record) (string, error) {
		return fmt.Sprintf("%f", r.Duration.Seconds()*1000), nil
	}},
}

func (r *Record) MarshalRecord() ([]string, error) {
	record := make([]string, len(Columns))
	for i, col := range Columns {
		val, err := col.MarshalValue(r)
		if err != nil {
			return nil, err
		}
		record[i] = val
	}
	return record, nil
}

func Headers() []string {
	headers := make([]string, len(Columns))
	for i, col := range Columns {
		headers[i] = col.Name
	}
	return headers
}
