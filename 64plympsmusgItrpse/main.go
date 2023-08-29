package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}
type fixedIncome struct {
	project   string
	bidAmount int
}
type timeAndMaterial struct {
	project     string
	totalHrs    int
	costPerHour int
}
type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}
func (p fixedIncome) source() string {
	return p.project
}
func (p fixedIncome) calculate() int {
	return p.bidAmount
}
func (p timeAndMaterial) source() string {
	return p.project
}
func (p timeAndMaterial) calculate() int {
	return p.costPerHour * p.totalHrs
}
func totalIncome(ti []Income) {
	netIncome := 0
	for _, v := range ti {
		fmt.Printf("Project name: %s, Total amount:$ %d \n", v.source(), v.calculate())
		netIncome += v.calculate()
	}
	fmt.Println("Total Income :", netIncome)
}
func main() {
	Project1 := fixedIncome{project: "Celito", bidAmount: 10000}
	Project2 := fixedIncome{"Sapling", 20000}
	Project3 := timeAndMaterial{"Kong", 7, 500}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	AllProjects := []Income{Project1, Project2, Project3, bannerAd, popupAd}
	totalIncome(AllProjects)
}
