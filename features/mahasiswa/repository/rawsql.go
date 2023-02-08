package repository

import (
	"be13/ca/features/mahasiswa"
	"database/sql"
	"log"
)

type mahasiswaRepository struct {
	db *sql.DB
}

func NewRaw(db *sql.DB) mahasiswa.RepositoryInterface {
	return &mahasiswaRepository{
		db: db,
	}
}

// Create implements user.RepositoryInterface
func (repo *mahasiswaRepository) Create(input mahasiswa.Core) (err error) {

	_, errExec := repo.db.Exec(("Insert into mahasiswa (Nama, Alamat) Values (?, ?)"), input.Nama, input.Alamat)
	if errExec != nil {
		return errExec
	}
	return nil
}

// DeleteUser implements mahasiswa.RepositoryInterface
func (repo *mahasiswaRepository) Delete(id int) (err error) {

	_, errExec := repo.db.Exec(("Delete from mahasiswa where id in (?)"), id)
	if errExec != nil {
		return errExec
	}
	return nil
}

// UpdateUser implements mahasiswa.RepositoryInterface
func (repo *mahasiswaRepository) Update(input mahasiswa.Core, id int) (err error) {

	_, errExec := repo.db.Exec(("Update mahasiswa set nama = ?, alamat = ? where id = ?"), input.Nama, input.Alamat, id)
	if errExec != nil {
		return errExec
	}
	return nil
}

// Read implements mahasiswa.RepositoryInterface
func (repo *mahasiswaRepository) Read(id int) (data []mahasiswa.NilaiMhs, err error) {

	result, errSelect := repo.db.Query("select u.nama, v.nama_matkul, avg(nilai) from nilai inner join mahasiswa u on nilai.mahasiswa_id = u.id inner join matkul v on nilai.matkul_id = v.id where nilai.mahasiswa_id = (?)", id)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
	}

	var mhs []mahasiswa.NilaiMhs
	for result.Next() {
		var mhsRow mahasiswa.NilaiMhs

		errScan := result.Scan(&mhsRow.Nama, &mhsRow.Nama_Matkul, &mhsRow.Nilai_Rerata)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}

		mhs = append(mhs, mhsRow)
	}
	return mhs, nil
}
