package main

type JsonResponse struct {
	Header struct {
		Description string   `json:"description"`
		Columns     []string `json:"columns"`
		ResultCode  string   `json:"resultCode"`
		ResultMsg   string   `json:"resultMsg"`
	} `json:"header"`
	Body struct {
		Items []struct {
			BizesID    string `json:"bizesId"`
			BizesNm    string `json:"bizesNm"`
			BrchNm     string `json:"brchNm"`
			IndsLclsCd string `json:"indsLclsCd"`
			IndsLclsNm string `json:"indsLclsNm"`
			IndsMclsCd string `json:"indsMclsCd"`
			IndsMclsNm string `json:"indsMclsNm"`
			IndsSclsCd string `json:"indsSclsCd"`
			IndsSclsNm string `json:"indsSclsNm"`
			KsicCd     string `json:"ksicCd"`
			KsicNm     string `json:"ksicNm"`
			CtprvnCd   string `json:"ctprvnCd"`
			CtprvnNm   string `json:"ctprvnNm"`
			SignguCd   string `json:"signguCd"`
			SignguNm   string `json:"signguNm"`
			AdongCd    string `json:"adongCd"`
			AdongNm    string `json:"adongNm"`
			LdongCd    string `json:"ldongCd"`
			LdongNm    string `json:"ldongNm"`
			LnoCd      string `json:"lnoCd"`
			PlotSctCd  string `json:"plotSctCd"`
			PlotSctNm  string `json:"plotSctNm"`
			LnoMnno    int    `json:"lnoMnno"`
			LnoSlno    int    `json:"lnoSlno"`
			LnoAdr     string `json:"lnoAdr"`
			RdnmCd     string `json:"rdnmCd"`
			Rdnm       string `json:"rdnm"`
			BldMnno    int    `json:"bldMnno"`
			// BldSlno    string  `json:"bldSlno"`
			BldMngNo string  `json:"bldMngNo"`
			BldNm    string  `json:"bldNm"`
			RdnmAdr  string  `json:"rdnmAdr"`
			OldZipcd string  `json:"oldZipcd"`
			NewZipcd string  `json:"newZipcd"`
			DongNo   string  `json:"dongNo"`
			FlrNo    string  `json:"flrNo"`
			HoNo     string  `json:"hoNo"`
			Lon      float64 `json:"lon"`
			Lat      float64 `json:"lat"`
		} `json:"items"`
	} `json:"body"`
}
