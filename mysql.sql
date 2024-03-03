create database backend_gdsc;

use backend_gdsc;

create table Mahasiswa(
	nim varchar(14) primary key,
    nama varchar(255),
    total_sks int(3)
);

desc Mahasiswa;
