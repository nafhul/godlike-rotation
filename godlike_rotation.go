package main

import (
	"fmt"
	"math"
)

type GivenData struct {
	NumbersOfPASTILGP   int
	NumbersOf2ndPremier int
	NumberOfMuggleData  int
}

const MAX_PRIMARY_ON_PAGE int = 5
const MAX_SECONDARY_ON_PAGE int = 20

// const MAX_PAGE int = 6

type Param struct {
	Page int
}

var Data GivenData
var PrimaryData []string
var PremierData []string
var MuggleData []string

func Reset() {
	PremierData = []string{}
	PrimaryData = []string{}
	MuggleData = []string{}
}

func GetCountPASTILGP() int {
	return Data.NumbersOfPASTILGP
}

func GetCountPremier() int {
	return Data.NumbersOf2ndPremier
}

func PopulateData(data GivenData) {
	Data = data
	for i := 0; i < Data.NumbersOfPASTILGP; i++ {
		PrimaryData = append(PrimaryData, fmt.Sprintf("PASTI/LGP %d", i+1))
	}

	for i := 0; i < Data.NumbersOf2ndPremier; i++ {
		PremierData = append(PremierData, fmt.Sprintf("Secondary Premier %d", i+1))
	}

	for i := 0; i < Data.NumberOfMuggleData; i++ {
		MuggleData = append(MuggleData, fmt.Sprintf("Muggle Data %d", i+1))
	}

	// fmt.Println(MuggleData)
}

func main() {
	p := NewPlay()
	data := GivenData{
		NumbersOfPASTILGP:   5,
		NumbersOf2ndPremier: 20,
		NumberOfMuggleData:  100,
	}
	p.Play(data)
}

type play struct{}

func NewPlay() *play {
	return &play{}
}

func (p *play) Play(data GivenData) {
	PopulateData(data)
	for i := 1; i <= int(GetTotalPage(data)); i++ {
		if i == 1 {
			continue
		}
		param := Param{
			Page: i,
		}
		fmt.Println("---------------------------------------------")
		fmt.Printf("Page: %d \n", param.Page)
		fmt.Printf("Data: %+v \n", Data)
		fmt.Printf("Total Page: %.0f \n", GetTotalPage(data))
		if IsRemainingPrimaryExist(param) {
			SlottingRule(param)
		} else if IsRemainingSecondaryPremierExist(param) {
			SikatAbisPrimary(param)
		} else {
			BurnAllMuggle(param)
		}
	}
	Reset()
}

func IsRemainingPrimaryExist(param Param) bool {
	if GetCountPASTILGP() == 0 {
		return false
	}
	fullhouse := GetCountPremier() > param.Page*MAX_SECONDARY_ON_PAGE
	if fullhouse {
		return fullhouse
	}

	remainingPrimary := GetCountPASTILGP() - ((param.Page - 1) * 20)
	if remainingPrimary > 0 {
		return true
	}
	// fmt.Printf("IsPremierFullhouse = %d > %d*%d = %t \n", GetCountPremier(), param.Page, MAX_SECONDARY_ON_PAGE, result)

	return fullhouse
}

func IsRemainingSecondaryPremierExist(param Param) bool {
	if GetCountPremier() == 0 {
		return false
	}
	result := GetCountPremier()-((param.Page-1)*MAX_SECONDARY_ON_PAGE) > 0
	fmt.Printf("IsRemainingSecondaryPremierExist = %t \n", result)
	return result
}

