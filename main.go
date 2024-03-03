package main

import (
	"backend-gdsc/controller"
	"fmt"
)

func main() {
	var choice int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Insert Mahasiswa")
		fmt.Println("2. List Mahasiswa")
		fmt.Println("3. Select Mahasiswa by NIM")
		fmt.Println("4. Delete Mahasiswa by NIM")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertMahasiswa()
		case 2:
			listMahasiswa()
		case 3:
			selectMahasiswaByNIM()
		case 4:
			deleteMahasiswaByNIM()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func insertMahasiswa() {
	var newMahasiswa controller.Mahasiswa
	fmt.Println("\nInsert Mahasiswa:")
	fmt.Print("NIM: ")
	fmt.Scanln(&newMahasiswa.NIM)
	fmt.Print("Nama: ")
	fmt.Scanln(&newMahasiswa.Nama)
	fmt.Print("Total SKS: ")
	fmt.Scanln(&newMahasiswa.TotalSKS)

	err := controller.Insert(newMahasiswa)
	if err != nil {
		fmt.Println("Error inserting Mahasiswa:", err)
	} else {
		fmt.Println("Mahasiswa inserted successfully.")
	}
}

func listMahasiswa() {
	fmt.Println("\nList Mahasiswa:")
	err := controller.Select()
	if err != nil {
		fmt.Println("Error selecting Mahasiswa:", err)
	}
}

func selectMahasiswaByNIM() {
	var nim string
	fmt.Print("\nInput NIM to select Mahasiswa: ")
	fmt.Scanln(&nim)
	err := controller.SelectByNIM(nim)
	if err != nil {
		fmt.Println("Error selecting Mahasiswa by NIM:", err)
	}
}

func deleteMahasiswaByNIM() {
	var deleteNIM string
	fmt.Print("\nInput NIM to delete Mahasiswa: ")
	fmt.Scanln(&deleteNIM)
	err := controller.Delete(deleteNIM)
	if err != nil {
		fmt.Println("Error deleting Mahasiswa by NIM:", err)
	} else {
		fmt.Println("Mahasiswa deleted successfully.")
	}
}
