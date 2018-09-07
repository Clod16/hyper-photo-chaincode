package main


type Description struct{
	ISO          string `json:"iso"`
	Exposition   string `json:"exposition"`
	ShutterSpeed string `json:"shutterSpeed"`
	Baffle       string `json:"Baffle"`
	
}

type Photo struct{
	ID          string `json:"id"`
	Author      string `json:"author"`
	Date        string `json:"date"`
	Format      string `json:"format"`
	Status      string `json:"status"`
	Description `json:"description"`
	

}