<h1> for first up mysql container</h1>
docker-compose -f docker-compose.yml build
<h1> for next attempts</h1>
docker-compose -f docker-compose.yml up -d db



dbeaver string connection
jdbc:mysql://localhost:3306/golang?allowPublicKeyRetrieval=true&useSSL=false



init sql code 

CREATE TABLE users (
	id INT auto_increment not null primary key,
	name varchar(100) NULL,
	age INT NULL
)

<br>
CREATE TABLE golang.products (
	Id INT auto_increment NULL,
	Model varchar(100) NULL,
	Company varchar(100) NULL,
	Price BIGINT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;


<br>
insert into products (model, company, price) 
values ('iPhone X', 'Apple', 74000),
('Pixel 2', 'Google', 62000),
('Galaxy S9', 'Samsung', 65000)