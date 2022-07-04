/*---------------------------------------*/
	//            CREDENCIALES DE LOCAL
	/*---------------------------------------*/
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "sasa"
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)

	/*---------------------------------------*/
	//            CREDENCIALES DE HEROKU
	/*---------------------------------------*/
	////mysql://b88d0a8e947584:823ea48f@us-cdbr-east-06.cleardb.net/heroku_8c0816ad800ce51?reconnect=true
	/*Driver := "mysql"
	Usuario := "b88d0a8e947584"
	Contrasenia := "823ea48f"
	Nombre := "heroku_8c0816ad800ce51"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(us-cdbr-east-06.cleardb.net)/"+Nombre)*/

/*---------------------------------------------------------*/
		BASE DE DATOS MYSQL
		Nota: para esta prueba de Go se utilizo myphpAdmin servidor xamp
		y se creo la BD con el nombre sitema
		y la tabla empleados con columnas
		id
		nombre
		correo

/*---------------------------------------------------------*/