func SlottingRule(param Param) {
	/**
	 * 1. Count Primary Project PASTI&LGP
	 * 2. Count Premier Secondary
	 * ^ Concurrent Boleh
	 * 3. Plating, based on slot rule. Primary di 1,5,10,15,20
	 */
	slot := []int{1, 5, 10, 15, 20}
	PrimaryOffset := (param.Page - 1) * MAX_PRIMARY_ON_PAGE
	PrimaryEndIndex := PrimaryOffset + MAX_PRIMARY_ON_PAGE
	fmt.Printf("Primary Offset: %d \n", PrimaryOffset)
	fmt.Printf("Primary End Index: %d \n", PrimaryEndIndex)

	remainingSecondaryPremier := GetCountPremier() - ((param.Page - 1) * 20)
	if remainingSecondaryPremier < MAX_SECONDARY_ON_PAGE {
		SikatAbisPrimary(param)
		return
	}
	if remainingSecondaryPremier > MAX_SECONDARY_ON_PAGE {
		remainingSecondaryPremier = MAX_SECONDARY_ON_PAGE
	}
	SecondaryOffset := (param.Page - 1) * MAX_SECONDARY_ON_PAGE
	SecondaryEndIndex := SecondaryOffset + remainingSecondaryPremier
	fmt.Printf("Secondary Premier Offset: %d \n", SecondaryOffset+1)
	fmt.Printf("Secondary Premier End Index: %d \n", SecondaryEndIndex)

	FinalPageResult := make([]string, MAX_SECONDARY_ON_PAGE)
	PrimaryResult := make([]string, MAX_PRIMARY_ON_PAGE)
	copy(FinalPageResult, PremierData[SecondaryOffset:SecondaryEndIndex])
	copy(PrimaryResult, PrimaryData[PrimaryOffset:PrimaryEndIndex])

	isAbisinPrimary := false
	for k, v := range slot {
		if v > remainingSecondaryPremier {
			isAbisinPrimary = true
			break
		}
		FinalPageResult = append(FinalPageResult, "new slot")
		insert(FinalPageResult, PrimaryResult[k], v-1)
	}

	if isAbisinPrimary {
		SikatAbisPrimary(param)
	}

	for k, v := range FinalPageResult {
		fmt.Printf("%d. %s \n", k+1, v)
	}
}

func SikatAbisPrimary(param Param) {
	/**
	 * Kondisi:
	 * Premier Listing Tidak mencukupi untuk fill in satu halaman penuh
	 *
	 * Prerequisites:
	 * 1. Count Primary Project PASTI&LGP
	 * 2. Count Premier Secondary
	 *
	 * Proses:
	 * 1. Cari Sisa Secondary Premier
	 * 2. Ambil semua sisa PASTI & LGP
	 * ^ Concurrent Boleh
	 * 3. Plating, based on slot rule. Primary di 1,5,10,15,20
	 */

	remainingSecondaryPremier := Data.NumbersOf2ndPremier - ((param.Page - 1) * MAX_SECONDARY_ON_PAGE)
	remainingPrimaryPASTILGP := Data.NumbersOfPASTILGP - ((param.Page - 1) * MAX_PRIMARY_ON_PAGE)
	fmt.Printf("Remaining secondary Premier = %d \n", remainingSecondaryPremier)
	fmt.Printf("Remaining primary project = %d \n", remainingPrimaryPASTILGP)

	prevLastIndexPrimary := (param.Page - 1) * MAX_PRIMARY_ON_PAGE
	sikatAbisPrimaryIndex := Data.NumbersOfPASTILGP
	// fmt.Printf("Prev Primary Index = %d \n", prevLastIndexPrimary)

	secondaryOffset := (param.Page - 1) * 20
	secondaryLastIndex := secondaryOffset + remainingSecondaryPremier
	// fmt.Printf("Prev Secondary Premier Index = %d \n", secondaryOffset)
	// fmt.Printf("Prev Last Index = %d \n", secondaryLastIndex)

	filledIn := remainingPrimaryPASTILGP + remainingSecondaryPremier
	var additionalData int
	if filledIn < MAX_SECONDARY_ON_PAGE {
		additionalData = MAX_SECONDARY_ON_PAGE - filledIn
	}
	// fmt.Printf("Additional Muggle Data = %d \n", additionalData)

	slot := []int{1, 5, 10, 15, 20}
	FinalPageResult := make([]string, remainingSecondaryPremier)
	PrimaryResult := make([]string, remainingPrimaryPASTILGP)
	MuggleResult := make([]string, additionalData)
	copy(FinalPageResult, PremierData[secondaryOffset:secondaryLastIndex])
	copy(PrimaryResult, PrimaryData[prevLastIndexPrimary:sikatAbisPrimaryIndex])
	copy(MuggleResult, MuggleData[0:additionalData])

	var usedSlot int
	if remainingPrimaryPASTILGP > 0 {
		for k, v := range slot {
			if v <= remainingSecondaryPremier {
				FinalPageResult = append(FinalPageResult, "new slot")
				insert(FinalPageResult, PrimaryResult[k], v-1)
				usedSlot++
			}
		}
	}

	sikatAbis := usedSlot + (remainingPrimaryPASTILGP - usedSlot)
	for i := usedSlot; i < sikatAbis; i++ {
		FinalPageResult = append(FinalPageResult, PrimaryResult[i])
	}

	lastIndex := len(FinalPageResult) - 1
	j := 0
	for i := lastIndex; i < MAX_SECONDARY_ON_PAGE-1; i++ {
		FinalPageResult = append(FinalPageResult, MuggleResult[j])
		j++
	}

	for k, v := range FinalPageResult {
		fmt.Printf("%d. %s \n", k+1, v)
	}
}

