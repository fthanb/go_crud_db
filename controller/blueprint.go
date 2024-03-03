package controller

import (
	"backend-gdsc/config"
	"context"
	"database/sql"
	"fmt"
)

type Mahasiswa struct {
	NIM      string
	Nama     string
	TotalSKS int
}

func Insert(mhs Mahasiswa) error {
	db := config.Connect_db()
	ctx := context.Background()

	state := "INSERT INTO Mahasiswa(nim, nama, total_sks) VALUES (?, ?, ?)"
	_, err := db.ExecContext(ctx, state, mhs.NIM, mhs.Nama, mhs.TotalSKS)
	if err != nil {
		return err
	}
	return nil
}

func Select() error {
	db := config.Connect_db()
	ctx := context.Background()

	state := "SELECT nim, nama, total_sks FROM Mahasiswa"
	rows, err := db.QueryContext(ctx, state)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nim, nama string
		var totalSKS int
		err := rows.Scan(&nim, &nama, &totalSKS)
		if err != nil {
			return err
		}
		fmt.Println("NIM:", nim)
		fmt.Println("Nama:", nama)
		fmt.Println("Total SKS:", totalSKS)
	}
	return nil
}

func SelectByNIM(nim string) error {
	db := config.Connect_db()
	ctx := context.Background()

	state := "SELECT nim, nama, total_sks FROM Mahasiswa WHERE nim = ?"
	row := db.QueryRowContext(ctx, state, nim)

	var mhs Mahasiswa
	err := row.Scan(&mhs.NIM, &mhs.Nama, &mhs.TotalSKS)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan")
			return nil
		}
		return err
	}

	fmt.Println("NIM:", mhs.NIM)
	fmt.Println("Nama:", mhs.Nama)
	fmt.Println("Total SKS:", mhs.TotalSKS)
	return nil
}

func Delete(nim string) error {
	db := config.Connect_db()
	ctx := context.Background()

	state := "DELETE FROM Mahasiswa WHERE nim = ?"
	_, err := db.ExecContext(ctx, state, nim)
	if err != nil {
		return err
	}
	fmt.Println("Data Mahasiswa dengan NIM", nim, "telah dihapus")
	return nil
}
