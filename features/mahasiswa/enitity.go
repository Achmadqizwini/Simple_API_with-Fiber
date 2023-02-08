package mahasiswa

type Core struct {
	ID     uint
	Nama   string
	Alamat string
}

type NilaiMhs struct {
	ID           uint
	Nama         string
	Nama_Matkul  string
	Nilai_Rerata uint
}

type ServiceInterface interface {
	Create(input Core) (err error)
	Delete(id int) (err error)
	Update(input Core, id int) (err error)
	Read(id int) (data []NilaiMhs, err error)
}

type RepositoryInterface interface {
	Create(input Core) (err error)
	Delete(id int) (err error)
	Update(input Core, id int) (err error)
	Read(id int) (data []NilaiMhs, err error)
}
