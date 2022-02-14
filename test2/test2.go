package repo

import (
	"gorm.io/gorm"
)
type (
	Area struct {
		ID int64 `gorm:"column:id;primaryKey;"`
		AreaValue int64 `gorm:"column:area_value"`
		AreaType string `gorm:"column:type"`
	}
)

type AreaRepository interface {
	InsertArea(param1 int32, param2 int64, tipe string, ar *Area) (err error)
}

type areaRepo struct {
	DB *gorm.DB
}
	
// Handler memiliki function seperti dibawah ini
func (_r *areaRepo) InsertArea(param1 int32, param2 int64, tipe string, ar *Area) (err error) {
	// inst := _r.DB.Model(ar)

	var area int64
	area = 0
	switch tipe {
		case "persegi panjang":
			area = int64(param1) * param2
			ar.AreaValue = area
			ar.AreaType = "persegi panjang"
			// err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
		case "persegi":
			area = int64(param1) * param2
			ar.AreaValue = area
			ar.AreaType = "persegi"
			// err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
	
		case "segitiga":

			areaSegitiga := 0.5 * (float64(param1) * float64(param2))
			ar.AreaValue = int64(areaSegitiga)
			ar.AreaType = "segitiga"
			// err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
		default:
			ar.AreaValue = 0
			ar.AreaType = "undefined data"
			// err = _r.DB.create(&ar).Error
			if err != nil {
				return err
			}
		
	}
	return nil
}

// err = _u.repository.InsertArea(10, 10, ‘persegi’)
// if err != nil {
// log.Error().Msg(err.Error())
// err = errors.New(en.ERROR_DATABASE)
// return err
// }

// 1. Penggunaan nama variable "type" pada parameter fungsi tersebut tidak diperbolehkandan type datanya seharusnya string bukan array string.
// 2. type data pada param 1 sebaiknya menggunakan int64, atau jika tetap menggunakan int32 pada saat perkalian di dalam switch case dikonversi terlebih dahulu ke int64
// 3. type data area juga seharusnya menggunakan int64 karena data tersebut akan di assign ke field AreaValue pada struct Area
// 4. di dalam case persegi panjang pendeklarasian ulang variable area tidak diperlukan, cukup dengan assignment area =  param1 * param2
// 5. di dalam case persegi deklarasi Var juga tidak diperlukan, cukup dengan assignment area =  param1 * param2
// 6. pada case segitiga, sebaiknya dibuatkan variable baru dengan type data float64 dikarenakan pengali segitiga bertype data float
//    dan variable param1 & param2 perlu dikonversi ke type float64. setelah itu assignment ke field AreaValue pada struct Area perlu dikonversi kembali ke type data int64 
// 7. sebelum memanggil fungsi tersebut perlu di deklarasikan var dengan type data Area
// 8. pada saat memanggil fungsi tersebut variable yg telah dibuat di masukkan ke dalam parameter ke 4 fungsi sebagai pointer.