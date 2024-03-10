# Simple Golang REST API with PostgreSQL
This is a project for creating APIs with.
 - `Login` and `Register` API Function with JWT Auth.
 - `Create` API Function - It can add an employee and fill the created_at column in the employee table.
 - `Read` API Function - It can view detailed information of the selected employee data.
 - `Update` API Function - It can modify employee data and update the updated_at column in the employee table when modifying employee data.
 - `Delete` API Function - It performs deletion of employee data using soft delete, allowing deleted data to be restored.

## Getting Started
You can download this repo or clone using below command. (folder-name will be project folder in which you want to start your project).
```
git clone https://github.com/latifis/go-restapi.git <folder-name>
```
or from **Download Zip**
```
https://github.com/latifis/go-restapi
```

### Installing And Running
1. Instalation

To install dependencies, run the following command in your terminal or command prompt:
```
> go mod download
```
2. Setup Database

After installing dependencies, you need to configure the database settings. Open the `setup` file in `models` folder and configure it according to the database you are using:
```
go-restapi/models/setup.go
```
3. Running The Program

Once dependencies are installed and the database is configured, start the application with the following command:
```
> go main.go
```

### API Documentation
To view documentation for using this program, you can go to the [documentation](https://documenter.getpostman.com/view/19724773/2sA2xfXsux), or at the following link:
```
https://documenter.getpostman.com/view/19724773/2sA2xfXsux
```

### Preview

- Login
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/login-berhasil.png?raw=true)
 - Login Failed
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/login-gagal.png?raw=true)
 - Register
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/register.png?raw=true)
- Failed Register when User Already Exists 
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/when-user-already-exists.png?raw=true)
 - Get Data Unauthorized
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/get-data-unauthorized.png?raw=true)
 - Create Karyawan
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/create-karyawan.png?raw=true)
 - Get All Data Karyawan
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/get-all-data-karyawan.png?raw=true)
 - Get Data Karyawan by ID
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/get-data-karyawan-by-id.png?raw=true)
 - Update Data Karyawan
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/update-data-karyawan.png?raw=true)
 - Soft Delete Karyawan
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/soft-delete-karyawan.png?raw=true)
 - Get Karyawan When They Have Been Soft Delete
 ![alt text](https://github.com/latifis/etc/blob/main/go-restapi/get-karyawan-when-they-have-been-soft-delete.png?raw=true)