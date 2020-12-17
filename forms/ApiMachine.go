package forms

type AddUser struct {
	Idk  string `json:"Idk" `
	Nama string `json:"Nama" `
	Sn   string `json:"sn" `
}

type Rambu struct {
	Id            string `json:"_id" bson:"_id,omitempty"`
	RuasJalan     string `json:"ruasJalan" bson:"ruasJalan"`
	StatusJalan   string `json:"statusJalan" bson:"statusJalan"`
	Km            string `json:"km" bson:"km"`
	Tanggal       string `json:"tanggal" bson:"tanggal"`
	DataLalin     string `json:"dataLalin" bson:"dataLalin"`
	Lat           string `json:"lat" bson:"lat"`
	Lng           string `json:"lng" bson:"lng"`
	NamaJalan     string `json:"namaJalan" bson:"namaJalan"`
	JenisRambu    string `json:"jenisRambu" bson:"jenisRambu"`
	TabelRambu    string `json:"tabelRambu" bson:"tabelRambu"`
	Rambu         string `json:"rambu" bson:"rambu"`
	KodeRambu     string `json:"kodeRambu" bson:"kodeRambu"`
	Kategori      string `json:"kategori" bson:"kategori"`
	AsetTetap     string `json:"asetTetap" bson:"asetTetap"`
	Kondisi       string `json:"kondisi" bson:"kondisi"`
	KondisiVisual string `json:"kondisiVisual" bson:"kondisiVisual"`
}

type Markateng struct {
	Id               string `json:"_id" bson:"_id,omitempty"`
	RuasJalan        string `json:"ruasJalan" bson:"ruasJalan"`
	StatusJalan      string `json:"statusJalan" bson:"statusJalan"`
	Km               string `json:"km" bson:"km"`
	Tanggal          string `json:"tanggal" bson:"tanggal"`
	DataLalin        string `json:"dataLalin" bson:"dataLalin"`
	Lat              string `json:"lat" bson:"lat"`
	Lng              string `json:"lng" bson:"lng"`
	NamaJalan        string `json:"namaJalan" bson:"namaJalan"`
	JenisMarkaTengah string `json:"jenisMarkaTengah" bson:"jenisMarkaTengah"`
	PanjangMarka     string `json:"panjangMarka" bson:"panjangMarka"`
	AsetTetap        string `json:"asetTetap" bson:"asetTetap"`
	Kondisi          string `json:"kondisi" bson:"kondisi"`
	KondisiVisual    string `json:"kondisiVisual" bson:"kondisiVisual"`
}
