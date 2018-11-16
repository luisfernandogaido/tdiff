package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	//ret, err := calcula(os.Args[1:]...)
	ret, err := calcula("15:04", "13:00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}

func calcula(pars ...string) (string, error) {
	var p1, p2 string
	switch len(pars) {
	case 1:
		p1 = time.Now().Format("15:04")
		p2 = pars[0]
	case 2:
		p1 = pars[0]
		p2 = pars[1]
	default:
		return "", errors.New("um ou dois parâmetros esperados")
	}
	if strings.Contains(p2, "h") || strings.Contains(p2, "m") {
		t, err := calcTim(p1, p2)
		return t.Format("2006-01-02 15:04"), err
	}
	dur, err := calcDur(p1, p2)
	return dur.String(), err
}

func calcDur(p1, p2 string) (time.Duration, error) {
	ini, err := analisaDataHora(p2)
	if err != nil {
		return 0, err
	}
	fim, err := analisaDataHora(p1)
	if err != nil {
		return 0, err
	}
	return fim.Sub(ini), nil
}

func calcTim(p1, p2 string) (time.Time, error) {
	t, err := analisaDataHora(p1)
	if err != nil {
		return time.Time{}, err
	}
	d, err := time.ParseDuration("-" + p2)
	if err != nil {
		return time.Time{}, err
	}
	return t.Add(d), nil

}

func analisaDataHora(s string) (time.Time, error) {
	var (
		t   time.Time
		err error
	)
	switch len(s) {
	case 5:
		t, err = time.Parse("2006-01-02T15:04", time.Now().Format("2006-01-02")+"T"+s)
	case 16:
		t, err = time.Parse("2006-01-02T15:04", s)
	default:
		err = errors.New("data/hora inválida")
	}
	return t, err
}
