package main

import (
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/rgmfn/scsboa-grid-data-tables/templates"
)

const port = 8008

//go:embed script.js
var scriptsJs []byte

type Row struct {
	GroupName   string `json:"group_name"`
	Division    string `json:"division"`
	Date        string `json:"date"`
	Comp        string `json:"comp"`
	Mp1Stt      string `json:"mp_1_stt"`
	Mp1Mus      string `json:"mp_1_mus"`
	Mp1Tot      string `json:"mp_1_tot"`
	Mp2Stt      string `json:"mp_2_stt"`
	Mp2Mus      string `json:"mp_2_mus"`
	Mp2Tot      string `json:"mp_2_tot"`
	MpTot       string `json:"mp_tot"`
	MeComp      string `json:"me_comp"`
	MeComm      string `json:"me_comm"`
	MeTot       string `json:"me_tot"`
	VeComp      string `json:"ve_comp"`
	VeComm      string `json:"ve_comm"`
	VeTot       string `json:"ve_tot"`
	GeTot       string `json:"ge_tot"`
	VpComp      string `json:"vp_comp"`
	VpExc       string `json:"vp_exc"`
	VpTot       string `json:"vp_tot"`
	SubTot      string `json:"sub_tot"`
	Pen         string `json:"pen"`
	PenTot      string `json:"pen_tot"`
	OverallBand string `json:"overall_band"`
	AuxComp     string `json:"aux_comp"`
	AuxAch      string `json:"aux_ach"`
	AuxTot      string `json:"aux_tot"`
	AuxAdjTot   string `json:"aux_adj_tot"`
	PercComp    string `json:"perc_comp"`
	PercAch     string `json:"perc_ach"`
	PercTot     string `json:"perc_tot"`
	PercAdjTot  string `json:"perc_adj_tot"`
}

func main() {
	r := chi.NewRouter()

	// files
	r.Get("/script.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/javascipt")
		w.Write(scriptsJs)
	})

	// resources
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templates.Home()).ServeHTTP(w, r)
	})

	// endpoints
	r.Route("/api", func(r chi.Router) {
		r.Get("/get-performances", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			file, err := os.Open("scsboa.csv")

			if err != nil {
				log.Fatal("Error reading file", err)
			}
			defer file.Close()

			reader := csv.NewReader(file)

			records, err := reader.ReadAll()

			rows := []Row{}
			for index, record := range records {
				if index == 0 {
					continue
				}

				row := Row{
					GroupName:   record[1],
					Division:    record[2],
					Date:        record[3],
					Comp:        record[4],
					Mp1Stt:      record[5],
					Mp1Mus:      record[6],
					Mp1Tot:      record[7],
					Mp2Stt:      record[8],
					Mp2Mus:      record[9],
					Mp2Tot:      record[10],
					MpTot:       record[11],
					MeComp:      record[12],
					MeComm:      record[13],
					MeTot:       record[14],
					VeComp:      record[15],
					VeComm:      record[16],
					VeTot:       record[17],
					GeTot:       record[18],
					VpComp:      record[19],
					VpExc:       record[20],
					VpTot:       record[21],
					SubTot:      record[22],
					Pen:         record[23],
					PenTot:      record[24],
					OverallBand: record[25],
					AuxComp:     record[26],
					AuxAch:      record[27],
					AuxTot:      record[28],
					AuxAdjTot:   record[29],
					PercComp:    record[30],
					PercAch:     record[31],
					PercTot:     record[32],
					PercAdjTot:  record[33],
				}
				rows = append(rows, row)
			}

// 			log.Printf("%+v", rows)

			json.NewEncoder(w).Encode(rows)
		})
	})

	log.Printf("Starting server on port %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		panic(err)
	}
}
