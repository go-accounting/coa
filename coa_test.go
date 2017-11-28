package coa

import (
	"testing"
	"time"
)

type store map[string][]byte

func (s store) Get(key []byte) ([]byte, error)     { return s[string(key)], nil }
func (s store) Put(key []byte, value []byte) error { s[string(key)] = value; return nil }

func TestSaveChartOfAccounts(t *testing.T) {
	r := NewCoaRepository(store{})
	coas, err := r.AllChartsOfAccounts()
	if err != nil {
		t.Error(err)
	}
	if len(coas) != 0 {
		t.Error("Repository is not empty")
	}
	coa := &ChartOfAccounts{Name: "coa"}
	coa, err = r.SaveChartOfAccounts(coa)
	if err != nil {
		t.Error(err)
	}
	if coa.Id == "" {
		t.Error("coa.Id is empty")
	}
	if coa.AsOf.IsZero() {
		t.Error("coa.AsOf is zero")
	}
	if coa.Created.IsZero() {
		t.Error("coa.Created is zero")
	}
	coas, err = r.AllChartsOfAccounts()
	if err != nil {
		t.Error(err)
	}
	if len(coas) != 1 {
		t.Errorf("Expected repository length 1 but was %v", len(coas))
	}
	if coa.Id != coas[0].Id {
		t.Errorf("Expected coa.Id==coas[0].Id %v %v", coa.Id, coas[0].Id)
	}
	coa.Name = "coacoa"
	oldId := coa.Id
	oldAsOf := coa.AsOf
	time.Sleep(1 * time.Millisecond)
	coa, err = r.SaveChartOfAccounts(coa)
	if err != nil {
		t.Error(err)
	}
	if coa.Id != oldId {
		t.Errorf("Expected coa.Id was not changed %v %v", coa.Id, oldId)
	}
	if !coa.AsOf.After(oldAsOf) {
		t.Errorf("Expected coa.AsOf is updated %v %v", coa.AsOf, oldAsOf)
	}
	coas, err = r.AllChartsOfAccounts()
	if err != nil {
		t.Error(err)
	}
	if len(coas) != 1 {
		t.Errorf("Expected repository length 1 but was %v", len(coas))
	}
	if coa.Id != coas[0].Id {
		t.Errorf("Expected coa.Id==coas[0].Id %v %v", coa.Id, coas[0].Id)
	}
	if coas[0].Name != "coacoa" {
		t.Errorf("Expected coa.Name==coacoa but was %v", coas[0].Name)
	}
}

func TestIfAllChartsOfAccountsIsSorted(t *testing.T) {
	r := NewCoaRepository(store{})
	r.SaveChartOfAccounts(&ChartOfAccounts{Name: "2"})
	r.SaveChartOfAccounts(&ChartOfAccounts{Name: "1"})
	coas, _ := r.AllChartsOfAccounts()
	if len(coas) != 2 {
		t.Errorf("Expected len(coas)==2 but was %v", len(coas))
	}
	if coas[0].Name != "1" || coas[1].Name != "2" {
		t.Errorf("Expected coas[0].Name==1 and coas[1].Name == 2 but was %v %v", coas[0].Name, coas[1].Name)
	}
}

func TestSaveAccount(t *testing.T) {
	r := NewCoaRepository(store{})
	coa, err := r.SaveChartOfAccounts(&ChartOfAccounts{Name: "coa"})
	if err != nil {
		t.Fatal(err)
	}
	a, err := r.SaveAccount(coa.Id, &Account{Number: "1", Name: "a1", Tags: []string{"balanceSheet", "increaseOnDebit"}})
	if err != nil {
		t.Fatal(err)
	}
	if a.Id == "" {
		t.Error("Id must not be null")
	}
	if a.Number != "1" {
		t.Errorf("The number (%v) must be 1", a.Number)
	}
	if a.Name != "a1" {
		t.Errorf("The name (%v) must be 'account'", a.Name)
	}
	accounts, err := r.AllAccounts(coa.Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) != 1 {
		t.Error("The account must be persisted")
	}
	if !a.Tags.Contains("detail") {
		t.Errorf("a.Tags %v does not contain detail", a.Tags)
	}
	_, err = r.SaveAccount(coa.Id, &Account{Number: "1.1", Name: "a1.1",
		Parent: a.Id, Tags: []string{"balanceSheet", "increaseOnDebit"}})
	if err != nil {
		t.Fatal(err)
	}
	accounts, err = r.AllAccounts(coa.Id)
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) != 2 {
		t.Error("The account must be persisted")
	}
	if !accounts[0].Tags.Contains("summary") {
		t.Errorf("accounts[0].Tags %v does not contains summary", accounts[0].Tags)
	}
	if !accounts[1].Tags.Contains("detail") {
		t.Errorf("accounts[1].Tags %v does not contains detail", accounts[1].Tags)
	}
	if coa.RetainedEarningsAccount != "" {
		t.Errorf("Expected empty but was %v", coa.RetainedEarningsAccount)
	}
	accounts[1].Tags = append(accounts[1].Tags, "retainedEarnings")
	_, err = r.SaveAccount(coa.Id, accounts[1])
	if err != nil {
		t.Fatal(err)
	}
	coa, err = r.GetChartOfAccounts(coa.Id)
	if err != nil {
		t.Fatal(err)
	}
	if coa.RetainedEarningsAccount != accounts[1].Id {
		t.Errorf("Expected %v but was %v", accounts[1].Id, coa.RetainedEarningsAccount)
	}
}

func TestIfAllAccountsIsSorted(t *testing.T) {
	r := NewCoaRepository(store{})
	coa, err := r.SaveChartOfAccounts(&ChartOfAccounts{Name: "coa"})
	check(t, err)
	_, err = r.SaveAccount(coa.Id, &Account{Number: "2", Name: "a2", Tags: []string{"balanceSheet", "increaseOnDebit"}})
	check(t, err)
	_, err = r.SaveAccount(coa.Id, &Account{Number: "1", Name: "a1", Tags: []string{"balanceSheet", "increaseOnDebit"}})
	check(t, err)
	accounts, err := r.AllAccounts(coa.Id)
	check(t, err)
	if accounts[0].Number != "1" || accounts[1].Number != "2" {
		t.Errorf("Expected sorted but was %v %v", accounts[0].Number, accounts[1].Number)
	}
	idx, err := r.Indexes(coa.Id, []string{accounts[0].Id, accounts[1].Id}, nil)
	check(t, err)
	if idx[0] != 1 || idx[1] != 0 {
		t.Errorf("Expected 1 0 but was %v %v", idx[0], idx[1])
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
