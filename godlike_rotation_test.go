package main

import (
	"testing"
)

func TestGetFirstFullMugglePage(t *testing.T) {
	type args struct {
		data GivenData
	}
	tests := []struct {
		name                     string
		args                     args
		wantFirstMugglePage      int
		wantMuggleFillInPrevPage int
	}{
		{
			name: "Primary 5, Secondary Premier 23. In Page 2 Sisa 3 Premier 0 Primary. Butuh 17 Muggle",
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   5,
					NumbersOf2ndPremier: 23,
					NumberOfMuggleData:  100,
				},
			},
			wantFirstMugglePage:      3,
			wantMuggleFillInPrevPage: 17,
		},
		{
			name: "Primary 15, Secondary Premier 23. In Page 2 Sisa 3 Premier 10 Primary. Butuh 7 Muggle",
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   15,
					NumbersOf2ndPremier: 23,
					NumberOfMuggleData:  100,
				},
			},
			wantFirstMugglePage:      3,
			wantMuggleFillInPrevPage: 7,
		},
		{
			name: "Primary 0, Secondary Premier 0. In Page 1 Sisa 0 Premier 0 Primary. Butuh 20 Muggle",
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   0,
					NumbersOf2ndPremier: 0,
					NumberOfMuggleData:  100,
				},
			},
			wantFirstMugglePage:      2,
			wantMuggleFillInPrevPage: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstMugglePage, gotMuggleFillInPrevPage := GetFirstFullMugglePage(tt.args.data)
			if gotFirstMugglePage != tt.wantFirstMugglePage {
				t.Errorf("GetFirstFullMugglePage() gotFirstMugglePage = %v, want %v", gotFirstMugglePage, tt.wantFirstMugglePage)
			}
			if gotMuggleFillInPrevPage != tt.wantMuggleFillInPrevPage {
				t.Errorf("GetFirstFullMugglePage() gotMuggleFillInPrevPage = %v, want %v", gotMuggleFillInPrevPage, tt.wantMuggleFillInPrevPage)
			}
		})
	}
}

func Test_play_Play(t *testing.T) {
	type args struct {
		data GivenData
	}
	tests := []struct {
		name string
		p    *play
		args args
	}{
		{
			name: "Happy Path 1 - Slotted Page, Mixed Page, Full Muggle",
			p:    &play{},
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   15,
					NumbersOf2ndPremier: 23,
					NumberOfMuggleData:  80,
				},
			},
		},
		{
			name: "Happy Path 2 - Slotted Page, Mixed Page, Full Muggle",
			p:    &play{},
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   5,
					NumbersOf2ndPremier: 23,
					NumberOfMuggleData:  40,
				},
			},
		},
		{
			name: "Happy Path 3 - Mixed Page, Full Muggle",
			p:    &play{},
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   5,
					NumbersOf2ndPremier: 5,
					NumberOfMuggleData:  40,
				},
			},
		},
		{
			name: "Full Muggle",
			p:    &play{},
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   0,
					NumbersOf2ndPremier: 0,
					NumberOfMuggleData:  40,
				},
			},
		},
		{
			name: "No PASTI & LGP",
			p:    &play{},
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   0,
					NumbersOf2ndPremier: 10,
					NumberOfMuggleData:  40,
				},
			},
		},
		{
			name: "No Premier",
			p:    &play{},
			args: args{
				data: GivenData{
					NumbersOfPASTILGP:   6,
					NumbersOf2ndPremier: 0,
					NumberOfMuggleData:  40,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &play{}
			p.Play(tt.args.data)
		})
	}
}