func BurnAllMuggle(param Param) {
	/**
	 * Kondisi:
	 * PASTI & LGP Habis
	 * Secondary Premier Habis
	 *
	 * Prerequisites:
	 * 1. Count Primary Project PASTI&LGP
	 * 2. Count Premier Secondary
	 *
	 * Proses:
	 *
	 */
	firstFullMugglePage, muggleFillInMixedPage := GetFirstFullMugglePage(Data)

	var Offset int
	if param.Page < firstFullMugglePage {
		Offset = 0
	} else if muggleFillInMixedPage == MAX_SECONDARY_ON_PAGE && firstFullMugglePage != param.Page {
		Offset = ((param.Page - firstFullMugglePage) * 20)
	} else {
		Offset = muggleFillInMixedPage + ((param.Page - firstFullMugglePage) * 20)
	}

	endIndex := Offset + MAX_SECONDARY_ON_PAGE
	// fmt.Println(Offset)
	// fmt.Println(endIndex)
	// fmt.Println(param.Page)
	// fmt.Println(firstFullMugglePage)

	FinalPageResult := make([]string, MAX_SECONDARY_ON_PAGE)
	copy(FinalPageResult, MuggleData[Offset:endIndex])

	for k, v := range FinalPageResult {
		if len(v) > 1 {
			fmt.Printf("%d. %s \n", k+1, v)
		}
	}
}

func GetFirstFullMugglePage(data GivenData) (firstMugglePage int, muggleFillInMixedPage int) {
	/**
	 * 1. Get Total page of Premier
	 * 2. Get Premier items of last mixed page
	 * 3. Get Primary items of last mixed page
	 * 4. Get muggle on page
	 */
	premierPage := data.NumbersOf2ndPremier / MAX_SECONDARY_ON_PAGE
	primaryFloorOver := data.NumbersOf2ndPremier % MAX_SECONDARY_ON_PAGE
	abisinPrimary := data.NumbersOfPASTILGP - (premierPage * MAX_PRIMARY_ON_PAGE)
	totalNonMuggleItems := primaryFloorOver + abisinPrimary
	muggleCount := MAX_SECONDARY_ON_PAGE - totalNonMuggleItems

	if muggleCount > 0 && premierPage > 0 {
		return premierPage + 2, muggleCount
	} else if totalNonMuggleItems < MAX_SECONDARY_ON_PAGE {
		return 2, muggleCount
	}

	return 1, muggleCount
}

func insert(array []string, object string, position int) []string {
	return append(array[:position], append([]string{object}, array[position:]...)...)
}

func GetTotalPage(data GivenData) float64 {
	totalData := data.NumbersOfPASTILGP + data.NumbersOf2ndPremier + data.NumberOfMuggleData
	totalPage := float64(totalData) / float64(MAX_SECONDARY_ON_PAGE)
	fmt.Printf("Total Data %d \n", totalData)
	return math.Ceil(totalPage)
}
